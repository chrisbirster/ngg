package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TursoBackend uses Turso's versioned SQL-over-HTTP pipeline. It intentionally
// has no CGO or database-driver dependency, which keeps the Fly image tiny.
type TursoBackend struct { endpoint, token string; client *http.Client }

func NewTursoBackend(databaseURL, token string, client *http.Client) (*TursoBackend,error) {
	if strings.TrimSpace(databaseURL)=="" { return nil,errors.New("TURSO_DATABASE_URL is required") }
	u,err:=url.Parse(databaseURL); if err!=nil{return nil,fmt.Errorf("parse Turso URL: %w",err)}
	if u.Scheme=="turso" { u.Scheme="https" }; if u.Scheme!="https" && u.Scheme!="http" { return nil,fmt.Errorf("unsupported Turso URL scheme %q",u.Scheme) }
	u.Path=strings.TrimRight(u.Path,"/")+"/v2/pipeline"
	if client==nil { client=&http.Client{Timeout:10*time.Second} }
	return &TursoBackend{endpoint:u.String(),token:token,client:client},nil
}

type tursoArg struct { Type string `json:"type"`; Value string `json:"value"` }
type tursoStmt struct { SQL string `json:"sql"`; Args []tursoArg `json:"args,omitempty"` }
type tursoRequest struct { Type string `json:"type"`; Stmt *tursoStmt `json:"stmt,omitempty"` }
type tursoPipeline struct { Requests []tursoRequest `json:"requests"` }
func textArg(v string) tursoArg { return tursoArg{Type:"text",Value:v} }
func (t *TursoBackend) run(ctx context.Context, statements ...tursoStmt) ([]json.RawMessage,error) {
	reqs:=make([]tursoRequest,0,len(statements)+1); for i:=range statements { st:=statements[i]; reqs=append(reqs,tursoRequest{Type:"execute",Stmt:&st}) }; reqs=append(reqs,tursoRequest{Type:"close"})
	body,_:=json.Marshal(tursoPipeline{Requests:reqs}); req,err:=http.NewRequestWithContext(ctx,http.MethodPost,t.endpoint,bytes.NewReader(body)); if err!=nil{return nil,err}
	req.Header.Set("Content-Type","application/json"); if t.token!="" { req.Header.Set("Authorization","Bearer "+t.token) }
	resp,err:=t.client.Do(req); if err!=nil{return nil,fmt.Errorf("Turso request: %w",err)}; defer resp.Body.Close(); raw,err:=io.ReadAll(io.LimitReader(resp.Body,4<<20)); if err!=nil{return nil,err}
	if resp.StatusCode/100!=2{return nil,fmt.Errorf("Turso returned %s: %s",resp.Status,strings.TrimSpace(string(raw)))}
	var envelope struct{ Results []json.RawMessage `json:"results"` }; if err:=json.Unmarshal(raw,&envelope);err!=nil{return nil,fmt.Errorf("decode Turso response: %w",err)}
	for _,item:=range envelope.Results { var result struct{ Type string `json:"type"`; Error *struct{Message string `json:"message"`} `json:"error"` }; _=json.Unmarshal(item,&result); if result.Type=="error" { if result.Error!=nil{return nil,errors.New(result.Error.Message)}; return nil,errors.New("Turso statement failed") } }
	return envelope.Results,nil
}

func (t *TursoBackend) Migrate(ctx context.Context) error {
	_,err:=t.run(ctx,
		tursoStmt{SQL:`CREATE TABLE IF NOT EXISTS platform_snapshots (key TEXT PRIMARY KEY, payload TEXT NOT NULL, updated_at TEXT NOT NULL)`},
		tursoStmt{SQL:`CREATE TABLE IF NOT EXISTS platform_events (id TEXT PRIMARY KEY, event_type TEXT NOT NULL, payload TEXT NOT NULL, created_at TEXT NOT NULL)`},
	); return err
}

func (t *TursoBackend) Load(ctx context.Context) ([]byte,error) {
	results,err:=t.run(ctx,tursoStmt{SQL:`SELECT payload FROM platform_snapshots WHERE key = ? LIMIT 1`,Args:[]tursoArg{textArg("community")}}); if err!=nil{return nil,err}; if len(results)==0{return nil,nil}
	var item struct{ Type string `json:"type"`; Response struct{ Result struct{ Rows [][]struct{Type,Value string} `json:"rows"` } `json:"result"` } `json:"response"` }; if err:=json.Unmarshal(results[0],&item);err!=nil{return nil,err}; if len(item.Response.Result.Rows)==0||len(item.Response.Result.Rows[0])==0{return nil,nil}; return []byte(item.Response.Result.Rows[0][0].Value),nil
}

func (t *TursoBackend) Save(ctx context.Context, payload []byte, eventType string, event []byte) error {
	now:=time.Now().UTC().Format(time.RFC3339Nano)
	_,err:=t.run(ctx,
		tursoStmt{SQL:`BEGIN IMMEDIATE`},
		tursoStmt{SQL:`INSERT INTO platform_snapshots(key,payload,updated_at) VALUES(?,?,?) ON CONFLICT(key) DO UPDATE SET payload=excluded.payload, updated_at=excluded.updated_at`,Args:[]tursoArg{textArg("community"),textArg(string(payload)),textArg(now)}},
		tursoStmt{SQL:`INSERT INTO platform_events(id,event_type,payload,created_at) VALUES(?,?,?,?)`,Args:[]tursoArg{textArg(newID("evt")),textArg(eventType),textArg(string(event)),textArg(now)}},
		tursoStmt{SQL:`COMMIT`},
	); return err
}
func (t *TursoBackend) Health(ctx context.Context) error { _,err:=t.run(ctx,tursoStmt{SQL:`SELECT 1`}); return err }

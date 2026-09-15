package platform

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTursoBackendPipeline(t *testing.T){
	var requestBody,authorization string
	server:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){authorization=r.Header.Get("Authorization");body,_:=io.ReadAll(r.Body);requestBody=string(body);w.Header().Set("Content-Type","application/json");_,_=w.Write([]byte(`{"results":[{"type":"ok","response":{"type":"execute","result":{"cols":[{"name":"payload","decltype":"TEXT"}],"rows":[[{"type":"text","value":"{\"ok\":true}"}]],"affected_row_count":0,"last_insert_rowid":null}}},{"type":"ok","response":{"type":"close"}}]}`))}))
	defer server.Close();b,err:=NewTursoBackend(server.URL,"secret",server.Client());if err!=nil{t.Fatal(err)}
	payload,err:=b.Load(context.Background());if err!=nil{t.Fatal(err)};if string(payload)!=`{"ok":true}`{t.Fatalf("payload %q",payload)}
	if authorization!="Bearer secret"{t.Fatalf("authorization %q",authorization)};if !strings.Contains(requestBody,"platform_snapshots")||!strings.Contains(requestBody,`"type":"text"`){t.Fatalf("unexpected request %s",requestBody)}
}

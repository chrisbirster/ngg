package platform

import (
	"encoding/json"
	"testing"
)

func TestCreatorProjectForumAndSDK(t *testing.T){
	s:=NewStore();u,err:=s.CreateUser("maker@example.com","Maker","Maker Person");if err!=nil{t.Fatal(err)}
	p,err:=s.CreateProject(u.ID,"game","tiny-game","Tiny Game","A tiny game",[]string{"arcade"});if err!=nil{t.Fatal(err)}
	upload,err:=s.CreateUpload(p.ID,"bundle","application/zip",123);if err!=nil{t.Fatal(err)}
	if _,err=s.CompleteUpload(upload.ID);err!=nil{t.Fatal(err)}
	if _,err=s.CreateRelease(p.ID,upload.ID,"1.0.0");err!=nil{t.Fatal(err)}
	thread,err:=s.CreateThread("creators",u.ID,"Hello creators","I shipped a game");if err!=nil{t.Fatal(err)}
	if _,err=s.AddPost(thread.ID,u.ID,"Play it!");err!=nil{t.Fatal(err)}
	if _,err=s.SubmitScore("tiny-game","high-score",u.ID,42,nil);err!=nil{t.Fatal(err)}
	if got:=s.Leaderboard("tiny-game","high-score");len(got)!=1||got[0].Score!=42{t.Fatalf("unexpected leaderboard: %#v",got)}
	save,err:=s.PutSave("tiny-game",u.ID,"main",json.RawMessage(`{"level":2}`),0);if err!=nil||save.Version!=1{t.Fatalf("unexpected save: %#v %v",save,err)}
}

func TestFiveArcadeGamesAndRoomVersioning(t *testing.T){
	s:=NewStore();if len(s.ArcadeGames())!=5{t.Fatalf("got %d arcade games",len(s.ArcadeGames()))}
	r,err:=s.CreateRoom("tic-tac-toe","one");if err!=nil{t.Fatal(err)}
	r,err=s.JoinRoom(r.ID,"two");if err!=nil||r.Status!="playing"{t.Fatalf("room: %#v %v",r,err)}
	if _,err=s.Move(r.ID,"one",r.Version,json.RawMessage(`{"cell":4}`));err!=nil{t.Fatal(err)}
	if _,err=s.Move(r.ID,"two",r.Version,json.RawMessage(`{"cell":5}`));err!=ErrConflict{t.Fatalf("expected conflict, got %v",err)}
}

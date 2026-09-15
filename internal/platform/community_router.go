package platform

import (
	"encoding/json"
	"net/http"
)

func (a *API) registerCommunity(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/accounts",a.createAccount)
	mux.HandleFunc("GET /api/v1/profiles/{handle}",a.profile)
	mux.HandleFunc("POST /api/v1/sessions",a.createSession)
	mux.HandleFunc("POST /api/v1/projects",a.createProject)
	mux.HandleFunc("GET /api/v1/projects/{id}",a.project)
	mux.HandleFunc("POST /api/v1/projects/{id}/members",a.projectMember)
	mux.HandleFunc("POST /api/v1/projects/{id}/submit",a.submitProject)
	mux.HandleFunc("POST /api/v1/projects/{id}/uploads",a.createUpload)
	mux.HandleFunc("POST /api/v1/uploads/{id}/complete",a.completeUpload)
	mux.HandleFunc("POST /api/v1/projects/{id}/releases",a.createRelease)
	mux.HandleFunc("POST /api/v1/releases/{id}/status",a.releaseStatus)
	mux.HandleFunc("GET /api/v1/content/{type}/{id}/comments",a.contentComments)
	mux.HandleFunc("POST /api/v1/content/{type}/{id}/comments",a.addContentComment)
	mux.HandleFunc("POST /api/v1/content/{type}/{id}/ratings",a.rate)
	mux.HandleFunc("POST /api/v1/content/{type}/{id}/favorite",a.favorite)
	mux.HandleFunc("GET /api/v1/forums/categories",a.categories)
	mux.HandleFunc("POST /api/v1/forums/threads",a.createThread)
	mux.HandleFunc("GET /api/v1/forums/threads/{id}",a.thread)
	mux.HandleFunc("POST /api/v1/forums/threads/{id}/posts",a.addPost)
	mux.HandleFunc("GET /api/v1/notifications/{userID}",a.notifications)
	mux.HandleFunc("POST /api/v1/reports",a.createCommunityReport)
	mux.HandleFunc("GET /api/v1/moderation/reports",a.moderationQueue)
	mux.HandleFunc("POST /api/v1/moderation/reports/{id}/actions",a.moderate)
	mux.HandleFunc("POST /api/v1/sdk/games/{gameID}/achievements/{achievementID}/unlock",a.unlock)
	mux.HandleFunc("GET /api/v1/sdk/games/{gameID}/leaderboards/{board}",a.leaderboard)
	mux.HandleFunc("POST /api/v1/sdk/games/{gameID}/leaderboards/{board}",a.submitScore)
	mux.HandleFunc("GET /api/v1/sdk/games/{gameID}/saves/{slot}",a.getSave)
	mux.HandleFunc("PUT /api/v1/sdk/games/{gameID}/saves/{slot}",a.putSave)
	mux.HandleFunc("GET /api/v1/arcade/games",a.arcadeGames)
	mux.HandleFunc("POST /api/v1/arcade/rooms",a.createRoom)
	mux.HandleFunc("GET /api/v1/arcade/rooms/{id}",a.room)
	mux.HandleFunc("POST /api/v1/arcade/rooms/{id}/join",a.joinRoom)
	mux.HandleFunc("POST /api/v1/arcade/rooms/{id}/moves",a.move)
}

func(a *API)createAccount(w http.ResponseWriter,r *http.Request){var v struct{Email,Handle,DisplayName string};if !decode(w,r,&v){return};x,e:=a.store.CreateUser(v.Email,v.Handle,v.DisplayName);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)profile(w http.ResponseWriter,r *http.Request){x,e:=a.store.Profile(r.PathValue("handle"));if e!=nil{writeError(w,e);return};x.Email="";writeJSON(w,http.StatusOK,x)}
func(a *API)createSession(w http.ResponseWriter,r *http.Request){var v struct{Email string};if !decode(w,r,&v){return};s,token,e:=a.store.CreateSession(v.Email);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,map[string]any{"session":s,"token":token})}
func(a *API)createProject(w http.ResponseWriter,r *http.Request){var v struct{OwnerID,Kind,Slug,Title,Description string;Tags []string};if !decode(w,r,&v){return};x,e:=a.store.CreateProject(v.OwnerID,v.Kind,v.Slug,v.Title,v.Description,v.Tags);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)project(w http.ResponseWriter,r *http.Request){x,e:=a.store.Project(r.PathValue("id"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)projectMember(w http.ResponseWriter,r *http.Request){var v struct{UserID,Role string};if !decode(w,r,&v){return};x,e:=a.store.AddProjectMember(r.PathValue("id"),v.UserID,v.Role);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)submitProject(w http.ResponseWriter,r *http.Request){x,e:=a.store.SubmitProject(r.PathValue("id"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)createUpload(w http.ResponseWriter,r *http.Request){var v struct{Kind,ContentType string;Size int64};if !decode(w,r,&v){return};x,e:=a.store.CreateUpload(r.PathValue("id"),v.Kind,v.ContentType,v.Size);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)completeUpload(w http.ResponseWriter,r *http.Request){x,e:=a.store.CompleteUpload(r.PathValue("id"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)createRelease(w http.ResponseWriter,r *http.Request){var v struct{UploadID,Version string};if !decode(w,r,&v){return};x,e:=a.store.CreateRelease(r.PathValue("id"),v.UploadID,v.Version);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)releaseStatus(w http.ResponseWriter,r *http.Request){var v struct{Status,PlayerURL string};if !decode(w,r,&v){return};x,e:=a.store.UpdateRelease(r.PathValue("id"),v.Status,v.PlayerURL);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)contentComments(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.Comments(r.PathValue("type"),r.PathValue("id")))}
func(a *API)addContentComment(w http.ResponseWriter,r *http.Request){var v struct{AuthorID,Body string};if !decode(w,r,&v){return};x,e:=a.store.AddContentComment(r.PathValue("type"),r.PathValue("id"),v.AuthorID,v.Body);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)rate(w http.ResponseWriter,r *http.Request){var v struct{UserID string;Value int};if !decode(w,r,&v){return};x,e:=a.store.Rate(r.PathValue("type"),r.PathValue("id"),v.UserID,v.Value);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,map[string]float64{"average":x})}
func(a *API)favorite(w http.ResponseWriter,r *http.Request){var v struct{UserID string};if !decode(w,r,&v){return};x,e:=a.store.Favorite(r.PathValue("type"),r.PathValue("id"),v.UserID);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,map[string]bool{"favorite":x})}
func(a *API)categories(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.Categories())}
func(a *API)createThread(w http.ResponseWriter,r *http.Request){var v struct{CategoryID,AuthorID,Title,Body string};if !decode(w,r,&v){return};x,e:=a.store.CreateThread(v.CategoryID,v.AuthorID,v.Title,v.Body);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)thread(w http.ResponseWriter,r *http.Request){t,p,e:=a.store.Thread(r.PathValue("id"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,map[string]any{"thread":t,"posts":p})}
func(a *API)addPost(w http.ResponseWriter,r *http.Request){var v struct{AuthorID,Body string};if !decode(w,r,&v){return};x,e:=a.store.AddPost(r.PathValue("id"),v.AuthorID,v.Body);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)notifications(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.UserNotifications(r.PathValue("userID")))}
func(a *API)createCommunityReport(w http.ResponseWriter,r *http.Request){var v struct{ReporterID,TargetType,TargetID,Reason string};if !decode(w,r,&v){return};x,e:=a.store.CreateReport(v.ReporterID,v.TargetType,v.TargetID,v.Reason);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)moderationQueue(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.ModerationQueue())}
func(a *API)moderate(w http.ResponseWriter,r *http.Request){var v struct{ModeratorID,Action,Note string};if !decode(w,r,&v){return};x,e:=a.store.Moderate(r.PathValue("id"),v.ModeratorID,v.Action,v.Note);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)unlock(w http.ResponseWriter,r *http.Request){var v struct{UserID,Name,Description string;Points int};if !decode(w,r,&v){return};x,e:=a.store.UnlockAchievement(r.PathValue("gameID"),v.UserID,r.PathValue("achievementID"),v.Name,v.Description,v.Points);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)leaderboard(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.Leaderboard(r.PathValue("gameID"),r.PathValue("board")))}
func(a *API)submitScore(w http.ResponseWriter,r *http.Request){var v struct{UserID string;Score int64;Metadata json.RawMessage};if !decode(w,r,&v){return};x,e:=a.store.SubmitScore(r.PathValue("gameID"),r.PathValue("board"),v.UserID,v.Score,v.Metadata);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)getSave(w http.ResponseWriter,r *http.Request){x,e:=a.store.GetSave(r.PathValue("gameID"),r.URL.Query().Get("userId"),r.PathValue("slot"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)putSave(w http.ResponseWriter,r *http.Request){var v struct{UserID string;Version int;Data json.RawMessage};if !decode(w,r,&v){return};x,e:=a.store.PutSave(r.PathValue("gameID"),v.UserID,r.PathValue("slot"),v.Data,v.Version);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)arcadeGames(w http.ResponseWriter,r *http.Request){writeJSON(w,http.StatusOK,a.store.ArcadeGames())}
func(a *API)createRoom(w http.ResponseWriter,r *http.Request){var v struct{GameID,HostID string};if !decode(w,r,&v){return};x,e:=a.store.CreateRoom(v.GameID,v.HostID);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusCreated,x)}
func(a *API)room(w http.ResponseWriter,r *http.Request){x,e:=a.store.Room(r.PathValue("id"));if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)joinRoom(w http.ResponseWriter,r *http.Request){var v struct{UserID string};if !decode(w,r,&v){return};x,e:=a.store.JoinRoom(r.PathValue("id"),v.UserID);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}
func(a *API)move(w http.ResponseWriter,r *http.Request){var v struct{UserID string;Version int;Move json.RawMessage};if !decode(w,r,&v){return};x,e:=a.store.Move(r.PathValue("id"),v.UserID,v.Version,v.Move);if e!=nil{writeError(w,e);return};writeJSON(w,http.StatusOK,x)}

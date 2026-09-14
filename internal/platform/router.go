package platform

import (
	"encoding/json"
	"errors"
	"net/http"
)

type API struct { store *Store }

func NewHandler(store *Store) http.Handler {
	api := &API{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/healthz", func(w http.ResponseWriter, _ *http.Request) { writeJSON(w, http.StatusOK, map[string]string{"status":"ok"}) })
	mux.HandleFunc("GET /api/v1/games/{slug}", api.game)
	mux.HandleFunc("GET /api/v1/games/{slug}/related", api.related)
	mux.HandleFunc("POST /api/v1/games/{slug}/reactions", api.react)
	mux.HandleFunc("POST /api/v1/games/{slug}/comments", api.comment)
	mux.HandleFunc("POST /api/v1/games/{slug}/playlists", api.playlist)
	mux.HandleFunc("POST /api/v1/games/{slug}/reports", api.report)
	return mux
}

func (a *API) game(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.Snapshot(r.PathValue("slug"))
	if err != nil { writeError(w, err); return }
	writeJSON(w, http.StatusOK, result)
}
func (a *API) related(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.Related(r.PathValue("slug"))
	if err != nil { writeError(w, err); return }
	writeJSON(w, http.StatusOK, result)
}
func (a *API) react(w http.ResponseWriter, r *http.Request) {
	var input struct { Type string `json:"type"` }
	if !decode(w,r,&input) { return }
	likes, err := a.store.React(r.PathValue("slug"), input.Type)
	if err != nil { writeError(w, err); return }
	writeJSON(w, http.StatusOK, map[string]int{"likes":likes})
}
func (a *API) comment(w http.ResponseWriter, r *http.Request) {
	var input struct { Author, Body string }
	if !decode(w,r,&input) { return }
	comment, err := a.store.AddComment(r.PathValue("slug"), input.Author, input.Body)
	if err != nil { writeError(w, err); return }
	writeJSON(w, http.StatusCreated, comment)
}
func (a *API) playlist(w http.ResponseWriter, r *http.Request) {
	var input struct { Name string }
	if !decode(w,r,&input) { return }
	if err := a.store.AddToPlaylist(r.PathValue("slug"), input.Name); err != nil { writeError(w, err); return }
	writeJSON(w, http.StatusCreated, map[string]bool{"saved":true})
}
func (a *API) report(w http.ResponseWriter, r *http.Request) {
	var input struct { Reason string }
	if !decode(w,r,&input) { return }
	if err := a.store.Report(r.PathValue("slug"), input.Reason); err != nil { writeError(w, err); return }
	w.WriteHeader(http.StatusNoContent)
}
func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	defer r.Body.Close()
	if err := json.NewDecoder(http.MaxBytesReader(w,r.Body,1<<20)).Decode(dst); err != nil {
		writeJSON(w,http.StatusBadRequest,map[string]string{"error":"invalid JSON body"}); return false
	}
	return true
}
func writeError(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	if errors.Is(err, ErrNotFound) { status = http.StatusNotFound }
	writeJSON(w,status,map[string]string{"error":err.Error()})
}
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

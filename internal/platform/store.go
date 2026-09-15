package platform

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var ErrNotFound = errors.New("game not found")

type Store struct {
	mu sync.RWMutex
	games map[string]Game
	likes map[string]int
	comments map[string][]Comment
	playlists map[string]map[string]bool
	reports []string
	state CommunityState
	backend StateBackend
}

type CommunityState struct {
	Users map[string]User `json:"users"`
	Sessions map[string]Session `json:"sessions"`
	Projects map[string]Project `json:"projects"`
	Uploads map[string]Upload `json:"uploads"`
	Releases map[string]Release `json:"releases"`
	ContentComments map[string][]ContentComment `json:"contentComments"`
	Ratings map[string]map[string]int `json:"ratings"`
	Favorites map[string]map[string]bool `json:"favorites"`
	CommunityPlaylists map[string]Playlist `json:"communityPlaylists"`
	ForumCategories map[string]ForumCategory `json:"forumCategories"`
	ForumThreads map[string]ForumThread `json:"forumThreads"`
	ForumPosts map[string][]ForumPost `json:"forumPosts"`
	Notifications map[string][]Notification `json:"notifications"`
	CommunityReports map[string]Report `json:"communityReports"`
	ModerationActions []ModerationAction `json:"moderationActions"`
	Achievements map[string]Achievement `json:"achievements"`
	Unlocks []AchievementUnlock `json:"unlocks"`
	Scores []LeaderboardScore `json:"scores"`
	Saves map[string]CloudSave `json:"saves"`
	ArcadeGames map[string]ArcadeGame `json:"arcadeGames"`
	ArcadeRooms map[string]ArcadeRoom `json:"arcadeRooms"`
}

type persistedState struct {
	Games map[string]Game `json:"games"`; Likes map[string]int `json:"likes"`; Comments map[string][]Comment `json:"comments"`
	Playlists map[string]map[string]bool `json:"playlists"`; Reports []string `json:"reports"`; Community CommunityState `json:"community"`
}

type StateBackend interface {
	Load(context.Context) ([]byte, error)
	Save(context.Context, []byte, string, []byte) error
	Health(context.Context) error
}

func NewStore() *Store {
	xo := Game{
		Slug: "xo-arena-football", Title: "XO Arena Football",
		Creator: Creator{Name: "XO Arena", Handle: "xo-arena", URL: "/play/xo-arena-football"},
		Description: "Call the plays. Coach your team. Build a dynasty.",
		Genres: []string{"Sports", "Strategy", "Simulation"}, Tags: []string{"football", "multiplayer", "html5", "free"},
		Runtime: "html5", Multiplayer: true, Mobile: true,
		PublishedAt: "2026-09-14", UpdatedAt: "2026-09-14", GameURL: "/play/xo-arena-football", Accent: "#ffd633",
	}
	s := &Store{
		games: map[string]Game{xo.Slug: xo},
		likes: map[string]int{xo.Slug: 2400},
		comments: map[string][]Comment{xo.Slug: {
			{ID: "c1", GameSlug: xo.Slug, Author: "pixelcoach", Body: "The X/O presentation makes play calling instantly readable.", Score: 42, CreatedAt: time.Date(2026,9,14,15,0,0,0,time.UTC)},
			{ID: "c2", GameSlug: xo.Slug, Author: "fourthandone", Body: "Give me one more drive. Then another.", Score: 27, CreatedAt: time.Date(2026,9,14,16,0,0,0,time.UTC)},
		}},
		playlists: map[string]map[string]bool{}, state: seedCommunity(),
	}
	return s
}

func seedCommunity() CommunityState {
	games := []ArcadeGame{
		{ID:"tic-tac-toe",Slug:"tic-tac-toe",Name:"Tic-Tac-Toe",Description:"Three in a row.",MinPlayers:2,MaxPlayers:2},
		{ID:"four-in-a-row",Slug:"four-in-a-row",Name:"Four in a Row",Description:"Connect four tokens.",MinPlayers:2,MaxPlayers:2},
		{ID:"checkers",Slug:"checkers",Name:"Checkers",Description:"Classic diagonal strategy.",MinPlayers:2,MaxPlayers:2},
		{ID:"chess",Slug:"chess",Name:"Chess",Description:"Classic chess.",MinPlayers:2,MaxPlayers:2},
		{ID:"yacht-dice",Slug:"yacht-dice",Name:"Yacht Dice",Description:"Five-dice scorecard game.",MinPlayers:1,MaxPlayers:6},
	}
	a := map[string]ArcadeGame{}; for _, g := range games { a[g.ID] = g }
	return CommunityState{
		Users:map[string]User{}, Sessions:map[string]Session{}, Projects:map[string]Project{}, Uploads:map[string]Upload{}, Releases:map[string]Release{},
		ContentComments:map[string][]ContentComment{}, Ratings:map[string]map[string]int{}, Favorites:map[string]map[string]bool{}, CommunityPlaylists:map[string]Playlist{},
		ForumCategories:map[string]ForumCategory{"general":{ID:"general",Slug:"general",Name:"General",Description:"Talk about games, movies, art, and NGG."},"creators":{ID:"creators",Slug:"creators",Name:"Creator Corner",Description:"Share work and find collaborators."}},
		ForumThreads:map[string]ForumThread{}, ForumPosts:map[string][]ForumPost{}, Notifications:map[string][]Notification{}, CommunityReports:map[string]Report{}, ModerationActions:[]ModerationAction{}, Achievements:map[string]Achievement{}, Unlocks:[]AchievementUnlock{}, Scores:[]LeaderboardScore{}, Saves:map[string]CloudSave{}, ArcadeGames:a, ArcadeRooms:map[string]ArcadeRoom{},
	}
}

func NewPersistentStore(ctx context.Context, backend StateBackend) (*Store, error) {
	s := NewStore(); s.backend = backend
	payload, err := backend.Load(ctx); if err != nil { return nil, err }
	if len(payload) == 0 { if err := s.persistLocked(ctx,"platform.seed",map[string]bool{"seeded":true}); err != nil { return nil, err }; return s,nil }
	var p persistedState; if err := json.Unmarshal(payload,&p); err != nil { return nil, fmt.Errorf("decode platform state: %w",err) }
	if p.Games != nil { s.games=p.Games }; if p.Likes != nil { s.likes=p.Likes }; if p.Comments != nil { s.comments=p.Comments }; if p.Playlists != nil { s.playlists=p.Playlists }; s.reports=p.Reports
	if p.Community.Users != nil { s.state=normalizeCommunity(p.Community) }
	return s,nil
}

func normalizeCommunity(v CommunityState) CommunityState {
	seed:=seedCommunity()
	if v.Users==nil{v.Users=seed.Users};if v.Sessions==nil{v.Sessions=seed.Sessions};if v.Projects==nil{v.Projects=seed.Projects};if v.Uploads==nil{v.Uploads=seed.Uploads};if v.Releases==nil{v.Releases=seed.Releases}
	if v.ContentComments==nil{v.ContentComments=seed.ContentComments};if v.Ratings==nil{v.Ratings=seed.Ratings};if v.Favorites==nil{v.Favorites=seed.Favorites};if v.CommunityPlaylists==nil{v.CommunityPlaylists=seed.CommunityPlaylists}
	if v.ForumCategories==nil{v.ForumCategories=seed.ForumCategories};if v.ForumThreads==nil{v.ForumThreads=seed.ForumThreads};if v.ForumPosts==nil{v.ForumPosts=seed.ForumPosts};if v.Notifications==nil{v.Notifications=seed.Notifications};if v.CommunityReports==nil{v.CommunityReports=seed.CommunityReports}
	if v.Achievements==nil{v.Achievements=seed.Achievements};if v.Saves==nil{v.Saves=seed.Saves};if v.ArcadeGames==nil{v.ArcadeGames=seed.ArcadeGames};if v.ArcadeRooms==nil{v.ArcadeRooms=seed.ArcadeRooms};return v
}

func (s *Store) persistLocked(ctx context.Context, event string, detail any) error {
	if s.backend == nil { return nil }
	p,err:=json.Marshal(persistedState{Games:s.games,Likes:s.likes,Comments:s.comments,Playlists:s.playlists,Reports:s.reports,Community:s.state}); if err!=nil{return err}
	e,_:=json.Marshal(detail); return s.backend.Save(ctx,p,event,e)
}
func (s *Store) Health(ctx context.Context) error { if s.backend==nil{return nil};return s.backend.Health(ctx) }
func newID(prefix string) string { var b [12]byte; _,_ = rand.Read(b[:]); return prefix+"_"+hex.EncodeToString(b[:]) }

func (s *Store) Snapshot(slug string) (Snapshot, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	game, ok := s.games[slug]
	if !ok { return Snapshot{}, ErrNotFound }
	comments := append([]Comment(nil), s.comments[slug]...)
	return Snapshot{Game: game, Likes: s.likes[slug], Comments: comments}, nil
}

func (s *Store) Related(slug string) ([]Game, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	if _, ok := s.games[slug]; !ok { return nil, ErrNotFound }
	result := make([]Game, 0, 3)
	for key, game := range s.games { if key != slug { result = append(result, game) } }
	return result, nil
}

func (s *Store) React(slug, reaction string) (int, error) {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return 0, ErrNotFound }
	if reaction == "like" { s.likes[slug]++ }
	if err := s.persistLocked(context.Background(),"game.reacted",map[string]string{"slug":slug,"reaction":reaction}); err != nil { return 0,err }
	return s.likes[slug], nil
}

func (s *Store) AddComment(slug, author, body string) (Comment, error) {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return Comment{}, ErrNotFound }
	body = strings.TrimSpace(body)
	if body == "" || len(body) > 1000 { return Comment{}, errors.New("comment must contain 1-1000 characters") }
	if strings.TrimSpace(author) == "" { author = "guest" }
	comment := Comment{ID: fmt.Sprintf("c%d", len(s.comments[slug])+1), GameSlug: slug, Author: author, Body: body, CreatedAt: time.Now().UTC()}
	s.comments[slug] = append([]Comment{comment}, s.comments[slug]...)
	if err := s.persistLocked(context.Background(),"game.commented",comment); err != nil { return Comment{},err }
	return comment, nil
}

func (s *Store) AddToPlaylist(slug, playlist string) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return ErrNotFound }
	if playlist == "" { playlist = "play-later" }
	if s.playlists[playlist] == nil { s.playlists[playlist] = map[string]bool{} }
	s.playlists[playlist][slug] = true
	return s.persistLocked(context.Background(),"game.playlisted",map[string]string{"slug":slug,"playlist":playlist})
}

func (s *Store) Report(slug, reason string) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return ErrNotFound }
	reason = strings.TrimSpace(reason)
	if reason == "" { return errors.New("reason is required") }
	s.reports = append(s.reports, slug+":"+reason)
	return s.persistLocked(context.Background(),"game.reported",map[string]string{"slug":slug,"reason":reason})
}

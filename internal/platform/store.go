package platform

import (
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
	return &Store{
		games: map[string]Game{xo.Slug: xo},
		likes: map[string]int{xo.Slug: 2400},
		comments: map[string][]Comment{xo.Slug: {
			{ID: "c1", GameSlug: xo.Slug, Author: "pixelcoach", Body: "The X/O presentation makes play calling instantly readable.", Score: 42, CreatedAt: time.Date(2026,9,14,15,0,0,0,time.UTC)},
			{ID: "c2", GameSlug: xo.Slug, Author: "fourthandone", Body: "Give me one more drive. Then another.", Score: 27, CreatedAt: time.Date(2026,9,14,16,0,0,0,time.UTC)},
		}},
		playlists: map[string]map[string]bool{},
	}
}

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
	return comment, nil
}

func (s *Store) AddToPlaylist(slug, playlist string) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return ErrNotFound }
	if playlist == "" { playlist = "play-later" }
	if s.playlists[playlist] == nil { s.playlists[playlist] = map[string]bool{} }
	s.playlists[playlist][slug] = true
	return nil
}

func (s *Store) Report(slug, reason string) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.games[slug]; !ok { return ErrNotFound }
	reason = strings.TrimSpace(reason)
	if reason == "" { return errors.New("reason is required") }
	s.reports = append(s.reports, slug+":"+reason)
	return nil
}

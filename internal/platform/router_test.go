package platform

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGameAndSocialFlow(t *testing.T) {
	handler := NewHandler(NewStore())
	get := httptest.NewRecorder()
	handler.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/api/v1/games/xo-arena-football", nil))
	if get.Code != http.StatusOK || !strings.Contains(get.Body.String(), "XO Arena Football") { t.Fatalf("unexpected game response: %d %s", get.Code, get.Body.String()) }

	like := httptest.NewRecorder()
	handler.ServeHTTP(like, httptest.NewRequest(http.MethodPost, "/api/v1/games/xo-arena-football/reactions", strings.NewReader(`{"type":"like"}`)))
	if like.Code != http.StatusOK || !strings.Contains(like.Body.String(), "2401") { t.Fatalf("unexpected like response: %d %s", like.Code, like.Body.String()) }

	comment := httptest.NewRecorder()
	handler.ServeHTTP(comment, httptest.NewRequest(http.MethodPost, "/api/v1/games/xo-arena-football/comments", strings.NewReader(`{"Author":"coach","Body":"Great game"}`)))
	if comment.Code != http.StatusCreated || !strings.Contains(comment.Body.String(), "Great game") { t.Fatalf("unexpected comment response: %d %s", comment.Code, comment.Body.String()) }
}

func TestUnknownGame(t *testing.T) {
	response := httptest.NewRecorder()
	NewHandler(NewStore()).ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/api/v1/games/nope", nil))
	if response.Code != http.StatusNotFound { t.Fatalf("got %d", response.Code) }
}

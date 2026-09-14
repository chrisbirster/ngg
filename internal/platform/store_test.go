package platform

import (
	"strings"
	"testing"
)

func TestXOArenaUsesNGGOwnedRelativeURLs(t *testing.T) {
	snapshot, err := NewStore().Snapshot("xo-arena-football")
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.Game.GameURL != "/play/xo-arena-football" {
		t.Fatalf("unexpected game URL %q", snapshot.Game.GameURL)
	}
	if snapshot.Game.Creator.URL != "/play/xo-arena-football" {
		t.Fatalf("unexpected creator URL %q", snapshot.Game.Creator.URL)
	}
	if strings.Contains(strings.ToLower(snapshot.Game.GameURL+snapshot.Game.Creator.URL), "vutadex") {
		t.Fatal("retired VutaDex domain must not appear in the NGG manifest")
	}
}

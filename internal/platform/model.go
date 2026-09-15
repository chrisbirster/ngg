package platform

import (
	"encoding/json"
	"time"
)

type Creator struct {
	Name string `json:"name"`
	Handle string `json:"handle"`
	URL string `json:"url"`
}

type Game struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Creator Creator `json:"creator"`
	Description string `json:"description"`
	Genres []string `json:"genres"`
	Tags []string `json:"tags"`
	Runtime string `json:"runtime"`
	Multiplayer bool `json:"multiplayer"`
	Mobile bool `json:"mobile"`
	PublishedAt string `json:"publishedAt"`
	UpdatedAt string `json:"updatedAt"`
	GameURL string `json:"gameUrl"`
	Accent string `json:"accent"`
}

type Comment struct {
	ID string `json:"id"`
	GameSlug string `json:"gameSlug"`
	Author string `json:"author"`
	Body string `json:"body"`
	Score int `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}

type Snapshot struct {
	Game Game `json:"game"`
	Likes int `json:"likes"`
	Comments []Comment `json:"comments"`
}

type User struct {
	ID string `json:"id"`
	Email string `json:"email,omitempty"`
	Handle string `json:"handle"`
	DisplayName string `json:"displayName"`
	Bio string `json:"bio"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	ID string `json:"id"`
	UserID string `json:"userId"`
	TokenHash string `json:"-"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Project struct {
	ID, OwnerID, Kind, Slug, Title, Description, Status string
	Tags []string
	Members map[string]string
	CreatedAt, UpdatedAt time.Time
}

type Upload struct {
	ID, ProjectID, Kind, ObjectKey, UploadURL, Status, ContentType string
	Size int64
	CreatedAt time.Time
}

type Release struct {
	ID, ProjectID, Version, ObjectKey, PlayerURL, Status string
	CreatedAt time.Time
}

type ContentComment struct {
	ID, ContentType, ContentID, AuthorID, Body string
	Score int
	CreatedAt time.Time
}

type Playlist struct { ID, OwnerID, Name string; Items []string; CreatedAt time.Time }
type ForumCategory struct { ID, Slug, Name, Description string }
type ForumThread struct { ID, CategoryID, AuthorID, Title string; Locked bool; CreatedAt, UpdatedAt time.Time }
type ForumPost struct { ID, ThreadID, AuthorID, Body string; CreatedAt, UpdatedAt time.Time }
type Notification struct { ID, UserID, Kind, Message, URL string; Read bool; CreatedAt time.Time }
type Report struct { ID, ReporterID, TargetType, TargetID, Reason, Status string; CreatedAt time.Time }
type ModerationAction struct { ID, ModeratorID, ReportID, Action, Note string; CreatedAt time.Time }
type Achievement struct { ID, GameID, Name, Description string; Points int }
type AchievementUnlock struct { AchievementID, UserID string; UnlockedAt time.Time }
type LeaderboardScore struct { GameID, Board, UserID string; Score int64; Metadata json.RawMessage; SubmittedAt time.Time }
type CloudSave struct { GameID, UserID, Slot string; Data json.RawMessage; Version int; UpdatedAt time.Time }
type ArcadeGame struct { ID, Slug, Name, Description string; MinPlayers, MaxPlayers int }
type ArcadeRoom struct { ID, GameID, HostID, Status string; Players []string; State json.RawMessage; Version int; CreatedAt, UpdatedAt time.Time }

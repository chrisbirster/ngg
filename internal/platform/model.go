package platform

import "time"

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

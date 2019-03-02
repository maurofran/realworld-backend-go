package article

import "time"

type Collection struct {
	Articles []*Representation `json:"articles"`
	Count    int               `json:"articlesCount"`
}

type Representation struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	Tags           []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         *Author   `json:"author"`
}

type Author struct {
	Username  string `json:"username"`
	Biography string `json:"bio"`
	Image     string `json:"image"`
	Following string `json:"following"`
}

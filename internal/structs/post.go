package structs

import "time"

// Post - Post Output post schema
type Post struct {
	ID        string             `json:"id,omitempty"`
	Title     string             `json:"title,omitempty"`
	Slug      string             `json:"slug,omitempty"`
	Body      string             `json:"body,omitempty"`
	Thumbnail string             `json:"thumbnail,omitempty"`
	Markdown  bool               `json:"markdown,omitempty"`
	Temp      bool               `json:"temp,omitempty"`
	Private   bool               `json:"private,omitempty"`
	User      *ReadUserSchema    `json:"user,omitempty"`
	Tags      []PostTagSchema    `json:"tags,omitempty"`
	Counter   *ReadCounterSchema `json:"counter,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// PostTagSchema - Output post tag schema
type PostTagSchema struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ReadUserSchema - Output user schema
type ReadUserSchema struct {
	ID           string         `json:"id,omitempty"`
	Username     string         `json:"username,omitempty"`
	Email        string         `json:"email,omitempty"`
	Phone        string         `json:"phone,omitempty"`
	Certified    bool           `json:"certified,omitempty"`
	Administered bool           `json:"administered,omitempty"`
	Profile      *ProfileSchema `json:"profile,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// ProfileSchema - Output user profile schema
type ProfileSchema struct {
	ID           string    `json:"id,omitempty"`
	DisplayName  string    `json:"display_name,omitempty"`
	ShortBio     string    `json:"short_bio,omitempty"`
	About        *string   `json:"about,omitempty"`
	Thumbnail    *string   `json:"thumbnail,omitempty"`
	ProfileLinks *[]string `json:"profile_links,omitempty"`
}

// ReadCounterSchema - Output counter schema
type ReadCounterSchema struct {
	Likes        int64     `json:"likes,omitempty"`
	LikeUsers    *[]string `json:"like_users,omitempty"`
	Dislikes     int64     `json:"dislikes,omitempty"`
	DislikeUsers *[]string `json:"dislike_users,omitempty"`
	Viewers      *int64    `json:"viewers,omitempty"`
	Shares       *int64    `json:"shares,omitempty"`
	ShareUsers   *[]string `json:"share_users,omitempty"`
	Relate       string    `json:"relate,omitempty"`
	RelateType   string    `json:"relate_type,omitempty"`
}

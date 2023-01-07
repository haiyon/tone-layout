package structs

import "time"

// FindGreeter - find greeter by user id and greeter id
type FindGreeter struct {
	User    string `json:"user,omitempty"`
	Greeter string `json:"greeter,omitempty"`
}

// Greeter - Greeter schema
type Greeter struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListGreeters - get user greeter list.
type ListGreeters struct {
	User   string `json:"user,omitempty"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int32  `json:"limit,omitempty"`
}

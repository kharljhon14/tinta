package data

import "time"

type Blog struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Version   int32     `json:"version"`
	Tags      []string  `json:"tags,omitempty"`
}

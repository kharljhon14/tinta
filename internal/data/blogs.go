package data

import (
	"time"

	"github.com/kharljhon14/tinta/internal/validator"
)

type Blog struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Version   int32     `json:"version"`
	Tags      []string  `json:"tags,omitzero"`
}

func ValidateBlog(v *validator.Validator, blog *Blog) {

	v.Check(blog.Title != "", "title", "must be provided")
	v.Check(len(blog.Title) <= 255, "title", "must not be more than 255 bytes long")

	v.Check(blog.Content != "", "content", "must be provided")
	v.Check(len(blog.Content) <= 500, "content", "must not be more than 500 bytes long")

	v.Check(blog.Tags != nil, "tags", "must be provided")
	v.Check(len(blog.Tags) >= 0, "tags", "must contain at least 1 tag")
	v.Check(len(blog.Tags) <= 5, "tags", "must not contain more than 5 genres")
	v.Check(validator.Unique(blog.Tags), "tags", "must not contain duplicate values")
}

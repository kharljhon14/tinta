package data

import (
	"database/sql"
	"time"

	"github.com/kharljhon14/tinta/internal/validator"
	"github.com/lib/pq"
)

type BlogModel struct {
	DB *sql.DB
}

type Blog struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Version   int32     `json:"version,omitzero"`
	Tags      []string  `json:"tags,omitzero"`
}

func (b *BlogModel) Insert(blog *Blog) error {
	query := `
		INSERT INTO blogs (title, content, author, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version;
	`

	args := []any{
		blog.Title,
		blog.Content,
		blog.Author,
		pq.Array(blog.Tags),
	}

	return b.DB.QueryRow(query, args...).Scan(
		&blog.ID,
		&blog.CreatedAt,
		&blog.Version,
	)
}

func (b *BlogModel) Get(id int64) (*Blog, error) {
	if id < 1 {
		return nil, sql.ErrNoRows
	}

	query := `
		SElECT id, title, content, author, tags, created_at, version
		FROM blogs
		WHERE id = $1;
	`

	var blog Blog

	err := b.DB.QueryRow(query, id).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.Author,
		pq.Array(&blog.Tags),
		&blog.CreatedAt,
		&blog.Version,
	)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (b *BlogModel) GetLatest() ([]*Blog, error) {
	return nil, nil
}

func (b *BlogModel) Update(blog *Blog) error {
	query := `
		UPDATE blogs 
		SET title = $1,
		content = $2,
		tags = $3,
		version = version + 1
		WHERE id = $4
		RETURNING version;
	`

	args := []any{
		blog.Title,
		blog.Content,
		pq.Array(blog.Tags),
		blog.ID,
	}

	return b.DB.QueryRow(query, args...).Scan(
		&blog.Version,
	)
}

func (b *BlogModel) Delete(id int64) error {
	if id < 1 {
		return sql.ErrNoRows
	}

	query := `
		DELETE FROM blogs
		WHERE id = $1;
	`

	result, err := b.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
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

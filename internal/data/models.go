package data

import (
	"database/sql"
	"errors"
)

var (
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Blogs BlogModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs: BlogModel{DB: db},
	}
}

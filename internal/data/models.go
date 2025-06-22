package data

import (
	"database/sql"
	"errors"
)

var (
	ErroRecordNotFound = errors.New("record not found")
)

type Models struct {
	Blogs BlogModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs: BlogModel{DB: db},
	}
}

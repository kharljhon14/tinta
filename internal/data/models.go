package data

import (
	"database/sql"
)

type Models struct {
	Blogs BlogModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs: BlogModel{DB: db},
	}
}

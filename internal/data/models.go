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
	Users UserModel
	Token TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs: BlogModel{DB: db},
		Users: UserModel{DB: db},
		Token: TokenModel{DB: db},
	}
}

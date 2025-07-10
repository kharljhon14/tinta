package data

import (
	"database/sql"
	"errors"
)

var (
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Blogs       BlogModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs:       BlogModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}

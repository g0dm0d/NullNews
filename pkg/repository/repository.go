package repository

import (
	"database/sql"

	"github.com/g0dm0d/nullnews/entity"
)

type News interface {
	CreateNews(entity.Article) (int, error)
}

type Auth interface {
	Register(entity.User)
	Login(entity.User) bool
}

type MainDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *MainDB {
	return &MainDB{db: db}
}

type Repository struct {
	News
	Auth
}

func NewRep(db *sql.DB) *Repository {
	return &Repository{
		News: NewDB(db),
		Auth: NewDB(db),
	}
}

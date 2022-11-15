package repository

import (
	"database/sql"
	"time"

	"github.com/g0dm0d/nullnews/entity"
)

type News interface {
	CreateNews(entity.Article) (int, error)
	DeleteNews(entity.Article)
}

type Auth interface {
	Register(entity.User)
	Login(string, string) (bool, entity.User)
	SaveSession(string, int, time.Time) (int, error)
	DeleteSession(int)
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

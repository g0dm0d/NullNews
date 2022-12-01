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
	Register(user entity.User) error
	Login(password, email string) (status bool, user entity.User)
	SaveSession(token string, userID int, timeEXP time.Time) (int, error)
}

type JWT interface {
	DeleteSession(sessionID int)
	SessionExist(token string) (user entity.User, sessionID int, err error)
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
	JWT
}

func NewRep(db *sql.DB) *Repository {
	return &Repository{
		News: NewDB(db),
		Auth: NewDB(db),
		JWT:  NewDB(db),
	}
}

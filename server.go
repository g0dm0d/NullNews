package nullnews

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Ctx struct {
	DB *sql.DB
}

type Config struct {
	Addr   string
	Port   string
	DB     *sql.DB
	Router func() chi.Router
}

func NewServer(cfg *Config) *http.Server {
	return &http.Server{
		Addr:              cfg.Addr + ":" + cfg.Port,
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           cfg.Router(),
	}
}

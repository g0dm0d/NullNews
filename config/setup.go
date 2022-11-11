package config

import "database/sql"

type App struct {
	DB     *sql.DB
	Config *ServerConfig
}

func New() *App {
	var app = new(App)
	vars := envParse()

	app.DB = setupDB(vars.DBDriver, vars.DBDSN)
	app.Config = &ServerConfig{
		Addr: vars.ApiAddr,
		Port: vars.ApiPort,
	}

	return app
}

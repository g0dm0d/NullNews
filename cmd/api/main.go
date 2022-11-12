package main

import (
	"github.com/g0dm0d/nullnews"
	"github.com/g0dm0d/nullnews/config"
	"github.com/g0dm0d/nullnews/pkg/handler"
	"github.com/g0dm0d/nullnews/pkg/middlewares"
	"github.com/g0dm0d/nullnews/pkg/repository"
	"github.com/g0dm0d/nullnews/pkg/service"
)

func main() {
	app := config.New()
	defer app.DB.Close()

	repository := repository.NewRep(app.DB)
	service := service.NewSer(repository, &service.Ctx{
		Secret: app.Config.Secret,
	})
	middlewares := middlewares.NewMid(service)
	handler := handler.New(service, middlewares)

	api := nullnews.NewServer(&nullnews.Config{
		Addr:   app.Config.Addr,
		Port:   app.Config.Port,
		DB:     app.DB,
		Router: handler.Router,
	})
	_ = api.ListenAndServe()
}

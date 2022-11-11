package db

import (
	"log"

	"github.com/g0dm0d/nullnews"
)

const createNews = "INSERT INTO news (title, text, user) VALUES ($1, $2, $3)"

type CreateNewsParams struct {
	Title string
	Text  string
	User  int
}

func CreateNews(c *nullnews.Ctx, argv CreateNewsParams) {
	_, err := c.DB.Exec(createNews, "Hello", "world!")
	if err != nil {
		log.Panicln(err)
	}
}

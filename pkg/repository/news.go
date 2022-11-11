package repository

import (
	"log"

	"github.com/g0dm0d/nullnews/entity"
)

func (r *MainDB) CreateNews(article entity.Article) (int, error) {
	res, err := r.db.Exec("INSERT INTO news (title, text) VALUES ($1, $2)", article.Title, article.Text)
	log.Println(res)
	if err != nil {
		log.Panicln(err)
	}
	return 0, err
}

package repository

import (
	"log"

	"github.com/g0dm0d/nullnews/entity"
)

func (r *MainDB) CreateNews(article entity.Article) (int, error) {
	_, err := r.db.Exec("INSERT INTO news (title, text, author) VALUES ($1, $2, $3)", article.Title, article.Text, article.Author)
	if err != nil {
		log.Panicln(err)
	}
	return 0, err
}

func (r *MainDB) DeleteNews(article entity.Article) {
	_, err := r.db.Exec("DELETE FROM news WHERE ID=$1", article.ID)
	if err != nil {
		log.Panicln(err)
	}
}

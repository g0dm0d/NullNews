package repository

import (
	"log"

	"github.com/g0dm0d/nullnews/entity"
	"golang.org/x/crypto/bcrypt"
)

func (r *MainDB) Register(article entity.User) {
	password, err := HashingPass(article.Password)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := r.db.Exec("INSERT INTO users (username, firstname, lastname, email, password) VALUES ($1, $2, $3, $4, $5)",
		article.Username, article.FirstName, article.LastName, article.Email, password)
	log.Println(res)
	if err != nil {
		log.Panicln(err)
		return
	}
}

func (r *MainDB) Login(article entity.User) bool {
	rows, err := r.db.Query("select password from users where email = $1", article.Email)
	if err != nil {
		log.Println(err)
	}
	var password string
	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return CheckPasswordHash(article.Password, password)
}

func HashingPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panicln(err)
		return "", err
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	log.Println(err)
	return err == nil
}

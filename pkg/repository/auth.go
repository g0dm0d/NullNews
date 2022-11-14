package repository

import (
	"log"
	"time"

	"github.com/g0dm0d/nullnews/entity"
	"golang.org/x/crypto/bcrypt"
)

func (r *MainDB) Register(user entity.User) {
	password, err := HashingPass(user.Password)
	if err != nil {
		log.Println(err)
	}
	_, err = r.db.Exec("INSERT INTO users (username, firstname, lastname, email, password) VALUES ($1, $2, $3, $4, $5)",
		user.Username, user.FirstName, user.LastName, user.Email, password)
	if err != nil {
		log.Println(err)
	}
}

func (r *MainDB) Login(password, email string) (bool, int) {
	row := r.db.QueryRow("SELECT password, id FROM users WHERE email = $1", email)
	var user entity.User
	err := row.Scan(&user.Password, &user.ID)
	if err != nil {
		log.Println(err)
		return false, 0
	}
	return CheckPasswordHash(password, user.Password), user.ID
}

func (r *MainDB) SaveSession(session string, userID int, time time.Time) (int, error) {
	row := r.db.QueryRow("INSERT INTO sessions (refresh_token, user_id, expires_time) VALUES ($1, $2, $3) RETURNING id", session, userID, time)
	var id int
	err := row.Scan(&id)
	return id, err
}

func (r *MainDB) DeleteSession(id float64) {
	_, err := r.db.Exec("DELETE FROM sessions WHERE id = $1", id)
	if err != nil {
		log.Println(err)
	}
}

func HashingPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package repository

import (
	"log"

	"github.com/g0dm0d/nullnews/entity"
)

func (r *MainDB) DeleteSession(id int) {
	_, err := r.db.Exec("DELETE FROM sessions WHERE id = $1", id)
	if err != nil {
		log.Println(err)
	}
}

func (r *MainDB) SessionExist(session string) (user entity.User, sessionID int, err error) {
	row := r.db.QueryRow("SELECT id, user_id FROM sessions WHERE refresh_token = $1 AND expires_time > CURRENT_TIMESTAMP", session)
	err = row.Scan(&sessionID, &user.ID)
	if err != nil {
		return
	}
	row = r.db.QueryRow("SELECT username, firstname, lastname, permission FROM users WHERE id = $1", user.ID)
	err = row.Scan(&user.Username, &user.FirstName, &user.LastName, &user.Permission)
	if err != nil {
		return
	}
	return user, sessionID, err
}

package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/g0dm0d/nullnews/entity"
	"github.com/g0dm0d/nullnews/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTService struct {
	repo    repository.JWT
	sectret string
}

func NewJWT(repo repository.JWT, sectret string) *JWTService {
	return &JWTService{
		repo:    repo,
		sectret: sectret,
	}
}

func (s *JWTService) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		return
	}
	SetCookie(w, "token", "")
	ParseToken, err := TokenParse(token.Value, s.sectret)
	if err != nil {
		return
	}
	s.repo.DeleteSession(ParseToken.SessionID)
}

type Session struct {
	Session string `json:"session"`
}

func (s *JWTService) RefreshJWT(w http.ResponseWriter, r *http.Request) {
	var session Session
	json.NewDecoder(r.Body).Decode(&session)
	user, sessionID, err := s.repo.SessionExist(session.Session)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := TokenGenerate(user, sessionID, s.sectret)
	SetCookie(w, "token", token)
}

const tokenTime = 15 * time.Minute
const sessionTime = 336 * time.Hour

func TokenGenerate(user entity.User, sessionID int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"username":   user.Username,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"session_id": sessionID,
		"permission": user.Permission,
		"exp":        time.Now().Add(tokenTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func TokenParse(tokenString, secret string) (*TokenJWT, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		return &TokenJWT{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenJWT{
			UserID:     claims["user_id"].(int),
			Username:   claims["username"].(string),
			FirstName:  claims["firstname"].(string),
			lastName:   claims["lastname"].(string),
			SessionID:  claims["session_id"].(int),
			Permission: claims["permission"].(int),
			Exp:        claims["exp"].(float64),
		}, err
	}
	return &TokenJWT{}, err
}

func SetCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: false,
	})
}

type TokenJWT struct {
	UserID     int
	Username   string
	FirstName  string
	lastName   string
	SessionID  int
	Permission int
	Exp        float64
}

func SessionGenerate() (string, time.Time) {
	cookie := uuid.New()
	return cookie.String(), time.Now().Add(sessionTime)
}

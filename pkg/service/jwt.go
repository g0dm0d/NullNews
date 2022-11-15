package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenTime = 15 * time.Minute
const sessionTime = 336 * time.Hour

func TokenGenerate(userID, sessionID, permission int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       userID,
		"session_id": sessionID,
		"permission": permission,
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
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		return &TokenJWT{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenJWT{
			UserID:     claims["user"].(int),
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
	SessionID  int
	Permission int
	Exp        float64
}

func SessionGenerate() (string, time.Time) {
	cookie := uuid.New()
	return cookie.String(), time.Now().Add(sessionTime)
}

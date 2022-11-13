package service

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenTime = 15 * time.Minute
const sessionTime = 336 * time.Hour

func TokenGenerate(userID, sessionID int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       userID,
		"session_id": sessionID,
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
			UserID:    claims["user"].(float64),
			SessionID: claims["session_id"].(float64),
			Exp:       claims["exp"].(float64),
		}, err
	}
	return &TokenJWT{}, err
}

type TokenJWT struct {
	UserID    float64
	SessionID float64
	Exp       float64
}

func SessionGenerate() (string, time.Time) {
	cookie := uuid.New()
	return cookie.String(), time.Now().Add(sessionTime)
}

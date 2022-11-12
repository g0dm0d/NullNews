package service

import (
	"log"
	"time"

	"github.com/g0dm0d/nullnews/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenTime = 15 * time.Minute
const sessionTime = 336 * time.Hour

type JWTService struct {
	repo repository.Auth
	ctx  string
}

func NewJWT(repo repository.Auth, ctx string) *JWTService {
	return &JWTService{
		repo: repo,
		ctx:  ctx,
	}
}

func TokenGenerate(id int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": id,
		"exp":  time.Now().Add(tokenTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func SessionGenerate() (string, int64) {
	cookie := uuid.New()
	return cookie.String(), time.Now().Add(sessionTime).Unix()
}

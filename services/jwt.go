package services

import (
	"time"

	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/security"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func makeToken(cred models.Credentials) (string, error) {
	jwtClaims := JwtClaims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 24 * 60 * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	ss, err := token.SignedString(security.MySigningKey)
	return ss, err
}

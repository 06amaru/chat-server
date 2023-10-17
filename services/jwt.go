package services

import (
	"time"

	"github.com/amaru0601/fluent/models"

	"github.com/golang-jwt/jwt"
)

var (
	MySigningKey []byte
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

	ss, err := token.SignedString(MySigningKey)
	return ss, err
}

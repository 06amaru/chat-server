package route

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	MySigningKey = []byte("Ja0ks0nE")
)

type JwtClaims struct {
	Username string
	jwt.StandardClaims
}

func makeToken(user *UserSignIn) (string, error) {
	jwtClaims := JwtClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	ss, err := jwtoken.SignedString(MySigningKey)
	return ss, err
}

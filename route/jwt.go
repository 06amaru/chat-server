package route

import "github.com/golang-jwt/jwt"

var (
	MySigningKey = []byte("Ja0ks0nE")
)

func makeToken(user *UserSignIn) (string, error) {
	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	ss, err := jwtoken.SignedString(MySigningKey)
	return ss, err
}

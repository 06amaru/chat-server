package route

import "github.com/golang-jwt/jwt"

var (
	mySigningKey = []byte("keysign")
)

func makeToken(user *UserSignIn) (string, error) {
	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	ss, err := jwtoken.SignedString(mySigningKey)
	return ss, err
}

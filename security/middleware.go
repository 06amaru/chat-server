package security

import (
	"errors"
	"reflect"

	"github.com/amaru0601/fluent/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.DefaultJWTConfig

/*func getSigningKey(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != config.SigningMethod {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
	}

	if len(config.SigningKeys) > 0 {
		// https://www.rfc-editor.org/rfc/rfc7515#section-4.1.4
		if kid, ok := token.Header["kid"].(string); ok {
			if key, ok := config.SigningKeys[kid]; ok {
				return key, nil
			}
		}
		return nil, fmt.Errorf("unexpected jwt key id=%v", token.Header["kid"])
	}

	return config.SigningKey, nil
}*/

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.QueryParam("jwt")
		config.SigningKey = services.MySigningKey
		//config.KeyFunc = getSigningKey

		token, err := parseToken(param)

		if err == nil {
			// Store user information from token into context.
			c.Set(config.ContextKey, token)
			if config.SuccessHandler != nil {
				config.SuccessHandler(c)
			}
			return next(c)
		}
		if config.ErrorHandler != nil {
			return config.ErrorHandler(err)
		}
		if config.ErrorHandlerWithContext != nil {
			return config.ErrorHandlerWithContext(err, c)
		}
		return &echo.HTTPError{
			Code:     middleware.ErrJWTInvalid.Code,
			Message:  middleware.ErrJWTInvalid.Message,
			Internal: err,
		}
	}
}

func parseToken(auth string) (interface{}, error) {
	token_ := new(jwt.Token)
	var err error
	if _, ok := config.Claims.(jwt.MapClaims); ok {
		token_, err = jwt.Parse(auth, config.KeyFunc)
	} else {
		t := reflect.ValueOf(config.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		token_, err = jwt.ParseWithClaims(auth, claims, config.KeyFunc)
	}
	if err != nil {
		return nil, err
	}
	if !token_.Valid {
		return nil, errors.New("invalid token")
	}
	return token_, nil
}

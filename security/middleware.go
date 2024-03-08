package security

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/amaru0601/channels/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.DefaultJWTConfig

func getSigningKey(token *jwt.Token) (interface{}, error) {
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
}

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data string
		for _, cookie := range c.Cookies() {
			if cookie != nil && (*cookie).Name == "token" {
				data = (*cookie).Value
			}
		}
		config.SigningKey = services.MySigningKey
		config.KeyFunc = getSigningKey

		token, err := parseToken(data)

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

func parseToken(tokenStr string) (interface{}, error) {
	token := new(jwt.Token)
	var err error
	if _, ok := config.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(tokenStr, config.KeyFunc)
	} else {
		t := reflect.ValueOf(config.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		token, err = jwt.ParseWithClaims(tokenStr, claims, config.KeyFunc)
	}
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

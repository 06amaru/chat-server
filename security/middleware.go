package security

import (
	"errors"
	"fmt"
	"github.com/amaru0601/fluent/services"
	"reflect"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	config := middleware.DefaultJWTConfig
	return func(c echo.Context) error {
		param := c.QueryParam("jwt")
		config.SigningKey = services.MySigningKey

		defaultKeyFunc := func(t *jwt.Token) (interface{}, error) {
			// Check the signing method
			if t.Method.Alg() != config.SigningMethod {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			if len(config.SigningKeys) > 0 {
				if kid, ok := t.Header["kid"].(string); ok {
					if key, ok := config.SigningKeys[kid]; ok {
						return key, nil
					}
				}
				return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
			}

			return config.SigningKey, nil
		}

		config.KeyFunc = defaultKeyFunc

		token, err := parseToken(param, config)

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

func parseToken(auth string, config middleware.JWTConfig) (interface{}, error) {
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

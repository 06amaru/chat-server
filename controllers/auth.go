package controllers

import (
	"github.com/amaru0601/channels/models"
	"github.com/amaru0601/channels/repository"
	"github.com/amaru0601/channels/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController() AuthController {
	repo := repository.NewRepository()

	return AuthController{
		service: services.NewAuthService(repo),
	}
}

func (ctrl AuthController) SignUp(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := ctrl.service.SignUp(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusAccepted)
}

func (ctrl AuthController) SignIn(c echo.Context) error {
	cred := new(models.Credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}

	token, err := ctrl.service.SignIn(*cred)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, token)
}

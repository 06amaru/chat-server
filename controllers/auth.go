package controllers

import (
	"net/http"

	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/repository"
	"github.com/jaox1/chat-server/services"
	"github.com/labstack/echo/v4"
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

func (ctrl AuthController) SignIn(c echo.Context) error {
	cred := new(models.Credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}

	token, err := ctrl.service.SignIn(*cred)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}

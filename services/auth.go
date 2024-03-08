package services

import (
	"fmt"

	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) AuthService {
	return AuthService{repo: repo}
}

func (svc AuthService) SignUp(user models.User) error {
	err := svc.repo.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (svc AuthService) SignIn(cred models.Credentials) (string, error) {
	user, err := svc.repo.FindUser(cred.Username)
	if err != nil {
		return "", err
	}

	if user.Password != cred.Password {
		return "", fmt.Errorf("password not match")
	}

	token, err := makeToken(cred)
	if err != nil {
		return "", err
	}

	return token, nil
}

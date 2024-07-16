package services

import (
	"fmt"

	"github.com/jaox1/chat-server/ent"
	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/repository"
	"github.com/jaox1/chat-server/security"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) AuthService {
	return AuthService{repo: repo}
}

func (svc AuthService) SignIn(cred models.Credentials) (*models.Credentials, error) {
	user, err := svc.repo.FindUser(cred.Username)
	if err != nil {
		switch errType := err.(type) {
		case *ent.NotFoundError:
			securePwd := security.CreateHash(cred.Password)
			publicKey := security.GenerateKey()
			privateKey := security.GenerateKey()
			secureKey := security.Encrypt(privateKey, cred.Password)

			err = svc.repo.SignUp(cred.Username, securePwd, publicKey, secureKey)
			if err != nil {
				return nil, err
			}

			return credentialWithToken(cred)
		default:
			return nil, errType
		}
	}

	securePwd := security.CreateHash(cred.Password)
	if user.Password != securePwd {
		return nil, fmt.Errorf("password not match")
	}

	return credentialWithToken(cred)
}

func credentialWithToken(cred models.Credentials) (*models.Credentials, error) {
	token, err := makeToken(cred)
	if err != nil {
		return nil, err
	}

	return &models.Credentials{
		Username: cred.Username,
		Token:    token,
	}, nil
}

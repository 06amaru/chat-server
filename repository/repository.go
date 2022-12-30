package repository

import (
	"context"
	"github.com/amaru0601/fluent/db"
	"github.com/amaru0601/fluent/ent"
	userEnt "github.com/amaru0601/fluent/ent/user"
	"github.com/amaru0601/fluent/models"
)

type Repository struct {
	client *ent.Client
}

func NewRepository() *Repository {
	postgres := db.GetPostgresClient()

	return &Repository{
		client: postgres,
	}
}
func (repo Repository) SignUp(user models.User) error {
	_, err := repo.client.User.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetPrivateKey(user.PrivateKey).
		SetPublicKey(user.PublicKey).
		Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) FindUser(username string) (*ent.User, error) {
	user, err := repo.client.User.
		Query().
		Where(userEnt.UsernameEQ(username)).
		First(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

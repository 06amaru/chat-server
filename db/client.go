package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/amaru0601/fluent/ent"
	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {

	// TODO definir variables de entorno para desarrollo local y dockercompose
	log.Println(os.Getenv("DB_HOST"))
	log.Println(os.Getenv("DB_PORT"))
	log.Println(os.Getenv("DB_NAME"))
	log.Println(os.Getenv("DB_USER"))
	log.Println(os.Getenv("DB_PASSWORD"))

	client, err := ent.Open("postgres", "host=postgresdb port=5432 user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client, nil
}

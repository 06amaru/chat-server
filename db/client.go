package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jaox1/chat-server/ent"
	_ "github.com/lib/pq"
)

//var postgresClient *ent.Client

func GetPostgresClient() *ent.Client {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Println(err)
		panic(err)
	}

	return client
}

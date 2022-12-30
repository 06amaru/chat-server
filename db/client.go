package db

import (
	"context"
	"log"
	"os"

	"github.com/amaru0601/fluent/ent"
	_ "github.com/lib/pq"
)

var postgresClient *ent.Client

func GetPostgresClient() *ent.Client {

	// TODO definir variables de entorno para desarrollo local y dockercompose
	log.Println(os.Getenv("DB_HOST"))
	log.Println(os.Getenv("DB_PORT"))
	log.Println(os.Getenv("DB_NAME"))
	log.Println(os.Getenv("DB_USER"))
	log.Println(os.Getenv("DB_PASSWORD"))

	client, err := ent.Open("postgres", "host=db.bcoyyczdpsaoazwywotk.supabase.co port=5432 user=postgres dbname=postgres password=G3ZcEtGQUr3Kn90A sslmode=disable")
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

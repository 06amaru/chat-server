package main

import (
	"context"
	"fluent/ent"
	"fluent/ent/user"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if _, err = CreateUser(ctx, client); err != nil {
		log.Fatal(err)
	}
	if _, err = QueryUserById(ctx, client, 1); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetUsername("jaoks").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)
	return u, nil
}

func QueryUserById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

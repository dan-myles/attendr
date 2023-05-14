package database

import (
	"attendr/watcher/ent"
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() {
	// Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseUrl := os.Getenv("DATABASE_URL")

	// Connecting to the database
	client, err := ent.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool
	schema_err := client.Schema.Create(context.Background())
	if schema_err != nil {
		log.Fatalf("Failed creating schema resources: %v", schema_err)
	}

	defer client.Close()
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.ASU_Watched_Class, error) {
	u, err := client.ASU_Watched_Class.
		Create().
		SetAge(20).
		SetName("Foo").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
		return nil, err
	}

	log.Println("user was created: ", u)
	return u, nil
}

package database

import (
	"attendr/watcher/ent"
	"attendr/watcher/log"
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client      *ent.Client
	databaseUrl string
)

func Init() {
	// Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Logger.Error("Error loading .env file")
	}

	databaseUrl = os.Getenv("DATABASE_URL")
	log.Logger.Debug("Loaded database env variables")

	// Connecting to the database
	Client, err = ent.Open("mysql", databaseUrl)
	if err != nil {
		log.Logger.Error("Failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool
	schema_err := Client.Schema.Create(context.Background())
	if schema_err != nil {
		log.Logger.Error("Failed creating schema resources: %v", schema_err)
	}

	log.Logger.Debug("Ran the database schema migration tool")

	defer Client.Close()
}

func openConnection() {
	// Connecting to the database
	Client, _ = ent.Open("mysql", databaseUrl)

	log.Logger.Debug("Opened a connection to the database")
}

func closeConnection() {
	// Closing the connection to the database
	Client.Close()

	log.Logger.Debug("Closed the connection to the database")
}

func AddToWatchlist(ctx context.Context) (*ent.ASU_Watched_Class, error) {
	openConnection()
	defer closeConnection()

	u, err := Client.ASU_Watched_Class.
		Create().
		SetTitle("Test Class").
		SetInstructor("Test Instructor").
		SetSubject("Test Subject").
		SetSubjectNumber("Test Subject Number").
		SetHasOpenSeats(false).
		SetClassNumber("Test Class Number").
		SetTerm("Test Term").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	log.Logger.Debug("Class was added watchlist database: ", u)
	return u, nil
}

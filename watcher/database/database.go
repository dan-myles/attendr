// TODO:
//  - Add email and phone number Class Tracking DB table (not sure if needed)
//      Clerk is handling phone numbers so prob not needed
//  - Add functions to add & remove user notification preferences to DB
//      In general track email and phone number to send notifications to
//  - Refactor close and open connection to happen in every function
//      This is because we want to multithread the database calls
//  - Integrate Clerk User ID's into DB

package database

import (
	"attendr/watcher/ent"
	"attendr/watcher/ent/asuwatchedclass"
	"attendr/watcher/utils"
	"attendr/watcher/utils/asu"
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client      *ent.Client
	databaseUrl string
)

/*
* Initializes the database connection & runs a schema migration tool
**/
func InitDbConnection() {
	// Grab the database url from the .env file
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error("Error loading .env file", "err", err)
	}
	databaseUrl = os.Getenv("DSN")

	// Test a connection to the database
	Client, err = ent.Open("mysql", databaseUrl)
	if err != nil {
		utils.Logger.Error("Failed opening connection to db", "err", err)
	}
	defer Client.Close()

	// Normally we would run the schema migration tool from Ent here
	// But we are using prisma to manage DB schema migrations

	utils.Logger.Debug("Finished initializing the database")
}

/*
* Opens a connection to the database
**/
func openConnection() {
	Client, _ = ent.Open("mysql", databaseUrl)
	utils.Logger.Debug("Opened a connection to the database")
}

/*
* Closes a current connection to the database
**/
func closeConnection() {
	Client.Close()
	utils.Logger.Debug("Closed the connection to the database")
}

/*
* Grabs the watchlist from the database
**/
func GrabWatchListFromDb(ctx context.Context) (
	[]*ent.ASUWatchedClass,
	error,
) {
	openConnection()
	defer closeConnection()

	watchlist, err := Client.ASUWatchedClass.Query().All(ctx)
	if err != nil {
		utils.Logger.Error("Error grabbing watchlist from db", "err", err)
	}

	utils.Logger.Debug("Grabbed watchlist from db")
	return watchlist, err
}

/*
* Adds a class to the database
**/
func AddClassToDb(ctx context.Context, class *asu.ApiResponse) (*ent.ASUWatchedClass, error) {
	openConnection()
	defer closeConnection()

	// Check if it is a recitation or not
	title := class.Classes[0].Clas.Title
	if class.Classes[0].Clas.Descr2 == "Recitation" {
		title = title + " Recitation"
	}

	// Check if class is already in DB
	existingClass, _ := Client.ASUWatchedClass.
		Query().
		Where(asuwatchedclass.ClassNumber(class.Classes[0].Clas.Classnbr)).
		First(ctx)
	if existingClass != nil {
		utils.Logger.Debug("Class already exists in DB", "class", existingClass)
		return nil, fmt.Errorf("Class already exists in DB")
	}

	// Add class to DB
	addedClass, err := Client.ASUWatchedClass.
		Create().
		SetTitle(title).
		SetUserID("DebugUserID"). // WARNING: GET RID OF THIS EVENTUALLY
		SetInstructor(class.Classes[0].Clas.Instructorslist[0]).
		SetSubject(class.Classes[0].Clas.Subject).
		SetSubjectNumber(class.Classes[0].Clas.Catalognbr).
		SetClassNumber(class.Classes[0].Clas.Classnbr).
		SetTerm(class.Classes[0].Clas.Strm).
		Save(ctx)
	if err != nil {
		utils.Logger.Error("Error adding to watchlist", "err", err)
		return nil, err
	}

	// There was no constraint error & no general error
	return addedClass, nil
}

/*
* Removes a class from the database
**/
func RemoveClassFromDb(ctx context.Context,
	classToDelete *ent.ASUWatchedClass,
) error {
	openConnection()
	defer closeConnection()

	err := Client.ASUWatchedClass.
		DeleteOne(classToDelete).
		Exec(ctx)
	if err != nil {
		utils.Logger.Error("Error removing class from watchlist", "err", err)
		return err
	}

	return nil
}

/*
* Get a class by ClassNumber
**/
func GetClassFromDbByClassNbr(ctx context.Context, nbr string) (
	*ent.ASUWatchedClass,
	error,
) {
	openConnection()
	defer closeConnection()

	class, err := Client.ASUWatchedClass.Query().
		Where(asuwatchedclass.ClassNumber(nbr)).
		First(ctx)
	if err != nil {
		utils.Logger.Error("Error grabbing class from db", "err", err)
		return nil, err
	}

	return class, nil
}

/*
* Remove a class from the database by ClassNumber
**/
func RemoveClassFromDbByClassNbr(ctx context.Context, nbr string) (
	int,
	error,
) {
	openConnection()
	defer closeConnection()

	classNbr, err := Client.ASUWatchedClass.Delete().Where(asuwatchedclass.ClassNumber(nbr)).Exec(ctx)
	if err != nil {
		utils.Logger.Error("Error deleting class from db", "err", err)
		return -1, err
	}

	return classNbr, nil
}

/*
* NOTE: DEPRECATED DO NOT USE
**/
func SyncWatchlistToDb(ctx context.Context, classes []ent.ASUWatchedClass) error {
	openConnection()
	defer closeConnection()

	var error error = nil

	// Querying to check if it exists
	// If it does not exist then create it
	for _, localClass := range classes {
		exists, err := Client.ASUWatchedClass.
			Query().
			Where(asuwatchedclass.ClassNumber(localClass.ClassNumber)).
			Exist(ctx)
		if err != nil {
			utils.Logger.Error("Error checking if class exists", "err", err)
			error = err
		}

		if !exists {
			_, err := Client.ASUWatchedClass.
				Create().
				SetTitle(localClass.Title).
				SetInstructor(localClass.Instructor).
				SetSubject(localClass.Subject).
				SetSubjectNumber(localClass.SubjectNumber).
				SetClassNumber(localClass.ClassNumber).
				SetTerm(localClass.Term).
				Save(ctx)
			if err != nil {
				utils.Logger.Error("Error syncing classes to db", "err", err)
				error = err
			}
		}
	}

	utils.Logger.Debug("Finished syncing watchlist to db")
	return error
}

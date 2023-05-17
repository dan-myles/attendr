// TODO:
//  - Add email and phone number Class Tracking DB table
//  - Add functions to add & remove user notification preferences to DB
//      In general track email and phone number to send notifications to

package database

import (
	"attendr/watcher/ent"
	"attendr/watcher/ent/asu_watched_class"
	"attendr/watcher/utils"
	"attendr/watcher/utils/asu"
	"context"
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

	// Open a connection to the database
	Client, err = ent.Open("mysql", databaseUrl)
	if err != nil {
		utils.Logger.Error("Failed opening connection to db", "err", err)
	}
	defer Client.Close()

	// Run the auto migration tool
	schema_err := Client.Schema.Create(context.Background())
	if schema_err != nil {
		utils.Logger.Error("Failed creating schema resources", "err", schema_err)
	}

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
* Syncs the watchlist to the database
* NOTE: Not sure if this is needed as classes are added to DB when POSTed
* Currently its not being called anywhere
**/
func SyncWatchlistToDb(ctx context.Context, classes []ent.ASU_Watched_Class) error {
	openConnection()
	defer closeConnection()

	var error error = nil

	// Querying to check if it exists
	// If it does not exist then create it
	for _, localClass := range classes {
		exists, err := Client.ASU_Watched_Class.
			Query().
			Where(asu_watched_class.ClassNumber(localClass.ClassNumber)).
			Exist(ctx)
		if err != nil {
			utils.Logger.Error("Error checking if class exists", "err", err)
			error = err
		}

		if !exists {
			_, err := Client.ASU_Watched_Class.
				Create().
				SetTitle(localClass.Title).
				SetInstructor(localClass.Instructor).
				SetSubject(localClass.Subject).
				SetSubjectNumber(localClass.SubjectNumber).
				SetHasOpenSeats(localClass.HasOpenSeats).
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

/*
* Grabs the watchlist from the database
**/
func GrabWatchListFromDb(ctx context.Context) (
	[]*ent.ASU_Watched_Class,
	error,
) {
	openConnection()
	defer closeConnection()

	watchlist, err := Client.ASU_Watched_Class.Query().All(ctx)
	if err != nil {
		utils.Logger.Error("Error grabbing watchlist from db", "err", err)
	}

	utils.Logger.Debug("Grabbed watchlist from db")
	return watchlist, err
}

/*
* Adds a class to the database
**/
func AddClassToDb(ctx context.Context, class *asu.ApiResponse) (*ent.ASU_Watched_Class, error) {
	openConnection()
	defer closeConnection()

	// Check if it is a recitation or not
	title := class.Classes[0].Clas.Title
	if class.Classes[0].Clas.Descr2 == "Recitation" {
		title = title + " Recitation"
	}

	// Add class to Db
	addedClass, err := Client.ASU_Watched_Class.
		Create().
		SetTitle(title).
		SetInstructor(class.Classes[0].Clas.Instructorslist[0]).
		SetSubject(class.Classes[0].Clas.Subject).
		SetSubjectNumber(class.Classes[0].Clas.Catalognbr).
		SetHasOpenSeats(false).
		SetClassNumber(class.Classes[0].Clas.Classnbr).
		SetTerm(class.Classes[0].Clas.Strm).
		Save(ctx)

		// Check constraint error (if class already exists)
	if ent.IsConstraintError(err) {
		utils.Logger.Debug("Class already exists in watchlist database")
		existingClass, _ := Client.ASU_Watched_Class.
			Query().
			Where(asu_watched_class.ClassNumber(class.Classes[0].Clas.Classnbr)).
			First(ctx)

		return existingClass, nil
	}

	// Check general error
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
	classToDelete *ent.ASU_Watched_Class,
) error {
	openConnection()
	defer closeConnection()

	err := Client.ASU_Watched_Class.
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
	*ent.ASU_Watched_Class,
	error,
) {
	openConnection()
	defer closeConnection()

	class, err := Client.ASU_Watched_Class.Query().
		Where(asu_watched_class.ClassNumber(nbr)).
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

	classNbr, err := Client.ASU_Watched_Class.Delete().Where(asu_watched_class.ClassNumber(nbr)).Exec(ctx)
	if err != nil {
		utils.Logger.Error("Error deleting class from db", "err", err)
		return -1, err
	}

	return classNbr, nil
}

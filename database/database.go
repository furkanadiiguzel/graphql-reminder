package database

import (
	"database/sql"
	"fmt"

	"github.com/furkanadiiguzel/graphql-reminder/graph/model"
	_ "github.com/lib/pq"
)

type DB struct {
	client *sql.DB
}

func Connect() *DB {
	connectionString := "postgres://username:password@furkanadiguzel:5432/dbname?sslmode=require"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return nil
	}
	defer db.Close()
	return &DB{client: db}
}
func (db *DB) GetReminder(id string) *model.ReminderListing {

	var reminderListing model.ReminderListing
	return &reminderListing
}
func (db *DB) GetReminders() []*model.ReminderListing {
	var reminderListings []*model.ReminderListing
	return reminderListings
}
func (db *DB) CreateReminderListing(reminderInfo model.CreateReminderListingInput) *model.ReminderListing {
	var returnReminderListing model.ReminderListing
	return &returnReminderListing
}
func (db *DB) UpdateReminderListing(id string, reminderInfo model.UpdateReminderListingInput) *model.ReminderListing {
	var ReminderListing model.ReminderListing
	return &ReminderListing
}
func (db *DB) DeleteReminderListing(id string) *model.DeleteReminderResponse {

	return &model.DeleteReminderResponse{DeleteReminderID: id}
}

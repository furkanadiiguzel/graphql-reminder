package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/furkanadiiguzel/graphql-reminder/graph/model"
)

var connectionString = "postgres://postgres:1234@localhost:5432/dbname?sslmode=require"

type DB struct {
	client *sql.DB
}

func Connect() *DB {
	client, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return &DB{client: client}
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

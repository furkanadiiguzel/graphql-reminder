package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/furkanadiiguzel/graphql-reminder/graph/generated"
	"github.com/furkanadiiguzel/graphql-reminder/graph/model"
)

// CreateReminderListing is the resolver for the createReminderListing field.
func (r *mutationResolver) CreateReminderListing(ctx context.Context, input model.CreateReminderListingInput) (*model.ReminderListing, error) {
	return db.CreateReminderListing(input), nil
}

// UpdateReminderListing is the resolver for the updateReminderListing field.
func (r *mutationResolver) UpdateReminderListing(ctx context.Context, id string, input model.UpdateReminderListingInput) (*model.ReminderListing, error) {
	return db.UpdateReminderListing(id, input), nil
}

// DeleteReminderListing is the resolver for the deleteReminderListing field.
func (r *mutationResolver) DeleteReminderListing(ctx context.Context, id string) (*model.DeleteReminderResponse, error) {
	return db.DeleteReminderListing(id), nil
}

// Reminders is the resolver for the reminers field.
func (r *queryResolver) Reminders(ctx context.Context) ([]*model.ReminderListing, error) {
	panic(fmt.Errorf("not implemented: Reminders - reminders"))
}

// Reminder is the resolver for the reminder field.
func (r *queryResolver) Reminder(ctx context.Context, id string) (*model.ReminderListing, error) {
	return db.GetReminder(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()

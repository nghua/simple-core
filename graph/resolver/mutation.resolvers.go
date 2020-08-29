package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"simple-core/graph/generated"
	"simple-core/graph/model"
	"simple-core/service/users"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, email string, password string, userMeta *model.UserMeta) (bool, error) {
	return users.RegisterUser(email, password, userMeta)
}

func (r *mutationResolver) InsertUser(ctx context.Context, email string, password string, role *int, userMeta *model.UserMeta) (bool, error) {
	return users.InsertUser(email, password, *role, userMeta)
}

func (r *mutationResolver) AlterUserInfo(ctx context.Context, id int64, email *string, password *string, role *int, userMeta *model.UserMeta) (bool, error) {
	return users.AlterUserInfo(id, *email, *password, *role, userMeta)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	return users.DeleteUser(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

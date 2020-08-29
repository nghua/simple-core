package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"simple-core/graph/generated"
	"simple-core/graph/model"
	"simple-core/service/users"
)

func (r *queryResolver) GetUserInfo(ctx context.Context, id int64) (*model.User, error) {
	cols := collectFields(ctx)
	return users.GetUserInfo(id, cols...)
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (string, error) {
	return users.UserLogin(email, password)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

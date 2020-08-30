package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"simple-core/graph/generated"
	"simple-core/graph/model"
	"simple-core/service/terms"
	"simple-core/service/users"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, email string, password string, userMeta *model.UserMeta) (bool, error) {
	return users.Register(email, password, userMeta)
}

func (r *mutationResolver) InsertUser(ctx context.Context, email string, password string, role *int, userMeta *model.UserMeta) (bool, error) {
	return users.Insert(email, password, *role, userMeta)
}

func (r *mutationResolver) AlterUser(ctx context.Context, id int64, email *string, password *string, role *int, userMeta *model.UserMeta) (bool, error) {
	return users.Alter(id, *email, *password, *role, userMeta)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	return users.Delete(id)
}

func (r *mutationResolver) AddTerm(ctx context.Context, termType int, name string, meta *model.TermMeta) (bool, error) {
	return terms.Add(termType, name, meta)
}

func (r *mutationResolver) AlterTerm(ctx context.Context, id int64, name *string, meta *model.TermMeta) (bool, error) {
	return terms.Alter(id, *name, meta)
}

func (r *mutationResolver) DeleteTerm(ctx context.Context, id int64) (bool, error) {
	return terms.Delete(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"simple-core/graph/generated"
	"simple-core/graph/model"
	"simple-core/service/terms"
	"simple-core/service/users"

	"github.com/99designs/gqlgen/graphql"
)

func (r *queryResolver) GetUser(ctx context.Context, id int64) (*model.User, error) {
	var cols []string
	fields := graphql.CollectFieldsCtx(ctx, nil)

	for _, f := range fields {
		if f.Name == "id" {
			continue
		}
		cols = append(cols, f.Name)
	}

	return users.Get(id, cols...)
}

func (r *queryResolver) GetUserList(ctx context.Context, offset *int, row *int) ([]*model.User, error) {
	var cols []string
	fields := graphql.CollectFieldsCtx(ctx, nil)

	for _, f := range fields {
		cols = append(cols, f.Name)
	}

	return users.GetList(*offset, *row, cols...)
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (string, error) {
	return users.Login(email, password)
}

func (r *queryResolver) GetTerm(ctx context.Context, id int64, termType *int) (*model.Term, error) {
	var cols []string
	fields := graphql.CollectFieldsCtx(ctx, nil)

	for _, f := range fields {
		if f.Name == "id" {
			continue
		}
		cols = append(cols, f.Name)
	}

	return terms.Get(id, *termType, cols...)
}

func (r *queryResolver) GetTermList(ctx context.Context, termType *int, offset *int, row *int, non *bool) ([]*model.Term, error) {
	var cols []string
	fields := graphql.CollectFieldsCtx(ctx, nil)

	for _, f := range fields {
		cols = append(cols, f.Name)
	}

	if !*non {
		return terms.GetList(*termType, *offset, *row, cols...)
	}

	return terms.GetNonNullList(*termType, *offset, *row, cols...)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

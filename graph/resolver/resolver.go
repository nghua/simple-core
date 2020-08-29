package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver 为外部handler调用
type Resolver struct{}

func collectFields(ctx context.Context) []string {
	var cols []string
	fields := graphql.CollectFieldsCtx(ctx, nil)

	for _, f := range fields {
		if f.Name == "id" {
			continue
		}
		cols = append(cols, f.Name)
	}

	return cols
}

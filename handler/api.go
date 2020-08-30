package handler

import (
	"context"
	"simple-core/graph/generated"
	"simple-core/graph/resolver"
	"simple-core/service/errmsg"
	"simple-core/service/token"
	"simple-core/utils/tool"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler api接口的handler
func GraphqlHandler() gin.HandlerFunc {
	conf := generated.Config{Resolvers: &resolver.Resolver{}}
	conf.Directives.HasRole = func(ctx context.Context, obj interface{},
		next graphql.Resolver, role int) (res interface{}, err error) {
		gc, err := tool.GetGinContext(ctx)
		if err != nil {
			return nil, err
		}

		ut := token.ParseToken(gc)
		if ut != nil {
			currentUserRole := ut.UserRole
			if currentUserRole >= role {
				return next(ctx)
			}
		}

		return nil, errmsg.TokenNotFoundError
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(conf))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Playground页面的handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

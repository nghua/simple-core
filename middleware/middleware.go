package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ContextKey 自定义Context
type ContextKey string

// GetGinContext 将*gin.Context绑定到Request，方便graphql使用
func GetGinContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ContextKey("ginContext"), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

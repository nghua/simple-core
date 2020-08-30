package middleware

import (
	"context"
	"simple-core/public/types"

	"github.com/gin-gonic/gin"
)

// GetGinContext 将*gin.Context绑定到Request，方便graphql使用
func GetGinContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), types.SimpleContextKey("ginContext"), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

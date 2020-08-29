package tool

import (
	"context"
	"errors"
	"simple-core/middleware"

	"github.com/gin-gonic/gin"
)

// GetGinContext 传入context判断是否为*gin.Context
func GetGinContext(ctx context.Context) (*gin.Context, error) {
	c := ctx.Value(middleware.ContextKey("ginContext"))
	if c == nil {
		return nil, errors.New("没有获取到上下文")
	}

	ginContext, ok := c.(*gin.Context)
	if !ok {
		return nil, errors.New("上下文错误，不是 *gin.Context")
	}

	return ginContext, nil
}

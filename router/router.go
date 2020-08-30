package router

import (
	"simple-core/handler"
	"simple-core/public/setting"

	"github.com/gin-gonic/gin"
)

func initRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middleware...)

	apiHandler(router)

	return router
}

func apiHandler(router *gin.Engine) {
	// 博客 api 路由
	router.POST("/api", handler.GraphqlHandler())

	if setting.Mode == gin.DebugMode {
		// playground
		router.GET("/playground", handler.PlaygroundHandler())
	}
}

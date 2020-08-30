package router

import (
	"context"
	"net/http"
	"simple-core/middleware"
	"simple-core/public/setting"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var server *http.Server

func Run() {
	gin.SetMode(setting.Mode)

	router := initRouter(gin.Recovery(), gin.Logger(), middleware.GetGinContext())

	server = &http.Server{
		Addr:           setting.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(setting.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << setting.MaxHeaderBytes,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] Http Server Stop err: %v\n", err)
	}

	log.Printf(" [INFO] Http Server Stop stopped\n")
	cancel()
}

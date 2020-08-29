package main

import (
	"os"
	"os/signal"
	"simple-core/router"
	"syscall"
)

func main() {
	router.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.Stop()
}

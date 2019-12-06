package main

import (
	"net/http"
	"time"
	"log"
	"github.com/willpower88/gotest/router"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: router.Router(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

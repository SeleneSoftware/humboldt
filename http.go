package main

import (
	"log"
	"net/http"
	"sync"
)

//This file contains some helper functions for starting and stopping the HTTP server.

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: "localhost:8080"}

	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServer(): %v", err)
		}
	}()

	return srv
}

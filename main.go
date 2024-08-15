package main

import (
	"context"
	"gateway/handlers"
	"gateway/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	pool := handlers.NewPool()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		handlers.StartTCPServer(ctx, pool)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		handlers.HandlePool(ctx, pool)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		services.CleanupOldEntries(ctx)
	}()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWebSocket(pool, w, r)
	})

	apiGateway := &http.Server{
		Addr:    ":8081",
		Handler: http.HandlerFunc(handlers.HandleAPI),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("API Gateway started on port 8081")
		if err := apiGateway.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API Gateway server failed: %v", err)
		}
	}()

	wsServer := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := wsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("WebSocket server failed: %v", err)
		}
	}()
	log.Println("WebSocket Server started on port 8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	// Initiate graceful shutdown
	cancel()

	// Shutdown the HTTP server with a timeout
	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := wsServer.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("WebSocket Server Shutdown Failed:%+v", err)
	}

	if err := apiGateway.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("API Gateway Server Shutdown Failed:%+v", err)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("Server gracefully stopped")
}

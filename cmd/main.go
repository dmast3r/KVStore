package main

import (
	"KVStore/internals/consts"
	"KVStore/internals/handlers"
	"KVStore/internals/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
    http.HandleFunc("/get", handlers.GetHandler)
    http.HandleFunc("/put", handlers.PutHandler)
    http.HandleFunc("/delete", handlers.DeleteHandler)

    srv :=  &http.Server{Addr: fmt.Sprintf(":%d", consts.PORT)}
    
    go func() {
        log.Printf("Starting the server at port %d\n", consts.PORT)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe error: %v", err)
        }
    }()

    stopChan := make(chan os.Signal, 1)
    signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

    <-stopChan
    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server Shutdown Failed:%+v", err)
    }

    log.Println("Server gracefully stopped!")

    utils.CloseDBConnections()
}

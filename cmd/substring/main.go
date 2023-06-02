package main

import (
	"fmt"
	"go-test-softweather/internal/config"
	"go-test-softweather/internal/handlers"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Printf("server start on %s port\n", cfg.ServerAddr)

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/api/substring", handlers.SubstringHandler)

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, nil))
}

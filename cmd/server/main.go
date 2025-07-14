package main

import (
	"log"
	"net/http"

	"github.com/Adit0507/autocomplete-search/config"
	"github.com/Adit0507/autocomplete-search/pkg/api"
)

func main() {
	cfg := config.NewConfig()
	router := api.NewRouter()

	log.Printf("Starting server on port %s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

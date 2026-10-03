package main

import (
	"log"
	"net/http"
	"os"

	"cafemovil-back/internal/api"
	"cafemovil-back/internal/catalog"
	"cafemovil-back/internal/user"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userService := user.NewService(user.NewMemoryRepository())
	catalogService := catalog.NewService(catalog.NewSeedRepository())
	router := api.NewRouter(userService, catalogService)

	log.Printf("cafemovil backend listening on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"kv-search-api/internal/api"
	"kv-search-api/internal/repository"
	"kv-search-api/internal/service"
	"kv-search-api/pkg/config"
	"log"
	"net/http"
)

/*
App runs on localhost, port 8000.
*/
func main() {
	err := run(8000, config.DBConnect())
	if err != nil {
		log.Printf("Connection failed.")
	}
}

func run(port int, mongoClient *mongo.Client) error {
	storageAPI := InitStorageAPI(mongoClient)
	log.Printf("Server running at http://localhost:%d/", port)
	http.HandleFunc("/api/create", storageAPI.Create)
	http.HandleFunc("/api/fetch", storageAPI.Fetch)
	http.HandleFunc("/api/search", storageAPI.Search)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}
	return nil
}

/*
the application works with different layers. when the API in the top layer is initialized, all layers work.
*/
func InitStorageAPI(mongoClient *mongo.Client) api.StorageAPI {
	storageRepository := repository.NewRepository(mongoClient)
	storageService := service.NewStorageService(storageRepository)
	storageAPI := api.NewStorageAPI(storageService)
	return storageAPI
}

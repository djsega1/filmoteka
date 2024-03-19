package main

import (
	"context"
	"log"
	"net/http"

	"github.com/djsega1/filmoteka/actors/internal/config"
	"github.com/djsega1/filmoteka/actors/internal/database"
	"github.com/djsega1/filmoteka/actors/internal/handlers"
	"github.com/djsega1/filmoteka/actors/internal/repository"
	"github.com/gorilla/mux"
)

func main() {

	//actorsSubR.HandleFunc("/actors", handlers.GetActorsList).Methods(http.MethodGet)
	//actorsSubR.HandleFunc("/actors", handlers.GetActorsList).Methods(http.MethodGet)

	cfg_path := "configs/config.yml"

	cfg := config.GetConfig(cfg_path)

	pool := database.NewPool(context.TODO(), cfg.Storage)
	repo := repository.NewActorsRepository(pool)
	actorsHandler := handlers.NewActorsHandler(repo)

	defer pool.Close()

	r := mux.NewRouter()
	actorsSubR := r.PathPrefix("/api/v1").Subrouter()
	actorsSubR.HandleFunc("/actors", actorsHandler.GetAllActors).Methods(http.MethodGet)
	actorsSubR.HandleFunc("/actors/{id:[0-9]+}", actorsHandler.GetActorByID).Methods(http.MethodGet)
	actorsSubR.HandleFunc("/actors/create", actorsHandler.CreateActor).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
	log.Println("Successfully started")
}

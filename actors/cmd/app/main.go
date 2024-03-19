package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/djsega1/filmoteka/actors/internal/config"
	"github.com/djsega1/filmoteka/actors/internal/database"
	"github.com/djsega1/filmoteka/actors/internal/handlers"
	"github.com/djsega1/filmoteka/actors/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	cfg_path := "configs/config.yml"

	cfg := config.GetConfig(cfg_path)

	pool := database.NewPool(context.TODO(), cfg.Storage)
	repo := repository.NewActorsRepository(pool)
	actorsHandler := handlers.NewActorsHandler(repo)

	defer pool.Close()

	r := mux.NewRouter()
	actorsSubR := r.PathPrefix("/api/v1").Subrouter()
	actorsSubR.HandleFunc("/actors", actorsHandler.GetAllActors).Methods(http.MethodGet)
	actorsSubR.HandleFunc("/actors", actorsHandler.CreateActor).Methods(http.MethodPost)
	actorsSubR.HandleFunc("/actors/{id:[0-9]+}", actorsHandler.GetActorByID).Methods(http.MethodGet)
	actorsSubR.HandleFunc("/actors/{id:[0-9]+}", actorsHandler.UpdateActorByID).Methods(http.MethodPatch)
	actorsSubR.HandleFunc("/actors/{id:[0-9]+}", actorsHandler.DeleteActorByID).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port), r))
	log.Println("Successfully started")
}

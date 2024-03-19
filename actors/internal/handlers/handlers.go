package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/djsega1/filmoteka/actors/internal/models"
	"github.com/djsega1/filmoteka/actors/internal/repository"
)

type actorsHandler struct {
	repository repository.Repository
}

func NewActorsHandler(repository repository.Repository) *actorsHandler {
	return &actorsHandler{
		repository: repository,
	}
}

func (h *actorsHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor

	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding request body at CreateActor", http.StatusBadRequest)
		return
	}

	if err := h.repository.Create(context.TODO(), &actor); err != nil {
		http.Error(w, "Error creating actor at CreateActor", http.StatusBadRequest)
		return
	}
}

/*
func (h *actorsHandler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	all, err := repository.Repository.FindAll(repo, context.TODO())
}*/

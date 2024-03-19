package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/djsega1/filmoteka/actors/internal/models"
	"github.com/djsega1/filmoteka/actors/internal/repository"
	"github.com/gorilla/mux"
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

func (h *actorsHandler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	all, err := h.repository.FindAll(context.TODO())

	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
}

func (h *actorsHandler) GetActorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	actor, err := h.repository.FindOne(context.TODO(), vars["id"])

	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	allBytes, err := json.Marshal(actor)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
}

func (h *actorsHandler) UpdateActorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	actor, err := h.repository.FindOne(context.TODO(), vars["id"])

	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		w.WriteHeader(500)
		log.Println(err)
		http.Error(w, "Error decoding request body at CreateActor", http.StatusBadRequest)
		return
	}

	if err := h.repository.Update(context.TODO(), actor); err != nil {
		w.WriteHeader(500)
		log.Println(err)
		http.Error(w, "Error updating actor at UpdateActorByID", http.StatusBadRequest)
		return
	}

	allBytes, err := json.Marshal(actor)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
}

func (h *actorsHandler) DeleteActorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := h.repository.Delete(context.TODO(), vars["id"])

	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

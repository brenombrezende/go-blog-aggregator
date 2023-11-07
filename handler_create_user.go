package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg apiConfig) handleCreateUsers(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request")
		return
	}

	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	})

	convertedUser := databaseUserToUser(newUser)

	respondWithJSON(w, 201, convertedUser)

}

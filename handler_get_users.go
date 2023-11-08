package main

import (
	"net/http"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
)

func (cfg apiConfig) handleGetUsers(w http.ResponseWriter, r *http.Request, user database.User) {

	convertedUser := databaseUserToUser(user)

	respondWithJSON(w, http.StatusAccepted, convertedUser)

}

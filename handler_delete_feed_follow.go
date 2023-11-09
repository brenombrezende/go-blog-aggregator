package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "feedFollowID")
	feedID, err := uuid.Parse(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid Feed ID - %s", err))
		return
	}

	rows, err := cfg.DB.DeleteFeedFollow(r.Context(), feedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to delete feed follow - %s", err))
		return
	}

	respondWithJSON(w, http.StatusAccepted, fmt.Sprintf("Ok - deleted %d rows", rows))
}

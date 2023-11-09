package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg apiConfig) handleCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	type requestBody struct {
		FeedID uuid.UUID
	}

	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request")
		return
	}

	if req.FeedID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "Unable to process request - Invalid feed ID")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    req.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create feed - %s", err))
		return
	}

	convertedFeedFollow := databaseFeedFollowtoFeedFollow(feedFollow)

	respondWithJSON(w, http.StatusAccepted, convertedFeedFollow)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg apiConfig) handleCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	type requestBody struct {
		Name string
		Url  string
	}

	type responseBody struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request")
		return
	}

	if req.Name == "" || req.Url == "" {
		respondWithError(w, http.StatusBadRequest, "Unable to process request - Invalid name or URL")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
		Url:       req.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create feed - %s", err))
		return
	}

	convertedFeed := databaseFeedtoFeed(feed)

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create feed follow - %s", err))
		return
	}

	convertedFeedFollow := databaseFeedFollowtoFeedFollow(feedFollow)

	respondWithJSON(w, http.StatusAccepted, responseBody{
		Feed:       convertedFeed,
		FeedFollow: convertedFeedFollow,
	})

}

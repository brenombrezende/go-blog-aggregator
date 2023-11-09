package main

import (
	"fmt"
	"net/http"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
)

func (cfg apiConfig) handleGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	allFeedFollows, err := cfg.DB.SelectAllFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to get feeds - %s", err))
		return
	}

	convertedFeedFollows := make([]FeedFollow, len(allFeedFollows))

	for i := range allFeedFollows {
		convertedFeedFollows[i] = databaseFeedFollowtoFeedFollow(allFeedFollows[i])
	}

	respondWithJSON(w, http.StatusAccepted, convertedFeedFollows)
}

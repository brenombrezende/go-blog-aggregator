package main

import (
	"fmt"
	"net/http"
)

func (cfg apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	allFeeds, err := cfg.DB.SelectAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to get feeds - %s", err))
		return
	}

	convertedFeeds := make([]Feed, len(allFeeds))

	for i := range allFeeds {
		convertedFeeds[i] = databaseFeedtoFeed(allFeeds[i])
	}

	respondWithJSON(w, http.StatusAccepted, convertedFeeds)
}

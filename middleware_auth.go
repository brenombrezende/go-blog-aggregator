package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		splitKey := strings.Split(apiKey, "ApiKey ")
		if len(splitKey) == 1 {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid Authorization Method"))
			return
		}

		apiKey = splitKey[1]
		if len(apiKey) != 64 {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Incorrect ApiKey Size"))
			return
		}

		user, err := cfg.DB.SelectByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to find user in database, %s", err))
			return
		}

		handler(w, r, user)
	}
}

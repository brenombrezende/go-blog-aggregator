package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/brenombrezende/go-blog-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment file - %s", err)
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("dbURL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to the database - %s", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1 := chi.NewRouter()
	v1.Get("/err", handleError)
	v1.Get("/readiness", handleReadiness)
	v1.Get("/feeds", apiCfg.handleGetFeeds)
	v1.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUsers))
	v1.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollow))

	v1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeeds))
	v1.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollows))
	v1.Post("/users", apiCfg.handleCreateUsers)

	v1.Delete("/feed_follows/{feedFollowID}", apiCfg.handlerDeleteFeedFollows)

	r.Mount("/v1", v1)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Starting server on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}

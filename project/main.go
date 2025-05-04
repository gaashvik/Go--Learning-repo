package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gaashvik/Go--Learning-repo/project/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("hello world")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("could not find")
	}
	con, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot connect")
	} else {
		fmt.Println("connected to db")
		fmt.Println("DB_URL:", dbURL)

	}
	queries := database.New(con)
	apiCFG := apiConfig{
		DB: queries,
	}
	chiz := chi.NewRouter()
	chiz.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/user", apiCFG.handleCreateUser)
	v1Router.Get("/getuser", apiCFG.handleGetUser)
	chiz.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: chiz,
		Addr:    ":" + portString,
	}

	log.Println("server running")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

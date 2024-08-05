package main

import (
	"fmt"
	"go-graphql-api/gql"
	"go-graphql-api/postgres"
	"go-graphql-api/server"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func main() {
	router, db := initialzeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initialzeAPI() (*chi.Mux, *postgres.Db) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := chi.NewRouter()

	db, err := postgres.New(os.Getenv("POSTGRES_URI"))

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		Gqlschema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())

	return router, db

}

package server

import (
	"encoding/json"
	"go-graphql-api/gql"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type Server struct {
	Gqlschema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.Gqlschema)
		render.JSON(w, r, result)
	}
}

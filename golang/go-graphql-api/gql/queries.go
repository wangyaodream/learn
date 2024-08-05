package gql

import (
	"go-graphql-api/postgres"

	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {
	// Create a new root that describes base query set up.
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}

package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{},
				},
			),
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema\nerror: %v", err)
	}

	return schema
}

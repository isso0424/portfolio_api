package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/mutations/skill"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Root",
					Fields: graphql.Fields{
						"FetchSkill": skill.AddSkill,
					},
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Mutation",
					Fields: graphql.Fields{
						"AddSkill": skill.AddSkill,
					},
				},
			),
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema\nerror: %v", err)
	}

	return schema
}

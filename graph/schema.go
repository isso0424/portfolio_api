package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/mutations/contact"
	"github.com/isso0424/portfolio_api/graph/mutations/product"
	"github.com/isso0424/portfolio_api/graph/mutations/skill"
	"github.com/isso0424/portfolio_api/graph/query"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Root",
					Fields: graphql.Fields{
						"Skills": query.FetchSkill,
						"Products": query.FetchProduct,
						"Contacts": query.FetchContact,
					},
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Mutation",
					Fields: graphql.Fields{
						"AddSkill": skill.AddSkill,
						"AddProduct": product.AddProduct,
						"AddContact": contact.AddContact,
						"DeleteContact": contact.DeleteContact,
						"DeleteProduct": product.DeleteProduct,
						"DeleteSkill": skill.DeleteSkill,
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

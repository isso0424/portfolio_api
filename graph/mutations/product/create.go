package product

import "github.com/graphql-go/graphql"

var AddProduct = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"tags": &graphql.ArgumentConfig{
			Type: graphql.NewList(
				graphql.String,
			),
		},
	},
}

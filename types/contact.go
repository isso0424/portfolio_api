package types

import "github.com/graphql-go/graphql"

var Contact = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "contact",
		Fields: graphql.Fields{
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

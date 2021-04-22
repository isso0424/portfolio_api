package types

import "github.com/graphql-go/graphql"

var Skill = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Skill",
		Fields: graphql.Fields{
			"icon": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

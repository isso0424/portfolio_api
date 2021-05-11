package product

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

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
	Type: types.Product,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		token := p.Context.Value("token").(string)
		result, err := auth.ValidateToken(token)
		if err != nil {
			return nil, errors.New("invalid token")
		}
		if !result {
			return nil, errors.New("invalid token")
		}

		title := p.Args["title"].(string)
		description := p.Args["description"].(string)
		tags := p.Args["tags"].([]string)

		product, err := variables.ProductDB.Add(title, description, tags)
		if product == nil {
			return nil, err
		}

		return product, err
	},
}

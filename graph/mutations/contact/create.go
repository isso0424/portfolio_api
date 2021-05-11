package contact

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var AddContact = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"text": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"link": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Type: types.Contact,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		token := p.Context.Value("token").(string)
		result, err := auth.ValidateToken(token)
		if err != nil || !result {
			return nil, errors.New("invalid token")
		}
		name := p.Args["name"].(string)
		text := p.Args["text"].(string)
		link := p.Args["link"].(string)

		contact, err := variables.ContactDB.Add(name, text, link)
		if contact == nil {
			return nil, err
		}

		return *contact, err
	},
}

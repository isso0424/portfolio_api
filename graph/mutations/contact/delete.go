package contact

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var DeleteContact = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
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
		id := p.Args["id"].(string)
		contact, err := variables.ContactDB.Delete(id)
		if contact == nil {
			return nil, err
		}

		return *contact, err
	},
}

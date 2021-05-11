package query

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var FetchContact = &graphql.Field{
	Type: graphql.NewList(types.Contact),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		contacts, err := variables.ContactDB.GetAll()
		if err != nil {
			log.Println("error: ", err)
			return nil, errors.New("internal server error")
		}

		return contacts, nil
	},
}

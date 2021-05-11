package query

import (
	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var FetchContact = &graphql.Field{
	Type: graphql.NewList(types.Contact),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return variables.ContactDB.GetAll()
	},
}

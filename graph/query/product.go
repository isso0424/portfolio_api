package query

import (
	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var FetchProduct = &graphql.Field{
	Type: graphql.NewList(types.Product),
	Resolve: func(p graphql.ResolveParams) (interface{}, error){
		return variables.ProductDB.GetAll()
	},
}

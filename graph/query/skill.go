package query

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var FetchSkill = &graphql.Field{
	Type: graphql.NewList(types.Skill),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		skills := variables.SkillDB.GetAll()
		log.Printf("%v\n", skills)
		return skills, nil
	},
}

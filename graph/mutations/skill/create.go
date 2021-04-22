package skill

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var AddSkill = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"icon": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Type: types.Skill,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		icon := p.Args["icon"].(string)
		name := p.Args["name"].(string)
		log.Printf("%v\n", p.Args)

		return variables.SkillDB.Add(name, icon), nil
	},
}

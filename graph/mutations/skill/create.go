package skill

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/auth"
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
		token := p.Context.Value("token").(string)
		result, err := auth.ValidateToken(token)
		if err != nil {
			return nil, errors.New("invalid token")
		}
		if !result {
			return nil, errors.New("invalid token")
		}
		icon := p.Args["icon"].(string)
		name := p.Args["name"].(string)
		log.Printf("%v\n", p.Args)

		skill, err := variables.SkillDB.Add(name, icon)
		if skill == nil {
			return nil, err
		}

		return *skill, err
	},
}

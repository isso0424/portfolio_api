package skill

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/types"
)

var DeleteSkill = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Type: types.Skill,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		token := p.Context.Value("token").(string)
		result, err := auth.ValidateToken(token)
		if err != nil || !result {
			return nil, errors.New("invalid token")
		}
		name := p.Args["name"].(string)

		skill, err := variables.SkillDB.Delete(name)
		if skill == nil {
			return nil, err
		}

		return *skill, err
	},
}

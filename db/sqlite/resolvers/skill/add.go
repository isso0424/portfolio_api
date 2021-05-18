package skill

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *SkillDB) Add(name, icon string) (*domain.Skill, types.APIError) {
	newData := &model.Skill{}
	newData.Name = name
	newData.Icon = icon
	err := db.db.Create(&newData).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &newData.Skill, nil
}

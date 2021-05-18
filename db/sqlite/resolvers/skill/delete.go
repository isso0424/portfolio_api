package skill

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *SkillDB) Delete(id string) (*domain.Skill, types.APIError) {
	data := &model.Skill{}
	err := db.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &data.Skill, nil
}

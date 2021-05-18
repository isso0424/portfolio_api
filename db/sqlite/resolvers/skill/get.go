package skill

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *SkillDB) GetAll() ([]domain.Skill, types.APIError) {
	products := []model.Skill{}
	err := db.db.Find(&products).Error
	if err != nil {
		return []domain.Skill{}, internal_error.New(err.Error())
	}

	data := []domain.Skill{}

	for _, d := range products {
		data = append(data, domain.Skill{
			Name: d.Name,
			Icon: d.Icon,
			ID: d.ID,
		})
	}

	return data, nil
}

func(db *SkillDB) GetByID(id string) (*domain.Skill, types.APIError) {
	skill := model.Skill{}
	err := db.db.Where("id = ?", id).Find(&skill).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &skill.Skill, nil
}

func(db *SkillDB) GetByName(name string) (*domain.Skill, types.APIError) {
	skill := model.Skill{}
	err := db.db.Where("name = ?", name).Find(&skill).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &skill.Skill, nil
}

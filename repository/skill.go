package repository

import (
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
)

type SkillRepository interface {
	Add(name, icon string) (*domain.Skill, types.APIError)
	Delete(name string) (*domain.Skill, types.APIError)
	GetAll() ([]domain.Skill, types.APIError)
	GetByName(name string) (*domain.Skill, types.APIError)
}

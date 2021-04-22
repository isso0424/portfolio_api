package repository

import "github.com/isso0424/portfolio_api/domain"

type SkillRepository interface {
	Add(name, icon string) domain.Skill
	Delete(name, icon string) domain.Skill
	GetAll() []domain.Skill
	GetByName(name string) domain.Skill
}

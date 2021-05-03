package repository

import "github.com/isso0424/portfolio_api/domain"

type SkillRepository interface {
	Add(name, icon string) (*domain.Skill, error)
	Delete(name string) (*domain.Skill, error)
	GetAll() ([]domain.Skill, error)
	GetByName(name string) (*domain.Skill, error)
}

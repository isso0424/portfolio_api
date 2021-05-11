package mem_db

import (
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/user_error"
)

type SkillDB struct {
	data []domain.Skill
}

func (db *SkillDB) Add(name, icon string) (*domain.Skill, types.APIError) {
	newSkill := domain.Skill{
		Name: name,
		Icon: icon,
	}

	if exist, index, _ := db.searchByName(name); exist {
		db.data[index] = newSkill
	} else {
		db.data = append(db.data, newSkill)
	}

	return &newSkill, nil
}

func (db *SkillDB) Delete(name string) (*domain.Skill, types.APIError) {
	if exist, index, skill := db.searchByName(name); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)

		return skill, nil
	}

	return nil, user_error.New("Skill: %s(name) is not found")
}

func (db *SkillDB) GetAll() ([]domain.Skill, types.APIError) {
	return db.data, nil
}

func (db *SkillDB) GetByName(name string) (*domain.Skill, types.APIError) {
	if exist, _, skill := db.searchByName(name); exist {
		return skill, nil
	}

	return nil, user_error.New("Skill: %s(name) is not found")
}

func (db *SkillDB) searchByName(name string) (bool, int, *domain.Skill) {
	for index, skill := range db.data {
		if skill.Name == name {
			return true, index, &skill
		}
	}

	return false, -1, nil
}

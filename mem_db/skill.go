package mem_db

import "github.com/isso0424/portfolio_api/domain"

type SkillDB struct {
	data []domain.Skill
}

func(db *SkillDB) Add(name, icon string) domain.Skill {
	newSkill := domain.Skill{
		Name: name,
		Icon: icon,
	}

	if exist, index, _ := db.searchByName(name); exist {
		db.data[index] = newSkill
	} else {
		db.data = append(db.data, newSkill)
	}

	return newSkill
}

func(db *SkillDB) Delete(name string) domain.Skill {
	if exist, index, skill := db.searchByName(name); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)

		return skill
	}

	return domain.Skill{}
}

func(db *SkillDB) GetAll() []domain.Skill {
	return db.data
}

func(db *SkillDB) GetByName(name string) domain.Skill {
	if exist, _, skill := db.searchByName(name); exist {
		return skill
	}

	return domain.Skill{}
}

func(db *SkillDB) searchByName(name string) (bool, int, domain.Skill) {
	for index, skill := range db.data {
		if skill.Name == name {
			return true, index, skill
		}
	}

	return false, -1, domain.Skill{}
}

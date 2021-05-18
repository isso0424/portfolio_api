package skill

import "gorm.io/gorm"

type SkillDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) SkillDB {
	return SkillDB{db}
}

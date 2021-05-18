package model

import (
	"github.com/isso0424/portfolio_api/domain"
	"gorm.io/gorm"
)

type Skill struct {
	domain.Skill
	gorm.Model
	ID string `gorm:"primarykey"`
}

package model

import (
	"github.com/isso0424/portfolio_api/domain"
	"gorm.io/gorm"
)

type Contact struct {
	domain.Contact
	gorm.Model
	ID string `gorm:"primarykey"`
}

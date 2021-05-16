package model

import (
	"github.com/isso0424/portfolio_api/domain"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	domain.Contact
}

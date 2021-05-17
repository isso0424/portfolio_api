package model

import (
	"github.com/isso0424/portfolio_api/domain"
	"gorm.io/gorm"
)

type Product struct {
	domain.Product
	gorm.Model
	Tags string
	ID string `gorm:"primarykey"`
}

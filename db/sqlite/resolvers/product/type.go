package product

import "gorm.io/gorm"

type ProductDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProductDB {
	return ProductDB{ db }
}

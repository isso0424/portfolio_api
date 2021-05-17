package contact

import "gorm.io/gorm"

type ContactDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) ContactDB {
	return ContactDB{ db }
}

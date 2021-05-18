package sqlite

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/db/sqlite/resolvers/contact"
	"github.com/isso0424/portfolio_api/db/sqlite/resolvers/product"
	"github.com/isso0424/portfolio_api/db/sqlite/resolvers/skill"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dbSet struct {
	Skill skill.SkillDB
	Product product.ProductDB
	Contact contact.ContactDB
}

func Initialize(dbPath string) (dbSet, error) {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		return dbSet{}, err
	}

	db.AutoMigrate(&model.Contact{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Skill{})

	return dbSet{
		Skill: skill.New(db),
		Product: product.New(db),
		Contact: contact.New(db),
	}, nil
}

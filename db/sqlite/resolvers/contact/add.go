package contact

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *ContactDB) Add(name, text, link string) (*domain.Contact, types.APIError) {
	newData := model.Contact{}
	newData.Name = name
	newData.Text = text
	newData.Link = link

	err := db.db.Create(&newData).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &newData.Contact, nil
}

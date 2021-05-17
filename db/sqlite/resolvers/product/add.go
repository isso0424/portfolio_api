package product

import (
	"encoding/json"

	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func (db *ProductDB) Add(title, description string, tags[]string) (*domain.Product, types.APIError) {
	product := model.Product{}
	product.Title = title
	product.Description = description
	data, err := json.Marshal(tags)
	if err != nil {
		return nil, internal_error.New(err.Error())
	}
	product.Tags = string(data)

	err = db.db.Create(&product).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &product.Product, nil
}

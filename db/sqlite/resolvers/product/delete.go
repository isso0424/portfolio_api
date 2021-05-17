package product

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func (db ProductDB) Delete(id string) (*domain.Product, types.APIError) {
	product := &model.Product{}
	err := db.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &product.Product, nil
}

package product

import (
	"encoding/json"

	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *ProductDB) GetAll() ([]domain.Product, types.APIError) {
	products := []model.Product{}
	err := db.db.Find(&products).Error
	if err != nil {
		return []domain.Product{}, internal_error.New(err.Error())
	}

	data := []domain.Product{}
	for _, d := range products {
		var tags []string
		err := json.Unmarshal([]byte(d.Tags), &tags)
		if err != nil {
			return []domain.Product{}, internal_error.New(err.Error())
		}
		data = append(data, domain.Product{
			Title: d.Title,
			ID: d.ID,
			Description: d.Description,
			Tags: tags,
		})
	}

	return data, nil
}

func(db *ProductDB) GetByID(id string) (*domain.Product, types.APIError) {
	data := model.Product{}
	err := db.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &data.Product, nil
}

func(db *ProductDB) GetByTitle(title string) (*domain.Product, types.APIError) {
	data := model.Product{}
	err := db.db.Where("title = ?", title).First(&data).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &data.Product, nil
}

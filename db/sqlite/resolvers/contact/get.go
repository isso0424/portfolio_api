package contact

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *ContactDB) GetAll() ([]domain.Contact, types.APIError) {
	contacts := []model.Contact{}
	err := db.db.Find(&contacts).Error
	if err != nil {
		return []domain.Contact{}, internal_error.New(err.Error())
	}

	data := []domain.Contact{}
	for _, d := range contacts {
		data = append(data, domain.Contact{
			Name: d.Name,
			ID: d.ID,
			Link: d.Link,
			Text: d.Text,
		})
	}

	return data, nil
}

func(db *ContactDB) GetByID(id string) (*domain.Contact, types.APIError) {
	data := model.Contact{}
	err := db.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &data.Contact, nil
}

func(db *ContactDB) GetByName(name string) (*domain.Contact, types.APIError) {
	data := model.Contact{}
	err := db.db.Where("name = ?", name).First(&data).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &data.Contact, nil
}

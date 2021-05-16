package mem_db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/repository"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/user_error"
)

type ProductDB struct {
	repository.ProductRepository
	data []domain.Product
}

func (db *ProductDB) Add(title, description string, tags []string) (*domain.Product, types.APIError) {
	id := uuid.New().String()
	newProduct := domain.Product{
		ID: id,
		Title:       title,
		Description: description,
		Tags:        tags,
	}

	if exist, index, _ := db.searchByID(id); exist {
		db.data[index] = newProduct
	} else {
		db.data = append(db.data, newProduct)
	}

	return &newProduct, nil
}

func (db *ProductDB) Delete(id string) (*domain.Product, types.APIError) {
	if exist, index, product := db.searchByID(id); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)

		return product, nil
	}

	return nil, user_error.New(fmt.Sprintf("Product: %s(id) is not found", id))
}

func (db *ProductDB) GetAll() ([]domain.Product, types.APIError) {
	return db.data, nil
}

func (db *ProductDB) GetByTitle(title string) (*domain.Product, types.APIError) {
	if exist, _, product := db.searchByTitle(title); exist {
		return product, nil
	}

	return nil, user_error.New(fmt.Sprintf("Product: %s(title) is not found", title))
}

func (db *ProductDB) GetByID(id string) (*domain.Product, types.APIError) {
	if exist, _, product := db.searchByID(id); exist {
		return product, nil
	}

	return nil, user_error.New(fmt.Sprintf("Product: %s(id) is not found", id))
}

func (db *ProductDB) searchByTitle(title string) (bool, int, *domain.Product) {
	for index, product := range db.data {
		if product.Title == title {
			return true, index, &product
		}
	}

	return false, -1, nil
}

func (db *ProductDB) searchByID(id string) (bool, int, *domain.Product) {
	for index, product := range db.data {
		if product.ID == id {
			return true, index, &product
		}
	}

	return false, -1, nil
}

package mem_db

import (
	"github.com/isso0424/portfolio_api/domain"
)

type ProductDB struct {
	data []domain.Product
}

func (db *ProductDB) Add(title, description string, tags []string) (*domain.Product, error) {
	newProduct := domain.Product{
		Title:       title,
		Description: description,
		Tags:        tags,
	}

	if exist, index, _ := db.searchByTitle(title); exist {
		db.data[index] = newProduct
	} else {
		db.data = append(db.data, newProduct)
	}

	return &newProduct, nil
}

func (db *ProductDB) Delete(title string) (*domain.Product, error) {
	if exist, index, product := db.searchByTitle(title); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)

		return product, nil
	}

	return nil, nil
}

func (db *ProductDB) GetAll() ([]domain.Product, error) {
	return db.data, nil
}

func (db *ProductDB) GetByTitle(title string) (*domain.Product, error) {
	if exist, _, product := db.searchByTitle(title); exist {
		return product, nil
	}

	return nil, nil
}

func (db *ProductDB) searchByTitle(title string) (bool, int, *domain.Product) {
	for index, product := range db.data {
		if product.Title == title {
			return true, index, &product
		}
	}

	return false, -1, nil
}

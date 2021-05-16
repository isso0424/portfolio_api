package repository

import (
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
)

type ProductRepository interface {
	Add(title, description string, tags []string) (*domain.Product, types.APIError)
	Delete(id string) (*domain.Product, types.APIError)
	GetAll() ([]domain.Product, types.APIError)
	GetByTitle(title string) (*domain.Product, types.APIError)
	GetByID(id string) (*domain.Product, types.APIError)
}

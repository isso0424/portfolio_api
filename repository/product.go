package repository

import "github.com/isso0424/portfolio_api/domain"

type ProductRepository interface {
	Add(title, description string, tags []string) (*domain.Product, error)
	Delete(title string) (*domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetByTitle(title string) (*domain.Product, error)
}

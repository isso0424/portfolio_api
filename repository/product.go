package repository

import "github.com/isso0424/portfolio_api/domain"

type ProductRepository interface {
	Add(title, description string, tags []string) domain.Product
	Delete(title string) domain.Product
	GetAll() []domain.Product
	GetByTitle(title string) domain.Product
}

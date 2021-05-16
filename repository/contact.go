package repository

import (
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
)

type ContactRepository interface {
	Add(name, text, link string) (*domain.Contact, types.APIError)
	Delete(id string) (*domain.Contact, types.APIError)
	GetAll() ([]domain.Contact, types.APIError)
	GetByName(name string) (*domain.Contact, types.APIError)
	GetByID(id string) (*domain.Contact, types.APIError)
}

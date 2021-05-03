package repository

import "github.com/isso0424/portfolio_api/domain"

type ContactRepository interface {
	Add(name, text, link string) (*domain.Contact, error)
	Delete(name string) (*domain.Contact, error)
	GetAll() ([]domain.Contact, error)
	GetByName(name string) (*domain.Contact, error)
}

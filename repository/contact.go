package repository

import "github.com/isso0424/portfolio_api/domain"

type ContactRepository interface {
	Add(name, text, link string) domain.Contact
	Delete(name string) domain.Contact
	GetAll() []domain.Contact
	GetByName(name string) domain.Contact
}

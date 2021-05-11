package mem_db

import (
	"fmt"

	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/user_error"
)

type ContactDB struct {
	data []domain.Contact
}

func (db *ContactDB) Add(name, text, link string) (*domain.Contact, types.APIError) {
	newContact := domain.Contact{
		Name: name,
		Text: text,
		Link: link,
	}

	if exist, index, _ := db.searchByName(name); exist {
		db.data[index] = newContact
	} else {
		db.data = append(db.data, newContact)
	}

	return &newContact, nil
}

func (db *ContactDB) searchByName(name string) (bool, int, *domain.Contact) {
	for index, contact := range db.data {
		if contact.Name == name {
			return true, index, &contact
		}
	}

	return false, -1, nil
}

func (db *ContactDB) Delete(name string) (*domain.Contact, types.APIError) {
	if exist, index, contact := db.searchByName(name); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)
		return contact, nil
	}

	return nil, user_error.New(fmt.Sprintf("Contact: %s(Name) is not found", name))
}

func (db *ContactDB) GetAll() ([]domain.Contact, types.APIError) {
	return db.data, nil
}

func (db *ContactDB) GetByName(name string) (*domain.Contact, types.APIError) {
	exist, _, contact := db.searchByName(name)

	if !exist {
		return nil, user_error.New(fmt.Sprintf("Contact: %s(Name) is not found", name))
	}

	return contact, nil
}

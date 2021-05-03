package mem_db

import "github.com/isso0424/portfolio_api/domain"

type ContactDB struct {
	data []domain.Contact
}

func (db *ContactDB) Add(name, text, link string) (*domain.Contact, error) {
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

func (db *ContactDB) Delete(name string) (*domain.Contact, error) {
	if exist, index, contact := db.searchByName(name); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)
		return contact, nil
	}

	return nil, nil
}

func (db *ContactDB) GetAll() ([]domain.Contact, error) {
	return db.data, nil
}

func (db *ContactDB) GetByName(name string) (*domain.Contact, error) {
	_, _, contact := db.searchByName(name)

	return contact, nil
}

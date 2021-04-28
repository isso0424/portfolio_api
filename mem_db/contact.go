package mem_db

import "github.com/isso0424/portfolio_api/domain"

type ContactDB struct {
	data []domain.Contact
}

func (db *ContactDB) Add(name, text, link string) domain.Contact {
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

	return newContact
}

func (db *ContactDB) searchByName(name string) (bool, int, domain.Contact) {
	for index, contact := range db.data {
		if contact.Name == name {
			return true, index, contact
		}
	}

	return false, -1, domain.Contact{}
}

func (db *ContactDB) Delete(name string) domain.Contact {
	if exist, index, contact := db.searchByName(name); exist {
		db.data = append(db.data[:index], db.data[index+1:]...)
		return contact
	}

	return domain.Contact{}
}

func (db *ContactDB) GetAll() []domain.Contact {
	return db.data
}

func (db *ContactDB) GetByName(name string) domain.Contact {
	_, _, contact := db.searchByName(name)

	return contact
}

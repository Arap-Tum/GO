package store

import "contactManagerApp/internal/models"

/*Uses slices/maps
Fast, temporary*/

type MemoryStore struct {
	contacts []models.Contact
}

// constructor
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: []models.Contact{},
	}
}

// Get all

func (s *MemoryStore) GetAll() []models.Contact {
	return s.contacts
}

// get by id
func (s *MemoryStore) GetByID(id string) (models.Contact, bool) {
	for _, contact := range s.contacts {
		if contact.ID == id {
			return contact, true
		}

	}
	return models.Contact{}, false
}

// CRETE
func (s *MemoryStore) Create(contact models.Contact) {
	s.contacts = append(s.contacts, contact)
}

// update (PUT)
func (s *MemoryStore) Update(id string, updated models.Contact) (models.Contact, bool) {
	for i, contact := range s.contacts {
		if contact.ID == id {
			updated.ID = id //ensure ID doesnt change
			s.contacts[i] = updated
			return updated, true
		}
	}
	return models.Contact{}, false
}

// 	DELETE
func (s *MemoryStore) Delete(id string) bool {
	for i, contact := range s.contacts {

		if contact.ID == id {
			// remove from theslice

			s.contacts = append(s.contacts[:i], s.contacts[1+1:]...)
			return true
		}

	}
	return false
}

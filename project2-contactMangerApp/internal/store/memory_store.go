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

// CRETE
func (s *MemoryStore) Create(contact models.Contact) {
	s.contacts = append(s.contacts, contact)
}

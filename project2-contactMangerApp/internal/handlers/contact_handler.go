package handlers

import (
	"contactManagerApp/internal/models"
	"contactManagerApp/internal/store"
	"encoding/json"
	"net/http"
)

/*
Where HTTP logic lives

👉 Each function:

Reads request
Calls store
Returns JSON response
*/

// this controlls how Go becomes Json

type ContactHandlers struct {
	store *store.MemoryStore
}

// constructor
func NewContactHandler(s *store.MemoryStore) *ContactHandlers {
	return &ContactHandlers{store: s}
}

// GET  /contacts
func (h *ContactHandlers) GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contacts := h.store.GetAll()

	json.NewEncoder(w).Encode(contacts)
}

// POST  /CONTACT
func (h *ContactHandlers) CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.store.Create(contact)

	json.NewEncoder(w).Encode(contact)
}

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

// GET /contacts/:id
func (h *ContactHandlers) GetContactByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract id from theurl
	id := r.URL.Path[len("/contacts/"):]

	contact, found := h.store.GetByID(id)
	if !found {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(contact)
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

// PUT /contact / : id
func (h *ContactHandlers) UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/contacts/"):]

	var updated models.Contact

	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	contact, found := h.store.Update(id, updated)

	if !found {
		http.Error(w, "Contacct not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

// DELETE /contact/:id
func (h *ContactHandlers) DeleteContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/contacts/"):]

	deleted := h.store.Delete(id)
	if !deleted {
		http.Error(w, "Contact not found ", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Contact deleted successfully",
	})
}

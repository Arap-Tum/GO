package router

import (
	"contactManagerApp/internal/handlers"
	"contactManagerApp/internal/store"
	"net/http"
)

/*Connects routes → handlers*/

func SetupRoutes() {
	store := store.NewMemoryStore()
	handler := handlers.NewContactHandler(store)

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetContacts(w, r)
		} else if r.Method == http.MethodPost {
			handler.CreateContact(w, r)
		}
	})

	http.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			handler.GetContactByID(w, r)

		case http.MethodPut:
			handler.UpdateContact(w, r)

		case http.MethodDelete:
			handler.DeleteContacts(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
}

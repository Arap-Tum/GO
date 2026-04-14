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
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

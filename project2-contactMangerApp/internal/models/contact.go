package models

// Data structure

// define
/*What is a Contact?
Fields: ID, Name, Email, Phone
*/

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

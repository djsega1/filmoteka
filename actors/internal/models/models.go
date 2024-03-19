package models

type Actor struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

package models

type Dog struct {
	ID       int    `json:"id"`
	Breed    string `json:"breed"`
	SubBreed string `json:"sub_breed"`
}

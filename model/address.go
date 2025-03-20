package model

import "github.com/google/uuid"

type CreateAddress struct {
	RecipentName    string    `json:"recipent_name"`
	PhoneNumber     string    `json:"phone_number"`
	RecipentAddress string    `json:"recipent_address"`
	PostalCode      string    `json:"postal_code"`
	UserID          uuid.UUID `json:"user_id"`
}

type UpdateAddress struct {
	RecipentName    string `json:"recipent_name"`
	PhoneNumber     string `json:"phone_number"`
	RecipentAddress string `json:"recipent_address"`
	PostalCode      string `json:"postal_code"`
}

type AddressResponse struct {
	UserID          string `json:"user_id"`
	RecipentName    string `json:"recipent_name"`
	PhoneNumber     string `json:"phone_number"`
	RecipentAddress string `json:"recipent_address"`
	PostalCode      string `json:"postal_code"`
}

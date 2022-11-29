package public

import (
	"github.com/google/uuid"
)

// AddressResponse represents address response
type AddressResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Country     string    `json:"country"`
	Province    string    `json:"province"`
	Region      string    `json:"region"`
	District    string    `json:"district"`
	Village     string    `json:"village"`
	FullAddress string    `json:"full_address"`
	ZipCode     string    `json:"zip_code"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	AddressName string    `json:"address_name"`
	IsDefault   bool      `json:"is_default"`
}

// ListAddressRequest represents params to get List address
type ListAddressRequest struct {
	Search  string       `json:"search"` // province, region, district, village, full_address, zip_code, address_name
	Page    int          `json:"page"`
	Limit   int          `json:"limit"`
	UserID  *uuid.UUID   `json:"user_id"`
	UserIDs *[]uuid.UUID `json:"user_ids"`
}

// CreateAddressRequest represents params to create new address
type CreateAddressRequest struct {
	Country     string `json:"country" validate:"required"`
	Province    string `json:"province" validate:"required"`
	Region      string `json:"region" validate:"required"`
	District    string `json:"district" validate:"required"`
	Village     string `json:"village" validate:"required"`
	FullAddress string `json:"full_address" validate:"required"`
	ZipCode     string `json:"zip_code" validate:"required"`
	Latitude    string `json:"latitude" validate:"required"`
	Longitude   string `json:"longitude" validate:"required"`
	AddressName string `json:"address_name" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
}

// UpdateAddressRequest represents params to create new address
type UpdateAddressRequest struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid4"`
	Country     string    `json:"country" validate:"required"`
	Province    string    `json:"province" validate:"required"`
	Region      string    `json:"region" validate:"required"`
	District    string    `json:"district" validate:"required"`
	Village     string    `json:"village" validate:"required"`
	FullAddress string    `json:"full_address" validate:"required"`
	ZipCode     string    `json:"zip_code" validate:"required"`
	Latitude    string    `json:"latitude" validate:"required"`
	Longitude   string    `json:"longitude" validate:"required"`
	AddressName string    `json:"address_name" validate:"required"`
	IsDefault   bool      `json:"is_default" validate:"required"`
}

// DeleteAddressRequest represents params to delete address
type DeleteAddressRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetAddressRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type SetDefaultAddressRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid4"`
}

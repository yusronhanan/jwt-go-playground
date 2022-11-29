package public

import (
	"time"

	"github.com/google/uuid"
)

type CreatePatientRequest struct {
	User CreateUserRequest `json:"user"`
}

type PatientResponse struct {
	ID          uuid.UUID    `json:"id"`
	User        UserResponse `json:"user"`
	IDNumber    *string      `json:"id_number"`
	DateOfBirth *time.Time   `json:"date_of_birth"`
	Gender      *string      `json:"gender"`
}

type UpdatePatientRequest struct {
	User        UpdateUserRequest `json:"user"`
	DateOfBirth time.Time         `json:"date_of_birth"`
	Gender      string            `json:"gender"`
}

type UpdatePatientIDNumberRequest struct {
	OTP      string `json:"otp" validate:"required"`
	IDNumber string `json:"id_number" validate:"required"`
}

// ListPatientsRequest represents params to get List patient
type ListPatientsRequest struct {
	Search  string      `json:"search"` //fullname, email, phone_number, id_number
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	UserID  uuid.UUID   `json:"user_id"`
	UserIDs []uuid.UUID `json:"user_ids"`
	ID      uuid.UUID   `json:"id"`
	IDs     []uuid.UUID `json:"ids"`
}

// GetPatientProfileRequest represents params to get patient data
type GetPatientProfileRequest struct {
	UserID uuid.UUID `qs:"user_id" validate:"required"`
}

// UpdatePatientAccountRequest represents params to update patient account
type UpdatePatientAccountRequest struct {
	ID          uuid.UUID         `json:"id" validate:"required,uuid4"`
	User        UpdateUserRequest `json:"user" validate:"required"`
	IDNumber    string            `json:"id_number" validate:"required"`
	DateOfBirth time.Time         `json:"date_of_birth" validate:"required"`
	Phone       string            `json:"phone" validate:"required"`
	Image       uuid.UUID         `json:"image" validate:"required,uuid4"`
	Role        string            `json:"role" validate:"required,oneof=patient admin"`
}

// UpdatePatientAccountResponse represents params to update patient account
type UpdatePatientAccountResponse struct {
	User    UserResponse    `json:"user"`
	Patient PatientResponse `json:"patient"`
}

type DeletePatientRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
}

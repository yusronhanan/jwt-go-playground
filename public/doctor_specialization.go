package public

import (
	"github.com/google/uuid"
)

type DoctorSpecializationResponse struct {
	ID             uuid.UUID `json:"id"`
	Specialization string    `json:"specialization"`
	Description    string    `json:"description"`
}

type CreateDoctorSpecializationRequest struct {
	Specialization string `json:"specialization"`
	Description    string `json:"description"`
}

type UpdateDoctorSpecializationRequest struct {
	ID             uuid.UUID `json:"id"`
	Specialization string    `json:"specialization"`
	Description    string    `json:"description"`
}

// ListDoctorSpecializationRequest represents params to get List specialization
type ListDoctorSpecializationRequest struct {
	Search string      `json:"search"`
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	ID     uuid.UUID   `json:"id"`
	IDs    []uuid.UUID `json:"ids"`
}

type DeleteDoctorSpecializationRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetDoctorSpecializationRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetDoctorSpecializationByDoctorIDRequest struct {
	DoctorID uuid.UUID `url_param:"doctor_id" validate:"required,uuid4"`
}

type GetDoctorSpecializationByUserIDRequest struct {
	UserID uuid.UUID `url_param:"user_id" validate:"required,uuid4"`
}

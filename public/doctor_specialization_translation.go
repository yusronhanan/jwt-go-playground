package public

import (
	"github.com/google/uuid"
)

type ListDoctorSpecializationTranslationRequest struct {
	Search            string      `json:"search"`
	Page              int         `json:"page"`
	Limit             int         `json:"limit"`
	SpecializationID  uuid.UUID   `json:"specialization_id"`
	SpecializationIDs []uuid.UUID `json:"specialization_ids"`
	LanguageCode      string      `json:"language_code" validate:"oneof=id en"`
}

type DoctorSpecializationTranslationResponse struct {
	ID                   uuid.UUID                     `json:"id"`
	SpecializationID     uuid.UUID                     `json:"specialization_id"`
	DoctorSpecialization *DoctorSpecializationResponse `json:"doctor_specialization"`
	Specialization       string                        `json:"specialization"`
	Description          string                        `json:"description"`
	LanguageCode         string                        `json:"language_code"`
}

type CreateDoctorSpecializationTranslationRequest struct {
	SpecializationID uuid.UUID `json:"specialization_id" validate:"required,uuid4"`
	Specialization   string    `json:"specialization" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	LanguageCode     string    `json:"language_code" validate:"required,oneof=id en"`
}
type UpdateDoctorSpecializationTranslationRequest struct {
	ID               uuid.UUID `json:"id" validate:"required,uuid4"`
	SpecializationID uuid.UUID `json:"specialization_id" validate:"required,uuid4"`
	Specialization   string    `json:"specialization" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	LanguageCode     string    `json:"language_code" validate:"required,oneof=id en"`
}

type DeleteDoctorSpecializationTranslationRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetDoctorSpecializationTranslationRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetDoctorSpecializationTranslationBySpecializationIDRequest struct {
	SpecializationID uuid.UUID `qs:"specialization_id" validate:"required,uuid4"`
	LanguageCode     string    `qs:"language_code" validate:"required,oneof=id en"`
}

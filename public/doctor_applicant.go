package public

import (
	"time"

	"github.com/google/uuid"
)

type DoctorApplicantResponse struct {
	ID                 uuid.UUID                                `json:"id"`
	FullName           string                                   `json:"full_name"`
	SpecializationID   uuid.UUID                                `json:"specialization_id"`
	Specialization     *DoctorSpecializationTranslationResponse `json:"specialization"`
	Email              string                                   `json:"email"`
	PhoneNumber        string                                   `json:"phone_number"`
	DoctorIDNumber     string                                   `json:"doctor_id_number"`
	HospitalID         *uuid.UUID                               `json:"hospital_id"`
	DateOfBirth        time.Time                                `json:"date_of_birth"`
	Gender             string                                   `json:"gender"`
	RegistrationLetter string                                   `json:"registration_letter"`
	PracticeLetter     string                                   `json:"practice_letter"`
	CV                 string                                   `json:"cv"`
	Status             string                                   `json:"status"`
	RejectedReason     *string                                  `json:"rejected_reason"`
}

type CreateDoctorApplicantRequest struct {
	FullName           string     `json:"full_name" validate:"required"`
	SpecializationID   uuid.UUID  `json:"specialization_id" validate:"required,uuid4"`
	Email              string     `json:"email" validate:"required"`
	PhoneNumber        string     `json:"phone_number" validate:"required"`
	DoctorIDNumber     string     `json:"doctor_id_number" validate:"required"`
	HospitalID         *uuid.UUID `json:"hospital_id"`
	DateOfBirth        time.Time  `json:"date_of_birth" validate:"required"`
	Gender             string     `json:"gender" validate:"required,oneof=male female"`
	RegistrationLetter string     `json:"registration_letter" validate:"required"`
	PracticeLetter     string     `json:"practice_letter" validate:"required"`
	CV                 string     `json:"cv" validate:"required"`
}

type UpdateDoctorApplicantRequest struct {
	ID                 uuid.UUID  `json:"id" validate:"required,uuid4"`
	FullName           string     `json:"full_name" validate:"required"`
	SpecializationID   uuid.UUID  `json:"specialization_id" validate:"required,uuid4"`
	Email              string     `json:"email" validate:"required"`
	PhoneNumber        string     `json:"phone_number" validate:"required"`
	DoctorIDNumber     string     `json:"doctor_id_number" validate:"required"`
	HospitalID         *uuid.UUID `json:"hospital_id"`
	DateOfBirth        time.Time  `json:"date_of_birth" validate:"required"`
	Gender             string     `json:"gender" validate:"required,oneof=male female"`
	RegistrationLetter string     `json:"registration_letter" validate:"required"`
	PracticeLetter     string     `json:"practice_letter" validate:"required"`
	CV                 string     `json:"cv" validate:"required"`
}

type DeleteDoctorApplicantRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required,uuid4"`
}

type GetDoctorApplicantRequest struct {
	ID           uuid.UUID `qs:"id" validate:"required,uuid4"`
	LanguageCode string    `qs:"language_code" validate:"required,oneof=id en"`
}

type ListDoctorApplicantRequest struct {
	Search            string      `json:"search"` //fullname, email, phone_number, doctor_id_number
	Page              int         `json:"page"`
	Limit             int         `json:"limit"`
	Status            *string     `json:"status"`
	ID                uuid.UUID   `json:"id"`
	IDs               []uuid.UUID `json:"ids"`
	SpecializationIDs []uuid.UUID `json:"specializationIDs"`
	LanguageCode      string      `json:"language_code" validate:"required,oneof=id en"`
}

type AcceptDoctorApplicantRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid4"`
}

type RejectDoctorApplicantRequest struct {
	ID             uuid.UUID `json:"id" validate:"required,uuid4"`
	RejectedReason string    `json:"rejected_reason"`
}

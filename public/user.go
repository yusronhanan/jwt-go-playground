package public

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents the user response object
type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	FullName      *string   `json:"full_name"`
	Email         *string   `json:"email"`
	PhoneNumber   *string   `json:"phone_number"`
	Image         *string   `json:"image"`
	Role          *string   `json:"role"`
	EmailVerified *bool     `json:"email_verified"`
	PhoneVerified *bool     `json:"phone_verified"`
	IsBlocked     *bool     `json:"is_blocked"`
}

// UpdateUserRequest represents the update user request object
type UpdateUserRequest struct {
	ID       uuid.UUID `url_param:"id" validate:"required,uuid4"`
	FullName string    `json:"full_name"`
	Image    string    `json:"image"`
}

// CreateUserRequest represents the create user request object
type CreateUserRequest struct {
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Role        string `json:"role" validate:"required,oneof=doctor patient admin"`
}

//ListUsersRequest params for find all user
type ListUsersRequest struct {
	Page     int         `json:"page" validate:"gt=0"`
	Search   string      `json:"search"` //fullname, email, phonenumber
	Limit    int         `json:"limit"`
	Location string      `json:"location"` // country, region, province
	ID       uuid.UUID   `json:"id"`
	IDs      []uuid.UUID `json:"ids"`
	Role     string      `json:"role" validate:"oneof=doctor patient admin"`
}

// RegisterUserRequest represents params to register a user
type RegisterUserRequest struct {
	FullName    string  `json:"full_name" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	Role        string  `json:"role" validate:"required,oneof=patient"`
	OTP         *string `json:"otp"`
}

// LoginResponse represents the response of login function
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiredAt    time.Time    `json:"expired_at"`
	User         UserResponse `json:"user"`
}

// LoginRequest represent the http request data for login user
type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=doctor patient admin"`
}

// LoginByOTPRequest represent the http request data for login by otp
type LoginByOTPRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	OTP        string `json:"otp" validate:"required"`
}

// ChangePasswordRequest represent the http request data for change password
type VerifyOTPChangePasswordRequest struct {
	OTP         string `json:"otp" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
type ChangePasswordOTPRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

// PasswordChecker represent the http request data for checking the password score
type PasswordChecker struct {
	Password string `json:"password" validate:"required"`
}

// GetUserRequest represent the http request data for get user
type GetUserRequest struct {
	UserID uuid.UUID `url_param:"user_id" validate:"required,uuid4"`
}

// GetUserByPhoneRequest represent the http request data for get user by phone
type GetUserByPhoneRequest struct {
	Phone string `url_param:"phone" validate:"required"`
}

type FilterUserRequest struct {
	Specialization string `json:"specialization"`
	Location       string `json:"location"`
}

type BlockUserRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
}

type DeleteUserRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
}

type VerifyOTPUpdateEmailRequest struct {
	OTP   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type VerifyOTPUpdatePhoneNumberRequest struct {
	OTP         string `json:"otp" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type ResetPasswordOTPRequest struct {
	Identifier string `json:"identifier" validate:"required"`
}
type VerifyOTPResetPasswordRequest struct {
	Identifier  string `json:"identifier" validate:"required"`
	OTP         string `json:"otp" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type UpdatePhoneNumberOTPRequest struct {
	NewPhoneNumber string `json:"new_phone_number" validate:"required"`
}

type UpdateEmailOTPRequest struct {
	Email string `json:"new_email" validate:"required"`
}

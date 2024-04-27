package types

// GetEmailVerificationRequest
type GetEmailVerificationRequest struct {
	Event    string `json:"event" binding:"required"`
	UserID   uint   `json:"userID" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// EmailVerifyRequest
type EmailVerifyRequest struct {
	Event            string `json:"event" binding:"required"`
	Email            string `json:"email" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}

// enum for the verification event, register and reset password
const (
	RegisterEvent      = "register"
	ResetPasswordEvent = "reset-password"
)

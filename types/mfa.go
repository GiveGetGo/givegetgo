package types

// MFAVerificationRequest
type MFAVerificationRequest struct {
	Email string `json:"email" binding:"required"`
}

// VerifyMFARequest is the request for verifying MFA
type VerifyMFARequest struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}

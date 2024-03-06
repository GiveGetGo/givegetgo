package types

// MFAVerificationRequest
type MFAVerificationRequest struct {
	UserID uint   `json:"userID" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

// VerifyMFARequest is the request for verifying MFA
type VerifyMFARequest struct {
	UserID           uint   `json:"userID" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}

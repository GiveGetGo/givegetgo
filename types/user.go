package types

// RegisterRequest - request body for the register endpoint
type RegisterRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest - request body for the login endpoint
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SetUserEmailVerifiedRequest - request body for setting user email to verified
type SetUserEmailVerifiedRequest struct {
	Email string `json:"email" binding:"required"`
}

// UserForgetPassRequest is the request body for the user forget password endpoint
type UserForgetPassRequest struct {
	Email string `json:"email" binding:"required"`
}

package types

// RegisterRequest - request body for the register endpoint
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Class    string `json:"class" binding:"required"`
	Major    string `json:"major" binding:"required"`
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

type UserResetPassRequest struct {
	Email       string `json:"email" binding:"required"`
	Newpassword string `json:"newpassword" binding:"required"`
}

// PostRequest - request body for the post endpoint
type PostRequest struct {
	//thinking UserID may not be neccessary when request a user to post something
	UserID      int64  `json:"userID" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

package types

// PostRequest - request body for the post endpoint
type PostRequest struct {
	//thinking UserID may not be neccessary when request a user to post something
	UserID      uint   `json:"userID" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

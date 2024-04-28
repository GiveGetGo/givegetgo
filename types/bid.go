package types

// BidRequest - request body for the bid endpoint
type BidRequest struct {
	PostID      uint   `json:"postid" binding:"required"`
    Description string `json:"description" binding:"required"`
}

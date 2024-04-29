package types

// BidRequest - request body for the bid endpoint
type BidRequest struct {
	Description string `json:"description" binding:"required"`
}

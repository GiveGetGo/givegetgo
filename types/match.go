package types

type MatchRequest struct {
	PostID uint `json:"postID" binding:"required"`
	BidID  uint `json:"bidID" binding:"required"`
}

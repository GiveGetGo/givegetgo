package types

type NotificationType string

const (
	NewMatch NotificationType = "newmatch"
	BidMatch NotificationType = "bidmatch"
	NewBid   NotificationType = "newbid"
)

type CreateNotificationRequest struct {
	UserID           uint             `json:"userID" binding:"required"`
	Description      string           `json:"description" binding:"required"`
	NotificationType NotificationType `json:"notificationType" binding:"required"`
}

package models

type Notification struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"created_at"`
}

type SendNotificationRequest struct {
	UserID  uint   `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

package model

import (
	"time"
)

// NotificationRequest defines the input data structure for a notification request
type NotificationRequest struct {
	User             string    `json:"user" `
	Type             string    `json:"type" `
	LastNotification time.Time `json:"last_notification" `
}

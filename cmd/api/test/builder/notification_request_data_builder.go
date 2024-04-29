package builder

import (
	"time"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
)

type NotificationRequestBuilder struct {
	user             string
	notificationType string
	lastNotification time.Time
}

func NewNotificationRequestBuilder() *NotificationRequestBuilder {
	return &NotificationRequestBuilder{}
}

func (b *NotificationRequestBuilder) WithUser(user string) *NotificationRequestBuilder {
	b.user = user
	return b
}

func (b *NotificationRequestBuilder) WithType(notificationType string) *NotificationRequestBuilder {
	b.notificationType = notificationType
	return b
}

func (b *NotificationRequestBuilder) WithLastNotification(lastNotification time.Time) *NotificationRequestBuilder {
	b.lastNotification = lastNotification
	return b
}

func (b *NotificationRequestBuilder) Build() model.NotificationRequest {
	return model.NotificationRequest{
		User:             b.user,
		Type:             b.notificationType,
		LastNotification: b.lastNotification,
	}
}

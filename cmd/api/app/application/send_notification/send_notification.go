package send_notification

import (
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/service/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/apierrors"
)

// SendNotificationApplication is the initial flow entry to check the rate limit
type SendNotificationApplication interface {
	// Handler is the input for access to send one send_notification
	Delegate(notificationRequest model.NotificationRequest) apierrors.ApiError
}
type SendNotification struct {
	SendNotificationService send_notification.SendNotificationService
}

func (sendNotification *SendNotification) Delegate(notificationRequest model.NotificationRequest) apierrors.ApiError {
	return sendNotification.SendNotificationService.SendNotification(notificationRequest)
}

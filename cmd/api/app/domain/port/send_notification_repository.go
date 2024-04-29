package port

import "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"

// SendNotificationRepository  use for all transactions about beer
type SendNotificationRepository interface {
	//SaveTimeNotification persist the Notification data
	SaveTimeNotification(notificationRequest model.NotificationRequest) (err error)
}

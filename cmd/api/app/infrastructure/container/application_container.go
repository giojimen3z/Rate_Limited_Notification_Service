package container

import (
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/application/send_notification"
)

func getSendNotificationApplication() send_notification.SendNotificationApplication {
	return &send_notification.SendNotification{SendNotificationService: getSendNotificationService()}
}

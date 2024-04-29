package container

import (
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/service/send_notification"
)

func getSendNotificationService() send_notification.SendNotificationService {
	return &send_notification.SendNotification{
		SendNotificationRepository: getSendNotificationRepository(),
	}
}

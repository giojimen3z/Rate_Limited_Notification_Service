package container

import (
	"database/sql"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/port"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/adapter/repository"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/config"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/controller/send_notification"
)

func GetSendNotificationController() *send_notification.SendNotificationController {
	return &send_notification.SendNotificationController{SendNotificationApplication: getSendNotificationApplication()}
}

func getSendNotificationRepository() port.SendNotificationRepository {
	return &repository.SendNotificationRepository{
		WriteClient:          getWriteConnectionClient(),
		ReadConnectionClient: getReadConnectionClient(),
	}
}

func getWriteConnectionClient() *sql.DB {
	conn, _ := config.GetWriteConnection()
	return conn
}

func getReadConnectionClient() *sql.DB {
	conn, _ := config.GetReadConnection()
	return conn
}

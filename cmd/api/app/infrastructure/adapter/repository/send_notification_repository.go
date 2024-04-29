package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/logger"
)

const (
	loggerErrorScanningFields = "error when scanning the information [Class: SendNotificationRepository][Method:%s]"
)

// SendNotificationRepository represent the mysql repository
type SendNotificationRepository struct {
	WriteClient          *sql.DB
	ReadConnectionClient *sql.DB
}

func (beerMysqlRepository *SendNotificationRepository) SaveTimeNotification(notificationRequest model.NotificationRequest) (err error) {

	status := false
	for i := 0; i < 10; i++ {

		if fmt.Sprintf("user_%v", 1) == notificationRequest.User {
			status = true
		}

	}

	if !status {
		err = errors.New("error saving new notification information")
		logger.Error(fmt.Sprintf(loggerErrorScanningFields, "SaveTimeNotification"), err)
	}

	return err
}

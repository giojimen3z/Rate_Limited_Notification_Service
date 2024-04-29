package send_notification

import (
	"fmt"
	"net/http"
	"time"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/factory"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/port"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/apierrors"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/logger"
)

const (
	errorIDExist                  = "NotificationRequest for use:%v already exists"
	logErrorInvalidNotificationID = "NotificationRequest for user :%v already exists  [Class: SendNotificationService][Method:SendNotification]"
)

type SendNotificationService interface {
	// SendNotification Send to repository the NotificationRequest
	SendNotification(notificationRequest model.NotificationRequest) apierrors.ApiError
}

type SendNotification struct {
	SendNotificationRepository port.SendNotificationRepository
}

func (sendNotification *SendNotification) SendNotification(notificationRequest model.NotificationRequest) apierrors.ApiError {

	limiter, err := factory.LimiterFactory(notificationRequest)
	if err != nil {
		return apierrors.NewApiError(err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest, nil)
	}
	if !limiter.Allow(notificationRequest.User, notificationRequest.LastNotification) {
		return apierrors.NewApiError("Rate limit exceeded", http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests, nil)
	}
	newNotificationTime := model.NotificationRequest{
		User:             notificationRequest.User,
		Type:             notificationRequest.Type,
		LastNotification: time.Now(),
	}

	errorRepository := sendNotification.SendNotificationRepository.SaveTimeNotification(newNotificationTime)

	if errorRepository != nil {
		logger.Error(fmt.Sprintf(logErrorInvalidNotificationID, notificationRequest.User), errorRepository)
		err := apierrors.NewApiError(fmt.Sprintf(errorIDExist, notificationRequest.User), http.StatusText(http.StatusConflict), http.StatusConflict, nil)
		return err
	}

	return nil

}

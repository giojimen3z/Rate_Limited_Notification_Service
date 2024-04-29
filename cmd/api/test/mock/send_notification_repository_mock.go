package mock

import (
	"time"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type SendNotificationRepositoryMock struct {
	mock.Mock
}

// Custom matcher para la validación de NotificationRequest con flexibilidad en el tiempo
func matchNotificationRequest(expected model.NotificationRequest, delta time.Duration) interface{} {
	return mock.MatchedBy(func(req model.NotificationRequest) bool {
		// Verificar user y type
		if req.User != expected.User || req.Type != expected.Type {
			return false
		}
		// Verificar tiempo dentro de un rango permitido
		lowerBound := expected.LastNotification.Add(-delta)
		upperBound := expected.LastNotification.Add(delta)
		return req.LastNotification.After(lowerBound) && req.LastNotification.Before(upperBound)
	})
}

func (mock *SendNotificationRepositoryMock) SaveTimeNotification(notificationModel model.NotificationRequest) (err error) {
	args := mock.Called(matchNotificationRequest(notificationModel, 5*time.Minute))
	return args.Error(0)
}

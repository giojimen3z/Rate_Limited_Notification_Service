package factory

import (
	"fmt"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	strategy2 "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/strategy"
)

func LimiterFactory(notificationRequest model.NotificationRequest) (strategy2.RateLimiterStrategy, error) {
	switch notificationRequest.Type {
	case "status":
		return strategy2.NewStatusNotificationLimiter(), nil
	case "news":
		return strategy2.NewNewsNotificationLimiter(), nil
	case "marketing":
		return strategy2.NewMarketingNotificationLimiter(), nil
	default:
		return nil, fmt.Errorf("unsupported notification type: %s", notificationRequest.Type)
	}
}

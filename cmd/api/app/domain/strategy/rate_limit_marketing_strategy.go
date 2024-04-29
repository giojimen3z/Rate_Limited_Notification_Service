package strategy

import (
	"time"

	"golang.org/x/time/rate"
)

type MarketingNotificationLimiter struct {
	limiter *rate.Limiter
}

func NewMarketingNotificationLimiter() *MarketingNotificationLimiter {
	return &MarketingNotificationLimiter{
		limiter: rate.NewLimiter(3, 3), // 3 solicitudes por hora
	}
}

func (m *MarketingNotificationLimiter) Allow(user string, lastNotification time.Time) bool {
	return time.Since(lastNotification) < time.Hour || m.limiter.Allow()
}

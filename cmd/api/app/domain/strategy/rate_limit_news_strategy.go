package strategy

import (
	"time"

	"golang.org/x/time/rate"
)

type NewsNotificationLimiter struct {
	limiter *rate.Limiter
}

func NewNewsNotificationLimiter() *NewsNotificationLimiter {
	return &NewsNotificationLimiter{
		limiter: rate.NewLimiter(rate.Every(24*time.Hour), 1), // 1 solicitud por día
	}
}

func (n *NewsNotificationLimiter) Allow(user string, lastNotification time.Time) bool {
	return time.Since(lastNotification) < 24*time.Hour || n.limiter.Allow()
}

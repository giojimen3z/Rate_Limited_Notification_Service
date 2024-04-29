package strategy

import (
	"time"

	"golang.org/x/time/rate"
)

type RateLimiterStrategy interface {
	Allow(user string, lastNotification time.Time) bool
}

type StatusNotificationLimiter struct {
	limiter *rate.Limiter
}

func NewStatusNotificationLimiter() *StatusNotificationLimiter {
	return &StatusNotificationLimiter{
		limiter: rate.NewLimiter(2, 2), // 2 solicitudes por minuto
	}
}

func (s *StatusNotificationLimiter) Allow(user string, lastNotification time.Time) bool {
	return time.Since(lastNotification) < time.Minute || s.limiter.Allow()
}

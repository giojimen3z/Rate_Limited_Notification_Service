package app

import (
	"os"

	"github.com/Rate_Limited_Notification_Service/pkg/logger"
	"github.com/Rate_Limited_Notification_Service/pkg/routerhandlers"
)

func StartApp() {
	router := routerhandlers.DefaultRouter()

	MapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}

package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/config"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/container"
	controller "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/controller/health"
)

func MapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping", controller.PingController.Ping)
	prefix := fmt.Sprintf("%s/Rate_Limited_Notification_Service/", prefixScope)

	baseUrl := router.Group(prefix)
	notification := baseUrl.Group("/notification")

	notification.POST("sendNotification", container.GetSendNotificationController().MakeSendNotification)

}

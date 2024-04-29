package send_notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/application/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/apierrors"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/logger"
)

var (
	invalidBodyErr = apierrors.NewBadRequestApiError("invalid request")
	successMassage = "the Notification  was sent successfully"
)

// SendNotificationController  used for inject the use case
type SendNotificationController struct {
	SendNotificationApplication send_notification.SendNotificationApplication
}

func (sendNotificationController *SendNotificationController) MakeSendNotification(context *gin.Context) {

	req := model.NotificationRequest{}
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(invalidBodyErr.Status(), invalidBodyErr)
		return
	}

	err := sendNotificationController.SendNotificationApplication.Delegate(req)

	if err != nil {
		logger.Error(err.Message(), err)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, successMassage)

}

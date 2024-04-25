package rate_notification

import (
	"fmt"
	"net/http"

	"github.com/Rate_Limited_Notification_Service/cmd/api/app/application/rate_notification"
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/Rate_Limited_Notification_Service/pkg/apierrors"
	"github.com/Rate_Limited_Notification_Service/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	invalidBodyErr = apierrors.NewBadRequestApiError("invalid request")
	successMassage = "the rate_notification %s was created successfully"
)

// CreateBeerController  used for inject the use case
type CreateBeerController struct {
	CreateBeerApplication rate_notification.CreateBeerApplication
}

func (createBeerController *CreateBeerController) MakeCreateBeer(context *gin.Context) {

	beer := model.Beer{}

	if err := context.ShouldBindJSON(&beer); err != nil {
		context.JSON(invalidBodyErr.Status(), invalidBodyErr)
		return
	}

	err := createBeerController.CreateBeerApplication.Handler(beer)

	if err != nil {
		logger.Error(err.Message(), err)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf(successMassage, beer.Name))

}

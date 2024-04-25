package rate_notification

import (
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/service/beer"
	"github.com/Rate_Limited_Notification_Service/pkg/apierrors"
)

// CreateBeerApplication is the initial flow entry to create one rate_notification
type CreateBeerApplication interface {
	// Handler is the input for access to create one rate_notification
	Handler(beer model.Beer) apierrors.ApiError
}
type CreateBeer struct {
	CreateBeerService beer.CreateBeerService
}

func (createBeer *CreateBeer) Handler(beer model.Beer) apierrors.ApiError {
	return createBeer.CreateBeerService.CreateBeer(beer)
}

package container

import (
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/application/rate_notification"
)

func getCreateBeerApplication() rate_notification.CreateBeerApplication {
	return &rate_notification.CreateBeer{CreateBeerService: getCreateBeerService()}
}

func getListBeerApplication() rate_notification.ListBeerApplication {
	return &rate_notification.ListBeer{ListBeerService: getListBeerService()}
}

func getBeerApplication() rate_notification.GetBeerApplication {
	return &rate_notification.GetBeer{GetBeerService: getBeerService()}
}
func getBeerBoxPriceApplication() rate_notification.GetBeerBoxPriceApplication {
	return &rate_notification.GetBeerBoxPrice{
		GetBeerService:         getBeerService(),
		ConvertCurrencyService: getConvertCurrencyService(),
		GetBeerBoxPriceService: getBeerBoxPriceService(),
	}
}

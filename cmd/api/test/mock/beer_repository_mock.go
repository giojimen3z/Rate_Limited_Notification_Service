package mock

import (
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type BeerRepositoryMock struct {
	mock.Mock
}

func (mock *BeerRepositoryMock) Save(beer model.Beer) (err error) {
	args := mock.Called(beer)
	return args.Error(0)
}

func (mock *BeerRepositoryMock) ListBeer() ([]model.Beer, error) {
	args := mock.Called()
	return args.Get(0).([]model.Beer), args.Error(1)
}

func (mock *BeerRepositoryMock) GetBeerByID(id int64) (beer model.Beer, err error) {
	args := mock.Called(id)
	return args.Get(0).(model.Beer), args.Error(1)
}

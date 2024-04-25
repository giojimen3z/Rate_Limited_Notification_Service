package rate_notification_test

import (
	"errors"
	"os"

	"github.com/Rate_Limited_Notification_Service/cmd/api/app/application/rate_notification"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/builder"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/mock"
)

var _ = Describe("Service", func() {
	Context("Create Beer", func() {
		var (
			repositoryMock    *mock.BeerRepositoryMock
			beerCreateService rate_notification.CreateBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			beerCreateService = rate_notification.CreateBeer{
				BeerRepository: repositoryMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid rate_notification request is received", func() {
			It("should return nil error", func() {

				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(nil)

				err := beerCreateService.CreateBeer(beer)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid rate_notification request is received with a invalid id", func() {
			It("should return error", func() {

				errorMock := errors.New("the id:1 is invalid")
				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(errorMock)
				errorExpected := "Message: Beer id:1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := beerCreateService.CreateBeer(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})

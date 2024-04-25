package rate_notification_test

import (
	"errors"

	beerApplicaciotn "github.com/Rate_Limited_Notification_Service/cmd/api/app/application/rate_notification"
	beerService "github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/service/rate_notification"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/builder"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/mock"
)

var _ = Describe("Handler", func() {
	Context("Create Beer", func() {
		var (
			repositoryMock    *mock.BeerRepositoryMock
			beerCreateUseCase beerApplicaciotn.CreateBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			beerCreateService := &beerService.CreateBeer{
				BeerRepository: repositoryMock,
			}
			beerCreateUseCase = beerApplicaciotn.CreateBeer{
				CreateBeerService: beerCreateService,
			}

		})

		When("a new valid rate_notification request is received", func() {
			It("should return nil error", func() {

				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(nil)

				err := beerCreateUseCase.Handler(beer)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid rate_notification request is received", func() {
			It("should return error", func() {

				errorMock := errors.New("Error 1062: Duplicate entry '1' for key 'rate_notification.PRIMARY'")
				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(errorMock)
				errorExpected := "Message: Beer id:1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := beerCreateUseCase.Handler(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})

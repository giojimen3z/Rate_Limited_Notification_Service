package rate_notification_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	beerApplication "github.com/Rate_Limited_Notification_Service/cmd/api/app/application/rate_notification"
	"github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	beerService "github.com/Rate_Limited_Notification_Service/cmd/api/app/domain/service/rate_notification"
	beerController "github.com/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/controller/rate_notification"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/builder"
	"github.com/Rate_Limited_Notification_Service/cmd/api/test/mock"
	"github.com/gin-gonic/gin"
)

var _ = Describe("Beer Controller", func() {
	Context("Create Beer", func() {
		var (
			beer                 model.Beer
			beerCreateController beerController.CreateBeerController
			context              *gin.Context
			repositoryMock       *mock.BeerRepositoryMock
			recorder             *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			beer = builder.NewBeerDataBuilder().Build()
			recorder = httptest.NewRecorder()
			context, _ = gin.CreateTestContext(recorder)
			repositoryMock = new(mock.BeerRepositoryMock)
			beerCreateService := &beerService.CreateBeer{
				BeerRepository: repositoryMock,
			}
			beerCreateUseCase := &beerApplication.CreateBeer{
				CreateBeerService: beerCreateService,
			}
			beerCreateController = beerController.CreateBeerController{
				CreateBeerApplication: beerCreateUseCase,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return 201 code", func() {
				body, _ := json.Marshal(beer)
				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string(body)))
				repositoryMock.On("Save", beer).Return(nil)
				expectMessage := "\"the rate_notification Golden was created successfully\""

				beerCreateController.MakeCreateBeer(context)

				Expect(http.StatusCreated).To(Equal(recorder.Code))
				Expect(expectMessage).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return 400 code", func() {

				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string("")))
				errorExpected := "{\"message\":\"invalid request\",\"error\":\"bad_request\",\"status\":400,\"cause\":[]}"

				beerCreateController.MakeCreateBeer(context)

				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new valid request is received but with  invalid id", func() {
			It("should return 409 code", func() {
				body, _ := json.Marshal(beer)
				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string(body)))
				errorMock := errors.New("the id:1 is invalid")
				repositoryMock.On("Save", beer).Return(errorMock)
				errorExpected := "{\"message\":\"Beer id:1 already exists\",\"error\":\"Conflict\",\"status\":409,\"cause\":null}"

				beerCreateController.MakeCreateBeer(context)

				Expect(http.StatusConflict).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

	})
})

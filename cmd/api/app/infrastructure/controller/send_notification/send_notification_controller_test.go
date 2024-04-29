package send_notification_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	sendNotificationApplication "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/application/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	sendNotificationService "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/service/send_notification"
	Controller "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/controller/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/builder"
	mockRepo "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("NotificationRequest Controller", func() {
	Context("Create NotificationRequest", func() {
		var (
			notificationRequest        model.NotificationRequest
			sendNotificationController Controller.SendNotificationController
			context                    *gin.Context
			repositoryMock             *mockRepo.SendNotificationRepositoryMock
			recorder                   *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			notificationRequest = builder.NewNotificationRequestBuilder().Build()
			recorder = httptest.NewRecorder()
			context, _ = gin.CreateTestContext(recorder)
			repositoryMock = new(mockRepo.SendNotificationRepositoryMock)
			sendNotificationService := &sendNotificationService.SendNotification{
				SendNotificationRepository: repositoryMock,
			}
			sendNotificationUseCase := &sendNotificationApplication.SendNotification{
				SendNotificationService: sendNotificationService,
			}
			sendNotificationController = Controller.SendNotificationController{

				SendNotificationApplication: sendNotificationUseCase,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return 200 code for status type notification", func() {
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest = builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()
				body, _ := json.Marshal(notificationRequest)
				recorder := httptest.NewRecorder()
				context, _ = gin.CreateTestContext(recorder)
				context.Request, _ = http.NewRequest("POST", "/Rate_Limited_Notification_Service/notification/sendNotification", strings.NewReader(string(body)))
				context.Request.Header.Set("Content-Type", "application/json")
				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(nil)

				sendNotificationController.MakeSendNotification(context)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				expectMessage := "\"the Notification  was sent successfully\""
				Expect(recorder.Body.String()).Should(Equal(expectMessage))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new valid request is received", func() {
			It("should return 200 code for news type notification", func() {
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("marketing").
					WithLastNotification(fixedTime).
					Build()
				body, _ := json.Marshal(notificationRequest)
				recorder := httptest.NewRecorder()
				context, _ := gin.CreateTestContext(recorder)
				context.Request, _ = http.NewRequest("POST", "/Rate_Limited_Notification_Service/notification/sendNotification", strings.NewReader(string(body)))
				context.Request.Header.Set("Content-Type", "application/json")
				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(nil)

				sendNotificationController.MakeSendNotification(context)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				expectMessage := "\"the Notification  was sent successfully\""
				Expect(recorder.Body.String()).Should(Equal(expectMessage))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new valid request is received", func() {
			It("should return 200 code for news type notification", func() {
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("news").
					WithLastNotification(fixedTime).
					Build()
				body, _ := json.Marshal(notificationRequest)
				recorder := httptest.NewRecorder()
				context, _ := gin.CreateTestContext(recorder)
				context.Request, _ = http.NewRequest("POST", "/Rate_Limited_Notification_Service/notification/sendNotification", strings.NewReader(string(body)))
				context.Request.Header.Set("Content-Type", "application/json")
				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(nil)

				sendNotificationController.MakeSendNotification(context)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				expectMessage := "\"the Notification  was sent successfully\""
				Expect(recorder.Body.String()).Should(Equal(expectMessage))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new invalid request is received", func() {
			It("should return 400 code", func() {

				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("test").
					WithLastNotification(fixedTime).
					Build()
				body, _ := json.Marshal(notificationRequest)
				recorder := httptest.NewRecorder()
				context, _ := gin.CreateTestContext(recorder)
				context.Request, _ = http.NewRequest("POST", "/Rate_Limited_Notification_Service/notification/sendNotification", strings.NewReader(string(body)))
				context.Request.Header.Set("Content-Type", "application/json")

				sendNotificationController.MakeSendNotification(context)

				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
				expectMessage := "{\"message\":\"unsupported notification type: test\",\"error\":\"Bad Request\",\"status\":400,\"cause\":null}"
				Expect(recorder.Body.String()).Should(Equal(expectMessage))
			})
		})

	})
})

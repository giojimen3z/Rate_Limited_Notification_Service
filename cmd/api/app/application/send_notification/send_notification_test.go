package send_notification_test

import (
	"errors"
	"os"
	"time"

	sendNotificationApplication "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/application/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/model"
	sendNotificationService "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/service/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/builder"
	mockRepo "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Delegate", func() {
	Context("Create Beer", func() {
		var (
			notificationRequest     model.NotificationRequest
			repositoryMock          *mockRepo.SendNotificationRepositoryMock
			sendNotificationUseCase sendNotificationApplication.SendNotificationApplication
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			notificationRequest = builder.NewNotificationRequestBuilder().Build()
			repositoryMock = new(mockRepo.SendNotificationRepositoryMock)
			sendNotificationService := &sendNotificationService.SendNotification{
				SendNotificationRepository: repositoryMock,
			}
			sendNotificationUseCase = &sendNotificationApplication.SendNotification{
				SendNotificationService: sendNotificationService,
			}

		})

		When("a new valid send_notification request is received", func() {
			It("should return nil error", func() {
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest = builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()
				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(nil)

				err := sendNotificationUseCase.Delegate(notificationRequest)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid send_notification request is received", func() {
			It("should return error", func() {
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest = builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()
				errorMock := errors.New("Error 1062: Duplicate entry '1' for key 'beer.PRIMARY'")
				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(errorMock)
				errorExpected := "Message: NotificationRequest for use:user_1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := sendNotificationUseCase.Delegate(notificationRequest)

				println(err.Error())

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

	})
})

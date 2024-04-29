package send_notification_test

import (
	"errors"
	"os"
	"time"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/domain/service/send_notification"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/builder"
	mockRepo "github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Service", func() {
	Context("Create Beer", func() {
		var (
			repositoryMock          *mockRepo.SendNotificationRepositoryMock
			sendNotificationService send_notification.SendNotification
		)
		BeforeEach(func() {
			repositoryMock = new(mockRepo.SendNotificationRepositoryMock)
			sendNotificationService = send_notification.SendNotification{
				SendNotificationRepository: repositoryMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid send_notification request is received", func() {
			It("should return nil error", func() {

				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()

				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(nil)

				err := sendNotificationService.SendNotification(notificationRequest)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid send_notification request is received with a invalid id", func() {
			It("should return error", func() {

				errorMock := errors.New("the id:1 is invalid")
				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()
				errorExpected := "Message: NotificationRequest for use:user_1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				repositoryMock.On("SaveTimeNotification", mock.Anything).Return(errorMock)

				err := sendNotificationService.SendNotification(notificationRequest)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})

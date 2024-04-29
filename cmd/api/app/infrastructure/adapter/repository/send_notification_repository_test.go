package repository_test

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/app/infrastructure/adapter/repository"
	"github.com/giojimen3z/Rate_Limited_Notification_Service/cmd/api/test/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	insertQueryBeer = "INSERT INTO beer"
	beerListQuery   = "SELECT (.+) FROM beer"
)
const (
	beerId = 1
)

var _ = Describe("Repository", func() {
	Context("NotificationRequest Mysql Repository", func() {
		var (
			db                         *sql.DB
			sendNotificationRepository repository.SendNotificationRepository
		)
		BeforeEach(func() {
			db, _, _ = sqlmock.New()
			sendNotificationRepository = repository.SendNotificationRepository{
				WriteClient:          db,
				ReadConnectionClient: db,
			}

		})
		When("a new valid beer request is received  and save in dba", func() {
			It("should return nil error", func() {

				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_1").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()

				err := sendNotificationRepository.SaveTimeNotification(notificationRequest)

				Expect(err).Should(BeNil())
			})
		})
		When("a new valid beer request and failed the transaction", func() {
			It("should return  error", func() {

				fixedTime := time.Date(2024, time.April, 29, 2, 45, 0, 0, time.UTC)
				notificationRequest := builder.NewNotificationRequestBuilder().
					WithUser("user_55").
					WithType("status").
					WithLastNotification(fixedTime).
					Build()

				err := sendNotificationRepository.SaveTimeNotification(notificationRequest)

				Expect(err).Should(Not(BeNil()))

			})
		})

	})
})

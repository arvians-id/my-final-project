package test_test

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

var _ = Describe("Test", func() {
	var (
		usercourseRepository repository.UserCourseRepository
		database             *sql.DB
		err                  error
		ctx                  *gin.Context
	)

	BeforeEach(func() {
		database, err = sql.Open("sqlite3", "../teenager.db")

		if err != nil {
			panic(err)
		}

		userCourseRepository := repository.NewUserCourseRepository()
		userCourseService := service.NewUserCourseService(&userCourseRepository, database)
		controller.NewUserCourseController(&userCourseService)
	})

	usercoursecreate := model.CreateUserCourseRequest{
		UserId:   5,
		CourseId: 3,
	}

	usercourse := entity.UserCourse{
		UserId:   5,
		CourseId: 3,
	}
	tx, _ := database.Begin()

	Describe("usercase_test", func() {
		When("usercourse create", func() {
			It("should create a new usercourse", func() {

				responses, err := usercourseRepository.Create(ctx, tx, usercourse)
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercoursecreate))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})

			It("Should get a usercourse by userid", func() {
				responses, err := usercourseRepository.FindByUserId(ctx, tx, utils.ToString(usercourse.UserId))
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercoursecreate))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})

			It("Should get a list of usercourse", func() {
				responses, err := usercourseRepository.FindAll(ctx, tx)
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercoursecreate))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})

			It("Should delete a usercourse", func() {
				err := usercourseRepository.Delete(ctx, tx, usercourse.UserId, usercourse.CourseId)
				Expect(err).To(BeNil())
			})
		})
	})
})

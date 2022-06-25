package unit_test

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

var _ = Describe("Test", func() {

	var (
		userController controller.UserController
		userService    service.UserServiceImplement
		userRepository repository.UserRepository
		database       *sql.DB
		err            error
		ctx            *gin.Context
	)

	BeforeEach(func() {
		database, err = sql.Open("sqlite3", "../teenager.db")

		if err != nil {
			panic(err)
		}

		userRepository = repository.NewUserRepository()
		userService = service.NewUserService(&userRepository, database)
		userController = controller.NewUserController(&userService)
	})

	Describe("User Login Register", func() {
		When("Data is Correct", func() {
			It("Should return list User", func() {

				userLoginResponse, err := userController.UserService.ListUser(ctx)

				Expect(err).To(BeNil())
				Expect(userLoginResponse).To(HaveLen(2))
			})
		})
	})
})

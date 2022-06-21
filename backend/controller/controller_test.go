package controller_test

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

var _ = Describe("Controller", func() {

	var (
		userController controller.UserController
		userService    service.UserServiceImplement
		userRepository repository.UserRepository
		ctx            *gin.Context
	)

	BeforeEach(func() {
		database, err := sql.Open("sqlite3", "../teenager.db")

		if err != nil {
			fmt.Println(err)
		}

		userRepository = repository.NewUserRepository()
		userService = service.NewUserService(&userRepository, database)
		userController = controller.NewUserController(&userService)
	})

	Describe("NewController", func() {
		It("should not be nil", func() {
			result, err := userController.UserService.ListUser(ctx)
			fmt.Println(result)
			Expect(err).To(BeNil())
			Expect(result).To(BeNil())
		})
	})
})

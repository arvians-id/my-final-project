package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

func NewInitializedServer(configuration config.Config) *gin.Engine {
	// Configuration
	router := gin.Default()
	database := config.NewSQLite(configuration)

	// setup gin cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour,
	}))

	// Setup Proxies (optional)
	// You can comment this section
	err := router.SetTrustedProxies([]string{configuration.Get("APP_URL")})
	if err != nil {
		panic(err)
	}

	// Course Setup
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(&courseRepository, database)

	// Module Articles Setup
	moduleArticlesRepository := repository.NewModuleArticlesRepository()
	moduleArticlesService := service.NewModuleArticlesService(&moduleArticlesRepository, &courseRepository, database)
	moduleArticlesController := controller.NewModuleArticlesController(&moduleArticlesService)

	// Module Submission Setup
	moduleSubmissionRepository := repository.NewModuleSubmissionsRepository()

	// User Submission Setup
	userSubmissionRepository := repository.NewUserSubmissionsRepository()
	userSubmissionService := service.NewUserSubmissionsService(&userSubmissionRepository, &moduleSubmissionRepository, &courseRepository, database)
	userSubmissionController := controller.NewUserSubmissionsController(&userSubmissionService)

	// UserCourse Setup
	userCourseRepository := repository.NewUserCourseRepository()
	userCourseService := service.NewUserCourseService(&userCourseRepository, &courseRepository, &moduleSubmissionRepository, &userSubmissionRepository, database)
	userCourseController := controller.NewUserCourseController(&userCourseService)

	// ---  Module Submission Setup
	moduleSubmissionService := service.NewModuleSubmissionsService(&moduleSubmissionRepository, &courseRepository, &userCourseRepository, &userSubmissionRepository, database)
	moduleSubmissionController := controller.NewModuleSubmissionsController(&moduleSubmissionService, &userCourseService)
	// ---  Course Setup
	courseController := controller.NewCourseController(&courseService, &userCourseService)
	// ---  User Setup
	userRepository := repository.NewUserRepository()

	// Question Setup
	questionRepository := repository.NewQuestionRepository()
	questionService := service.NewQuestionService(&questionRepository, &userRepository, database)
	questionController := controller.NewQuestionController(&questionService)

	// Answer Setup
	answerRepository := repository.NewAnswerRepository()
	answerService := service.NewAnswerService(&answerRepository, &questionRepository, database)
	answerController := controller.NewAnswerController(&answerService)

	// Email Verification Setup
	emailVerificationRepository := repository.NewEmailVerificationRepository()
	emailVerificationService := service.NewEmailService(&emailVerificationRepository, &userRepository, database)

	// User Setup
	userService := service.NewUserService(&userRepository, database, &emailVerificationRepository)
	userController := controller.NewUserController(&userService, &userCourseService, &emailVerificationService)

	// Routing
	userController.Route(router)
	courseController.Route(router)
	moduleArticlesController.Route(router)
	moduleSubmissionController.Route(router)
	userSubmissionController.Route(router)
	userCourseController.Route(router)
	questionController.Route(router)
	answerController.Route(router)

	return router
}

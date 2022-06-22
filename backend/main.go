package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/config"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

func main() {
	// Configuration
	configuration := config.New()
	router := gin.Default()
	database := config.NewSQLite(configuration)

	// Setup Proxies (optional)
	// You can comment this section
	err := router.SetTrustedProxies([]string{configuration.Get("APP_URL")})
	if err != nil {
		panic(err)
	}

	// User Setup
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	// Course Setup
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(&courseRepository, database)
	courseController := controller.NewCourseController(&courseService)

	// Module Articles Setup
	moduleArticlesRepository := repository.NewModuleArticlesRepository()
	moduleArticlesService := service.NewModuleArticlesService(&moduleArticlesRepository, &courseRepository, database)
	moduleArticlesController := controller.NewModuleArticlesController(&moduleArticlesService)

	// Module Submission Setup
	moduleSubmissionRepository := repository.NewModuleSubmissionsRepository()
	moduleSubmissionService := service.NewModuleSubmissionsService(&moduleSubmissionRepository, &courseRepository, database)
	moduleSubmissionController := controller.NewModuleSubmissionsController(&moduleSubmissionService)

	// User Submission Setup
	userSubmissionRepository := repository.NewUserSubmissionsRepository()
	userSubmissionService := service.NewUserSubmissionsService(&userSubmissionRepository, &moduleSubmissionRepository, &courseRepository, database)
	userSubmissionController := controller.NewUserSubmissionsController(&userSubmissionService)

	// Routing
	userController.Route(router)
	courseController.Route(router)
	moduleArticlesController.Route(router)
	moduleSubmissionController.Route(router)
	userSubmissionController.Route(router)

	// Run
	PORT := fmt.Sprintf(":%v", configuration.Get("APP_PORT"))
	teenager(PORT)

	err = router.Run(PORT)
	if err != nil {
		panic(err)
	}
}

func teenager(port string) {
	fmt.Print(`

	
┏━━━━┓
┃┏┓┏┓┃
┗┛┃┃┣┻━┳━━┳━┓┏━━┳━━┳━━┳━┓
╋╋┃┃┃┃━┫┃━┫┏┓┫┏┓┃┏┓┃┃━┫┏┛
╋╋┃┃┃┃━┫┃━┫┃┃┃┏┓┃┗┛┃┃━┫┃
╋╋┗┛┗━━┻━━┻┛┗┻┛┗┻━┓┣━━┻┛
╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋┏━┛┃
╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋┗━━┛
`)
}

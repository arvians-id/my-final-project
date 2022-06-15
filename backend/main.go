package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/config"
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

	// Module Submission Setup
	moduleSubmissionRepository := repository.NewModuleSubmissionsRepository()
	moduleSubmissionService := service.NewModuleSubmissionsService(&moduleSubmissionRepository, database)
	moduleSubmissionController := controller.NewModuleSubmissionsController(&moduleSubmissionService)

	// Routing
	userController.Route(router)
	courseController.Route(router)
	moduleSubmissionController.Route(router)

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

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

	// Routing
	userController.Route(router)
	courseController.Route(router)

	// Run
	PORT := fmt.Sprintf(":%v", configuration.Get("PORT"))
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

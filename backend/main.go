package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/gin-contrib/cors"
)

func main() {
	// Configuration
	configuration := config.New()
	router := gin.Default()
	database := config.NewSQLite(configuration)

	// setup gin cors
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET","POST","PUT","DELETE","OPTIONS","PATCH"},
    AllowHeaders:     []string{"Authorization", "Content-Type"},
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
	
	// User Setup
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	// Course Setup
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(&courseRepository, database)
	courseController := controller.NewCourseController(&courseService)

	// Question Setup
	questionRepository := repository.NewQuestionRepository()
	questionService := service.NewQuestionService(&questionRepository, database)
	questionController := controller.NewQuestionController(&questionService)

	// Answer Setup
	answerRepository := repository.NewAnswerRepository()
	answerService := service.NewAnswerService(&answerRepository, &questionRepository, database)
	answerController := controller.NewAnswerController(&answerService)

	// Routing
	userController.Route(router)
	courseController.Route(router)
	questionController.Route(router)
	answerController.Route(router)

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

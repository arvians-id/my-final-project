package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

func main() {
	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}
	const PORT = ":8080"
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	routing := userController.Route()
	teenager(PORT)
	routing.Run(PORT)
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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

type UserController struct {
	UserService service.UserServiceImplement
}

func NewUserController(userService *service.UserServiceImplement) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route() *gin.Engine {
	router := gin.Default()
	router.GET("/api/users", controller.getUser)
	router.POST("/api/users", controller.userRegister)
	return router
}

func (controller *UserController) getUser(ctx *gin.Context) {
	responses, err := controller.UserService.GetAllUser()

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get All User Successfull",
		Data:   responses,
	})
}

func (controller *UserController) userRegister(ctx *gin.Context) {
	var user model.UserRegister

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	responses, err := controller.UserService.RegisterUser(user)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, model.WebResponse{
		Code:   201,
		Status: "User Register Succesfully",
		Data:   responses,
	})
}

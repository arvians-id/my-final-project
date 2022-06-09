package controller

import (
	"net/http"
	"strconv"

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
	router.GET("/api/users/:id", controller.getUserByID)
	router.DELETE("/api/users/:id", controller.deleteUser)
	router.PUT("/api/users/:id", controller.updateUser)
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

func (controller *UserController) deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	err = controller.UserService.DeleteUser(id)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":   200,
		"Status": "Delete User Successfull",
	})
}

func (controller *UserController) updateUser(ctx *gin.Context) {
	var user model.UserRegister

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	responses, err := controller.UserService.UpdateUser(id, user)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":   200,
		"Status": "Update User Successfull",
		"Data":   responses,
	})
}

func (controller *UserController) getUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	response, err := controller.UserService.GetUserbyID(id)

	if err != nil {
		return
	}

	if response.Name == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get User By ID Successfull",
		Data:   response,
	})
}

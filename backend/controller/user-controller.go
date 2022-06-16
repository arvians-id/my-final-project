package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/middleware"
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

func (controller *UserController) Route(router *gin.Engine) *gin.Engine {

	// router.Use(func(c *gin.Context) {
	// 	c.Header("Content-Type", "application/json")
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
	// 	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// })

	api := router.Group("/api")
	{
		api.POST("/users", controller.userRegister)
		api.POST("/users/login", controller.userLogin)
		api.GET("/userstatus", middleware.UserHandler(controller.userStatus))
		api.POST("/users/logout", middleware.UserHandler(controller.userLogout))
		api.GET("/users/:id", middleware.UserHandler(controller.getUserByID))
		api.GET("/users", middleware.AdminHandler(controller.listUser))
		api.PUT("/users/:id", middleware.UserHandler(controller.updateUser))
		api.DELETE("/users/:id", middleware.AdminHandler(controller.deleteUser))
	}
	return router
}

//Function to register new user
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

//Function to login user
func (controller *UserController) userLogin(ctx *gin.Context) {
	var user model.UserRegister

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	response, err := controller.UserService.UserLogin(user.Email, user.Password)

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

	token := service.JWTAuthService().GenerateToken(entity.Users{
		Id:   response.Id,
		Name: response.Name,
		Role: response.Role,
	})

	ctx.Header("Authorization", token)

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Login Successfull",
		Token:  token,
		Data:   response,
	})
}

//Function to get user status
func (controller *UserController) userStatus(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})
		return
	}

	if ok := service.JWTAuthService().CheckToken(token); ok != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Invalid Token",
		})
		return
	}

	tokenClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your secret api key"), nil
	},
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Cannot parse token",
			Data:   "",
		})
		return
	}

	id := tokenClaims["id"].(float64)

	user, err := controller.UserService.GetUserbyID(int(id))

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "User Already Logged In",
		Data:   user,
	})
}

//Function to logout user
func (controller *UserController) userLogout(ctx *gin.Context) {
	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
		})
		return
	}

	if ok := service.JWTAuthService().CheckToken(token); ok != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})
		return
	}

	err := service.JWTAuthService().DeleteToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Logout Successfull",
	})
}

//Function to get user by id
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

//Function to show list user
func (controller *UserController) listUser(ctx *gin.Context) {
	responses, err := controller.UserService.ListUser()

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get All User Successfull",
		Data:   responses,
	})
}

//Function to update user
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

//Function to delete user
func (controller *UserController) deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	err = controller.UserService.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":   200,
		"Status": "Delete User Successfull",
	})
}

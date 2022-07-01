package controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/middleware"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

type UserController struct {
	UserService       service.UserServiceImplement
	UserCourseService service.UserCourseService
	EmailService      service.EmailService
}

func NewUserController(userService *service.UserServiceImplement, userCourseService *service.UserCourseService, emailService *service.EmailService) UserController {
	return UserController{
		UserService:       *userService,
		UserCourseService: *userCourseService,
		EmailService:      *emailService,
	}
}

func (controller *UserController) Route(router *gin.Engine) *gin.Engine {

	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
	})

	api := router.Group("/api")
	{
		api.POST("/users", controller.UserRegister)                                                // done
		api.POST("/users/login", controller.userLogin)                                             // done
		api.GET("/userstatus", middleware.UserHandler(controller.userStatus))                      // done
		api.POST("/users/logout", middleware.UserHandler(controller.userLogout))                   // done
		api.PUT("/users/roleupdate/:id/:role", middleware.AdminHandler(controller.userRoleUpdate)) // done
		api.GET("/users/:id", middleware.UserHandler(controller.getUserByID))                      // done
		api.GET("/users", middleware.AdminHandler(controller.listUser))                            // done
		api.PUT("/users/:id", middleware.UserHandler(controller.updateUser))                       // done
		api.DELETE("/users/:id", middleware.AdminHandler(controller.deleteUser))
		api.GET("/users/submissions", middleware.UserHandler(controller.StudentSubmission))
		api.GET("/users/verify", controller.VerifyEmail)
	}
	return router
}

//Function to register new user
func (controller *UserController) UserRegister(ctx *gin.Context) {
	var user model.UserRegisterResponse

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	// Data email verification
	timestamp := time.Now().Add(1 * time.Hour).Unix()
	timestampString := strconv.Itoa(int(timestamp))
	signature := md5.Sum([]byte(user.Email + timestampString))
	signatureString := fmt.Sprintf("%x", signature)

	responses, err := controller.UserService.RegisterUser(ctx, user, signatureString, int(timestamp))
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: err.Error(),
		})
		return
	}

	// Send email to user
	message := fmt.Sprintf("visite this link to verification your email : http://localhost:8080/verification?email=%v&signature=%v", user.Email, signatureString)
	err = controller.EmailService.SendEmailWithText(user.Email, message)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, model.WebResponse{
			Code:   400,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, model.WebResponse{
		Code:   201,
		Status: "User Register Successfully",
		Data:   responses,
	})
}

//Function to login user
func (controller *UserController) userLogin(ctx *gin.Context) {
	var user model.GetUserLogin

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Please Check Your Input",
		})
		return
	}

	response, err := controller.UserService.UserLogin(ctx, user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Please Check Your Input",
		})
		return
	}

	if response.Name == "" {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   400,
			Status: "User Not Found",
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
		})
		return
	}

	id := tokenClaims["id"].(float64)

	user, err := controller.UserService.GetUserbyID(ctx, int(id))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Cannot get user",
		})
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

//Function to update user role
func (controller *UserController) userRoleUpdate(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	roleUpdate, err := strconv.Atoi(ctx.Param("role"))

	if err != nil {
		return
	}

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
	_, err = jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your secret api key"), nil
	},
	)

	iduser := tokenClaims["id"].(float64)
	role := tokenClaims["role"].(string)
	iduserint := int(iduser)

	if iduserint != id && role != "1" {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "You are not authorized to update role this user",
		})
		return
	}

	if err != nil {
		return
	}

	response, err := controller.UserService.UpdateUserRole(ctx, id, roleUpdate)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	if response.Name == "" {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   404,
			Status: "User Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Update User Role Successfull",
		Data:   response,
	})
}

//Function to get user by id
func (controller *UserController) getUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	response, err := controller.UserService.GetUserbyID(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

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
	_, err = jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your secret api key"), nil
	},
	)

	iduser := tokenClaims["id"].(float64)
	role := tokenClaims["role"].(string)
	iduserint := int(iduser)

	if iduserint != id && role != "1" {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "You are not authorized to view this user",
		})
		return
	}

	if err != nil {
		return
	}

	if response.Name == "" {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   404,
			Status: "User Not Found",
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
	responses, err := controller.UserService.ListUser(ctx)
	if err != nil {
		panic(err)
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get All User Successfull",
		Data:   responses,
	})
}

//Function to update user
func (controller *UserController) updateUser(ctx *gin.Context) {
	var user model.GetUserDetailUpdate

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

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
	_, err = jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your secret api key"), nil
	},
	)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	iduser := tokenClaims["id"].(float64)
	role := tokenClaims["role"].(string)
	iduserint := int(iduser)

	if iduserint != id && role != "1" {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "You are not authorized to update this user",
		})
		return
	}

	responses, err := controller.UserService.UpdateUser(ctx, id, user)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Update User Successfull",
		Data:   responses,
	})
}

//Function to delete user
func (controller *UserController) deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	err = controller.UserService.DeleteUser(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Delete User Successfull",
	})
}

func (controller *UserController) StudentSubmission(ctx *gin.Context) {
	limit := -1
	if ctx.Query("limit") != "" {
		limits, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: err.Error(),
				Data:   nil,
			})
			return
		}
		limit = limits
	}

	idUser, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "user not found",
			Data:   nil,
		})
		return
	}

	id := int(idUser.(float64))
	studentSubmissions, err := controller.UserCourseService.FindAllStudentSubmissions(ctx, id, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   studentSubmissions,
	})
}

func (controller *UserController) VerifyEmail(ctx *gin.Context) {
	var request model.GetEmailVerificationRequest
	signature := ctx.Query("signature")
	email := ctx.Query("email")

	request.Email = email
	request.Signature = signature
	err := controller.EmailService.VerifyEmail(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "email successfully verified",
		Data:   nil,
	})
}

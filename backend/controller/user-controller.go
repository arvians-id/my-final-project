package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
	})

	router.GET("/api/users", controller.getUser)
	router.POST("/api/users", controller.userRegister)
	router.GET("/api/users/:id", controller.getUserByID)
	router.DELETE("/api/users/:id", controller.deleteUser)
	router.PUT("/api/users/:id", controller.updateUser)
	router.POST("/api/users/login", controller.userLogin)
	router.GET("/api/user", controller.user)
	router.POST("/api/user/logout", controller.userLogout)
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

func (controller *UserController) userLogin(ctx *gin.Context) {
	var user model.UserRegister

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	response, err := controller.UserService.UserLogin(user.Email, user.Password)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(response.Id),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte("your secret api key"))

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"Status":  "Internal Server Error",
			"message": "Error Generate Token",
		})
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		MaxAge:   3600,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour * 1),
		Secure:   false,
		HttpOnly: true,
	}

	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Login Successfull",
		Data:   cookie,
	})
}

func (controller *UserController) user(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("jwt")
	var user model.UserRegister

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your secret api key"), nil
	},
	)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Id, _ = strconv.Atoi(claims["iss"].(string))

		response, err := controller.UserService.GetUserbyID(user.Id)

		if err != nil {
			return
		}

		ctx.Header("Accept", "application/json")
		ctx.Header("Content-Type", "application/json")

		ctx.JSON(http.StatusOK, model.WebResponse{
			Code:   200,
			Status: "Login Successfull",
			Data:   response,
		})
	}
}

func (controller *UserController) userLogout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("jwt")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})
	}

	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-1 * time.Hour)

	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Logout Successfull",
		Data:   "",
	})
}

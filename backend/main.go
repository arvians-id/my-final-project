package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
)

var users = []entity.User{
	{
		ID:            1,
		Name:          "Adi Syahputra",
		Unique_number: "2kjb5mjk23b2m",
		Phone:         "081234567890",
		Email:         "adisyahputra@gmail.com",
		Password:      "2kjbk2jkh24j2h42b42",
		Role:          2,
		Image:         "www.reezyx.com",
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	},
	{
		ID:            2,
		Name:          "Lucinta Luna",
		Unique_number: "435b2h6j342v12",
		Phone:         "080987654321",
		Email:         "lucintaluna@gmail.com",
		Password:      "2kjb423hj5g42b4j2",
		Role:          1,
		Image:         "www.llsebenarnyacowok.com",
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	},
}

func getUserbyID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("value"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	for _, user := range users {
		if user.ID == id {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.IndentedJSON(http.StatusOK, model.WebResponse{
				Code:   200,
				Status: "User Found",
				Data:   user,
			})
			return
		}
	}
}

func userRegister(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	users = append(users, user)

	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, model.WebResponse{
		Code:   201,
		Status: "User Register Succesfully",
		Data:   user,
	})
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	return
}

func getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get All User Successfull",
		Data:   users,
	})
}

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", hello)
	router.GET("/user/:value", getUserbyID)
	router.GET("/user", getUser)
	router.POST("/register", userRegister)
	return router
}

func main() {
	routing := getRouter()
	routing.Run(":8080")
}

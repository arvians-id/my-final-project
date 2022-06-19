package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/middleware"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type UserCourseController struct {
	UserCourseService service.UserCourseService
}

func NewUserCourseController(usercourseService *service.UserCourseService) UserCourseController {
	return UserCourseController{
		UserCourseService: *usercourseService,
	}
}

func (controller *UserCourseController) Route(router *gin.Engine) *gin.Engine {

	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
	})

	api := router.Group("/api")
	{
		api.POST("/usercourse", middleware.AdminHandler(controller.usercourseCreate))
		api.GET("/usercourse/:id", middleware.UserHandler(controller.getUserCourseByUserId))
		api.GET("/usercourse", middleware.AdminHandler(controller.listUserCourse))
		api.DELETE("/usercourse/:userid/:courseid", middleware.AdminHandler(controller.deleteUserCourse))
	}
	return router
}

func (controller *UserCourseController) usercourseCreate(ctx *gin.Context) {
	var usercourse model.CreateUserCourseRequest

	if err := ctx.BindJSON(&usercourse); err != nil {
		return
	}

	responses, err := controller.UserCourseService.Create(ctx, usercourse)

	if err != nil {
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, model.WebResponse{
		Code:   201,
		Status: "User Course Create Succesfully",
		Data:   responses,
	})
}

func (controller *UserCourseController) getUserCourseByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return
	}

	response, err := controller.UserCourseService.FindByUserId(ctx, utils.ToString(id))

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
			Data:   "You are not authorized to view this user",
		})
		return
	}

	if err != nil {
		return
	}

	if response.UserId < 0 {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   404,
			Status: "User Course Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get User By User Id Successfull",
		Data:   response,
	})
}

func (controller *UserCourseController) listUserCourse(ctx *gin.Context) {
	responses, err := controller.UserCourseService.FindAll(ctx)

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "Get All User Course Successfull",
		Data:   responses,
	})
}

func (controller *UserCourseController) deleteUserCourse(ctx *gin.Context) {
	userid, _ := strconv.Atoi(ctx.Param("userid"))
	courseid, err := strconv.Atoi(ctx.Param("courseid"))

	if err != nil {
		return
	}

	err = controller.UserCourseService.Delete(ctx, userid, courseid)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Course Not Found",
		})
		return
	}

	ctx.Header("Accept", "application/json")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":   200,
		"Status": "Delete User Course Successfull",
	})
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"github.com/rg-km/final-project-engineering-12/backend/middleware"	
)

type QuestionController struct {
	QuestionService service.QuestionService
}

func NewQuestionController(questionService *service.QuestionService) *QuestionController {
	return &QuestionController{
		QuestionService: *questionService,
	}
}

func (controller *QuestionController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/questions/all", middleware.UserHandler(controller.FindAll))
		authorized.POST("/questions/create", middleware.UserHandler(controller.Create))
		authorized.PUT("/questions/update/:questionId", middleware.UserHandler(controller.Update))
		authorized.DELETE("/questions/:questionId", middleware.UserHandler(controller.Delete))
		authorized.GET("/questions/by-user/:userId", middleware.UserHandler(controller.FindByUserId))
	}

	return router
}

func (controller *QuestionController) FindAll(ctx *gin.Context) {
	questions, err := controller.QuestionService.FindAll(ctx.Request.Context())
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
		Data:   questions,
	})
}

func (controller *QuestionController) Create(ctx *gin.Context) {
	var request model.CreateQuestionRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	question, err := controller.QuestionService.Create(ctx, request)
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
		Status: "question successfully created",
		Data:   question,
	})
}

func (controller *QuestionController) Delete(ctx *gin.Context) {
	questionId := utils.ToInt(ctx.Param("questionId"))

	err := controller.QuestionService.Delete(ctx.Request.Context(), questionId)
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
		Status: "question successfully deleted",
		Data:   nil,
	})
}



func (controller *QuestionController) Update(ctx *gin.Context) {
	var request model.UpdateQuestionRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	questionId := utils.ToInt(ctx.Param("questionId"))
	
	question, err := controller.QuestionService.Update(ctx, request, questionId)
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
		Status: "question successfully update",
		Data:   question,
	})
}

func (controller *QuestionController) FindByUserId(ctx *gin.Context) {
	userId := utils.ToInt(ctx.Param("userId"))
	questions, err := controller.QuestionService.FindByUserId(ctx.Request.Context(), userId)
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
		Data:   questions,
	})
}

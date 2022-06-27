package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/middleware"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"net/http"
)

type AnswerController struct {
	AnswerService service.AnswerService
}

func NewAnswerController(answerService *service.AnswerService) *AnswerController {
	return &AnswerController{
		AnswerService: *answerService,
	}
}

func (controller *AnswerController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/answers/all", middleware.UserHandler(controller.FindAll))
		authorized.POST("/answers/create", middleware.UserHandler(controller.Create))
		authorized.PUT("/answers/update/:answerId", middleware.UserHandler(controller.Update))
		authorized.DELETE("/answers/:answerId", middleware.UserHandler(controller.Delete))
		authorized.GET("/answers/by-user/:userId", middleware.UserHandler(controller.FindByUserId))
		authorized.GET("/answers/by-question-id/:questionId", middleware.UserHandler(controller.FindByQuestionId))		
	}

	return router
}

func (controller *AnswerController) FindAll(ctx *gin.Context) {
	answers, err := controller.AnswerService.FindAll(ctx.Request.Context())
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
		Data:   answers,
	})
}

func (controller *AnswerController) Create(ctx *gin.Context) {
	var request model.CreateAnswerRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	answer, err := controller.AnswerService.Create(ctx, request)
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
		Status: "answer successfully created",
		Data:   answer,
	})
}

func (controller *AnswerController) Delete(ctx *gin.Context) {
	answerId := utils.ToInt(ctx.Param("answerId"))

	err := controller.AnswerService.Delete(ctx.Request.Context(), answerId)
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
		Status: "answer successfully deleted",
		Data:   nil,
	})
}

func (controller *AnswerController) Update(ctx *gin.Context) {
	var request model.UpdateAnswerRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	answerId := utils.ToInt(ctx.Param("answerId"))

	answer, err := controller.AnswerService.Update(ctx, request, answerId)
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
		Status: "answer successfully update",
		Data:   answer,
	})
}

func (controller *AnswerController) FindByUserId(ctx *gin.Context) {
	userId := utils.ToInt(ctx.Param("userId"))
	answers, err := controller.AnswerService.FindByUserId(ctx.Request.Context(), userId)
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
		Data:   answers,
	})
}

func (controller *AnswerController) FindByQuestionId(ctx *gin.Context) {
	questionId := utils.ToInt(ctx.Param("questionId"))
	answers, err := controller.AnswerService.FindByUserId(ctx.Request.Context(), questionId)
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
		Data:   answers,
	})
}
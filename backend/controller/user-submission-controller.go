package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"net/http"
	"strconv"
	"strings"
)

type UserSubmissionsController struct {
	UserSubmissionsService service.UserSubmissionsService
}

func NewUserSubmissionsController(userSubmissionService *service.UserSubmissionsService) *UserSubmissionsController {
	return &UserSubmissionsController{
		UserSubmissionsService: *userSubmissionService,
	}
}

func (controller *UserSubmissionsController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api/courses/:code/submissions/:articleId")
	{
		authorized.GET("/user-submit", controller.FindAll)
		authorized.POST("/user-submit", controller.Create)
		authorized.PATCH("/user-submit/:userSubmissionId", controller.UpdateGrade)
		authorized.GET("/user-submit/:userSubmissionId", controller.FindUserSubmissionById)
	}

	return router
}

func (controller *UserSubmissionsController) FindAll(ctx *gin.Context) {
	moduleSubmissionId, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	Modsubs, err := controller.UserSubmissionsService.FindAll(ctx.Request.Context(), moduleSubmissionId)
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
		Data:   Modsubs,
	})
}

func (controller *UserSubmissionsController) FindUserSubmissionById(ctx *gin.Context) {
	moduleSubmissionId, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	userSubmissionId, err := strconv.Atoi(ctx.Param("userSubmissionId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	UserSubmission, err := controller.UserSubmissionsService.FindUserSubmissionById(ctx, userSubmissionId, 2, moduleSubmissionId)
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
		Status: "user submission successfully updated",
		Data:   UserSubmission,
	})
}

func (controller *UserSubmissionsController) Create(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	submissionId, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	//claim, _ := ctx.Get("user_id")
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.WebResponse{
	//		Code:   http.StatusBadRequest,
	//		Status: err.Error(),
	//		Data:   nil,
	//	})
	//	return
	//}

	ct := strings.Split(file.Filename, ".")
	file.Filename = utils.RandomString(20) + "." + ct[len(ct)-1]
	path, err := utils.GetPath("/assets/", file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	var request model.CreateUserSubmissionsRequest
	request.UserId = 2
	request.ModuleSubmissionId = submissionId
	request.File = file.Filename

	UserSubmission, err := controller.UserSubmissionsService.SubmitFile(ctx, request)
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
		Status: "user submission successfully created",
		Data:   UserSubmission,
	})
}

func (controller *UserSubmissionsController) UpdateGrade(ctx *gin.Context) {
	var request model.UpdateUserGradeRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	userSubmissionId, err := strconv.Atoi(ctx.Param("userSubmissionId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	request.Id = userSubmissionId

	err = controller.UserSubmissionsService.UpdateGrade(ctx, request)
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
		Status: "user submission successfully updated",
		Data:   nil,
	})
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"io"
	"net/http"
	"os"
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
	authorized := router.Group("/api/courses/:code/submissions/:submissionId")
	{
		authorized.GET("/user-submit/:userSubmissionId", controller.FindUserSubmissionById)
		authorized.POST("/user-submit", controller.Create)
		authorized.PATCH("/user-submit/:userSubmissionId", controller.UpdateGrade)
		authorized.POST("/user-submit/:userSubmissionId/download", controller.Download)
	}

	return router
}

func (controller *UserSubmissionsController) FindUserSubmissionById(ctx *gin.Context) {
	userSubmissionId, err := strconv.Atoi(ctx.Param("userSubmissionId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	userSubmission, err := controller.UserSubmissionsService.FindUserSubmissionById(ctx, userSubmissionId)
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
		Data:   userSubmission,
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

	submissionId, err := strconv.Atoi(ctx.Param("submissionId"))
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

	userSubmission, err := controller.UserSubmissionsService.SubmitFile(ctx, request)
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
		Data:   userSubmission,
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

func (controller *UserSubmissionsController) Download(ctx *gin.Context) {
	userSubmissionId, err := strconv.Atoi(ctx.Param("userSubmissionId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	userSubmission, err := controller.UserSubmissionsService.FindUserSubmissionById(ctx, userSubmissionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}
	if userSubmission.File == "" {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "no file",
			Data:   nil,
		})
		return
	}

	path, err := utils.GetPath("/assets/", userSubmission.File)
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", userSubmission.File)
	ctx.Header("Content-Disposition", contentDisposition)

	if _, err := io.Copy(ctx.Writer, f); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}
}

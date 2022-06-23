package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"strconv"
)

type CourseController struct {
	CourseService     service.CourseService
	UserCourseService service.UserCourseService
}

func NewCourseController(courseService *service.CourseService, userCourseService *service.UserCourseService) *CourseController {
	return &CourseController{
		CourseService:     *courseService,
		UserCourseService: *userCourseService,
	}
}

func (controller *CourseController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api/courses")
	{
		authorized.GET("", controller.FindAll)
		authorized.GET("/:code", controller.FindById)
		authorized.GET("/:code/users", controller.FindAllUserByCourseId)
		authorized.POST("", controller.Create)
		authorized.PATCH("/:code", controller.Update)
		authorized.DELETE("/:code", controller.Delete)
		authorized.PATCH("/:code/status", controller.ChangeStatus)
	}

	return router
}

func (controller *CourseController) FindAll(ctx *gin.Context) {
	status := true
	if ctx.Query("status") != "" {
		statuses, err := strconv.ParseBool(ctx.Query("status"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: err.Error(),
				Data:   nil,
			})
			return
		}
		status = statuses
	}

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

	courses, err := controller.CourseService.FindAll(ctx.Request.Context(), status, limit)
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
		Data:   courses,
	})
}

func (controller *CourseController) FindById(ctx *gin.Context) {
	code := ctx.Param("code")
	course, err := controller.CourseService.FindByCode(ctx.Request.Context(), code)
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
		Data:   course,
	})
}

func (controller *CourseController) Create(ctx *gin.Context) {
	var request model.CreateCourseRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	course, err := controller.CourseService.Create(ctx, request)
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
		Status: "course successfully created",
		Data:   course,
	})
}

func (controller *CourseController) FindAllUserByCourseId(ctx *gin.Context) {
	code := ctx.Param("code")
	responses, err := controller.UserCourseService.FindAllUserByCourseId(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *CourseController) Update(ctx *gin.Context) {
	var request model.UpdateCourseRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	code := ctx.Param("code")
	course, err := controller.CourseService.Update(ctx, request, code)
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
		Status: "course successfully updated",
		Data:   course,
	})
}

func (controller *CourseController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")
	err := controller.CourseService.Delete(ctx.Request.Context(), code)
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
		Status: "course successfully deleted",
		Data:   nil,
	})
}

func (controller *CourseController) ChangeStatus(ctx *gin.Context) {
	var request model.UpdateStatusCourseRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	code := ctx.Param("code")
	err = controller.CourseService.ChangeActiveCourse(ctx, request, code)
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
		Status: "course successfully updated",
		Data:   nil,
	})
}

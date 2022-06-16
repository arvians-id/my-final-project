package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"strconv"
)

type ModuleController struct {
	ModuleService service.ModuleService
}

func NewModuleController(moduleService *service.ModuleService) *ModuleController {
	return &ModuleController{
		ModuleService: *moduleService,
	}
}

func (controller *ModuleController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/modules", controller.FindAll)
		authorized.GET("/courses/:code/modules", controller.FindAllByRelation)
		authorized.GET("/courses/:code/modules/:id", controller.FindById)
		authorized.POST("/modules", controller.Create)
		authorized.POST("/courses/:code/modules", controller.CreateByCourse)
		authorized.PATCH("/courses/:code/modules/:id", controller.Update)
		authorized.DELETE("/courses/:code/modules/:id", controller.Delete)
	}

	return router
}

func (controller *ModuleController) FindAll(ctx *gin.Context) {
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

	modules, err := controller.ModuleService.FindAll(ctx.Request.Context(), limit)
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
		Data:   modules,
	})
}

func (controller *ModuleController) FindAllByRelation(ctx *gin.Context) {
	code := ctx.Param("code")
	modules, err := controller.ModuleService.FindAllByRelation(ctx.Request.Context(), code)
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
		Data:   modules,
	})
}

func (controller *ModuleController) FindById(ctx *gin.Context) {
	code := ctx.Param("code")
	paramModule := ctx.Param("id")
	id, err := strconv.Atoi(paramModule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	module, err := controller.ModuleService.FindById(ctx.Request.Context(), code, id)
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
		Data:   module,
	})
}

func (controller *ModuleController) Create(ctx *gin.Context) {
	var request model.CreateModuleRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	module, err := controller.ModuleService.Create(ctx, request)
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
		Status: "module successfully created",
		Data:   module,
	})
}

func (controller *ModuleController) CreateByCourse(ctx *gin.Context) {
	var request model.CreateModuleByCourseRequest
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
	request.CodeCourse = code

	module, err := controller.ModuleService.CreateByCourse(ctx, request)
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
		Status: "module successfully created",
		Data:   module,
	})
}

func (controller *ModuleController) Update(ctx *gin.Context) {
	var request model.UpdateModuleRequest
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
	paramModule := ctx.Param("id")
	id, err := strconv.Atoi(paramModule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	module, err := controller.ModuleService.Update(ctx, request, code, id)
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
		Status: "module successfully updated",
		Data:   module,
	})
}

func (controller *ModuleController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")
	paramModule := ctx.Param("id")
	id, err := strconv.Atoi(paramModule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	err = controller.ModuleService.Delete(ctx.Request.Context(), code, id)
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
		Status: "module successfully deleted",
		Data:   nil,
	})
}

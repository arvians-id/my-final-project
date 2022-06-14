package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

type ModuleSubmissionsController struct {
	ModuleSubmissionsService service.ModuleSubmissionsService
}

func NewModuleSubmissionsController(moduleSubmissionsService *service.ModuleSubmissionsService) *ModuleSubmissionsController {
	return &ModuleSubmissionsController{
		ModuleSubmissionsService: *moduleSubmissionsService,
	}
}

func (controller *ModuleSubmissionsController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/submissions", controller.FindAll)
		authorized.GET("/submissions/:code", controller.FindByCode)
		authorized.POST("/submissions", controller.Create)
		authorized.PATCH("/submissions/:code", controller.Update)
		authorized.DELETE("/submissions/:code", controller.Delete)
	}

	return router
}

func (controller *ModuleSubmissionsController) FindAll(ctx *gin.Context) {
	Modsubs, err := controller.ModuleSubmissionsService.FindAll(ctx.Request.Context())
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

func (controller *ModuleSubmissionsController) FindByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	Modsubs, err := controller.ModuleSubmissionsService.FindByModId(ctx.Request.Context(), code)
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

func (controller *ModuleSubmissionsController) Create(ctx *gin.Context) {
	var request model.CreateModuleSubmissionsRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	Modsubs, err := controller.ModuleSubmissionsService.Create(ctx, request)
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
		Data:   Modsubs,
	})
}

func (controller *ModuleSubmissionsController) Update(ctx *gin.Context) {
	var request model.UpdateModuleSubmissionsRequest
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
	Modsubs, err := controller.ModuleSubmissionsService.Update(ctx, request, code)
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
		Data:   Modsubs,
	})
}

func (controller *ModuleSubmissionsController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")

	err := controller.ModuleSubmissionsService.Delete(ctx.Request.Context(), code)
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

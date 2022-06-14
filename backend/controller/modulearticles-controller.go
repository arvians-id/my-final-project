package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
)

type ModuleArticlesController struct {
	ModuleArticlesRepository service.ModuleArticlesService
}

func NewModuleArticlesController(moduleArticlesService *service.ModuleArticlesService) *ModuleArticlesController {
	return &ModuleArticlesController{
		ModuleArticlesRepository: *moduleArticlesService,
	}
}

func (controller *ModuleArticlesController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/articles", controller.FindAll)
		authorized.GET("/articles/:code", controller.FindByCode)
		authorized.POST("/articles", controller.Create)
		authorized.PATCH("/articles/:code", controller.Update)
		authorized.DELETE("/articles/:code", controller.Delete)
	}

	return router
}

func (controller *ModuleArticlesController) FindAll(ctx *gin.Context) {
	ModArs, err := controller.ModuleArticlesRepository.FindAll(ctx.Request.Context())
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
		Data:   ModArs,
	})
}

func (controller *ModuleArticlesController) FindByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	ModArs, err := controller.ModuleArticlesRepository.FindByModId(ctx.Request.Context(), code)
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
		Data:   ModArs,
	})
}

func (controller *ModuleArticlesController) Create(ctx *gin.Context) {
	var request model.CreateModuleArticlesRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ModArs, err := controller.ModuleArticlesRepository.Create(ctx, request)
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
		Data:   ModArs,
	})
}

func (controller *ModuleArticlesController) Update(ctx *gin.Context) {
	var request model.UpdateModuleArticlesRequest
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
	ModArs, err := controller.ModuleArticlesRepository.Update(ctx, request, code)
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
		Data:   ModArs,
	})
}

func (controller *ModuleArticlesController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")

	err := controller.ModuleArticlesRepository.Delete(ctx.Request.Context(), code)
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

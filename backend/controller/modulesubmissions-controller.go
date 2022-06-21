package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"strconv"
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
	authorized := router.Group("/api/courses/:code")
	{
		authorized.GET("/submissions", controller.FindAll)
		authorized.GET("/submissions/:articleId", controller.FindByCode)
		authorized.POST("/submissions", controller.Create)
		authorized.PATCH("/submissions/:articleId", controller.Update)
		authorized.DELETE("/submissions/:articleId", controller.Delete)
		authorized.GET("/submissions/:articleId/next", controller.Next)
		authorized.GET("/submissions/:articleId/previous", controller.Previous)
	}

	return router
}

func (controller *ModuleSubmissionsController) FindAll(ctx *gin.Context) {
	codeCourse := ctx.Param("code")
	Modsubs, err := controller.ModuleSubmissionsService.FindAll(ctx.Request.Context(), codeCourse)
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
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}
	Modsubs, err := controller.ModuleSubmissionsService.FindByModId(ctx.Request.Context(), code, idArticle)
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

	code := ctx.Param("code")
	Modsubs, err := controller.ModuleSubmissionsService.Create(ctx, request, code)
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
		Status: "module submission successfully created",
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
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	Modsubs, err := controller.ModuleSubmissionsService.Update(ctx, request, code, idArticle)
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
		Status: "module submission successfully updated",
		Data:   Modsubs,
	})
}

func (controller *ModuleSubmissionsController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	err = controller.ModuleSubmissionsService.Delete(ctx.Request.Context(), code, idArticle)
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
		Status: "module submission successfully deleted",
		Data:   nil,
	})
}

func (controller *ModuleSubmissionsController) Next(ctx *gin.Context) {
	code := ctx.Param("code")
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	nextModule, err := controller.ModuleSubmissionsService.Next(ctx.Request.Context(), code, idArticle)
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
		Data:   nextModule,
	})
}

func (controller *ModuleSubmissionsController) Previous(ctx *gin.Context) {
	code := ctx.Param("code")
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	previousModule, err := controller.ModuleSubmissionsService.Previous(ctx.Request.Context(), code, idArticle)
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
		Data:   previousModule,
	})
}

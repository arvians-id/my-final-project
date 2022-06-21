package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"strconv"
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
	authorized := router.Group("/api/courses/:code")
	{
		authorized.GET("/articles", controller.FindAll)
		authorized.GET("/articles/:articleId", controller.FindByCode)
		authorized.POST("/articles", controller.Create)
		authorized.PATCH("/articles/:articleId", controller.Update)
		authorized.DELETE("/articles/:articleId", controller.Delete)
		authorized.GET("/articles/:articleId/next", controller.Next)
		authorized.GET("/articles/:articleId/previous", controller.Previous)
	}

	return router
}

func (controller *ModuleArticlesController) FindAll(ctx *gin.Context) {
	codeCourse := ctx.Param("code")
	ModArs, err := controller.ModuleArticlesRepository.FindAll(ctx.Request.Context(), codeCourse)
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
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ModArs, err := controller.ModuleArticlesRepository.FindByModId(ctx.Request.Context(), code, idArticle)
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

	code := ctx.Param("code")
	ModArs, err := controller.ModuleArticlesRepository.Create(ctx, request, code)
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
		Status: "module article successfully created",
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
	idArticle, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ModArs, err := controller.ModuleArticlesRepository.Update(ctx, request, code, idArticle)
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
		Status: "module article successfully updated",
		Data:   ModArs,
	})
}

func (controller *ModuleArticlesController) Delete(ctx *gin.Context) {
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

	err = controller.ModuleArticlesRepository.Delete(ctx.Request.Context(), code, idArticle)
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
		Status: "module article successfully deleted",
		Data:   nil,
	})
}

func (controller *ModuleArticlesController) Next(ctx *gin.Context) {
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

	nextModule, err := controller.ModuleArticlesRepository.Next(ctx.Request.Context(), code, idArticle)
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

func (controller *ModuleArticlesController) Previous(ctx *gin.Context) {
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

	previousModule, err := controller.ModuleArticlesRepository.Previous(ctx.Request.Context(), code, idArticle)
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

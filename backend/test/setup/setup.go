package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/route"
)

func ModuleSetup(configuration config.Config) *gin.Engine {
	initialized := route.NewInitializedServer(configuration)
	return initialized
}

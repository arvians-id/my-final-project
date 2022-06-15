package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
)

func UserHandler(handler func(ctx *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "Unauthorized",
			})
			return
		}

		err := service.JWTAuthService().CheckToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Please Login First",
			})
			return
		}

		handler(ctx)
	}
}

func AdminHandler(handler func(ctx *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "Token is required",
				Data:   "",
			})
			return
		}

		err := service.JWTAuthService().CheckToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Please Login First",
			})
			return
		}

		tokenClaims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your secret api key"), nil
		},
		)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "Cannot parse token",
				Data:   "",
			})
			return
		}

		if tokenClaims["role"] != "1" {
			ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   401,
				Status: "You are not admin",
				Data:   "",
			})
			return
		}

		handler(ctx)
	}
}

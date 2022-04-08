package router

import (
	"net/http"

	"github.com/NwokoyeChigozie/quik_task/pkg/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(validate *validator.Validate) *gin.Engine {
	r := gin.New()

	// Middlewares
	// r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	ApiVersion := "v1"
	AuthUrl(r, validate, ApiVersion)
	WalletUrl(r, validate, ApiVersion)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    400,
			"status":  http.StatusNotFound,
		})
	})

	return r
}

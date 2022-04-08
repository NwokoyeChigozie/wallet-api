package router

import (
	"fmt"

	"github.com/NwokoyeChigozie/quik_task/pkg/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AuthUrl(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {

	auth := user.Controller{Validate: validate}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/create_user", auth.CreateUser)
		authUrl.GET("/get_user/:id", auth.GetUser)
		authUrl.POST("/login", auth.Login)
	}
	return r
}

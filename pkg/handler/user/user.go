package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/internal/user"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository"
	"github.com/NwokoyeChigozie/quik_task/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
}

func (base *Controller) CreateUser(c *gin.Context) {
	var (
		ctx = context.Background()
		req = model.CreateUserRequestModel{}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := user.NewUserService(repository.NewUser())
	user, code, err := service.CreateUser(ctx, req)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": user})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetUser(c *gin.Context) {
	var (
		ctx = context.Background()
	)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "wrong id data type, should be of type int", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := user.NewUserService(repository.NewUser())
	user, code, err := service.GetUser(ctx, userID)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", gin.H{"user": user})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) Login(c *gin.Context) {
	var (
		ctx = context.Background()
		req = model.LoginRequestModel{}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := user.NewUserService(repository.NewUser())
	user, code, err := service.Login(ctx, req)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": user})
	c.JSON(http.StatusOK, rd)

}

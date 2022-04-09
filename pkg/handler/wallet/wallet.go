package wallet

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/internal/wallet"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository"
	"github.com/NwokoyeChigozie/quik_task/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
}

var (
	wrongIdType = "wrong id data type, should be of type int"
	parseFailed = "Failed to parse request body"
)

func (base *Controller) GetBalance(c *gin.Context) {
	var (
		ctx = context.Background()
	)

	walletID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", wrongIdType, err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := wallet.NewWalletService(repository.NewWallet())
	wallet, code, err := service.GetWallet(ctx, walletID)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", gin.H{"wallet": wallet})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) CreditBalance(c *gin.Context) {
	var (
		ctx = context.Background()
		req = model.DebitOrCreditWalletRequest{}
	)

	walletID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", wrongIdType, err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", parseFailed, err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := wallet.NewWalletService(repository.NewWallet())
	wallet, code, err := service.CreditWallet(ctx, req, walletID)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", gin.H{"wallet": wallet})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DebitBalance(c *gin.Context) {
	var (
		ctx = context.Background()
		req = model.DebitOrCreditWalletRequest{}
	)

	walletID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", wrongIdType, err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", parseFailed, err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	service := wallet.NewWalletService(repository.NewWallet())
	wallet, code, err := service.DebitWallet(ctx, req, walletID)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", gin.H{"wallet": wallet})
	c.JSON(http.StatusOK, rd)

}

package router

import (
	"fmt"

	"github.com/NwokoyeChigozie/quik_task/pkg/handler/wallet"
	"github.com/NwokoyeChigozie/quik_task/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func WalletUrl(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {

	wallet := wallet.Controller{Validate: validate}

	walletUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion), middleware.Authorize())
	{
		walletUrl.GET("/wallets/:id/balance", wallet.GetBalance)
		walletUrl.POST("/wallets/:id/debit", wallet.DebitBalance)
		walletUrl.POST("/wallets/:id/credit", wallet.CreditBalance)
	}

	return r
}

package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/pkg/middleware"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage/mysql"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage/redis"
	mRedis "github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	db  mysql.MysqlDB
	rdb redis.Redis
}

func NewWallet() *Wallet {
	return &Wallet{
		db:  *mysql.NewMysqlDB(),
		rdb: *redis.GetRedisDb(),
	}
}

func (w *Wallet) GetWallet(ctx context.Context, walletID int) (*model.Wallet, int, error) {
	var (
		wallet  = model.Wallet{}
		keyName = getWalletKeyName(walletID)
	)

	serialized, err := w.rdb.RedisGet(keyName)
	if err == mRedis.Nil {
		_, err := w.db.GetWithCondition("id = ?", &wallet, walletID)
		if err != nil {
			return &wallet, 500, err
		}
		wallet, err = ResolveWallet(wallet)
		if err != nil {
			return &wallet, 500, err
		}
		err = w.rdb.RedisSet(keyName, wallet)
		if err != nil {
			return &wallet, 500, err
		}

		if !middleware.ValidateRequestUser(int(wallet.UserID)) {
			return &wallet, http.StatusUnauthorized, fmt.Errorf("access Denied")
		}

		return &wallet, 200, nil
	}
	err = json.Unmarshal(serialized, &wallet)
	if err != nil {
		return &wallet, 500, err
	}
	if !middleware.ValidateRequestUser(int(wallet.UserID)) {
		return &wallet, http.StatusUnauthorized, fmt.Errorf("access Denied")
	}

	go func(walletID int, keyName string, rdb redis.Redis) {
		gWallet := model.Wallet{}
		_, err := w.db.GetWithCondition("id = ?", &gWallet, walletID)
		if err != nil {
			fmt.Println(err)
		}
		gWallet, err = ResolveWallet(gWallet)
		if err != nil {
			fmt.Println(err)
		}
		err = rdb.RedisSet(keyName, gWallet)
		if err != nil {
			fmt.Println(err)
		}
	}(walletID, keyName, w.rdb)

	return &wallet, 200, nil
}

func (w *Wallet) CreditWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	var (
		wallet = model.Wallet{}
	)
	_, err := w.db.GetWithCondition("id = ?", &wallet, walletID)
	if err != nil {
		return &wallet, 500, err
	}

	wallet, err = ResolveWallet(wallet)
	if err != nil {
		return &wallet, 500, err
	}

	reqAmount := decimal.NewFromFloat32(float32(form.Amount))
	wallet.WalletBalance = wallet.WalletBalance.Add(reqAmount)
	wallet.Balance = wallet.WalletBalance.String()

	err = w.db.UpdateWithCondition("id = ?", model.Wallet{Balance: wallet.Balance}, &wallet, walletID)
	if err != nil {
		return &wallet, 500, err
	}
	go w.GetWallet(ctx, walletID)

	return &wallet, 200, nil
}

func (w *Wallet) DebitWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	var (
		wallet = model.Wallet{}
	)
	_, err := w.db.GetWithCondition("id = ?", &wallet, walletID)
	if err != nil {
		return &wallet, 500, err
	}

	wallet, err = ResolveWallet(wallet)
	if err != nil {
		return &wallet, 500, err
	}

	reqAmount := decimal.NewFromFloat32(float32(form.Amount))

	if wallet.WalletBalance.LessThan(reqAmount) {
		return &wallet, http.StatusBadRequest, fmt.Errorf("insufficient balance")
	}

	wallet.WalletBalance = wallet.WalletBalance.Sub(reqAmount)
	wallet.Balance = wallet.WalletBalance.String()

	err = w.db.UpdateWithCondition("id = ?", model.Wallet{Balance: wallet.Balance}, &wallet, walletID)
	if err != nil {
		return &wallet, 500, err
	}
	go w.GetWallet(ctx, walletID)

	return &wallet, 200, nil

}

func getWalletKeyName(walletID int) string {
	return `wallet_` + strconv.Itoa(int(walletID))
}

func ResolveWallet(wallet model.Wallet) (model.Wallet, error) {
	balance, err := decimal.NewFromString(wallet.Balance)
	if err != nil {
		return wallet, err
	}
	wallet.WalletBalance = balance
	return wallet, nil
}

package repository

import (
	"context"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/stretchr/testify/mock"
)

type Wallet struct {
	mock.Mock
	TestData *model.Wallet
}

func (u *Wallet) DebitWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	args := u.Called(ctx, form, walletID)
	return u.TestData, args.Int(1), args.Error(2)
}

func (u *Wallet) GetWallet(ctx context.Context, walletID int) (*model.Wallet, int, error) {
	args := u.Called(ctx, walletID)
	return u.TestData, args.Int(1), args.Error(2)
}

func (u *Wallet) CreditWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	args := u.Called(ctx, form, walletID)
	return u.TestData, args.Int(1), args.Error(2)
}

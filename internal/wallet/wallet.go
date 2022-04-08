package wallet

import (
	"context"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
)

type WalletRepository interface {
	DebitWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error)
	GetWallet(ctx context.Context, walletID int) (*model.Wallet, int, error)
	CreditWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error)
}

type Service interface {
	DebitWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error)
	GetWallet(ctx context.Context, walletID int) (*model.Wallet, int, error)
	CreditWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error)
}

type wallet struct {
	walletRepository WalletRepository
}

func NewWalletService(walletRepository WalletRepository) Service {
	return &wallet{walletRepository: walletRepository}
}

func (p *wallet) DebitWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	return p.walletRepository.DebitWallet(ctx, form, walletID)
}

func (p *wallet) GetWallet(ctx context.Context, walletID int) (*model.Wallet, int, error) {
	return p.walletRepository.GetWallet(ctx, walletID)
}

func (p *wallet) CreditWallet(ctx context.Context, form model.DebitOrCreditWalletRequest, walletID int) (*model.Wallet, int, error) {
	return p.walletRepository.CreditWallet(ctx, form, walletID)
}

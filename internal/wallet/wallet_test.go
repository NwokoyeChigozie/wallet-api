package wallet_test

import (
	"context"
	"testing"
	"time"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/internal/wallet"
	"github.com/NwokoyeChigozie/quik_task/pkg/mocks/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var (
	greaterThan0 = "amount must be greater than 0"
)

func Test_CreditWallet(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Credit Wallet", func(t *testing.T) {
		balance, _ := decimal.NewFromString("100")
		testData := model.Wallet{
			ID:            1,
			UserID:        1,
			Balance:       "100",
			WalletBalance: balance,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		requestData := model.DebitOrCreditWalletRequest{
			Amount: 10,
		}
		walletRepository := &repository.Wallet{TestData: &testData}

		walletRepository.On("CreditWallet", context.Background(), requestData, 1).Return(testData, 0, nil).Once()

		service := wallet.NewWalletService(walletRepository)
		wallet, code, err := service.CreditWallet(context.Background(), requestData, 1)
		ast.NoError(err)
		ast.Equal(wallet, &testData)
		ast.Equal(code, 0)

		walletRepository.AssertExpectations(t)
	})

	t.Run("testing greater than 0", func(t *testing.T) {
		testData := model.Wallet{}
		requestData := model.DebitOrCreditWalletRequest{
			Amount: -1,
		}
		walletRepository := &repository.Wallet{TestData: &testData}

		walletRepository.On("CreditWallet", context.Background(), requestData, 1).Return(testData, 0, nil).Once()

		service := wallet.NewWalletService(walletRepository)
		user, code, err := service.CreditWallet(context.Background(), requestData, 1)
		ast.Error(err)
		ast.EqualError(err, greaterThan0)
		ast.Equal(user, &testData)
		ast.Equal(code, 400)

	})
}

func Test_DebitWallet(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Credit Wallet", func(t *testing.T) {
		balance, _ := decimal.NewFromString("100")
		testData := model.Wallet{
			ID:            1,
			UserID:        1,
			Balance:       "100",
			WalletBalance: balance,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		requestData := model.DebitOrCreditWalletRequest{
			Amount: 10,
		}
		walletRepository := &repository.Wallet{TestData: &testData}

		walletRepository.On("DebitWallet", context.Background(), requestData, 1).Return(testData, 0, nil).Once()

		service := wallet.NewWalletService(walletRepository)
		wallet, code, err := service.DebitWallet(context.Background(), requestData, 1)
		ast.NoError(err)
		ast.Equal(wallet, &testData)
		ast.Equal(code, 0)

		walletRepository.AssertExpectations(t)
	})

	t.Run("testing greater than 0", func(t *testing.T) {
		testData := model.Wallet{}
		requestData := model.DebitOrCreditWalletRequest{
			Amount: -1,
		}
		walletRepository := &repository.Wallet{TestData: &testData}

		walletRepository.On("DebitWallet", context.Background(), requestData, 1).Return(testData, 0, nil).Once()

		service := wallet.NewWalletService(walletRepository)
		user, code, err := service.DebitWallet(context.Background(), requestData, 1)
		ast.Error(err)
		ast.EqualError(err, greaterThan0)
		ast.Equal(user, &testData)
		ast.Equal(code, 400)

	})
}

func Test_GetWallet(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Credit Wallet", func(t *testing.T) {
		balance, _ := decimal.NewFromString("100")
		testData := model.Wallet{
			ID:            1,
			UserID:        1,
			Balance:       "100",
			WalletBalance: balance,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		walletRepository := &repository.Wallet{TestData: &testData}

		walletRepository.On("GetWallet", context.Background(), 1).Return(testData, 0, nil).Once()

		service := wallet.NewWalletService(walletRepository)
		wallet, code, err := service.GetWallet(context.Background(), 1)
		ast.NoError(err)
		ast.Equal(wallet, &testData)
		ast.Equal(code, 0)

		walletRepository.AssertExpectations(t)
	})

}

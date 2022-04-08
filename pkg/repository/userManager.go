package repository

import (
	"context"
	"fmt"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage/mysql"
	"github.com/NwokoyeChigozie/quik_task/utility"
)

type User struct {
	db mysql.MysqlDB
}

func NewUser() *User {
	return &User{
		db: *mysql.NewMysqlDB(),
	}
}

func (u *User) CreateUser(ctx context.Context, form model.CreateUserRequestModel) (*model.User, int, error) {

	// checking if a user with this email already exists
	var (
		service   = storage.NewMysqlStorageService(&u.db)
		checkUser = model.User{}
	)

	err1, _ := service.GetWithCondition("email = ?", &checkUser, form.Email)
	if err1 == nil {
		return &checkUser, 400, fmt.Errorf("user already exists with this email")
	}

	// generating a password hash
	hashPassword, err := utility.GenPasswordHash(form.Password)
	fmt.Print("hashed", hashPassword)
	if err != nil {
		return &model.User{}, 500, err
	}

	// creating a user with the data
	form.Password = hashPassword
	user := model.User{FirstName: form.FirstName,
		LastName: form.LastName,
		Email:    form.Email,
		Password: hashPassword,
	}
	fmt.Println(user)

	err = service.Create(&user)
	if err != nil {
		return &checkUser, 500, err
	}

	wallet := model.Wallet{
		UserID:  user.ID,
		Balance: "0.00",
	}
	err = service.Create(&wallet)
	if err != nil {
		return &checkUser, 500, err
	}

	transaction := model.Transactions{
		UserID: user.ID,
		Type:   "wallet creation",
		Amount: "0.00",
	}

	err = service.Create(&transaction)
	if err != nil {
		return &checkUser, 500, err
	}

	user.WalletID = wallet.ID

	return &user, 201, nil
}

func (u *User) GetUser(ctx context.Context, userID int) (*model.User, int, error) {

	var (
		user    = model.User{}
		wallet  = model.Wallet{}
		service = storage.NewMysqlStorageService(&u.db)
	)

	// checking if a user with this email already exists
	err1, _ := service.GetWithCondition("id = ?", &user, userID)
	if err1 != nil {
		return &user, 400, fmt.Errorf("user does not exist")
	}

	_, err := service.GetWithCondition("user_id = ?", &wallet, userID)
	if err != nil {
		return &user, 500, err
	}

	user.WalletID = wallet.ID
	return &user, 200, nil
}

func (u *User) Login(ctx context.Context, form model.LoginRequestModel) (*model.User, int, error) {
	var (
		service = storage.NewMysqlStorageService(&u.db)
		user    = model.User{}
		wallet  = model.Wallet{}
	)

	err1, _ := service.GetWithCondition("email = ?", &user, form.Email)
	if err1 != nil {
		return &user, 400, fmt.Errorf("invalid credentials")
	}

	if !utility.CheckPassword(form.Password, user.Password) {
		return &user, 400, fmt.Errorf("invalid credentials")
	}

	token, expiry, err := utility.CreateToken(int(user.ID))
	if err != nil {
		return &user, 500, err
	}

	err = service.UpdateWithCondition("id = ?", model.User{Token: token, TokenExpires: expiry}, &user, user.ID)
	if err != nil {
		return &user, 500, err
	}

	_, err = service.GetWithCondition("user_id = ?", &wallet, user.ID)
	if err != nil {
		return &user, 500, err
	}

	user.WalletID = wallet.ID

	return &user, 200, nil
}

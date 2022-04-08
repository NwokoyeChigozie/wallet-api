package mysql

import (
	"fmt"

	"github.com/NwokoyeChigozie/quik_task/internal/config"
	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection gets connection of mysqlDB database
func Connection() (db *gorm.DB) {
	return DB
}

func ConnectToDB() (db *gorm.DB) {
	dbC := config.GetConfig().Database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&charset=utf8mb4", dbC.Username, dbC.Password, dbC.Host, dbC.Port, dbC.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")
	_ = db.AutoMigrate(MigrationModels()...)
	DB = db
	return db
}

func MigrationModels() []interface{} {
	return []interface{}{
		model.User{},
		model.Wallet{},
		model.Transactions{},
	}
}

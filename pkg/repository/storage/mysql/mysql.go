package mysql

import (
	"errors"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

func NewMysqlDB() *MysqlDB {
	return &MysqlDB{DB: DB}
}

func (db *MysqlDB) GetWithCondition(whereString string, model interface{}, values ...interface{}) (error, error) {
	tx := db.DB.Where(whereString, values...).First(model)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		// first error is nil error
		return tx.Error, tx.Error
	}

	return nil, tx.Error
}

func (db *MysqlDB) Create(modelData interface{}) error {
	tx := db.DB.Create(modelData)
	return tx.Error
}

func (db *MysqlDB) UpdateWithCondition(whereString, data interface{}, model interface{}, values ...interface{}) error {

	tx := db.DB.Model(model).Where(whereString, values...).Updates(data)
	return tx.Error
}

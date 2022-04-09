package storage

// repositories
type MysqlRepository interface {
	GetWithCondition(whereString string, model interface{}, values ...interface{}) (error, error)
	Create(modelData interface{}) error
	UpdateWithCondition(whereString, data interface{}, model interface{}, values ...interface{}) error
}

type RedisRepository interface {
	RedisSet(key string, value interface{}) error
	RedisGet(key string) ([]byte, error)
	RedisDelete(key string) (int64, error)
}

// repositories

package storage

// services
type MysqlStorageService interface {
	GetWithCondition(whereString string, model interface{}, values ...interface{}) (error, error)
	Create(modelData interface{}) error
	UpdateWithCondition(whereString, data interface{}, model interface{}, values ...interface{}) error
}

type RedisStorageService interface {
	RedisSet(key string, value interface{}) error
	RedisGet(key string) ([]byte, error)
	RedisDelete(key string) (int64, error)
}

// services

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

// struct types
type mSql struct {
	mysqlRepository MysqlRepository
}

type redis struct {
	redisRepository RedisRepository
}

// struct types

// instantiation functions
func NewMysqlStorageService(mysqlRepository MysqlRepository) MysqlStorageService {
	return &mSql{mysqlRepository: mysqlRepository}
}

func NewRedisStorageService(redisRepository RedisRepository) RedisStorageService {
	return &redis{redisRepository: redisRepository}
}

// instantiation functions

// service functions

// mysql
func (m *mSql) GetWithCondition(whereString string, model interface{}, values ...interface{}) (error, error) {
	return m.mysqlRepository.GetWithCondition(whereString, model, values...)
}
func (m *mSql) Create(modelData interface{}) error {
	return m.mysqlRepository.Create(modelData)
}
func (m *mSql) UpdateWithCondition(whereString, data interface{}, model interface{}, values ...interface{}) error {
	return m.mysqlRepository.UpdateWithCondition(whereString, data, model, values...)
}

// mysql

// redis
func (r *redis) RedisSet(key string, value interface{}) error {
	return r.redisRepository.RedisSet(key, value)
}
func (r *redis) RedisGet(key string) ([]byte, error) {
	return r.redisRepository.RedisGet(key)
}
func (r *redis) RedisDelete(key string) (int64, error) {
	return r.redisRepository.RedisDelete(key)
}

// redis
// service functions

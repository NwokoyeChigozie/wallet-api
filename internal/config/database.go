package config

type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

type RedisConfiguration struct {
	Redishost string
	Redisport string
}

// type ParamsConfiguration struct {
// 	MasterPassword       string
// 	LiveClassSecretToken string

// 	//BBB
// 	LiveClassClient  string
// 	BBBSecret        string
// 	BBBServerBaseUrl string

// 	//Logging
// 	Sentrydsn   string
// 	Environment string
// 	Debug       bool
// 	Release     string
// 	StackTract  bool
// 	AppBaseUrl  string
// }

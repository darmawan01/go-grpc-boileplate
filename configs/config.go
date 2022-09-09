package configs

import (
	"go_grpc_boileplate/common/constant"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	Env  string
	Port string
	// Database connection info
	DB *ConnInfo
	// Redus connection info
	Redis *ConnInfo
	// JWT
	JWT *JWT
}

type ConnInfo struct {
	Host string
	Port string
	User string
	Pass string
	// Eg: Database name
	Name string
}

type JWT struct {
	SecretKey string
}

// Load configs from env
func LoadFromEnv() *Configs {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Configs{
		Env:  os.Getenv(constant.ENV),
		Port: os.Getenv(constant.PORT),
		DB: &ConnInfo{
			Host: os.Getenv(constant.DB_HOST),
			Port: os.Getenv(constant.DB_PORT),
			User: os.Getenv(constant.DB_USER),
			Pass: os.Getenv(constant.DB_PASS),
			Name: os.Getenv(constant.DB_NAME),
		},
		Redis: &ConnInfo{
			Host: os.Getenv(constant.REDIS_HOST),
			Port: os.Getenv(constant.REDIS_PORT),
			User: os.Getenv(constant.REDIS_USER),
			Pass: os.Getenv(constant.REDIS_PASS),
		},
		JWT: &JWT{
			SecretKey: os.Getenv(constant.JWT_SECRET_KEY),
		},
	}
}

func LoadFromVault() *Configs {
	return &Configs{}
}

func (conf *Configs) IsDevelopment() bool {
	return conf.Env == "development"
}

func (conf *Configs) IsStaging() bool {
	return conf.Env == "staging"
}

func (conf *Configs) IsProduction() bool {
	return conf.Env == "production"
}

package configs

import (
	"context"
	"log"
	"os"

	"go_grpc_boileplate/common/constant"

	"github.com/bytedance/sonic"
	"github.com/joho/godotenv"

	vault "github.com/hashicorp/vault/api"
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

func LoadFromVault(address, token string) *Configs {
	config := vault.DefaultConfig()
	config.Address = address

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(token)

	ctx := context.Background()

	secret, err := client.KVv2(os.Getenv("SERVICE_NAME")).Get(ctx, os.Getenv("SECRET_NAME"))
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	b, err := sonic.Marshal(secret.Data)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	var configs Configs
	err = sonic.Unmarshal(b, &configs)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	return &configs
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

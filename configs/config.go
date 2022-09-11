package configs

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"go_grpc_boileplate/common/constant"

	"github.com/bytedance/sonic"
	"github.com/joho/godotenv"

	vault "github.com/hashicorp/vault/api"
)

type Configs struct {
	Env  string `json:"env"`
	Port string `json:"port"`
	// Database connection info
	DB ConnInfo `json:"db"`
	// Redus connection info
	Redis ConnInfo `json:"redis"`
	// JWT
	JWT JWT `json:"jwt"`
}

type ConnInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	// Eg: Database name
	Name string `json:"name"`

	MaxOpenConn int `json:"max_open_conn"`
	MaxIdleConn int `json:"max_idle_conn"`
	MaxLifeTime int `json:"max_life_time"` // will convert to minutes
}

type JWT struct {
	SecretKey string `json:"secret_key"`
}

var (
	Config Configs
	Mutex  sync.RWMutex
)

// Load configs from env
func LoadFromEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = Configs{
		Env:  envDefaultValueString("development", os.Getenv(constant.ENV)),
		Port: envDefaultValueString("8080", os.Getenv(constant.PORT)),
		DB: ConnInfo{
			Host:        envDefaultValueString("localhost", os.Getenv(constant.DB_HOST)),
			Port:        envDefaultValueString("5432", os.Getenv(constant.DB_PORT)),
			User:        envDefaultValueString("postgres", os.Getenv(constant.DB_USER)),
			Pass:        envDefaultValueString("secret", os.Getenv(constant.DB_PASS)),
			Name:        envDefaultValueString("postgres", os.Getenv(constant.DB_NAME)),
			MaxOpenConn: envDefaultValueInt(100, os.Getenv(constant.DB_MAX_OPEN_CONN)),
			MaxIdleConn: envDefaultValueInt(5, os.Getenv(constant.DB_MAX_IDLE_CONN)),
			MaxLifeTime: envDefaultValueInt(15, os.Getenv(constant.DB_MAX_LIFE_TIME)),
		},
		Redis: ConnInfo{
			Host: envDefaultValueString("localhost", os.Getenv(constant.REDIS_HOST)),
			Port: envDefaultValueString("6379", os.Getenv(constant.REDIS_PORT)),
			User: envDefaultValueString("admin", os.Getenv(constant.REDIS_USER)),
			Pass: envDefaultValueString("secret", os.Getenv(constant.REDIS_PASS)),
		},
		JWT: JWT{
			SecretKey: envDefaultValueString("secretJwtKey", os.Getenv(constant.JWT_SECRET_KEY)),
		},
	}
}

func LoadFromVault(address, token string) {
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

	var temp Configs
	err = sonic.Unmarshal(b, &temp)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	// remove mutex if you don't load the config anymore
	Mutex.Lock()
	Config = temp
	Mutex.Unlock()
}

func (conf Configs) IsDevelopment() bool {
	return strings.ToLower(conf.Env) == "development"
}

func (conf Configs) IsStaging() bool {
	return strings.ToLower(conf.Env) == "staging"
}

func (conf Configs) IsProduction() bool {
	return strings.ToLower(conf.Env) == "production"
}

func envDefaultValueString(defaultValue, data string) string {
	if data == "" {
		return defaultValue
	}
	return data
}

func envDefaultValueInt(defaultValue int, data string) int {
	if data == "" {
		return defaultValue
	}

	dataInt, err := strconv.Atoi(data)
	if err != nil {
		return defaultValue
	}

	return dataInt
}

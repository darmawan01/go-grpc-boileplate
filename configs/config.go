package configs

import (
	"context"
	"log"
	"os"
	"strings"

	"go_grpc_boileplate/common/constant"
	"go_grpc_boileplate/common/util"

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
	JWT `json:"jwt"`
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

type Vault struct {
	Address     string `json:"-"`
	Token       string `json:"-"`
	ServiceName string `json:"-"`
	SecretName  string `json:"-"`
}

var Config Configs

func (conf Configs) IsDevelopment() bool {
	return strings.ToLower(conf.Env) == "development"
}

func (conf Configs) IsStaging() bool {
	return strings.ToLower(conf.Env) == "staging"
}

func (conf Configs) IsProduction() bool {
	return strings.ToLower(conf.Env) == "production"
}

// Load configs from env or vault
func LoadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	isConfigFromVault := util.DefaultValueBool(false, os.Getenv(constant.USE_VAULT_CONFIG))

	if isConfigFromVault {
		vaultConfig = Vault{
			Address:     os.Getenv(constant.VAULT_ADDRESS),
			Token:       os.Getenv(constant.VAULT_TOKEN),
			ServiceName: os.Getenv(constant.SERVICE_NAME),
			SecretName:  os.Getenv(constant.SECRET_NAME),
		}

		vaultConfig.load()
		return
	}

	Config = Configs{
		Env:  util.DefaultValueString("development", os.Getenv(constant.ENV)),
		Port: util.DefaultValueString("8080", os.Getenv(constant.PORT)),
		DB: ConnInfo{
			Host:        util.DefaultValueString("localhost", os.Getenv(constant.DB_HOST)),
			Port:        util.DefaultValueString("5432", os.Getenv(constant.DB_PORT)),
			User:        util.DefaultValueString("postgres", os.Getenv(constant.DB_USER)),
			Pass:        util.DefaultValueString("secret", os.Getenv(constant.DB_PASS)),
			Name:        util.DefaultValueString("postgres", os.Getenv(constant.DB_NAME)),
			MaxOpenConn: util.DefaultValueInt(100, os.Getenv(constant.DB_MAX_OPEN_CONN)),
			MaxIdleConn: util.DefaultValueInt(5, os.Getenv(constant.DB_MAX_IDLE_CONN)),
			MaxLifeTime: util.DefaultValueInt(15, os.Getenv(constant.DB_MAX_LIFE_TIME)),
		},
		Redis: ConnInfo{
			Host: util.DefaultValueString("localhost", os.Getenv(constant.REDIS_HOST)),
			Port: util.DefaultValueString("6379", os.Getenv(constant.REDIS_PORT)),
			User: util.DefaultValueString("admin", os.Getenv(constant.REDIS_USER)),
			Pass: util.DefaultValueString("secret", os.Getenv(constant.REDIS_PASS)),
		},
		JWT: JWT{
			SecretKey: util.DefaultValueString("secretJwtKey", os.Getenv(constant.JWT_SECRET_KEY)),
		},
	}
}

var vaultConfig Vault

func (vc Vault) load() {
	config := vault.DefaultConfig()
	config.Address = vc.Address

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(vc.Token)

	ctx := context.Background()

	secret, err := client.KVv2(vc.ServiceName).Get(ctx, vc.SecretName)
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

	Config = temp
}

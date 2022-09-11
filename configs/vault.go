package configs

import (
	"context"
	"log"

	"github.com/bytedance/sonic"

	vault "github.com/hashicorp/vault/api"
)

type Vault struct {
	Address     string `json:"-"`
	Token       string `json:"-"`
	ServiceName string `json:"-"`
	SecretName  string `json:"-"`
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

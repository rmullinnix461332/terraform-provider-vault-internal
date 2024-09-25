package vaultinternal

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rmullinnix461332/terraform-provider-vault-internal/vaultclient"
	"github.com/rmullinnix461332/logger"
)

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL for Vault",
				DefaultFunc: schema.EnvDefaultFunc("VAULT_URL", nil),
			},
            "role": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vault role for access",
				DefaultFunc: schema.EnvDefaultFunc("VAULT_ROLE", nil),
			},
            "jwt_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path in filesystem for jwt.",
				DefaultFunc: schema.EnvDefaultFunc("JWT_PATH", nil),
            },
		},
		ResourcesMap: map[string]*schema.Resource{
		},
        DataSourcesMap: map[string]*schema.Resource{
			"vaultinternal_secret": datasourceSecret(),
        },
		ConfigureContextFunc: providerConfigure,
	}

	return p
}

type vaultConfig struct {
	server string
    role string
    jwtPath string
	client *vaultclient.VaultClient
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
    logger.Init("info")

	server := data.Get("server").(string)
	role := data.Get("role").(string)
	jwt_path := data.Get("jwt_path").(string)

	client, err := vaultclient.NewVaultClient(server, role, jwt_path)

	if err != nil {
		fmt.Println("config error", err)
	}
	return vaultConfig{
		server: server,
        role: role,
        jwtPath: jwt_path,
		client: &client,
	}, diag.Diagnostics{}
}

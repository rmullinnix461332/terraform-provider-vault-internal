package vaultinternal

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSecret() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
            "id": {
                Type:         schema.TypeString,
                Computed:     true,
            },
            "path": {
                Type:         schema.TypeString,
                ForceNew:     true,
                Required:     true,
            },
            "secret": {
                Type:         schema.TypeMap,
                Computed:     true,
                Sensitive:     true,
            },
        },
        Read: datasourceSecretRead,
    }
}

func datasourceSecretRead(data *schema.ResourceData, meta interface{}) error {
	clientConfig := meta.(vaultConfig)
	client := clientConfig.client
	secretPath := data.Get("path").(string)

	secretValue, err := client.GetSecrets(secretPath)
	if err != nil {
		return err
	}

    fmt.Println("  secret value", secretValue)
	err = data.Set("secret", secretValue)
	if err != nil {
        return fmt.Errorf("Could not get secret for path %s: %s", secretPath, err)
	}

	data.SetId(secretPath)

	return nil
}

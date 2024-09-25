package vaultinternal

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVaultSecret_multi(t *testing.T) {
	sPath := "slus-dcp/data/infrastructure-system/terraform/VariableSets/DCP-Holocron-OIDC"
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            testAccVaultSecret(sPath),
			},
		},
	})
}

func testAccVaultSecret(sPath string) string {
    return fmt.Sprintf(`
terraform {
    required_providers {
        vaultinternal = {
            source = "vault-internal"
        }
    }
}

data "vaultinternal_secret" "test1" {
    path  = %q
}
`, sPath)
}

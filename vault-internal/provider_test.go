package vaultinternal

import (
	"os"
    "sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testInitOnce = sync.Once{}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

const providerName = "vault-internal"

var providerFactories = map[string]func() (*schema.Provider, error){
	providerName: func() (*schema.Provider, error) {
		initTestProvider()
		return testAccProvider, nil
	},
}

func initTestProvider() {
    testInitOnce.Do(
        func() {
            if testAccProvider == nil {
            	testAccProvider = New()
            	testAccProviders = map[string]*schema.Provider{
            		"vault-internal": testAccProvider,
                }
	            testAccProvider.Configure(nil, terraform.NewResourceConfigRaw(nil))
            }
        },
	)
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("VAULT_URL") == "" {
		t.Fatal("VAULT_URL must be set for acceptance tests")
	}
	if os.Getenv("VAULT_ROLE") == "" {
		t.Fatal("VAULT_URL must be set for acceptance tests")
	}
	if os.Getenv("JWT_PATH") == "" {
		t.Fatal("JWT_PATH must be set for acceptance tests")
	}
}

func TestProvider(t *testing.T) {
	if err := New().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = New()
}

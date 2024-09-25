package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/rmullinnix461332/terraform-provider-vault-internal/vault-internal"
)

func main() {

    plugin.Serve(
        &plugin.ServeOpts{
            ProviderFunc: vaultinternal.New,
            ProviderAddr: "app.terraform.io/SLUS-DCP/vault-internal",
        },
    )
}

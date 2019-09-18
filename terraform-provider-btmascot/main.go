package main

import (
	"github.com/billtrust/meetup-terraform-provider/terraform-provider-btmascot/btmascot"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: btmascot.Provider,
	})
}

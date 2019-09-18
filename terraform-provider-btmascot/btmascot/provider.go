package btmascot

import (
	billtrustclient "github.com/billtrust/meetup-terraform-provider/api-client/client"

	"github.com/go-openapi/strfmt"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider : terraform main provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"root_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BT_ROOT_URL", nil),
				Description: "Billtrust Services root URL",
			},
		},
		// All resources handled by this provider must be made public in this map for TF to be able to use them
		ResourcesMap: map[string]*schema.Resource{
			"btmascot_mascot": resourceMascot(),
		},
		ConfigureFunc: configureProvider,
	}
}

// Main entrypoint for terraform. Setup the client that will be used by all resources. DO NOT instanciate a new client in each resource call unless absolutely necessary
func configureProvider(d *schema.ResourceData) (interface{}, error) {
	baseURL := d.Get("root_url").(string)

	transport := httptransport.New(baseURL, "/", nil)
	transport.Producers["application/json"] = runtime.JSONProducer()
	transport.SetDebug(true)
	client := billtrustclient.New(transport, strfmt.Default)

	return client, nil
}

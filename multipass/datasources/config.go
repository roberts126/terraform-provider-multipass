package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func ConfigType() *schema.Resource {
	return &schema.Resource{
		ReadContext: provider.LoadConfig,
		Schema: map[string]*schema.Schema{
			"flag": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}

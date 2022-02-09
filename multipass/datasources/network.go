package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func NetworkType() *schema.Resource {
	return &schema.Resource{
		ReadContext: provider.LoadNetwork,
		Schema: map[string]*schema.Schema{
			"fuzzy": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeBool,
			},
			"description": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"name": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"search": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeString,
			},
			"type": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}

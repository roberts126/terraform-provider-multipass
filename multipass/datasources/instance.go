package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func InstanceType() *schema.Resource {
	return &schema.Resource{
		ReadContext: provider.LoadInstance,
		Schema: map[string]*schema.Schema{
			"fuzzy": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeBool,
			},
			"alias": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"command": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"instance": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"search": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}

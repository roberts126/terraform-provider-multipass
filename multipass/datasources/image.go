package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func ImageType() *schema.Resource {
	return &schema.Resource{
		ReadContext: provider.LoadImage,
		Schema: map[string]*schema.Schema{
			"aliases": {
				Computed: true,
				Type:     schema.TypeList,
			},
			"fuzzy": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeBool,
			},
			"name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"os": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"release": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"remote": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"search": {
				Computed: false,
				Optional: true,
				Type:     schema.TypeString,
			},
			"version": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

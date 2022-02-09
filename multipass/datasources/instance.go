package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func InstanceType() *schema.Resource {
	return &schema.Resource{
		ReadContext: provider.LoadInstance,
		Schema: map[string]*schema.Schema{
			"instances": {
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disks": {
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"Total": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeString,
									},
									"Used": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeString,
									},
								},
							},
						},
						"image_hash": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
						"image_release": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
						"ipv4": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeList,
						},
						"load": {
							Computed: false,
							Optional: true,
							Type:     schema.TypeFloat,
						},
						"memory": {
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"Total": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeInt,
									},
									"Used": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeInt,
									},
								},
							},
						},
						"mounts": {
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gid_mappings": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeList,
									},
									"source_path": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeString,
									},
									"path": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeString,
									},
									"uid_mappings": {
										Computed: true,
										Optional: true,
										Type:     schema.TypeList,
									},
								},
							},
						},
						"release": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
						"state": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

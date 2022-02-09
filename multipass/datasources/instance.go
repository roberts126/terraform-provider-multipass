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
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: InstanceData(),
				},
			},
		},
	}
}

func InstanceData() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disks": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeSet,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"device": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeString,
					},
					"total": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeString,
					},
					"used": {
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
			Elem:     schema.TypeString,
		},
		"load": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeFloat,
		},
		"name": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeString,
		},
		"memory": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeSet,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"total": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeInt,
					},
					"used": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeInt,
					},
				},
			},
		},
		"mounts": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeSet,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"gid_mappings": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeList,
						Elem:     schema.TypeString,
					},
					"source_path": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeString,
					},
					"mount_path": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeString,
					},
					"uid_mappings": {
						Computed: true,
						Optional: true,
						Type:     schema.TypeList,
						Elem:     schema.TypeString,
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
	}
}

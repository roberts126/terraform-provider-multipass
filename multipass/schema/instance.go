package schema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var InstanceDiskSchema = &schema.Resource{
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
}

var InstanceMemorySchema = &schema.Resource{
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
}

var InstanceMountSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"gid_mappings": {
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"source_path": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"mount_path": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"uid_mappings": {
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

var InstanceNetworkSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Required: true,
			Type:     schema.TypeString,
		},
		"mac": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"mode": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"required": {
			Optional: true,
			Type:     schema.TypeBool,
		},
	},
}

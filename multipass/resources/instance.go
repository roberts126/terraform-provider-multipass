package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func InstanceType() *schema.Resource {
	i := &Instance{}

	return &schema.Resource{
		CreateContext: i.Create,
		DeleteContext: i.Delete,
		ReadContext:   i.Read,
		UpdateContext: i.Update,
		Importer: &schema.ResourceImporter{
			StateContext: i.ImportState,
		},
		Schema: creationSchema(),
	}
}

func creationSchema() map[string]*schema.Schema {
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
			Default:  nil,
			Type:     schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"load": {
			Computed: true,
			Optional: true,
			Default:  nil,
			Type:     schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
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
					"local_path": {
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
		"bridged": {
			Optional: true,
			Type:     schema.TypeBool,
		},
		"cloud_init": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"cpus": {
			Optional: true,
			Type:     schema.TypeInt,
		},
		"disk": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"image": {
			Required: true,
			Type:     schema.TypeString,
		},
		"mem": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"name": {
			Required: true,
			Type:     schema.TypeString,
		},
		"network": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
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
			},
		},
	}
}

type Instance struct {
	name  string
	image string
}

func (i Instance) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)
	i.image = d.Get("image").(string)
	i.name = d.Get("name").(string)

	_, err := p.Launch(i.image, i.name, i.buildFlags(d)...)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = i.attachMounts(p, d); err != nil {
		if _, err = p.Delete(i.name); err != nil {
			return diag.FromErr(err)
		}
	}

	diags := provider.LoadSingleInstance(ctx, d, m)
	if diags.HasError() {
		if _, err = p.Delete(i.name); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func (i Instance) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return provider.LoadSingleInstance(ctx, d, m)
}

func (i Instance) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Can't update
	return nil
}

func (i Instance) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)

	name := d.Get("name").(string)

	_, err := p.Delete(name)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func (i Instance) ImportState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

func (i *Instance) attachMounts(p *provider.Provider, d *schema.ResourceData) error {
	var err error

	for _, mount := range i.buildMounts(d) {
		lPath, lOk := mount["local_path"]
		mPath, mOk := mount["mount_path"]

		if lOk && mOk {
			if _, err = p.Mount(i.name, lPath, mPath); err != nil {
				if _, err = p.Delete(i.name); err != nil {
					return err
				}

				return err
			}
		}
	}

	return nil
}

func (i *Instance) buildFlags(d *schema.ResourceData) []string {
	flags := make([]string, 0)
	if _, bridgedOk := d.GetOk("bridged"); bridgedOk {
		flags = append(flags, "--bridged")
	}

	simpleFlags := []string{"cloud_init", "cpus", "disk", "mem"}
	for _, f := range simpleFlags {
		if v, ok := d.GetOk(f); ok {
			flags = append(flags, "--"+f, fmt.Sprintf("%v", v))
		}
	}

	return append(flags, i.buildNetworks(d)...)
}

func (i *Instance) buildMounts(d *schema.ResourceData) []map[string]string {
	mounts := make([]map[string]string, 0)

	if v, ok := d.GetOk("mounts"); ok {
		vl := v.(*schema.Set).List()
		for _, m := range vl {
			iPaths := m.(map[string]interface{})
			paths := make(map[string]string, 0)
			for k, p := range iPaths {
				paths[k] = fmt.Sprintf("%v", p)
			}

			mounts = append(mounts, paths)
		}
	}

	return mounts
}

func (i *Instance) buildNetworks(d *schema.ResourceData) []string {
	flags := make([]string, 0)

	v, ok := d.GetOk("network")
	if !ok {
		return flags
	}

	list := v.(*schema.Set).List()
	if len(list) < 1 {
		return flags
	}

	for _, iNet := range list {
		var net map[string]interface{}
		if net, ok = iNet.(map[string]interface{}); ok {
			var name, mac, mode interface{}
			netFlags := make([]string, 0)

			if name, ok = net["name"]; ok {
				netFlags = append(netFlags, fmt.Sprintf("name=%v", name))
			}

			if mac, ok = net["mac"]; ok {
				netFlags = append(netFlags, fmt.Sprintf("mac=%v", mac))
			}

			if mode, ok = net["mode"]; ok {
				netFlags = append(netFlags, fmt.Sprintf("mode=%v", mode))
			}

			if len(netFlags) > 0 {
				flags = append(flags, []string{"--network", strings.Join(netFlags, ",")}...)
			}
		}
	}

	return flags
}

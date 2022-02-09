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
		Schema: map[string]*schema.Schema{
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
			"mounts": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_path": {
							Required: true,
							Type:     schema.TypeString,
						},
						"mount_path": {
							Required: true,
							Type:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

type Instance struct{}

func (i Instance) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)

	image := d.Get("image").(string)
	name := d.Get("name").(string)

	launchFlags := make([]string, 0)
	simpleFlags := []string{"bridged", "cloud_init", "cpus", "disk", "mem"}
	for _, f := range simpleFlags {
		v, ok := d.GetOk(f)
		if ok {
			if f == "bridged" {
				launchFlags = append(launchFlags, "--"+f)
			} else {
				launchFlags = append(launchFlags, "--"+f, fmt.Sprintf("%v", v))
			}
		}
	}

	launchFlags = append(launchFlags, i.buildNetwork(d)...)

	_, err := p.Launch(image, name, launchFlags...)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, mount := range i.buildMounts(d) {
		lPath, lOk := mount["local_path"]
		mPath, mOk := mount["local_path"]

		if lOk && mOk {
			if _, err = p.Mount(name, lPath, mPath); err != nil {
				if _, err = p.Delete(name); err != nil {
					return diag.FromErr(err)
				}

				return diag.FromErr(err)
			}
		}
	}

	return provider.LoadInstance(ctx, d, m)
}

func (i Instance) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return provider.LoadInstance(ctx, d, m)
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

func (i Instance) buildMounts(d *schema.ResourceData) []map[string]string {
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

func (i Instance) buildNetwork(d *schema.ResourceData) []string {
	flags := make([]string, 0)

	v, ok := d.GetOk("network")
	if !ok {
		return flags
	}

	vl := v.(*schema.Set).List()
	if len(vl) < 1 {
		return flags
	}

	iNetwork := vl[0].(map[string]interface{})

	name, ok := iNetwork["name"]
	if ok {
		flags = append(flags, fmt.Sprintf("name=%v", name))
	}

	mac, ok := iNetwork["mac"]
	if ok {
		flags = append(flags, fmt.Sprintf("mac=%v", mac))
	}

	mode, ok := iNetwork["mode"]
	if ok {
		flags = append(flags, fmt.Sprintf("mode=%v", mode))
	}

	return append([]string{"--network"}, strings.Join(flags, ","))
}

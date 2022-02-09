package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func ConfigType() *schema.Resource {
	a := &Config{}

	return &schema.Resource{
		CreateContext: a.Create,
		DeleteContext: a.Delete,
		ReadContext:   a.Read,
		UpdateContext: a.Update,
		Importer: &schema.ResourceImporter{
			StateContext: a.ImportState,
		},
		Schema: map[string]*schema.Schema{
			"flag": {
				Computed: false,
				Optional: false,
				Type:     schema.TypeString,
			},
			"value": {
				Computed: false,
				Optional: false,
				Type:     schema.TypeString,
			},
		},
	}
}

type Config struct {
}

func (c Config) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)

	if err := c.setFlag(p, d); err != nil {
		return diag.FromErr(err)
	}

	return provider.LoadConfig(ctx, d, m)
}

func (c Config) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return provider.LoadConfig(ctx, d, m)
}

func (c Config) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)

	if err := c.setFlag(p, d); err != nil {
		return diag.FromErr(err)
	}

	return provider.LoadConfig(ctx, d, m)
}

func (c Config) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (c Config) ImportState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

func (c Config) setFlag(p *provider.Provider, d *schema.ResourceData) error {
	flag := d.Get("flag").(string)
	val := d.Get("value").(string)

	if _, err := p.Set(flag, val); err != nil {
		return err
	}

	d.SetId(flag)

	return nil
}

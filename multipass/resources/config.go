package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func ConfigType() *schema.Resource {
	c := &Config{}

	return &schema.Resource{
		CreateContext: c.Create,
		DeleteContext: c.Delete,
		ReadContext:   c.Read,
		UpdateContext: c.Update,
		Importer: &schema.ResourceImporter{
			StateContext: c.ImportState,
		},
		Schema: map[string]*schema.Schema{
			"flag": {
				Required: true,
				Type:     schema.TypeString,
			},
			"value": {
				Required: true,
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

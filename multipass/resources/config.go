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
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: provider.GetSchema(),
	}
}

type Config struct {
}

func (c Config) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)
	p.ConfigureMultipass(d)

	return provider.LoadConfig(ctx, d, m)
}

func (c Config) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return provider.LoadConfig(ctx, d, m)
}

func (c Config) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return c.Create(ctx, d, m)
}

func (c Config) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (c Config) ImportState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

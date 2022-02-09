package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/multipass/provider"
)

func AliasType() *schema.Resource {
	a := &Alias{}

	return &schema.Resource{
		CreateContext: a.Create,
		DeleteContext: a.Delete,
		ReadContext:   a.Read,
		UpdateContext: a.Update,
		Importer: &schema.ResourceImporter{
			StateContext: a.ImportState,
		},
		Schema: map[string]*schema.Schema{
			"alias": {
				Required: true,
				Type:     schema.TypeString,
			},
			"command": {
				Required: true,
				Type:     schema.TypeString,
			},
			"instance": {
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}

type Alias struct {
}

func (a Alias) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)
	if err := a.alias(p, d); err != nil {
		return diag.FromErr(err)
	}

	return provider.LoadAlias(ctx, d, m)
}

func (a Alias) Read(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return provider.LoadAlias(c, d, m)
}

func (a Alias) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)
	if err := a.unalias(p, d); err != nil {
		return diag.FromErr(err)
	}

	if err := a.alias(p, d); err != nil {
		return diag.FromErr(err)
	}

	return provider.LoadAlias(ctx, d, m)
}

func (a Alias) Delete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	p := m.(*provider.Provider)
	if err := a.unalias(p, d); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func (a Alias) ImportState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

func (a *Alias) alias(p *provider.Provider, d *schema.ResourceData) error {
	alias := d.Get("alias").(string)
	command := d.Get("command").(string)
	instance := d.Get("instance").(string)

	_, err := p.Alias(instance, command, alias)
	if err != nil {
		return err
	}

	d.SetId(alias)

	return nil
}

func (a *Alias) unalias(p *provider.Provider, d *schema.ResourceData) error {
	alias := d.Get("alias").(string)

	_, err := p.Unalias(alias)
	if err != nil {
		return err
	}

	return nil
}

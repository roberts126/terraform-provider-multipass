package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/cli"
)

type Provider struct {
	configured bool
	client     *cli.Client
}

func NewProvider(c *cli.Client) *Provider {
	return &Provider{
		configured: true,
		client:     c,
	}
}

func GetFuzzy(d *schema.ResourceData, def bool) bool {
	iFuzzy, beenSet := d.GetOk("fuzzy")
	if !beenSet {
		return def
	}

	var fuzzy, ok bool
	fuzzy, ok = iFuzzy.(bool)

	if !ok {
		return def
	}

	return fuzzy
}

func AddError(diags diag.Diagnostics, msg string, err error) diag.Diagnostics {
	return append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  msg,
		Detail:   err.Error(),
	})
}

func (p *Provider) Alias(instance, command, alias string) ([]byte, error) {
	return p.client.Alias(instance, command, alias)
}

func (p *Provider) Aliases() ([]byte, error) {
	return p.client.Aliases()
}

func (p *Provider) Delete(name string) ([]byte, error) {
	return p.client.Delete(name)
}

func (p *Provider) Get(flag string) ([]byte, error) {
	return p.client.Get(flag)
}

func (p *Provider) Find() ([]byte, error) {
	return p.client.Find()
}

func (p *Provider) Info(name string) ([]byte, error) {
	return p.client.Info(name)
}

func (p *Provider) Launch(image, name string, args ...string) ([]byte, error) {
	return p.client.Launch(image, name, args...)
}

func (p *Provider) Mount(instance, local, mount string) ([]byte, error) {
	return p.client.Mount(instance, local, mount)
}

func (p *Provider) List() ([]byte, error) {
	return p.client.List()
}

func (p *Provider) Networks() ([]byte, error) {
	return p.client.Networks()
}

func (p *Provider) Set(flag, value string) ([]byte, error) {
	return p.client.Set(flag, value)
}

func (p *Provider) Unalias(alias string) ([]byte, error) {
	return p.client.Unalias(alias)
}

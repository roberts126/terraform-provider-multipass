package multipass

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/cli"
	"terraform-multipass-provider/multipass/datasources"
	"terraform-multipass-provider/multipass/provider"
	"terraform-multipass-provider/multipass/resources"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: provider.GetSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"multipass_alias":    resources.AliasType(),
			"multipass_config":   resources.ConfigType(),
			"multipass_instance": resources.InstanceType(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"multipass_alias":    datasources.AliasType(),
			"multipass_config":   datasources.ConfigType(),
			"multipass_image":    datasources.ImageType(),
			"multipass_instance": datasources.InstanceType(),
			"multipass_network":  datasources.NetworkType(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var runner cli.Runner
	var err error
	var diags diag.Diagnostics

	runner, err = cli.NewMultipassDefaultRunner()
	if err != nil {
		return nil, provider.AddError(diags, "error locating multipass command", err)
	}

	p := provider.NewProvider(cli.NewClient(runner))
	p.ConfigureMultipass(d)

	return p, nil
}

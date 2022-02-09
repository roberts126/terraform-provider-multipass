package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/models"
)

func LoadAlias(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	aliasName := d.Get("search").(string)
	fuzzy := GetFuzzy(d, false)

	b, err := p.Aliases()
	if err != nil {
		return AddError(diags, "error getting aliases", err)
	}

	aliases, err := models.NewAliasDetailsFromOutput(b)
	if err != nil {
		return AddError(diags, "error parsing aliases", err)
	}

	alias, err := aliases.FindAlias(aliasName, fuzzy)
	if err != nil {
		return AddError(diags, "error looking up alias", err)
	}

	if err = d.Set("alias", alias.Alias); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("command", alias.Command); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("instance", alias.Instance); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aliasName)

	return diags
}

func LoadConfig(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	flagName := d.Get("flag").(string)

	b, err := p.Get(flagName)
	if err != nil {
		return AddError(diags, "error getting flag", err)
	}

	var flagDetails = struct {
		Flag  string
		Value string
	}{
		Flag:  flagName,
		Value: string(b),
	}

	if err = d.Set("flag", flagDetails); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(flagName)

	return diags
}

func LoadImage(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	search := d.Get("search").(string)
	fuzzy := GetFuzzy(d, false)

	b, err := p.Find()
	if err != nil {
		return AddError(diags, "error getting images", err)
	}

	images, err := models.NewImageDetailsFromOutput(b)
	if err != nil {
		return AddError(diags, "error parsing images", err)
	}

	image, err := images.FindImage(search, fuzzy)
	if err != nil {
		return AddError(diags, "error loading image", err)
	}

	if err = d.Set("name", image.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("os", image.OS); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("release", image.Release); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("remote", image.Remote); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("version", image.Version); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s:%v", image.Name, image.Version))

	return diags
}

func LoadInstance(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	name := d.Get("name").(string)

	b, err := p.Info(name)
	if err != nil {
		return AddError(diags, "error getting instance", err)
	}

	instances, err := models.NewInstanceDetailsFromOutput(b)
	if err != nil {
		return AddError(diags, "error parsing instance", err)
	}

	var data = struct {
		Instances []*models.Instance
	}{
		Instances: instances.List,
	}

	if err = d.Set("instances", data); err != nil {
		return AddError(diags, "error setting instance list", err)
	}

	if name == "" {
		d.SetId("AllInstances")
	} else {
		d.SetId(name)
	}

	return diags
}

func LoadSingleInstance(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	name := d.Get("name").(string)

	b, err := p.Info(name)
	if err != nil {
		return AddError(diags, "error getting instance", err)
	}

	instances, err := models.NewInstanceDetailsFromOutput(b)
	if err != nil {
		return AddError(diags, "error parsing instance", err)
	}

	if len(instances.List) != 1 {
		return AddError(diags, "incorrect number of instance returned", errors.New("invalid number of instances"))
	}

	instance := instances.List[0].AsMap()
	for k, v := range instance {
		if err = d.Set(k, v); err != nil {
			fmt.Printf("Error setting field %s to %v\n", k, v)
			return AddError(diags, "error setting instance list", err)
		}
	}

	d.SetId(name)

	return diags
}

func LoadNetwork(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := m.(*Provider)
	search := d.Get("search").(string)
	fuzzy := GetFuzzy(d, false)

	b, err := p.Find()
	if err != nil {
		return AddError(diags, "error getting networks", err)
	}

	networks, err := models.NewNetworkDetailsFromOutput(b)
	if err != nil {
		return AddError(diags, "error parsing networks", err)
	}

	image, err := networks.FindNetwork(search, fuzzy)
	if err != nil {
		return AddError(diags, "error loading network", err)
	}

	if err = d.Set("name", image.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("description", image.Description); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("type", image.Type); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(image.Name)

	return diags
}

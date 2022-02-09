package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceType() *schema.Resource {
	return nil
}

type Instance struct {
	p Repository
}

func (i Instance) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (i Instance) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (i Instance) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (i Instance) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (i Instance) ImportState(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

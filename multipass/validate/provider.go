package validate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"terraform-multipass-provider/buildopts"
	"terraform-multipass-provider/cli"
	"terraform-multipass-provider/models"
)

func WindowsTerminalProfile(i interface{}, path cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	profile := i.(string)
	if len(profile) == 0 {
		return diags
	}

	if profile != "primary" && profile != "none" {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Bad Windows Terminal profile",
			Detail:        "windows terminal profile must be set to primary or none",
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("client_apps_windows_terminal_profiles")}),
		})
	}

	return diags
}

func BridgedNetwork(i interface{}, path cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	runner, err := cli.NewMultipassDefaultRunner()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Unable To Validate Networks",
			Detail:        err.Error(),
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("local_bridged_network")}),
		})
	}

	b, err := runner.Run("networks", "--format", "json")
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Unable To Validate Networks",
			Detail:        err.Error(),
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("local_bridged_network")}),
		})
	}

	networks, err := models.NewNetworkDetailsFromOutput(b)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Unable To Validate Networks",
			Detail:        err.Error(),
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("local_bridged_network")}),
		})
	}

	if _, err = networks.FindNetwork(i.(string), false); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Unable To Validate Networks",
			Detail:        err.Error(),
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("local_bridged_network")}),
		})
	}

	return diags
}

func Driver(i interface{}, path cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if !driverAllowed(i.(string)) {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Bad Local Driver",
			Detail:        fmt.Sprintf("local driver must be set to one of %s, but was set too %v", strings.Join(buildopts.AllowedDrivers, ", "), i),
			AttributePath: append(path, cty.IndexStep{Key: cty.StringVal("local_driver")}),
		})
	}

	return diags
}

func driverAllowed(driver string) bool {
	for _, allowed := range buildopts.AllowedDrivers {
		if driver == allowed {
			return true
		}
	}

	return false
}

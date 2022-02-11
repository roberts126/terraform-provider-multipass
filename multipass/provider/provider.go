package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-multipass-provider/build"
	"terraform-multipass-provider/cli"
	"terraform-multipass-provider/multipass/validate"
)

const (
	KeyWindowsTerminalProfiles = "client_apps_windows_terminal_profiles"
	KeyGuiAutostart            = "client_gui_autostart"
	KeyGuiHotkey               = "client_gui_hotkey"
	KeyPrimaryName             = "client_primary_name"
	KeyBridgedNetwork          = "local_bridged_network"
	KeyDriver                  = "local_driver"
	KeyPrivilegedMounts        = "local_privileged_mounts"
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

func GetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		KeyWindowsTerminalProfiles: {
			Optional:         true,
			Type:             schema.TypeString,
			ValidateDiagFunc: validate.WindowsTerminalProfile,
		},
		KeyGuiAutostart: {
			Optional: true,
			Type:     schema.TypeBool,
		},
		KeyGuiHotkey: {
			Optional: true,
			Type:     schema.TypeString,
		},
		KeyPrimaryName: {
			Optional: true,
			Type:     schema.TypeString,
		},
		KeyBridgedNetwork: {
			Optional:         true,
			Type:             schema.TypeString,
			ValidateDiagFunc: validate.BridgedNetwork,
		},
		KeyDriver: {
			Optional:         true,
			Type:             schema.TypeString,
			ValidateDiagFunc: validate.Driver,
		},
		KeyPrivilegedMounts: {
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func GetSchemaKeys() map[string]string {
	keys := map[string]string{
		KeyGuiAutostart:     "client.apps.windows-terminal.profiles",
		KeyGuiHotkey:        "client.gui.autostart",
		KeyPrimaryName:      "client.gui.hotkey",
		KeyBridgedNetwork:   "client.primary-name",
		KeyDriver:           "local.bridged-network",
		KeyPrivilegedMounts: "local.driver",
	}

	if build.SetWindowsTerminal {
		keys[KeyWindowsTerminalProfiles] = ""
	}

	return keys
}

func (p *Provider) ConfigureMultipass(d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	for key, flag := range GetSchemaKeys() {
		if v, ok := d.GetOk(key); ok {
			if b, err := p.Set(flag, fmt.Sprintf("%v", v)); err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Error Setting Config",
					Detail:   err.Error() + "\n" + string(b),
				})

				return diags
			}
		}
	}

	return diags
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

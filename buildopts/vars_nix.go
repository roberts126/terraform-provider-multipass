//go:build linux || freebsd

package buildopts

const (
	MultipassExecutable = "multipass"
	SetWindowsTerminal  = false
)

var AllowedDrivers = []string{"qemu", "libvirt", "lxd"}

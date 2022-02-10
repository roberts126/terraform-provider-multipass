//go:build linux || freebsd

package build

const (
	MultipassExecutable = "multipass"
	SetWindowsTerminal  = false
)

var AllowedDrivers = []string{"qemu", "libvirt", "lxd"}

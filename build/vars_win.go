//go:build windows

package build

const (
	MultipassExecutable = "multipass.exe"
	SetWindowsTerminal  = true
)

var AllowedDrivers = []string{"hyperv", "virtualbox"}

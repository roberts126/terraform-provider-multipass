//go:build windows

package buildopts

const (
	MultipassExecutable = "multipass.exe"
	SetWindowsTerminal  = true
)

var AllowedDrivers = []string{"hyperv", "virtualbox"}

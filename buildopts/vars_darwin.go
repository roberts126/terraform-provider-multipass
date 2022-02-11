//go:build darwin

package buildopts

const (
	MultipassExecutable = "multipass"
	SetWindowsTerminal  = false
)

var AllowedDrivers = []string{"hyperkit", "virtualbox"}

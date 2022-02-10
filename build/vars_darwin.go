//go:build darwin

package build

const (
	MultipassExecutable = "multipass"
	SetWindowsTerminal  = false
)

var AllowedDrivers = []string{"hyperkit", "virtualbox"}

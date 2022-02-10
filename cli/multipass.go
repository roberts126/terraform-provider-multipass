package cli

import (
	"bytes"
	"os/exec"

	"terraform-multipass-provider/build"
)

type MultipassRunner struct {
	baseCommand string
}

func NewMultipassDefaultRunner() (*MultipassRunner, error) {
	c, err := exec.LookPath(build.MultipassExecutable)
	if err != nil {
		return nil, err
	}

	return NewMultipassRunner(c), nil
}

func NewMultipassRunner(command string) *MultipassRunner {
	return &MultipassRunner{
		baseCommand: command,
	}
}

func (r *MultipassRunner) Run(args ...string) ([]byte, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(r.baseCommand, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return stderr.Bytes(), err
	}

	return stdout.Bytes(), nil
}

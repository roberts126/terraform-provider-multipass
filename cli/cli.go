package cli

import (
	"fmt"
)

const (
	argFmt  = "--format"
	argJson = "json"
)

type Runner interface {
	Run(args ...string) ([]byte, error)
}

// Client represents the entry point for the multipass cli
type Client struct {
	runner Runner
}

// NewClient returns a pointer to a new client
func NewClient(r Runner) *Client {
	return &Client{
		runner: r,
	}
}

// Alias executes the multipass alias command
func (c *Client) Alias(instance, command, alias string) ([]byte, error) {
	return c.runner.Run("alias", fmt.Sprintf("%s:%s", instance, command), alias)
}

// Aliases executes the multipass aliases command with json output
func (c *Client) Aliases() ([]byte, error) {
	return c.runner.Run("aliases", argFmt, argJson)
}

// Delete executes the multipass delete command with the --purge flag
func (c *Client) Delete(name string) ([]byte, error) {
	return c.runner.Run("delete", "--purge", name)
}

// Get executes the multipass get command
func (c *Client) Get(flag string) ([]byte, error) {
	return c.runner.Run("get", flag)
}

// Find executes the multipass find command with json output
func (c *Client) Find() ([]byte, error) {
	return c.runner.Run("find", argFmt, argJson)
}

// Info executes the multipass info command with json output.
// If name is empty it return info for all instances.
func (c *Client) Info(name string) ([]byte, error) {
	if name == "" {
		return c.runner.Run("info", "--all", argFmt, argJson)
	}

	return c.runner.Run("info", argFmt, argJson, name)
}

// Launch executes the multipass launch command
func (c *Client) Launch(image, name string, args ...string) ([]byte, error) {
	args = append([]string{"launch"}, args...)
	args = append(args, "--name", name, image)

	return c.runner.Run(args...)
}

// List executes the multipass list command with json output
func (c *Client) List() ([]byte, error) {
	return c.runner.Run("list", argFmt, argJson)
}

// Networks executes the multipass networks command with json output
func (c *Client) Networks() ([]byte, error) {
	return c.runner.Run("networks", argFmt, argJson)
}

// Set executes the multipass set command
func (c *Client) Set(flag, value string) ([]byte, error) {
	return c.runner.Run("set", flag, value)
}

// Unalias executes the multipass unalias command
func (c *Client) Unalias(alias string) ([]byte, error) {
	return c.runner.Run("unalias", alias)
}

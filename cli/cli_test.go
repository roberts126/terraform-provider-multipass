package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	noErrFmt     = "%s command must not error"
	incorrectFmt = "incorrect %s command result"
)

type mockRunner struct {
	err error
}

func (c *mockRunner) Run(args ...string) ([]byte, error) {
	e := c.err

	if c.err != nil {
		c.err = nil
	}

	return []byte("noop " + strings.Join(args, " ")), e
}

func TestCli(t *testing.T) {
	runner := &mockRunner{}

	client := NewClient(runner)

	t.Run("TestAlias", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop alias instance:command alias")
		actual, err := client.Alias("instance", "command", "alias")

		if assert.NoErrorf(t, err, noErrFmt, "alias") {
			assert.Equalf(t, expected, actual, incorrectFmt, "alias")
		}
	})

	t.Run("TestAliases", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop aliases --format json")
		actual, err := client.Aliases()

		if assert.NoErrorf(t, err, noErrFmt, "aliases") {
			assert.Equalf(t, expected, actual, incorrectFmt, "aliases")
		}
	})

	t.Run("TestGet", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop get flag")
		actual, err := client.Get("flag")

		if assert.NoErrorf(t, err, noErrFmt, "get") {
			assert.Equalf(t, expected, actual, incorrectFmt, "flag")
		}
	})

	t.Run("TestFind", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop find --format json")
		actual, err := client.Find()

		if assert.NoErrorf(t, err, noErrFmt, "find") {
			assert.Equalf(t, expected, actual, incorrectFmt, "find")
		}
	})

	t.Run("TestInfo", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop info --format json name")
		actual, err := client.Info("name")

		if assert.NoErrorf(t, err, noErrFmt, "info") {
			assert.Equalf(t, expected, actual, incorrectFmt, "info")
		}
	})

	t.Run("TestInfoAll", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop info --all --format json")
		actual, err := client.Info("")

		if assert.NoError(t, err, noErrFmt, "info with empty string") {
			assert.Equalf(t, expected, actual, incorrectFmt, "info with empty string")
		}
	})

	t.Run("TestLaunch", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop launch --arg value --name name image")
		actual, err := client.Launch("image", "name", "--arg", "value")

		if assert.NoErrorf(t, err, noErrFmt, "launch") {
			assert.Equalf(t, expected, actual, incorrectFmt, "launch")
		}
	})

	t.Run("TestList", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop list --format json")
		actual, err := client.List()

		if assert.NoErrorf(t, err, noErrFmt, "list") {
			assert.Equalf(t, expected, actual, incorrectFmt, "list")
		}
	})

	t.Run("TestNetworks", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop networks --format json")
		actual, err := client.Networks()

		if assert.NoErrorf(t, err, noErrFmt, "networks") {
			assert.Equalf(t, expected, actual, incorrectFmt, "networks")
		}
	})

	t.Run("TestSet", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop set flag value")
		actual, err := client.Set("flag", "value")

		if assert.NoErrorf(t, err, noErrFmt, "set") {
			assert.Equalf(t, expected, actual, incorrectFmt, "set")
		}
	})

	t.Run("TestUnalias", func(t *testing.T) {
		runner.err = nil

		expected := []byte("noop unalias alias")
		actual, err := client.Unalias("alias")

		if assert.NoErrorf(t, err, noErrFmt, "unalias") {
			assert.Equalf(t, expected, actual, incorrectFmt, "unalias")
		}
	})
}

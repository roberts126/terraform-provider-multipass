package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AliasDetails struct {
	Aliases []Alias `json:"aliases" tfsdk:"aliases"`
}

type Alias struct {
	Alias    string `json:"alias" tfsdk:"alias"`
	Command  string `json:"command" tfsdk:"command"`
	Instance string `json:"instance" tfsdk:"instance"`
}

func NewAliasDetailsFromOutput(b []byte) (*AliasDetails, error) {
	d := &AliasDetails{}
	if err := json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *AliasDetails) FindAlias(alias string, fuzzy bool) (*Alias, error) {
	for _, a := range d.Aliases {
		if a.Alias == alias || (fuzzy && strings.Contains(a.Alias, alias)) {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("could not locate alias: %s", alias)
}

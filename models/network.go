package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NetworkDetails struct {
	List []Network `json:"list" tfsdk:"list"`
}

type Network struct {
	Description string `json:"description" tfsdk:"description"`
	Name        string `json:"name" tfsdk:"name"`
	Type        string `json:"type" tfsdk:"type"`
}

func NewNetworkDetailsFromOutput(b []byte) (*NetworkDetails, error) {
	d := &NetworkDetails{}
	if err := json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *NetworkDetails) FindNetwork(nw string, fuzzy bool) (*Network, error) {
	for _, net := range d.List {
		if net.Name == nw || (fuzzy && strings.Contains(net.Name, nw)) {
			return &net, nil
		}
	}

	return nil, fmt.Errorf("could not locate network: %s", nw)
}

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ImageDetails struct {
	Errors []string         `json:"errors" tfsdk:"errors"`
	Images map[string]Image `json:"images" tfsdk:"images"`
}

type Image struct {
	Name    string   `json:"name,omitempty" tfsdk:"name"`
	Aliases []string `json:"aliases" tfsdk:"aliases"`
	OS      string   `json:"os" tfsdk:"os"`
	Release string   `json:"release" tfsdk:"release"`
	Remote  string   `json:"remote" tfsdk:"remote"`
	Version string   `json:"version" tfsdk:"version"`
}

func NewImageDetailsFromOutput(b []byte) (*ImageDetails, error) {
	d := &ImageDetails{}
	if err := json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *ImageDetails) FindImage(name string, fuzzy bool) (*Image, error) {
	var err error
	var match *Image

	for n, img := range d.Images {
		if n == name || (fuzzy && strings.Contains(n, name)) || img.HasAlias(name, fuzzy) {
			if match == nil {
				t := img // img address keeps getting overwritten
				match = &t
				match.Name = n
			} else {
				return nil, fmt.Errorf("more than one result returned for image: %s", name)
			}
		}
	}

	if match == nil {
		err = fmt.Errorf("could not locate image: %s", name)
		match = nil
	}

	return match, err
}

func (i *Image) HasAlias(alias string, fuzzy bool) bool {
	for _, a := range i.Aliases {
		if a == alias || (fuzzy && strings.Contains(a, alias)) {
			return true
		}
	}

	return false
}

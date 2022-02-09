package models

import "encoding/json"

type InstanceDetails struct {
	Errors []string            `json:"errors" tfsdk:"errors"`
	List   map[string]Instance `json:"info" tfsdk:"info"`
}

type Instance struct {
	Name         string           `json:"-" tfsdk:"name"`
	Disks        map[string]Disk  `json:"disks" tfsdk:"disks"`
	ImageHash    string           `json:"image_hash" tfsdk:"image_hash"`
	ImageRelease string           `json:"image_release" tfsdk:"image_release"`
	Ipv4         []string         `json:"ipv4" tfsdk:"ipv4"`
	Load         []float64        `json:"load" tfsdk:"load"`
	Memory       Memory           `json:"memory" tfsdk:"memory"`
	Mounts       map[string]Mount `json:"mounts" tfsdk:"mounts"`
	Release      string           `json:"release" tfsdk:"release"`
	State        string           `json:"state" tfsdk:"state"`
}

type Disk struct {
	Total string `json:"total" tfsdk:"total"`
	Used  string `json:"used" tfsdk:"used"`
}

type Memory struct {
	Total int64 `json:"total" tfsdk:"total"`
	Used  int64 `json:"used" tfsdk:"used"`
}

type Mount struct {
	GidMappings []string `json:"gid_mappings" tfsdk:"gid_mappings"`
	SourcePath  string   `json:"source_path" tfsdk:"source_path"`
	UIDMappings []string `json:"uid_mappings" tfsdk:"uid_mappings"`
}

func NewInstanceDetailsFromOutput(b []byte) (*InstanceDetails, error) {
	d := &InstanceDetails{}
	if err := json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	return d, nil
}

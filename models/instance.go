package models

type InstanceDetails struct {
	Errors []string            `json:"errors" tfsdk:"errors"`
	Info   map[string]Instance `json:"info" tfsdk:"info"`
}

type Instance struct {
	Disks        map[string]TotalUsed `json:"disks" tfsdk:"disks"`
	ImageHash    string               `json:"image_hash" tfsdk:"image_hash"`
	ImageRelease string               `json:"image_release" tfsdk:"image_release"`
	Ipv4         []string             `json:"ipv4" tfsdk:"ipv4"`
	Load         []float64            `json:"load" tfsdk:"load"`
	Memory       TotalUsed            `json:"memory" tfsdk:"memory"`
	Mounts       map[string]Mount     `json:"mounts" tfsdk:"mounts"`
	Release      string               `json:"release" tfsdk:"release"`
	State        string               `json:"state" tfsdk:"state"`
}

type TotalUsed struct {
	Total string `json:"total" tfsdk:"total"`
	Used  string `json:"used" tfsdk:"used"`
}

type Mount struct {
	GidMappings []string `json:"gid_mappings" tfsdk:"gid_mappings"`
	SourcePath  string   `json:"source_path" tfsdk:"source_path"`
	UIDMappings []string `json:"uid_mappings" tfsdk:"uid_mappings"`
}

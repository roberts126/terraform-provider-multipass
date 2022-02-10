package models

import (
	"encoding/json"
)

type InstanceDetails struct {
	Errors []string            `json:"errors" tfsdk:"errors"`
	Info   map[string]Instance `json:"info"  tfsdk:"-"`
	List   []*Instance         `json:"-"  tfsdk:"instances"`
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
	GIDMappings []string `json:"gid_mappings" tfsdk:"gid_mappings"`
	SourcePath  string   `json:"source_path" tfsdk:"source_path"`
	UIDMappings []string `json:"uid_mappings" tfsdk:"uid_mappings"`
}

func NewInstanceDetailsFromOutput(b []byte) (*InstanceDetails, error) {
	d := &InstanceDetails{}
	if err := json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	i := 0
	d.List = make([]*Instance, len(d.Info))
	for name, instancePtr := range d.Info {
		instance := instancePtr
		instance.Name = name

		d.List[i] = &instance

		delete(d.Info, name)
		i++
	}

	return d, nil
}

func (i *Instance) AsMap() map[string]interface{} {
	disks := make([]map[string]interface{}, 0)
	for device, disk := range i.Disks {
		disks = append(disks, disk.AsMap(device))
	}

	mounts := make([]map[string]interface{}, 0)
	for path, mount := range i.Mounts {
		mounts = append(mounts, mount.AsMap(path))
	}

	memory := []map[string]interface{}{
		{
			"total": i.Memory.Total,
			"used":  i.Memory.Used,
		},
	}

	return map[string]interface{}{
		"name":          i.Name,
		"disks":         disks,
		"image_hash":    i.ImageHash,
		"image_release": i.ImageRelease,
		"memory":        memory,
		"ipv4":          i.Ipv4,
		"mounts":        mounts,
		"release":       i.Release,
		"state":         i.State,
	}
}

func (d *Disk) AsMap(device string) map[string]interface{} {
	return map[string]interface{}{
		"device": device,
		"total":  d.Total,
		"used":   d.Used,
	}
}

func (m *Memory) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"total": m.Total,
		"used":  m.Used,
	}
}

func (m *Mount) AsMap(path string) map[string]interface{} {
	return map[string]interface{}{
		"gid_mappings": m.GIDMappings,
		"mount_path":   path,
		"source_path":  m.SourcePath,
		"uid_mappings": m.UIDMappings,
	}
}

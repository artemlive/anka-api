package anka_api

import "time"

var NodeStates = []string{
	"Offline",
	"Inactive (Invalid License)",
	"Active",
	"Updating",
}

var InstanceStates = []string{
	"Scheduling",
	"Pulling",
	"Started",
	"Stopping",
	"Stopped",
	"Terminating",
	"Terminated",
	"Error",
	"Pushing",
}

type Node struct {
	NodeID         string      `json:"node_id"`
	NodeName       string      `json:"node_name"`
	CPU            uint        `json:"cpu_count,omitempty"`
	RAM            uint        `json:"ram,omitempty"`
	VMCount        uint        `json:"vm_count,omitempty"`
	VCPUCount      uint        `json:"vcpu_count,omitempty"`
	VRAM           uint        `json:"vram,omitempty"`
	CPUUtilization float32     `json:"cpu_util,omitempty"`
	RAMUtilization float32     `json:"ram_util,omitempty"`
	FreeDiskSpace  uint        `json:"free_disk_space,omitempty"`
	AnkaDiskUsage  uint        `json:"anka_disk_usage,omitempty"`
	DiskSize       uint        `json:"disk_size,omitempty"`
	State          string      `json:"state"`
	Capacity       uint        `json:"capacity"`
	Groups         []NodeGroup `json:"groups,omitempty"`
	Templates      []Template  `json:"templates,omitempty"`
	AnkaVersion    AnkaVersion `json:"anka_version"`
}

type NodeGroup struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	FallBackGroupId string `json:"fallback_group_id"`
}

type RegistryDisk struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

type Instance struct {
	InstanceID string `json:"instance_id"`
	Vm         VmData `json:"vm"`
}

type VmData struct {
	AnkaRegistry  string    `json:"anka_registry"`
	Ts            time.Time `json:"ts"`
	Progress      int       `json:"progress"`
	Tag           string    `json:"tag"`
	InstanceState string    `json:"instance_state"`
	CrTime        time.Time `json:"cr_time"`
	Vmid          string    `json:"vmid"`
	VMInfo        VMInfo    `json:"vminfo"`
}

type VMInfo struct {
	IP                  string         `json:"ip"`
	Status              string         `json:"status"`
	CPUCores            int            `json:"cpu_cores"`
	Name                string         `json:"name"`
	RAM                 string         `json:"ram"`
	VncPort             int            `json:"vnc_port"`
	NodeID              string         `json:"node_id"`
	PortForwarding      PortForwarding `json:"port_forwarding"`
	VncConnectionString string         `json:"vnc_connection_string"`
	HostIP              string         `json:"host_ip"`
	UUID                string         `json:"uuid"`
}

type PortForwarding []struct {
	Name      string `json:"Name"`
	Protocol  string `json:"protocol"`
	HostPort  int    `json:"host_port"`
	GuestPort int    `json:"guest_port"`
}

type Response interface {
	GetStatus() string
	GetMessage() string
	GetBody() interface{}
}

type DefaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (r *DefaultResponse) GetStatus() string {
	return r.Status
}

func (r *DefaultResponse) GetMessage() string {
	return r.Message
}

type NodesResponse struct {
	DefaultResponse
	Body []Node `json:"body,omitempty"`
}

func (r *NodesResponse) GetBody() interface{} {
	return r.Body
}

type RegistryDiskResponse struct {
	DefaultResponse
	Body RegistryDisk `json:"body,omitempty"`
}

func (r *RegistryDiskResponse) GetBody() interface{} {
	return r.Body
}

type Template struct {
	UUID string `json:"id"`
	Name string `json:"name"`
	Size uint   `json:"size"`
	Tags []TemplateTag
}
type RegistryTemplateResponse struct {
	DefaultResponse
	Body []Template `json:"body,omitempty"`
}

func (r *RegistryTemplateResponse) GetBody() interface{} {
	return r.Body
}

type TemplateTag struct {
	Name string `json:"tag"`
	Size uint   `json:"size"`
}

type RegistryTemplateTags struct {
	Versions []TemplateTag `json:"versions,omitempty"`
}

type RegistryTemplateTagsResponse struct {
	DefaultResponse
	Body RegistryTemplateTags `json:"body,omitempty"`
}

func (r *RegistryTemplateTagsResponse) GetBody() interface{} {
	return r.Body
}

type InstancesResponse struct {
	DefaultResponse
	Body []Instance `json:"body,omitempty"`
}

func (r *InstancesResponse) GetBody() interface{} {
	return r.Body
}

type VMResponse struct {
}

type AnkaVersion struct {
	Build   string `json:"build"`
	Product string `json:"product"`
	Version string `json:"version"`
	License string `json:"license"`
}

type VMRegistryConcreteTemplate struct {
	Name     string                      `json:"name"`
	ID       string                      `json:"id"`
	Versions []VMRegistryTemplateVersion `json:"versions"`
}

type VMRegistryTemplateVersion struct {
	ConfigFile  string   `json:"config_file"`
	StateFiles  []string `json:"state_files"`
	Number      int      `json:"number"`
	Size        int64    `json:"size"`
	Description string   `json:"description"`
	Nvram       string   `json:"nvram"`
	Images      []string `json:"images"`
	Tag         string   `json:"tag"`
}

type StatusBodyResponse struct {
	DefaultResponse
	StatusBody Response `json:"body"`
}

type StatusBody struct {
	Status          string `json:"status"`
	Version         string `json:"version"`
	RegistryAddress string `json:"registry_address"`
	RegistryStatus  string `json:"registry_status"`
	License         string `json:"license"`
}

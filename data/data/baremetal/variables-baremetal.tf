variable "ironic_uri" {
  type =  "string"
  description = "ironic connection URI"
}

variable "libvirt_uri" {
  type        = "string"
  description = "libvirt connection URI"
}

variable "os_image" {
  type        = "string"
  description = "The URL of the OS disk image"
}

variable "baremetal_bridge" {
  type        = "string"
  description = "The name of the baremetal bridge"
}

variable "overcloud_bridge" {
  type        = "string"
  description = "The name of the overcloud bridge"
}

variable "nodes" {
  type        = "list"
  description = "Baremetal node details for provisioning masters"
}

variable "master_configuration" {
  type        = "map"
  description = "Configuration information for masters such as image location"
}
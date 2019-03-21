variable "image_source" {
  description = "The URL of the OS disk image"
  type        = "string"
}

variable "image_checksum" {
  type        = "string"
  description = "The URL or checksum value of the image"
}

variable "root_gb" {
  type        = "string"
  description = "Size of the root disk"
}

variable "root_disk" {
  type        = "string"
  description = "Location of the root disk"
}

variable "ignition" {
  type        = "string"
  description = "The content of the master ignition file."
}

variable "nodes" {
  type        = "map"
  description = "Bare metal node provisioning information"
}

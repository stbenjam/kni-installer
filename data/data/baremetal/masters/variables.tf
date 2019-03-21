variable "image" {
  description = "The URL of the OS disk image"
  type        = "string"
}

variable "image_checksum" {
  type        = "string"
  description = "The URL or checksum value of the image"
}

variable "ignition" {
  type        = "string"
  description = "The content of the master ignition file."
}

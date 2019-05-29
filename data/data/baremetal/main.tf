provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

provider "ironic" {
  url          = "${var.ironic_uri}"
  microversion = "1.52"
}

module "bootstrap" {
  source = "./bootstrap"

  cluster_id          = "${var.cluster_id}"
  image               = "${var.os_image}"
  ignition            = "${var.ignition_bootstrap}"
  external_bridge     = "${var.external_bridge}"
  provisioning_bridge = "${var.provisioning_bridge}"
}

module "masters" {
  source = "./masters"

  ignition       = "${var.ignition_master}"
  image_source   = "${var.master_configuration["image_source"]}"
  image_checksum = "${var.master_configuration["image_checksum"]}"
  root_gb        = "${var.master_configuration["root_gb"]}"
  root_disk      = "${var.master_configuration["root_disk"]}"

  control_plane  = "${var.control_plane}"
  workers        = "${var.workers}"
}

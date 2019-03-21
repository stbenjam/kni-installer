provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

provider "ironic" {
  url = "${var.ironic_uri}"
  microversion = "1.50"
}

module "bootstrap" {
  source = "./bootstrap"

  cluster_id       = "${var.cluster_id}"
  image            = "${var.os_image}"
  ignition         = "${var.ignition_bootstrap}"
  baremetal_bridge = "${var.baremetal_bridge}"
  overcloud_bridge = "${var.overcloud_bridge}"
}

module "masters" {
  source          = "./masters"

  image           = "${var.master_image}"
  image_checksum  = "${var.master_image_checksum}"
  ignition        = "${var.ignition_master}"
}

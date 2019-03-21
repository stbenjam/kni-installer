resource "ironic_node_v1" "openshift-master-0" {
  name = "${var.nodes[0].name}"

  target_provision_state = "active"
  user_data = "${var.ignition}"

  ports = "${var.nodes[0].ports}"

  properties = "${var.nodes[0].properties}"

  instance_info = {
    "image_source" = "${var.image_source}"
    "image_checksum" = "${var.image_checksum}"
    "root_gb" = "${var.root_gb}"
    "root_device" = "${var.root_device}"
  }

  driver = "ipmi"
  driver_info  = "${var.nodes[0].driver_info}"
}

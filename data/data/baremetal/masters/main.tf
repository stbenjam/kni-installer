# FIXME: This Terraform HCL file defines the 3 master nodes.  It uses the original ironic_nodes.json format,
# flattened because Terraform v0.11 does not support nested data structures.  Maps may only be key/value
# pairs.  We could use terraform's resource `count` provider to have just one resource declaration, but
# the data would have to be structured differently.

resource "ironic_node_v1" "openshift-master-0" {
  name = "${var.nodes["nodes_0_name"]}"

  target_provision_state = "active"
  user_data              = "${var.ignition}"

  ports = [
    {
      address     = "${var.nodes["nodes_0_ports_0_address"]}"
      pxe_enabled = "${var.nodes["nodes_0_ports_0_pxe_enabled"]}"
    },
  ]

  properties {
    local_gb = "${var.nodes["nodes_0_properties_local_gb"]}"
    cpu_arch = "${var.nodes["nodes_0_properties_cpu_arch"]}"
  }

  instance_info = {
    image_source   = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb        = "${var.root_gb}"
    root_device    = "${var.root_disk}"
  }

  driver = "ipmi"

  driver_info {
    ipmi_port      = "${var.nodes["nodes_0_driver_info_ipmi_port"]}"
    ipmi_username  = "${var.nodes["nodes_0_driver_info_ipmi_username"]}"
    ipmi_password  = "${var.nodes["nodes_0_driver_info_ipmi_password"]}"
    ipmi_address   = "${var.nodes["nodes_0_driver_info_ipmi_address"]}"
    deploy_kernel  = "${var.nodes["nodes_0_driver_info_deploy_kernel"]}"
    deploy_ramdisk = "${var.nodes["nodes_0_driver_info_deploy_ramdisk"]}"
  }
}

resource "ironic_node_v1" "openshift-master-1" {
  name = "${var.nodes["nodes_1_name"]}"

  target_provision_state = "active"
  user_data              = "${var.ignition}"

  ports = [
    {
      address     = "${var.nodes["nodes_1_ports_0_address"]}"
      pxe_enabled = "${var.nodes["nodes_1_ports_0_pxe_enabled"]}"
    },
  ]

  properties {
    local_gb = "${var.nodes["nodes_1_properties_local_gb"]}"
    cpu_arch = "${var.nodes["nodes_1_properties_cpu_arch"]}"
  }

  instance_info = {
    image_source   = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb        = "${var.root_gb}"
    root_device    = "${var.root_disk}"
  }

  driver = "ipmi"

  driver_info {
    ipmi_port      = "${var.nodes["nodes_1_driver_info_ipmi_port"]}"
    ipmi_username  = "${var.nodes["nodes_1_driver_info_ipmi_username"]}"
    ipmi_password  = "${var.nodes["nodes_1_driver_info_ipmi_password"]}"
    ipmi_address   = "${var.nodes["nodes_1_driver_info_ipmi_address"]}"
    deploy_kernel  = "${var.nodes["nodes_1_driver_info_deploy_kernel"]}"
    deploy_ramdisk = "${var.nodes["nodes_1_driver_info_deploy_ramdisk"]}"
  }
}

resource "ironic_node_v1" "openshift-master-2" {
  name = "${var.nodes["nodes_2_name"]}"

  target_provision_state = "active"
  user_data              = "${var.ignition}"

  ports = [
    {
      address     = "${var.nodes["nodes_2_ports_0_address"]}"
      pxe_enabled = "${var.nodes["nodes_2_ports_0_pxe_enabled"]}"
    },
  ]

  properties {
    local_gb = "${var.nodes["nodes_2_properties_local_gb"]}"
    cpu_arch = "${var.nodes["nodes_2_properties_cpu_arch"]}"
  }

  instance_info = {
    image_source   = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb        = "${var.root_gb}"
    root_device    = "${var.root_disk}"
  }

  driver = "ipmi"

  driver_info {
    ipmi_port      = "${var.nodes["nodes_2_driver_info_ipmi_port"]}"
    ipmi_username  = "${var.nodes["nodes_2_driver_info_ipmi_username"]}"
    ipmi_password  = "${var.nodes["nodes_2_driver_info_ipmi_password"]}"
    ipmi_address   = "${var.nodes["nodes_2_driver_info_ipmi_address"]}"
    deploy_kernel  = "${var.nodes["nodes_2_driver_info_deploy_kernel"]}"
    deploy_ramdisk = "${var.nodes["nodes_2_driver_info_deploy_ramdisk"]}"
  }
}

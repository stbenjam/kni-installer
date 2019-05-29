resource "ironic_node_v1" "control-plane-host" {
  count = "${length(var.control_plane)}"
  name = "${lookup(var.control_plane["$count.index"], "name")}"
  resource_class = "baremetal"

  inspect = true
  clean = true
  available = true

  ports = [
    {
      address = "${lookup(var.control_plane["$count.index"], "port_address")}"
      pxe_enabled = "true"
    },
  ]

  properties = "${lookup(var.control_plane["$count.index"], "properties")}"
  root_device = "${lookup(var.control_plane["$count.index"], "root_device")}"

  driver = "${lookup(var.control_plane["$count.index"], "driver")}"
  driver_info = "${lookup(var.control_plane["$count.index"], "driver_info")}"

  vendor_interface = "no-vendor"
}

resource "ironic_allocation_v1" "control-plane-allocation" {
  name = "master-${count.index}"
  count = 3
  resource_class = "baremetal"

  candidate_nodes = [
    "${ironic_node_v1.control-plane-host.*.id}",
  ]
}

resource "ironic_deployment" "control-plane-deployment" {
  count = 3
  node_uuid = "${element(ironic_allocation_v1.control-plane-allocation.*.node_uuid, count.index)}"

  instance_info = {
    image_source = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb = "${var.root_gb}"
    address = "${lookup(var.control_plane["$count.index"], "root_gb")}"
  }

  user_data = "${var.ignition}"
}

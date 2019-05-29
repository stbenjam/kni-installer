resource "ironic_node_v1" "control-plane-host" {
  count = "${length(var.control_plane)}"
  name = "${element(var.control_plane.*.name, count.index)}"
  resource_class = "baremetal"

  inspect = true
  clean = true
  available = true

  ports = [
    {
      address = "${element(var.control_plane.*.port_address, count.index)}"
      pxe_enabled = "true"
    },
  ]

  properties = "${element(var.control_plane.*.properties, count.index)}"
  root_device = "${element(var.control_plane.*.root_device, count.index)}"

  driver = "${element(var.control_plane.*.driver, count.index)}"
  driver_info = "${element(var.control_plane.*.driver_info, count.index)}"

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
    address = "${element(var.control_plane.*.root_gb, count.index)}"
  }

  user_data = "${var.ignition}"
}

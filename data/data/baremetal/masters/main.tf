# This null resource is used to wait for the Ironic API to become available before trying to create
# the master hosts.
resource "null_resource" "api-wait" {
  provisioner "local-exec" {
    command = "timeout 900 bash -c 'while [[ \"$(curl -s -o /dev/null -w \"%%{http_code}\" ${var.ironic_uri})\" != \"200\" ]]; do sleep 5; done'"
  }
}

resource "ironic_node_v1" "openshift-master-host" {
  count          = var.master_count
  name           = var.hosts[count.index]["name"]
  resource_class = "baremetal"

  inspect   = true
  clean     = true
  available = true

  ports = [
    {
      address     = var.hosts[count.index]["port_address"]
      pxe_enabled = "true"
    },
  ]

  properties  = var.properties[count.index]
  root_device = var.root_devices[count.index]

  driver      = var.hosts[count.index]["driver"]
  driver_info = var.driver_infos[count.index]
  depends_on  = [
	null_resource.api-wait,
  ]
}

resource "ironic_allocation_v1" "openshift-master-allocation" {
  name           = "master-${count.index}"
  count          = var.master_count
  resource_class = "baremetal"

  candidate_nodes = ironic_node_v1.openshift-master-host.*.id
}

resource "ironic_deployment" "openshift-master-deployment" {
  count = var.master_count
  node_uuid = element(
    ironic_allocation_v1.openshift-master-allocation.*.node_uuid,
    count.index,
  )

  instance_info = var.instance_infos[count.index]
  user_data     = var.ignition
}


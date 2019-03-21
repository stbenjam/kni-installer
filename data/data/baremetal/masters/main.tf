# Possibly refactor to be based on counts https://blog.gruntwork.io/terraform-tips-tricks-loops-if-statements-and-gotchas-f739bbae55f9

resource "ironic_node_v1" "openshift-master-0" {
  name = "$name"

  target_provision_state = "active"
  user_data = "${var.ignition_master}"

  ports = [
    {
      "address" = "${"
      "pxe_enabled" = "true"
    }
  ]

  properties {
    "local_gb" = "${local_gb}"
    "cpu_arch" =  "${cpu_arch}"
  }

  instance_info = {
    "image_source" = "${image_source}"
    "image_checksum" = "${image_checksum}"
    "root_gb" = "${root_gb}"
    "root_device" = "${root_device}"
  }

  driver = "ipmi"
  driver_info {
    ${ipmi_port}
    "ipmi_username"=  "${ipmi_username}"
    "ipmi_password"=  "${ipmi_password}"
    "ipmi_address"=   "${ipmi_address}"
    "deploy_kernel"=  "${deploy_kernel}"
    "deploy_ramdisk"= "${deploy_ramdisk}"
  }
}


resource "ironic_node_v1" "openshift-master-1" {
  name = "$name"

  target_provision_state = "active"
  user_data = "\${file("master.ign")}"

  ports = [
    {
      "address" = "${mac}"
      "pxe_enabled" = "true"
    }
  ]

  properties {
    "local_gb" = "${local_gb}"
    "cpu_arch" =  "${cpu_arch}"
  }

  instance_info = {
    "image_source" = "${image_source}"
    "image_checksum" = "${image_checksum}"
    "root_gb" = "${root_gb}"
    "root_device" = "${root_device}"
  }

  driver = "ipmi"
  driver_info {
    ${ipmi_port}
    "ipmi_username"=  "${ipmi_username}"
    "ipmi_password"=  "${ipmi_password}"
    "ipmi_address"=   "${ipmi_address}"
    "deploy_kernel"=  "${deploy_kernel}"
    "deploy_ramdisk"= "${deploy_ramdisk}"
  }
}


resource "ironic_node_v1" "openshift-master-2" {
  name = "$name"

  target_provision_state = "active"
  user_data = "\${file("master.ign")}"

  ports = [
    {
      "address" = "${mac}"
      "pxe_enabled" = "true"
    }
  ]

  properties {
    "local_gb" = "${local_gb}"
    "cpu_arch" =  "${cpu_arch}"
  }

  instance_info = {
    "image_source" = "${image_source}"
    "image_checksum" = "${image_checksum}"
    "root_gb" = "${root_gb}"
    "root_device" = "${root_device}"
  }

  driver = "ipmi"
  driver_info {
    ${ipmi_port}
    "ipmi_username"=  "${ipmi_username}"
    "ipmi_password"=  "${ipmi_password}"
    "ipmi_address"=   "${ipmi_address}"
    "deploy_kernel"=  "${deploy_kernel}"
    "deploy_ramdisk"= "${deploy_ramdisk}"
  }
}

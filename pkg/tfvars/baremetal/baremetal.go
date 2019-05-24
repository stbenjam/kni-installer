// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"github.com/metal3-io/baremetal-operator/pkg/bmc"
	"github.com/metal3-io/baremetal-operator/pkg/hardware"
	libvirttfvars "github.com/openshift-metalkube/kni-installer/pkg/tfvars/libvirt"
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	"github.com/pkg/errors"
	"github.com/rodaine/hclencoder"
)

type config struct {
	LibvirtURI         string `hcl:"libvirt_uri,omitempty"`
	IronicURI          string `hcl:"ironic_uri,omitempty"`
	Image              string `hcl:"os_image,omitempty"`
	ExternalBridge     string `hcl:"external_bridge,omitempty"`
	ProvisioningBridge string `hcl:"provisioning_bridge,omitempty"`

	ControlPlane []map[string]interface{} `hcl:"control_plane"`
	Workers      []map[string]interface{} `hcl:"workers"`

	MasterConfiguration map[string]interface{} `hcl:"master_configuration"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, ironicURI, osImage, externalBridge, provisioningBridge string, hosts *[]baremetal.Host, configuration map[string]interface{}) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	var controlPlane []map[string]interface{}
	var workers []map[string]interface{}

	for _, host := range *hosts {
		// Hardware profile
		profile, err := hardware.GetProfile(host.HardwareProfile)
		if err != nil {
			return nil, err
		}

		// BMC Driver Info
		accessDetails, _ := bmc.NewAccessDetails(host.BMC.Address)
		credentials := bmc.Credentials{
			Username: host.BMC.Username,
			Password: host.BMC.Password,
		}
		driverInfo := accessDetails.DriverInfo(credentials)
		driverInfo["deploy_kernel"] = host.DeployKernel
		driverInfo["deploy_ramdisk"] = host.DeployRamdisk

		hostMap := map[string]interface{}{
			"name":         host.Name,
			"port_address": host.BootMACAddress,
			"properties": map[string]interface{}{
				"local_gb": profile.LocalGB,
				"cpu_arch": profile.CPUArch,
			},
			"root_device": profile.RootDeviceHints,
			"root_gb":     profile.RootGB,
			"driver":      accessDetails.Type(),
			"driver_info": driverInfo,
		}

		if host.Role == "master" {
			controlPlane = append(controlPlane, hostMap)
		} else {
			workers = append(workers, hostMap)
		}
	}

	cfg := &config{
		LibvirtURI:          libvirtURI,
		IronicURI:           ironicURI,
		Image:               osImage,
		ExternalBridge:      externalBridge,
		ProvisioningBridge:  provisioningBridge,
		ControlPlane:        controlPlane,
		Workers:             workers,
		MasterConfiguration: configuration,
	}

	return hclencoder.Encode(cfg)
}

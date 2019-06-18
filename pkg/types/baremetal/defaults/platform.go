package defaults

import (
	"fmt"
	"net"

	"github.com/openshift-metalkube/kni-installer/pkg/types"
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
)

const (
	LibvirtURI         = "qemu:///system"
	IronicURI          = "http://localhost:6385/v1"
	ExternalBridge     = "baremetal"
	ProvisioningBridge = "provisioning"
	HardwareProfile    = "default"
	ApiVIP = ""
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *baremetal.Platform, c *types.InstallConfig) {
	if p.LibvirtURI == "" {
		p.LibvirtURI = LibvirtURI
	}

	if p.IronicURI == "" {
		p.IronicURI = IronicURI
	}

	if p.ExternalBridge == "" {
		p.ExternalBridge = ExternalBridge
	}

	if p.ProvisioningBridge == "" {
		p.ProvisioningBridge = ProvisioningBridge
	}

	for _, host := range p.Hosts {
		if host.HardwareProfile == "" {
			host.HardwareProfile = HardwareProfile
		}
	}

	if p.ApiVIP == ApiVIP {
		// This name should resolve to exactly one address
		vip, err := net.LookupHost("api." + c.ClusterDomain())
		if err != nil {
			// This will fail validation and abort the install
			p.ApiVIP = fmt.Sprintf("DNS lookup failure: %s", err.Error())
		} else {
			p.ApiVIP = vip[0]
		}
	}
}

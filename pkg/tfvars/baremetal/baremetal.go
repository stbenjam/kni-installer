// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"

	libvirttfvars "github.com/openshift-metalkube/kni-installer/pkg/tfvars/libvirt"
	"github.com/pkg/errors"
)

type config struct {
	LibvirtURI      string `json:"libvirt_uri,omitempty"`
	IronicURI       string `json:"ironic_uri,omitempty"`
	Image           string `json:"os_image,omitempty"`
	BareMetalBridge string `json:"baremetal_bridge,omitempty"`
	OverCloudBridge string `json:"overcloud_bridge,omitempty"`

	// Data required for masters deployment
	Nodes               map[string]interface{}      `json:"nodes,omitempty"`
	MasterConfiguration baremetal.MasterConfiguration `json:"master_configuration"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, ironicURI, osImage, baremetalBridge, overcloudBridge string, nodes map[string]interface{}, configuration baremetal.MasterConfiguration) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	cfg := &config{
		LibvirtURI:          libvirtURI,
		IronicURI:           ironicURI,
		Image:               osImage,
		BareMetalBridge:     baremetalBridge,
		OverCloudBridge:     overcloudBridge,
		Nodes:               nodes,
		MasterConfiguration: configuration,
	}

	return json.MarshalIndent(cfg, "", "  ")
}

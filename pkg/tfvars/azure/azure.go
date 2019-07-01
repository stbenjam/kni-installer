package azure

import (
	"encoding/json"

	"github.com/openshift-metalkube/kni-installer/pkg/types/azure/defaults"
	azureprovider "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

// Auth is the collection of credentials that will be used by terrform.
type Auth struct {
	SubscriptionID string `json:"azure_subscription_id,omitempty"`
	ClientID       string `json:"azure_client_id,omitempty"`
	ClientSecret   string `json:"azure_client_secret,omitempty"`
	TenantID       string `json:"azure_tenant_id,omitempty"`
}

type config struct {
	Auth                        `json:",inline"`
	ExtraTags                   map[string]string `json:"azure_extra_tags,omitempty"`
	BootstrapInstanceType       string            `json:"azure_bootstrap_vm_type,omitempty"`
	MasterInstanceType          string            `json:"azure_master_vm_type,omitempty"`
	VolumeSize                  int32             `json:"azure_master_root_volume_size,omitempty"`
	VMImageID                   string            `json:"azure_image_id,omitempty"`
	Region                      string            `json:"azure_region,omitempty"`
	BaseDomainResourceGroupName string            `json:"azure_base_domain_resource_group_name,omitempty"`
}

// TFVars generates Azure-specific Terraform variables launching the cluster.
func TFVars(auth Auth, baseDomainResourceGroupName string, masterConfigs []*azureprovider.AzureMachineProviderSpec) ([]byte, error) {
	masterConfig := masterConfigs[0]
	region := masterConfig.Location
	cfg := &config{
		Auth:                        auth,
		Region:                      region,
		BaseDomainResourceGroupName: baseDomainResourceGroupName,
		BootstrapInstanceType:       defaults.BootstrapInstanceType(region),
		MasterInstanceType:          masterConfig.VMSize,
		VolumeSize:                  masterConfig.OSDisk.DiskSizeGB,
		VMImageID:                   masterConfig.Image.ResourceID,
	}

	return json.MarshalIndent(cfg, "", "  ")
}

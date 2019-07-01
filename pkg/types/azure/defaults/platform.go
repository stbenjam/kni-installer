package defaults

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types/azure"
)

var (
	// Overrides
	defaultMachineClass = map[string]string{}
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *azure.Platform) {
}

// getInstanceClass returns the instance "class" we should use for a given region.
func getInstanceClass(region string) string {
	if class, ok := defaultMachineClass[region]; ok {
		return class
	}

	return "Standard"
}

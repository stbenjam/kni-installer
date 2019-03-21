package baremetal

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// LibvirtURI is the identifier for the libvirtd connection.  It must be
	// reachable from the host where the installer is run.
	// +optional
	// Default is qemu:///system
	LibvirtURI string `json:"libvirt_uri,omitempty"`

	// IronicURI is the identifier for the Ironic connection.  It must be
	// reachable from the host where the installer is run.
	// +optional
	IronicURI string `json:"ironic_uri,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on bare metal for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`
}

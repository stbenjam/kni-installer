package aws

import "github.com/openshift-metalkube/kni-installer/pkg/destroy"

func init() {
	destroy.Registry["aws"] = New
}

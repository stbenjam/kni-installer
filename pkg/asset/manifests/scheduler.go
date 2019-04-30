package manifests

import (
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset"
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	schedulerCfgFilename = filepath.Join(manifestDir, "cluster-scheduler-02-config.yml")
)

// Scheduler generates the cluster-scheduler-*.yml files.
type Scheduler struct {
	FileList []*asset.File
}

var _ asset.WritableAsset = (*Scheduler)(nil)

// Name returns a human friendly name for the asset.
func (*Scheduler) Name() string {
	return "Scheduler Config"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*Scheduler) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate generates the scheduler config and its CRD.
func (s *Scheduler) Generate(dependencies asset.Parents) error {
	config := &configv1.Scheduler{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.SchemeGroupVersion.String(),
			Kind:       "Scheduler",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
			// not namespaced
		},
		Spec: configv1.SchedulerSpec{},
	}

	configData, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", s.Name())
	}

	s.FileList = []*asset.File{
		{
			Filename: schedulerCfgFilename,
			Data:     configData,
		},
	}

	return nil
}

// Files returns the files generated by the asset.
func (s *Scheduler) Files() []*asset.File {
	return s.FileList
}

// Load returns false since this asset is not written to disk by the installer.
func (s *Scheduler) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}

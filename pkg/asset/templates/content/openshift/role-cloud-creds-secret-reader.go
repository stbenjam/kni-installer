package openshift

import (
	"os"
	"path/filepath"

	"github.com/openshift-metalkube/kni-installer/pkg/asset"
	"github.com/openshift-metalkube/kni-installer/pkg/asset/templates/content"
)

const (
	roleCloudCredsSecretReaderFileName = "role-cloud-creds-secret-reader.yaml.template"
)

var _ asset.WritableAsset = (*RoleCloudCredsSecretReader)(nil)

// RoleCloudCredsSecretReader is the variable to represent contents of corresponding file
type RoleCloudCredsSecretReader struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *RoleCloudCredsSecretReader) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *RoleCloudCredsSecretReader) Name() string {
	return "RoleCloudCredsSecretReader"
}

// Generate generates the actual files by this asset
func (t *RoleCloudCredsSecretReader) Generate(parents asset.Parents) error {
	fileName := roleCloudCredsSecretReaderFileName
	data, err := content.GetOpenshiftTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *RoleCloudCredsSecretReader) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *RoleCloudCredsSecretReader) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, roleCloudCredsSecretReaderFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}

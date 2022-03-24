package api

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	"github.com/stretchr/testify/require"
)

func (suite *apiTestSuite) TestCreateCustomResourceQuota() {
	assert := require.New(suite.T())
	path := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName,
		constants.NamespacePath, "testproject", "kustomization.yaml",
	)
	// should fail if no namespace is present when attempting to append the kustomization
	err := suite.api.CreateCustomResourceQuota("test", models.CustomResourceQuota{}, true)
	assert.EqualError(err, "directory should already have been created for the namespace. directory test does not exist")

	// creates namespace directory
	err = os.MkdirAll(filepath.Dir(path), 0755)
	assert.Nil(err)

	// should fail because kustomization file in namespace directory does not exist, and so cannot append resource quota value to it
	err = suite.api.CreateCustomResourceQuota("testproject", models.CustomResourceQuota{}, true)
	assert.EqualError(err, "failed to append ResourceQuota resource to kustomization file")

	// should succeed
	kustom := models.NewKustomization([]string{"namespace.yaml"}, []string{"../../../../components/limitranges/default", "../../../../components/project-admin-rolebindings/testgroup"}, "testgroup")
	err = utils.WriteKustomization(filepath.Dir(path), kustom)
	assert.Nil(err)

}

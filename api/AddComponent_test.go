package api

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (suite *apiTestSuite) TestAddComponent() {
	assert := require.New(suite.T())

	// Should fail if project does not exist
	err := suite.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	err = suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
	)
	assert.Nil(err)

	// Should fail if component does not exist
	err = suite.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.EqualError(err, "component testcomponent does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		suite.dir, suite.api.AppName, constants.ComponentPath, "testcomponent",
	), 0755)
	assert.Nil(err)

	// Should succeed
	err = suite.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/AddComponent", suite.dir, expectedPaths)
}

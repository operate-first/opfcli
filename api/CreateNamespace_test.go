package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateNamespace() {
	assert := require.New(suite.T())

	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", suite.dir, expectedPaths)
}

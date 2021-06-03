package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateRoleBinding() {
	assert := require.New(suite.T())

	err := suite.api.CreateRoleBinding(
		"testgroup",
		"admin",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateRoleBinding", suite.dir, expectedPaths)
}

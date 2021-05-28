package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateProject() {
	assert := require.New(suite.T())

	err := suite.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
	}

	compareWithExpected(assert, "testdata/CreateProject", suite.dir, expectedPaths)
}

package api

import (
	"path/filepath"

	"github.com/stretchr/testify/require"
)

func (suite *apiTestSuite) TestAddGroupRBAC() {
	assert := require.New(suite.T())

	// Should fail if role does not exist
	err := suite.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"testrole",
	)
	assert.EqualError(err, "role testrole does not exist")

	// ---

	// Should fail if project does not exist
	err = suite.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	// Should fail if group does not exist
	err = suite.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
	)
	assert.Nil(err)

	err = suite.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.EqualError(err, "group testgroup2 does not exist")

	// ---

	// Should work if both project and group exist
	err = suite.api.CreateGroup(
		"testgroup2",
		false,
	)
	assert.Nil(err)
	assert.FileExists(filepath.Join(
		suite.dir,
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
	))

	err = suite.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/rbac.yaml",
	}

	compareWithExpected(assert, "testdata/AddGroupRBAC", suite.dir, expectedPaths)
}

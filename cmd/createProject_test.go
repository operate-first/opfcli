package cmd

import (
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestCreateProjectCmd() {
	assert := require.New(ctx.T())

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err := rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
	}

	compareWithExpected(assert, "testdata/createProject", ctx.dir, expectedPaths)

	// ---

	// Should fail if project already exists
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err = rootCmd.Execute()
	assert.EqualError(err, "namespace testproject already exists")
}

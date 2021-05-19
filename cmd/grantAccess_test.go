package cmd

import (
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestGrantAccessCmd() {
	assert := require.New(ctx.T())

	// Should fail if project does not exist
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	err := rootCmd.Execute()
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err = rootCmd.Execute()
	assert.Nil(err)

	// ---

	// Should fail if target group does not exist
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	err = rootCmd.Execute()
	assert.EqualError(err, "group testgroup2 does not exist")

	// ---

	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "create-group", "testgroup2"})
	err = rootCmd.Execute()
	assert.Nil(err)

	// ---

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	err = rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/rbac.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
	}

	compareWithExpected(assert, "testdata/grantAccess", ctx.dir, expectedPaths)
}

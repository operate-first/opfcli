package cmd

import (
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestCreateGroupCmd() {
	assert := require.New(ctx.T())

	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "create-group", "testgroup"})
	err := rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/createGroup", ctx.dir, expectedPaths)

	// ---

	// Command should fail if group already exists
	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "create-group", "testgroup"})
	err = rootCmd.Execute()
	assert.EqualError(err, "group testgroup already exists")
}

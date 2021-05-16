package cmd

import (
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestEnableMonitoringCmd() {
	assert := require.New(ctx.T())

	// Should fail if namespace does not exist
	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "enable-monitoring", "testproject"})
	err := rootCmd.Execute()
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	// Create project
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err = rootCmd.Execute()
	assert.Nil(err)

	// ---

	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "enable-monitoring", "testproject"})
	err = rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/enableMonitoring", ctx.dir, expectedPaths)
}

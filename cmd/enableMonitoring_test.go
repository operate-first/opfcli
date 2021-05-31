package cmd

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (suite *commandTestSuite) TestEnableMonitoring() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdEnableMonitoring(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdEnableMonitoring(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 2")

	// ---

	// Should fail because namespace does not exist
	cmd = NewCmdEnableMonitoring(suite.api)
	cmd.SetArgs([]string{"arg1"})
	err = cmd.Execute()
	assert.EqualError(err, "namespace arg1 does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.NamespacePath,
		"arg1",
	), 0755)
	assert.Nil(err)

	// Should fail because component does not exist
	cmd = NewCmdEnableMonitoring(suite.api)
	cmd.SetArgs([]string{"arg1"})
	err = cmd.Execute()
	assert.EqualError(err, "component monitoring-rbac does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.ComponentPath,
		"monitoring-rbac",
	), 0755)
	assert.Nil(err)

	// Should fail due to missing kustomization.yaml
	cmd = NewCmdEnableMonitoring(suite.api)
	cmd.SetArgs([]string{"arg1"})
	err = cmd.Execute()
	assert.NotNil(err)
	assert.Contains(err.Error(), "kustomization.yaml: no such file or directory")
}

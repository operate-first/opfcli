package cmd

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (suite *commandTestSuite) TestCreateProjectBasic() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 2 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 2 arg(s), received 3")

	// ---

	// Should fail with unknown option
	cmd = NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should succeed
	cmd = NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2"})
	err = cmd.Execute()
	assert.Nil(err)

	// ---

	// Should fail because group already exists
	cmd = NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2"})
	err = cmd.Execute()
	assert.EqualError(err, "group arg2 already exists")
}

func (suite *commandTestSuite) TestCreateProjectQuota() {
	assert := require.New(suite.T())

	// Should fail because quota does not exist
	cmd := NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"-q", "testquota", "arg1", "arg2"})
	err := cmd.Execute()
	assert.EqualError(err, "quota testquota does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.ComponentPath,
		"resourcequotas",
		"testquota",
	), 0755)
	assert.Nil(err)

	// Should succeed
	cmd = NewCmdCreateProject(suite.api)
	cmd.SetArgs([]string{"-q", "testquota", "arg1", "arg2"})
	err = cmd.Execute()
	assert.Nil(err)
}

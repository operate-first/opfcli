package cmd

import (
	"testing"

	"github.com/operate-first/opfcli/api"
	"github.com/operate-first/opfcli/utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type commandTestSuite struct {
	suite.Suite
	dir string
	api *api.API
}

func TestCommands(t *testing.T) {
	utils.ConfigureLogging()
	suite.Run(t, new(commandTestSuite))
}

func (suite *commandTestSuite) SetupTest() {
	suite.dir = suite.T().TempDir()
	suite.api = api.New("cluster-scope", suite.dir)
}

func (suite *commandTestSuite) TestRoot() {
	assert := require.New(suite.T())

	// should fail with unknown command
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"testcommand"})
	err := cmd.Execute()
	assert.Contains(err.Error(), "unknown command")

	// ---

	// should fail with unknown option
	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"--failure"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// --repo-dir should set repo-dir
	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"--repo-dir", suite.dir})
	err = cmd.Execute()
	assert.Nil(err)
	val, err := cmd.Flags().GetString("repo-dir")
	assert.Nil(err)
	assert.Equal(val, suite.dir)

	// ---

	// --app-name should set app-name
	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"--app-name", "testname"})
	err = cmd.Execute()
	assert.Nil(err)
	val, err = cmd.Flags().GetString("app-name")
	assert.Nil(err)
	assert.Equal(val, "testname")

	// ---

	// Should succeed with known command
	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"version"})
	err = cmd.Execute()
	assert.Nil(err)
}

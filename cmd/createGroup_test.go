package cmd

import "github.com/stretchr/testify/require"

func (suite *commandTestSuite) TestCreateGroup() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdCreateGroup(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdCreateGroup(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 2")

	// ---

	// Should fail with unknown option
	cmd = NewCmdCreateGroup(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should succeed
	cmd = NewCmdCreateGroup(suite.api)
	cmd.SetArgs([]string{"arg1"})
	err = cmd.Execute()
	assert.Nil(err)
}

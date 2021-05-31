package cmd

import "github.com/stretchr/testify/require"

func (suite *commandTestSuite) TestGrantAccess() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdGrantAccess(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 3 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdGrantAccess(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3", "arg4"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 3 arg(s), received 4")

	// ---

	// Should fail with unknown option
	cmd = NewCmdGrantAccess(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2", "arg3", "arg4"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should fail because role is missing
	cmd = NewCmdGrantAccess(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "role arg3 does not exist")
}

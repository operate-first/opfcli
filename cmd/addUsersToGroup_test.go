package cmd

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	"github.com/stretchr/testify/require"
)

func (suite *commandTestSuite) TestAddUsersToGroup() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdAddUsersToGroup(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 2 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdAddUsersToGroup(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 2 arg(s), received 5")

	// ---

	// Should fail with unknown option
	cmd = NewCmdAddUsersToGroup(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should fail if group does not exist
	cmd = NewCmdAddUsersToGroup(suite.api)
	cmd.SetArgs([]string{"arg1", "user1, user2, user3"})
	err = cmd.Execute()
	assert.EqualError(err, "group arg1 does not exist")

	//  create group

	err = suite.api.CreateGroup(
		"testgroup",
		[]string{},
		true,
	)
	assert.Nil(err)

	// create core common kustomization
	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	err = os.MkdirAll(commonOverlayPath, 0755)
	assert.Nil(err)
	kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
	err = utils.WriteKustomization(commonOverlayPath, kustom)
	assert.Nil(err)

	cmd = NewCmdAddUsersToGroup(suite.api)
	cmd.SetArgs([]string{"testgroup", "user1, user2, user3"})
	err = cmd.Execute()
	assert.Nil(err)

}

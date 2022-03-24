package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	"github.com/stretchr/testify/require"
)

func (suite *commandTestSuite) TestOnboardWithQuota() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 3")

	// ---

	// Should fail with unknown option
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should fail with quota option used but no quota
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, "quota testquota does not exist")

	err = os.MkdirAll(filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.ComponentPath,
		"resourcequotas",
		"testquota",
	), 0755)
	assert.Nil(err)

	// Should fail if lacks 2 overlay and common kustomizations
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in cluster overlay path: %s/%s/%s/%s/%s/%s", suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug", "kustomization.yaml"))

	prodOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug")
	err = os.MkdirAll(prodOverlayPath, 0755)
	assert.Nil(err)
	kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
	err = utils.WriteKustomization(prodOverlayPath, kustom)
	assert.Nil(err)

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	err = os.MkdirAll(commonOverlayPath, 0755)
	assert.Nil(err)
	kustom = models.NewKustomization([]string{"../../../base"}, nil, "")
	err = utils.WriteKustomization(commonOverlayPath, kustom)
	assert.Nil(err)

	// Should succeed if group and kustomization files already exists
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithQuota.yaml"})
	err = cmd.Execute()
	assert.Nil(err)
}

func (suite *commandTestSuite) TestOnboardWithoutQuota() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 3")

	// ---

	// Should fail with unknown option
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithoutQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in cluster overlay path: %s/%s/%s/%s/%s/%s", suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug", "kustomization.yaml"))

	prodOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug")
	err = os.MkdirAll(prodOverlayPath, 0755)
	assert.Nil(err)
	kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
	err = utils.WriteKustomization(prodOverlayPath, kustom)
	assert.Nil(err)

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithoutQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in common overlay path: %s/%s/%s/%s", suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath, "kustomization.yaml"))

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	err = os.MkdirAll(commonOverlayPath, 0755)
	assert.Nil(err)
	kustom = models.NewKustomization([]string{"../../../base"}, nil, "")
	err = utils.WriteKustomization(commonOverlayPath, kustom)
	assert.Nil(err)

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithoutQuota.yaml"})
	err = cmd.Execute()
	assert.Nil(err)
}

func (suite *commandTestSuite) TestOnboardWithCustomQuota() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 1 arg(s), received 3")

	// ---

	// Should fail with unknown option
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithCustomQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in cluster overlay path: %s/%s/%s/%s/%s/%s", suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug", "kustomization.yaml"))

	prodOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug")
	err = os.MkdirAll(prodOverlayPath, 0755)
	assert.Nil(err)
	kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
	err = utils.WriteKustomization(prodOverlayPath, kustom)
	assert.Nil(err)

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithCustomQuota.yaml"})
	err = cmd.Execute()
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in common overlay path: %s/%s/%s/%s", suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath, "kustomization.yaml"))

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	err = os.MkdirAll(commonOverlayPath, 0755)
	assert.Nil(err)
	kustom = models.NewKustomization([]string{"../../../base"}, nil, "")
	err = utils.WriteKustomization(commonOverlayPath, kustom)
	assert.Nil(err)

	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"testdata/OnboardConfigs/sampleOnboardConfigWithCustomQuota.yaml"})
	err = cmd.Execute()
	assert.Nil(err)
}

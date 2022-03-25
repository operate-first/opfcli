package api

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func (suite *apiTestSuite) TestAddUsersToGroup() {
	assert := require.New(suite.T())

	err := suite.api.AddUsersToGroup("testgroup", []string{"testuser1", "testuser2"})
	assert.EqualError(err, "group testgroup does not exist")

	err = suite.api.CreateGroup("testgroup", []string{"testuser1"}, true)
	assert.Nil(err)

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")

	err = suite.api.AddUsersToGroup("testgroup", []string{"testuser1", "testuser2"})
	assert.EqualError(err, fmt.Sprintf("kustomization does not exist in common overlay path: %s", commonOverlayKustomizationPath))

	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Common overlay kustomization must exist, to deploy group to cluster.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.AddUsersToGroup("testgroup", []string{"testuser2", "testuser3"})
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
	}

	compareWithExpected(assert, "testdata/AddUsersToGroup", suite.dir, expectedPaths)

}

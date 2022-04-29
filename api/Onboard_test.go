package api

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func (suite *apiTestSuite) TestOnboardWithoutQuota() {
	assert := require.New(suite.T())

	prodOverlayPath := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug",
	)

	prodOverlayPathExists, err := utils.PathExists(prodOverlayPath)
	assert.Nil(err)
	if !prodOverlayPathExists {
		log.Printf("Onboard needs moc/smaug overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(prodOverlayPath, 0755)
		assert.Nil(err)
	}
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	prodOverlayKustomizationExists, err := utils.PathExists(prodOverlayKustomizationPath)
	assert.Nil(err)
	if !prodOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
		err = utils.WriteKustomization(prodOverlayPath, kustom)
		assert.Nil(err)
	}

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Onboard needs common overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.Onboard(
		"./testdata/Onboard/sampleOnboardConfigWithoutQuota.yaml",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",

		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",

		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",

		"cluster-scope/overlays/prod/common/kustomization.yaml",
		"cluster-scope/overlays/prod/moc/smaug/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/Onboard/withoutQuota", suite.dir, expectedPaths)

}

func (suite *apiTestSuite) TestOnboardWithQuota() {

	assert := require.New(suite.T())

	quotaPath := filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.ComponentPath,
		"resourcequotas",
		"testquota",
		"resourcequotas.yaml",
	)

	quotaExists, err := utils.PathExists(filepath.Dir(quotaPath))
	assert.Nil(err)
	assert.False(quotaExists)
	if !quotaExists {
		log.Printf("Onboard tested with quota but doesn't exist. Creating for test purposes.")

		err = os.MkdirAll(filepath.Dir(quotaPath), 0755)
		assert.Nil(err)
		kustom := models.NewKomponent([]string{"resourcequotas.yaml"})
		err = utils.WriteKustomization(filepath.Dir(quotaPath), kustom)
		assert.Nil(err)

		byteArr := []byte("kind: ResourceQuota\napiVersion: v1\nmetadata:\n    name: testquota\nspec:\n    hard:\n        requests.cpu: '1'\n        \n        requests.memory: 4Gi\n        limits.cpu: '1'\n        limits.memory: 4Gi\n        requests.storage: 20Gi\n        count/objectbucketclaims.objectbucket.io: 1")
		err = ioutil.WriteFile(quotaPath, byteArr, 0644)
		assert.Nil(err)
	}

	prodOverlayPath := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug",
	)

	prodOverlayPathExists, err := utils.PathExists(prodOverlayPath)
	assert.Nil(err)
	if !prodOverlayPathExists {
		log.Printf("Onboard needs moc/smaug overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(prodOverlayPath, 0755)
		assert.Nil(err)
	}
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	prodOverlayKustomizationExists, err := utils.PathExists(prodOverlayKustomizationPath)
	assert.Nil(err)
	if !prodOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
		err = utils.WriteKustomization(prodOverlayPath, kustom)
		assert.Nil(err)
		// err = utils.PathExists()
	}

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Onboard needs common overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.Onboard(
		"./testdata/Onboard/sampleOnboardConfigWithQuota.yaml")
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",

		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",

		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",

		"cluster-scope/components/resourcequotas/testquota/kustomization.yaml",
		"cluster-scope/components/resourcequotas/testquota/resourcequotas.yaml",

		"cluster-scope/overlays/prod/common/kustomization.yaml",
		"cluster-scope/overlays/prod/moc/smaug/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/Onboard/withQuota", suite.dir, expectedPaths)

}

func (suite *apiTestSuite) TestOnboardWithCustomQuota() {

	assert := require.New(suite.T())

	prodOverlayPath := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug",
	)

	prodOverlayPathExists, err := utils.PathExists(prodOverlayPath)
	assert.Nil(err)
	if !prodOverlayPathExists {
		log.Printf("Onboard needs moc/smaug overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(prodOverlayPath, 0755)
		assert.Nil(err)
	}
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	prodOverlayKustomizationExists, err := utils.PathExists(prodOverlayKustomizationPath)
	assert.Nil(err)
	if !prodOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
		err = utils.WriteKustomization(prodOverlayPath, kustom)
		assert.Nil(err)
	}

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Onboard needs common overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.Onboard(
		"./testdata/Onboard/sampleOnboardConfigWithCustomQuota.yaml",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/core/namespaces/testproject/resourcequota.yaml",

		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",

		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",

		"cluster-scope/overlays/prod/common/kustomization.yaml",
		"cluster-scope/overlays/prod/moc/smaug/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/Onboard/withCustomQuota", suite.dir, expectedPaths)

}

func (suite *apiTestSuite) TestOnboardDisableLimitRangeAndDisplayName() {
	assert := require.New(suite.T())

	prodOverlayPath := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug",
	)

	prodOverlayPathExists, err := utils.PathExists(prodOverlayPath)
	assert.Nil(err)
	if !prodOverlayPathExists {
		log.Printf("Onboard needs moc/smaug overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(prodOverlayPath, 0755)
		assert.Nil(err)
	}
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	prodOverlayKustomizationExists, err := utils.PathExists(prodOverlayKustomizationPath)
	assert.Nil(err)
	if !prodOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
		err = utils.WriteKustomization(prodOverlayPath, kustom)
		assert.Nil(err)
	}

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Onboard needs common overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.Onboard(
		"./testdata/Onboard/sampleOnboardConfigDisableLimitRangeAndDisplayName.yaml",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",

		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",

		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",

		"cluster-scope/overlays/prod/common/kustomization.yaml",
		"cluster-scope/overlays/prod/moc/smaug/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/Onboard/disableLimitRangeAndDisplayName", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestOnboardWithProjectDetails() {
	assert := require.New(suite.T())

	prodOverlayPath := filepath.Join(
		suite.api.RepoDirectory, suite.api.AppName, constants.ProdOverlayPath, "moc", "smaug",
	)

	prodOverlayPathExists, err := utils.PathExists(prodOverlayPath)
	assert.Nil(err)
	if !prodOverlayPathExists {
		log.Printf("Onboard needs moc/smaug overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(prodOverlayPath, 0755)
		assert.Nil(err)
	}
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	prodOverlayKustomizationExists, err := utils.PathExists(prodOverlayKustomizationPath)
	assert.Nil(err)
	if !prodOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../../base"}, nil, "")
		err = utils.WriteKustomization(prodOverlayPath, kustom)
		assert.Nil(err)
	}

	commonOverlayPath := filepath.Join(suite.api.RepoDirectory, suite.api.AppName, constants.CommonOverlayPath)
	commonOverlayPathExists, err := utils.PathExists(commonOverlayPath)
	assert.Nil(err)
	if !commonOverlayPathExists {
		log.Printf("Onboard needs common overlay path for testing (based on sampleOnboardingConfig), as it is meant for use in operate-first/apps. Creating.")
		err = os.MkdirAll(commonOverlayPath, 0755)
		assert.Nil(err)
	}
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	commonOverlayKustomizationExists, err := utils.PathExists(commonOverlayKustomizationPath)
	assert.Nil(err)
	if !commonOverlayKustomizationExists {
		kustom := models.NewKustomization([]string{"../../../base"}, nil, "")
		err = utils.WriteKustomization(commonOverlayPath, kustom)
		assert.Nil(err)
	}

	err = suite.api.Onboard(
		"./testdata/Onboard/sampleOnboardConfigWithProjectDetails.yaml",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",

		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",

		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",

		"cluster-scope/overlays/prod/common/kustomization.yaml",
		"cluster-scope/overlays/prod/moc/smaug/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/Onboard/withProjectDetails", suite.dir, expectedPaths)

}

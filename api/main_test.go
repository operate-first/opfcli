package api

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/operate-first/opfcli/utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type apiTestSuite struct {
	suite.Suite
	dir string
	api *API
}

func TestAPI(t *testing.T) {
	utils.ConfigureLogging()
	suite.Run(t, new(apiTestSuite))
}

func (suite *apiTestSuite) SetupTest() {
	suite.dir = suite.T().TempDir()
	suite.api = New("cluster-scope", suite.dir)
}

func compareWithExpected(assert *require.Assertions, expectedRoot, actualRoot string, paths []string) {
	for _, path := range paths {
		expectedPath := filepath.Join(expectedRoot, path)
		actualPath := filepath.Join(actualRoot, path)

		assert.FileExists(actualPath)

		actualData, err := ioutil.ReadFile(actualPath)
		assert.Nil(err)

		expectedData, err := ioutil.ReadFile(expectedPath)
		assert.Nil(err)

		assert.YAMLEq(string(expectedData), string(actualData))
	}
}

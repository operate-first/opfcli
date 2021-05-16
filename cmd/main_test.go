package cmd

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Context struct {
	suite.Suite
	dir string
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Context))
}

func (ctx *Context) SetupTest() {
	ctx.dir = ctx.T().TempDir()
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

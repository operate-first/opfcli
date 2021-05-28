package api

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestAddComponent() {
	assert := require.New(ctx.T())

	// Should fail if project does not exist
	err := ctx.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	err = ctx.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
	)
	assert.Nil(err)

	// Should fail if component does not exist
	err = ctx.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.EqualError(err, "component testcomponent does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		ctx.dir, ctx.api.AppName, constants.ComponentPath, "testcomponent",
	), 0755)
	assert.Nil(err)

	// Should succeed
	err = ctx.api.AddComponent(
		"testproject",
		"testcomponent",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/AddComponent", ctx.dir, expectedPaths)
}

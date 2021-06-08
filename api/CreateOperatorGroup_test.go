package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateOperatorGroup() {
	assert := require.New(suite.T())

	err := suite.api.CreateOperatorGroup(
		"testproject",
		true,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/operators.coreos.com/operatorgroups/testproject/kustomization.yaml",
		"cluster-scope/base/operators.coreos.com/operatorgroups/testproject/operatorgroup.yaml",
	}

	compareWithExpected(assert, "testdata/CreateOperatorGroup", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateOperatorGroupSingleNamespace() {
	assert := require.New(suite.T())

	err := suite.api.CreateOperatorGroup(
		"testproject",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/operators.coreos.com/operatorgroups/testproject/kustomization.yaml",
		"cluster-scope/base/operators.coreos.com/operatorgroups/testproject/operatorgroup.yaml",
	}

	compareWithExpected(assert, "testdata/CreateOperatorGroupSingleNamespace", suite.dir, expectedPaths)
}

package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateGroup() {
	assert := require.New(suite.T())

	err := suite.api.CreateGroup(
		"testgroup",
		[]string{},
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateGroup/noUsers", suite.dir, expectedPaths)

	// ---

	// Should fail if group already exists and existsOk is false
	err = suite.api.CreateGroup(
		"testgroup",
		[]string{},
		false,
	)
	assert.EqualError(err, "group testgroup already exists")

	// ---

	// Should work if group already exists and existsOk is true
	err = suite.api.CreateGroup(
		"testgroup",
		[]string{},
		true,
	)
	assert.Nil(err)
}

func (suite *apiTestSuite) TestCreateGroupWithUsers() {
	assert := require.New(suite.T())

	err := suite.api.CreateGroup(
		"testgroup",
		[]string{"jared", "jason"},
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateGroup/Users", suite.dir, expectedPaths)

	// ---

	// Should fail if group already exists and existsOk is false
	err = suite.api.CreateGroup(
		"testgroup",
		[]string{"jared", "jason"},
		false,
	)
	assert.EqualError(err, "group testgroup already exists")

	// ---

	// Should work if group already exists and existsOk is true
	err = suite.api.CreateGroup(
		"testgroup",
		[]string{"jared", "jason"},
		true,
	)
	assert.Nil(err)
}

package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateSubscription() {
	assert := require.New(suite.T())

	err := suite.api.CreateSubscription(
		"testoperator",
		"testcatalog",
		"testproject",
		"",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/kustomization.yaml",
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/subscription.yaml",
	}

	compareWithExpected(assert, "testdata/CreateSubscription", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateSubscriptionManual() {
	assert := require.New(suite.T())

	err := suite.api.CreateSubscription(
		"testoperator",
		"testcatalog",
		"testproject",
		"",
		true,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/kustomization.yaml",
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/subscription.yaml",
	}

	compareWithExpected(assert, "testdata/CreateSubscriptionManual", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateSubscriptionChannel() {
	assert := require.New(suite.T())

	err := suite.api.CreateSubscription(
		"testoperator",
		"testcatalog",
		"testproject",
		"stable",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/kustomization.yaml",
		"cluster-scope/base/operators.coreos.com/subscriptions/testoperator/subscription.yaml",
	}

	compareWithExpected(assert, "testdata/CreateSubscriptionChannel", suite.dir, expectedPaths)
}

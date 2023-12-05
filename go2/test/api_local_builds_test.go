/*
Bitrise API

Testing LocalBuildsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/staffbase/bitrise-api"
)

func Test_openapi_LocalBuildsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LocalBuildsAPIService LocalBuildList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.LocalBuildsAPI.LocalBuildList(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
/*
Bitrise API

Testing UserAPIService

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

func Test_openapi_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService UserMachineTypeUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userSlug string

		resp, httpRes, err := apiClient.UserAPI.UserMachineTypeUpdate(context.Background(), userSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UserPlan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.UserPlan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UserProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.UserProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UserShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userSlug string

		resp, httpRes, err := apiClient.UserAPI.UserShow(context.Background(), userSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

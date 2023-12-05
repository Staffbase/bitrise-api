/*
Bitrise API

Testing AppSetupAPIService

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

func Test_openapi_AppSetupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppSetupAPIService AppConfigCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.AppSetupAPI.AppConfigCreate(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService AppCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppSetupAPI.AppCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService AppFinish", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.AppSetupAPI.AppFinish(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService AppSetupBitriseYmlConfigGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.AppSetupAPI.AppSetupBitriseYmlConfigGet(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService AppSetupBitriseYmlConfigUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		httpRes, err := apiClient.AppSetupAPI.AppSetupBitriseYmlConfigUpdate(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService AppWebhookCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.AppSetupAPI.AppWebhookCreate(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppSetupAPIService SshKeyCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appSlug string

		resp, httpRes, err := apiClient.AppSetupAPI.SshKeyCreate(context.Background(), appSlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
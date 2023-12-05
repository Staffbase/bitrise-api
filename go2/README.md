# Go API client for openapi

Official REST API for Bitrise.io

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.bitrise.io/contact](https://www.bitrise.io/contact)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/staffbase/bitrise-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.bitrise.io/v0.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivityAPI* | [**ActivityList**](docs/ActivityAPI.md#activitylist) | **Get** /me/activities | Get list of Bitrise activity events
*AddonsAPI* | [**AddonListByApp**](docs/AddonsAPI.md#addonlistbyapp) | **Get** /apps/{app-slug}/addons | Get list of the addons for apps
*AddonsAPI* | [**AddonListByOrganization**](docs/AddonsAPI.md#addonlistbyorganization) | **Get** /organizations/{organization-slug}/addons | Get list of the addons for organization
*AddonsAPI* | [**AddonListByUser**](docs/AddonsAPI.md#addonlistbyuser) | **Get** /users/{user-slug}/addons | Get list of the addons for user
*AddonsAPI* | [**AddonsList**](docs/AddonsAPI.md#addonslist) | **Get** /addons | Get list of available Bitrise addons
*AddonsAPI* | [**AddonsShow**](docs/AddonsAPI.md#addonsshow) | **Get** /addons/{addon-id} | Get a specific Bitrise addon
*AndroidKeystoreFileAPI* | [**AndroidKeystoreFileConfirm**](docs/AndroidKeystoreFileAPI.md#androidkeystorefileconfirm) | **Post** /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug} | Confirm an android keystore file upload
*AndroidKeystoreFileAPI* | [**AndroidKeystoreFileCreate**](docs/AndroidKeystoreFileAPI.md#androidkeystorefilecreate) | **Post** /apps/{app-slug}/android-keystore-files | Create an Android keystore file
*AndroidKeystoreFileAPI* | [**AndroidKeystoreFileDelete**](docs/AndroidKeystoreFileAPI.md#androidkeystorefiledelete) | **Delete** /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug} | Delete an android keystore file
*AndroidKeystoreFileAPI* | [**AndroidKeystoreFileList**](docs/AndroidKeystoreFileAPI.md#androidkeystorefilelist) | **Get** /apps/{app-slug}/android-keystore-files | Get a list of the android keystore files
*AppSetupAPI* | [**AppConfigCreate**](docs/AppSetupAPI.md#appconfigcreate) | **Post** /apps/{app-slug}/bitrise.yml | Upload a new bitrise.yml for your application.
*AppSetupAPI* | [**AppCreate**](docs/AppSetupAPI.md#appcreate) | **Post** /apps/register | Add a new app
*AppSetupAPI* | [**AppFinish**](docs/AppSetupAPI.md#appfinish) | **Post** /apps/{app-slug}/finish | Save the application at the end of the app registration process
*AppSetupAPI* | [**AppSetupBitriseYmlConfigGet**](docs/AppSetupAPI.md#appsetupbitriseymlconfigget) | **Get** /apps/{app-slug}/bitrise.yml/config | Getting the location of the application&#39;s bitrise.yaml
*AppSetupAPI* | [**AppSetupBitriseYmlConfigUpdate**](docs/AppSetupAPI.md#appsetupbitriseymlconfigupdate) | **Put** /apps/{app-slug}/bitrise.yml/config | Changing the location of the application&#39;s bitrise.yaml
*AppSetupAPI* | [**AppWebhookCreate**](docs/AppSetupAPI.md#appwebhookcreate) | **Post** /apps/{app-slug}/register-webhook | Register an incoming webhook for a specific application
*AppSetupAPI* | [**SshKeyCreate**](docs/AppSetupAPI.md#sshkeycreate) | **Post** /apps/{app-slug}/register-ssh-key | Add an SSH-key to a specific app
*AppleApiCredentialsAPI* | [**AppleApiCredentialList**](docs/AppleApiCredentialsAPI.md#appleapicredentiallist) | **Get** /users/{user-slug}/apple-api-credentials | List Apple API credentials for a specific user
*ApplicationAPI* | [**AppConfigDatastoreShow**](docs/ApplicationAPI.md#appconfigdatastoreshow) | **Get** /apps/{app-slug}/bitrise.yml | Get bitrise.yml of a specific app
*ApplicationAPI* | [**AppDelete**](docs/ApplicationAPI.md#appdelete) | **Delete** /apps/{app-slug} | Deletes an app
*ApplicationAPI* | [**AppList**](docs/ApplicationAPI.md#applist) | **Get** /apps | Get list of the apps
*ApplicationAPI* | [**AppListByOrganization**](docs/ApplicationAPI.md#applistbyorganization) | **Get** /organizations/{org-slug}/apps | Get list of the apps for an organization
*ApplicationAPI* | [**AppListByUser**](docs/ApplicationAPI.md#applistbyuser) | **Get** /users/{user-slug}/apps | Get list of the apps for a user
*ApplicationAPI* | [**AppNotifications**](docs/ApplicationAPI.md#appnotifications) | **Patch** /apps/{app-slug}/update-email-notifications | Updates the app&#39;s notification settings
*ApplicationAPI* | [**AppRolesQuery**](docs/ApplicationAPI.md#approlesquery) | **Get** /apps/{app-slug}/roles/{role-name} | Lists group roles for an app
*ApplicationAPI* | [**AppRolesUpdate**](docs/ApplicationAPI.md#approlesupdate) | **Put** /apps/{app-slug}/roles/{role-name} | Replaces group roles for an app
*ApplicationAPI* | [**AppShow**](docs/ApplicationAPI.md#appshow) | **Get** /apps/{app-slug} | Get a specific app
*ApplicationAPI* | [**AppUpdate**](docs/ApplicationAPI.md#appupdate) | **Patch** /apps/{app-slug} | Updates an app
*ApplicationAPI* | [**BranchList**](docs/ApplicationAPI.md#branchlist) | **Get** /apps/{app-slug}/branches | List the branches with existing builds of an app&#39;s repository
*BuildArtifactAPI* | [**ArtifactDelete**](docs/BuildArtifactAPI.md#artifactdelete) | **Delete** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Delete a build artifact
*BuildArtifactAPI* | [**ArtifactList**](docs/BuildArtifactAPI.md#artifactlist) | **Get** /apps/{app-slug}/builds/{build-slug}/artifacts | Get a list of all build artifacts
*BuildArtifactAPI* | [**ArtifactShow**](docs/BuildArtifactAPI.md#artifactshow) | **Get** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Get a specific build artifact
*BuildArtifactAPI* | [**ArtifactUpdate**](docs/BuildArtifactAPI.md#artifactupdate) | **Patch** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Update a build artifact
*BuildCertificateAPI* | [**BuildCertificateConfirm**](docs/BuildCertificateAPI.md#buildcertificateconfirm) | **Post** /apps/{app-slug}/build-certificates/{build-certificate-slug}/uploaded | Confirm a build certificate upload
*BuildCertificateAPI* | [**BuildCertificateCreate**](docs/BuildCertificateAPI.md#buildcertificatecreate) | **Post** /apps/{app-slug}/build-certificates | Create a build certificate
*BuildCertificateAPI* | [**BuildCertificateDelete**](docs/BuildCertificateAPI.md#buildcertificatedelete) | **Delete** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Delete a build certificate
*BuildCertificateAPI* | [**BuildCertificateList**](docs/BuildCertificateAPI.md#buildcertificatelist) | **Get** /apps/{app-slug}/build-certificates | Get a list of the build certificates
*BuildCertificateAPI* | [**BuildCertificateShow**](docs/BuildCertificateAPI.md#buildcertificateshow) | **Get** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Get a specific build certificate
*BuildCertificateAPI* | [**BuildCertificateUpdate**](docs/BuildCertificateAPI.md#buildcertificateupdate) | **Patch** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Update a build certificate
*BuildRequestAPI* | [**BuildRequestList**](docs/BuildRequestAPI.md#buildrequestlist) | **Get** /apps/{app-slug}/build-requests | List the open build requests for an app
*BuildRequestAPI* | [**BuildRequestUpdate**](docs/BuildRequestAPI.md#buildrequestupdate) | **Patch** /apps/{app-slug}/build-requests/{build-request-slug} | Update a build request
*BuildsAPI* | [**ArchivedBuildsList**](docs/BuildsAPI.md#archivedbuildslist) | **Get** /apps/{app-slug}/archived-builds | List 1000 archived builds of an app
*BuildsAPI* | [**BuildAbort**](docs/BuildsAPI.md#buildabort) | **Post** /apps/{app-slug}/builds/{build-slug}/abort | Abort a specific build
*BuildsAPI* | [**BuildBitriseYmlShow**](docs/BuildsAPI.md#buildbitriseymlshow) | **Get** /apps/{app-slug}/builds/{build-slug}/bitrise.yml | Get the bitrise.yml of a build
*BuildsAPI* | [**BuildList**](docs/BuildsAPI.md#buildlist) | **Get** /apps/{app-slug}/builds | List all builds of an app
*BuildsAPI* | [**BuildListAll**](docs/BuildsAPI.md#buildlistall) | **Get** /builds | List all builds
*BuildsAPI* | [**BuildLog**](docs/BuildsAPI.md#buildlog) | **Get** /apps/{app-slug}/builds/{build-slug}/log | Get the build log of a build
*BuildsAPI* | [**BuildShow**](docs/BuildsAPI.md#buildshow) | **Get** /apps/{app-slug}/builds/{build-slug} | Get a build of a given app
*BuildsAPI* | [**BuildTrigger**](docs/BuildsAPI.md#buildtrigger) | **Post** /apps/{app-slug}/builds | Trigger a new build/pipeline
*BuildsAPI* | [**BuildWorkflowList**](docs/BuildsAPI.md#buildworkflowlist) | **Get** /apps/{app-slug}/build-workflows | List the workflows of an app
*GenericProjectFileAPI* | [**GenericProjectFileConfirm**](docs/GenericProjectFileAPI.md#genericprojectfileconfirm) | **Post** /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded | Confirm a generic project file upload
*GenericProjectFileAPI* | [**GenericProjectFileDelete**](docs/GenericProjectFileAPI.md#genericprojectfiledelete) | **Delete** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Delete a generic project file
*GenericProjectFileAPI* | [**GenericProjectFileList**](docs/GenericProjectFileAPI.md#genericprojectfilelist) | **Get** /apps/{app-slug}/generic-project-files | Get a list of the generic project files
*GenericProjectFileAPI* | [**GenericProjectFileShow**](docs/GenericProjectFileAPI.md#genericprojectfileshow) | **Get** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Get a specific generic project file
*GenericProjectFileAPI* | [**GenericProjectFileUpdate**](docs/GenericProjectFileAPI.md#genericprojectfileupdate) | **Patch** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Update a generic project file
*GenericProjectFileAPI* | [**GenericProjectFilesCreate**](docs/GenericProjectFileAPI.md#genericprojectfilescreate) | **Post** /apps/{app-slug}/generic-project-files | Create a generic project file
*KeyValueCacheAPI* | [**CacheItemDelete**](docs/KeyValueCacheAPI.md#cacheitemdelete) | **Delete** /apps/{app-slug}/cache-items/{cache-item-id} | Deletes a key-value cache item
*KeyValueCacheAPI* | [**CacheItemDeleteAll**](docs/KeyValueCacheAPI.md#cacheitemdeleteall) | **Delete** /apps/{app-slug}/cache-items | Deletes all key-value cache items belonging to an app
*KeyValueCacheAPI* | [**CacheItemDownload**](docs/KeyValueCacheAPI.md#cacheitemdownload) | **Get** /apps/{app-slug}/cache-items/{cache-item-id}/download | Gets the download URL of a key-value cache item
*KeyValueCacheAPI* | [**CacheList**](docs/KeyValueCacheAPI.md#cachelist) | **Get** /apps/{app-slug}/cache-items | List the key-value cache items belonging to an app
*LocalBuildsAPI* | [**LocalBuildList**](docs/LocalBuildsAPI.md#localbuildlist) | **Get** /apps/{app-slug}/local-builds | List local builds of an app
*OrganizationsAPI* | [**OrgList**](docs/OrganizationsAPI.md#orglist) | **Get** /organizations | List the organizations that the user is part of
*OrganizationsAPI* | [**OrgShow**](docs/OrganizationsAPI.md#orgshow) | **Get** /organizations/{org-slug} | Get a specified organization.
*OrganizationsAPI* | [**OrganizationMachineTypeUpdate**](docs/OrganizationsAPI.md#organizationmachinetypeupdate) | **Patch** /organizations/{org-slug}/apps/machine_types | Migrate machine types
*OrganizationsAPI* | [**OrganzationGroupsList**](docs/OrganizationsAPI.md#organzationgroupslist) | **Get** /organizations/{org-slug}/groups | List organizations groups
*OutgoingWebhookAPI* | [**OutgoingWebhookCreate**](docs/OutgoingWebhookAPI.md#outgoingwebhookcreate) | **Post** /apps/{app-slug}/outgoing-webhooks | Create an outgoing webhook for an app
*OutgoingWebhookAPI* | [**OutgoingWebhookDelete**](docs/OutgoingWebhookAPI.md#outgoingwebhookdelete) | **Delete** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug} | Delete an outgoing webhook of an app
*OutgoingWebhookAPI* | [**OutgoingWebhookList**](docs/OutgoingWebhookAPI.md#outgoingwebhooklist) | **Get** /apps/{app-slug}/outgoing-webhooks | List the outgoing webhooks of an app
*OutgoingWebhookAPI* | [**OutgoingWebhookUpdate**](docs/OutgoingWebhookAPI.md#outgoingwebhookupdate) | **Put** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug} | Update an outgoing webhook of an app
*PipelinesAPI* | [**PipelineAbort**](docs/PipelinesAPI.md#pipelineabort) | **Post** /apps/{app-slug}/pipelines/{pipeline-id}/abort | Aborts a pipeline
*PipelinesAPI* | [**PipelineList**](docs/PipelinesAPI.md#pipelinelist) | **Get** /apps/{app-slug}/pipelines | List all pipelines and standalone builds of an app
*PipelinesAPI* | [**PipelineListAll**](docs/PipelinesAPI.md#pipelinelistall) | **Get** /pipelines | List all pipelines/standalone builds
*PipelinesAPI* | [**PipelineRebuild**](docs/PipelinesAPI.md#pipelinerebuild) | **Post** /apps/{app-slug}/pipelines/{pipeline-id}/rebuild | Rebuilds a pipeline
*PipelinesAPI* | [**PipelineShow**](docs/PipelinesAPI.md#pipelineshow) | **Get** /apps/{app-slug}/pipelines/{pipeline-id} | Get a pipeline of a given app
*ProvisioningProfileAPI* | [**ProvisioningProfileConfirm**](docs/ProvisioningProfileAPI.md#provisioningprofileconfirm) | **Post** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}/uploaded | Confirm a provisioning profile upload
*ProvisioningProfileAPI* | [**ProvisioningProfileCreate**](docs/ProvisioningProfileAPI.md#provisioningprofilecreate) | **Post** /apps/{app-slug}/provisioning-profiles | Create a provisioning profile
*ProvisioningProfileAPI* | [**ProvisioningProfileDelete**](docs/ProvisioningProfileAPI.md#provisioningprofiledelete) | **Delete** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Delete a provisioning profile
*ProvisioningProfileAPI* | [**ProvisioningProfileList**](docs/ProvisioningProfileAPI.md#provisioningprofilelist) | **Get** /apps/{app-slug}/provisioning-profiles | Get a list of the provisioning profiles
*ProvisioningProfileAPI* | [**ProvisioningProfileShow**](docs/ProvisioningProfileAPI.md#provisioningprofileshow) | **Get** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Get a specific provisioning profile
*ProvisioningProfileAPI* | [**ProvisioningProfileUpdate**](docs/ProvisioningProfileAPI.md#provisioningprofileupdate) | **Patch** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Update a provisioning profile
*ReleasesAPI* | [**ReleaseCreateAppStore**](docs/ReleasesAPI.md#releasecreateappstore) | **Post** /apps/{app-slug}/releases/app-store | Create a new Apple App Store release for the app.
*ReleasesAPI* | [**ReleaseCreateGooglePlay**](docs/ReleasesAPI.md#releasecreategoogleplay) | **Post** /apps/{app-slug}/releases/google-play | Create a new Google Play Store release for the app.
*SecretsAPI* | [**SecretDelete**](docs/SecretsAPI.md#secretdelete) | **Delete** /apps/{app-slug}/secrets/{secret-name} | Delete an application secret
*SecretsAPI* | [**SecretList**](docs/SecretsAPI.md#secretlist) | **Get** /apps/{app-slug}/secrets | List the application secrets (with no values)
*SecretsAPI* | [**SecretUpsert**](docs/SecretsAPI.md#secretupsert) | **Put** /apps/{app-slug}/secrets/{secret-name} | Upsert an application secret
*SecretsAPI* | [**SecretValueGet**](docs/SecretsAPI.md#secretvalueget) | **Get** /apps/{app-slug}/secrets/{secret-name}/value | Get the value of an (unprotected) application secrets
*TestDevicesAPI* | [**TestDeviceList**](docs/TestDevicesAPI.md#testdevicelist) | **Get** /apps/{app-slug}/test-devices | List the test devices for an app
*UserAPI* | [**UserMachineTypeUpdate**](docs/UserAPI.md#usermachinetypeupdate) | **Patch** /user/{user-slug}/apps/machine_types | Migrate machine types
*UserAPI* | [**UserPlan**](docs/UserAPI.md#userplan) | **Get** /me/plan | The subscription plan of the user
*UserAPI* | [**UserProfile**](docs/UserAPI.md#userprofile) | **Get** /me | Get your profile info
*UserAPI* | [**UserShow**](docs/UserAPI.md#usershow) | **Get** /users/{user-slug} | Get a specific user
*WebhookDeliveryItemAPI* | [**WebhookDeliveryItemList**](docs/WebhookDeliveryItemAPI.md#webhookdeliveryitemlist) | **Get** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items | List the webhook delivery items of an app
*WebhookDeliveryItemAPI* | [**WebhookDeliveryItemRedeliver**](docs/WebhookDeliveryItemAPI.md#webhookdeliveryitemredeliver) | **Post** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver | Re-deliver the webhook delivery items of an app
*WebhookDeliveryItemAPI* | [**WebhookDeliveryItemShow**](docs/WebhookDeliveryItemAPI.md#webhookdeliveryitemshow) | **Get** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug} | Get a specific delivery item of a webhook


## Documentation For Models

 - [AddonsAddon](docs/AddonsAddon.md)
 - [AddonsDeveloperLink](docs/AddonsDeveloperLink.md)
 - [AddonsFeature](docs/AddonsFeature.md)
 - [AddonsPlan](docs/AddonsPlan.md)
 - [AddonsSetupGuide](docs/AddonsSetupGuide.md)
 - [AddonsSetupInstruction](docs/AddonsSetupInstruction.md)
 - [AppRolesQuery200Response](docs/AppRolesQuery200Response.md)
 - [NullsInt64](docs/NullsInt64.md)
 - [NullsString](docs/NullsString.md)
 - [OrganzationGroupsList200ResponseInner](docs/OrganzationGroupsList200ResponseInner.md)
 - [PipelineabledomainArtifactMeta](docs/PipelineabledomainArtifactMeta.md)
 - [PipelineabledomainBuildToolInvocation](docs/PipelineabledomainBuildToolInvocation.md)
 - [PipelineabledomainEnvironments](docs/PipelineabledomainEnvironments.md)
 - [PipelineabledomainLocalConfig](docs/PipelineabledomainLocalConfig.md)
 - [PipelineabledomainTriggerParams](docs/PipelineabledomainTriggerParams.md)
 - [ServiceProxyErrorRespModel](docs/ServiceProxyErrorRespModel.md)
 - [ServiceStandardErrorRespModel](docs/ServiceStandardErrorRespModel.md)
 - [V0ActivityEventListResponseModel](docs/V0ActivityEventListResponseModel.md)
 - [V0ActivityEventResponseItemModel](docs/V0ActivityEventResponseItemModel.md)
 - [V0AddOnAppResponseItemModel](docs/V0AddOnAppResponseItemModel.md)
 - [V0AddonsListResponseModel](docs/V0AddonsListResponseModel.md)
 - [V0AddonsShowResponseModel](docs/V0AddonsShowResponseModel.md)
 - [V0AndroidKeystoreFileUploadParams](docs/V0AndroidKeystoreFileUploadParams.md)
 - [V0AppAddOnResponseItemModel](docs/V0AppAddOnResponseItemModel.md)
 - [V0AppAddOnsListResponseModel](docs/V0AppAddOnsListResponseModel.md)
 - [V0AppConfigRequestParam](docs/V0AppConfigRequestParam.md)
 - [V0AppDeleteRespModel](docs/V0AppDeleteRespModel.md)
 - [V0AppFinishParams](docs/V0AppFinishParams.md)
 - [V0AppFinishRespModel](docs/V0AppFinishRespModel.md)
 - [V0AppListResponseModel](docs/V0AppListResponseModel.md)
 - [V0AppNotificationSettingsParams](docs/V0AppNotificationSettingsParams.md)
 - [V0AppNotificationSettingsUpdateResponse](docs/V0AppNotificationSettingsUpdateResponse.md)
 - [V0AppRespModel](docs/V0AppRespModel.md)
 - [V0AppResponseItemModel](docs/V0AppResponseItemModel.md)
 - [V0AppSecret](docs/V0AppSecret.md)
 - [V0AppSecretGetValueResponse](docs/V0AppSecretGetValueResponse.md)
 - [V0AppSecretListResponse](docs/V0AppSecretListResponse.md)
 - [V0AppSecretUpsertParams](docs/V0AppSecretUpsertParams.md)
 - [V0AppShowResponseModel](docs/V0AppShowResponseModel.md)
 - [V0AppUpdateParams](docs/V0AppUpdateParams.md)
 - [V0AppUpdateRespModel](docs/V0AppUpdateRespModel.md)
 - [V0AppUploadParams](docs/V0AppUploadParams.md)
 - [V0AppWebhookCreateParams](docs/V0AppWebhookCreateParams.md)
 - [V0AppWebhookCreatedResponseModel](docs/V0AppWebhookCreatedResponseModel.md)
 - [V0AppWebhookDeletedResponseModel](docs/V0AppWebhookDeletedResponseModel.md)
 - [V0AppWebhookListResponseModel](docs/V0AppWebhookListResponseModel.md)
 - [V0AppWebhookResponseItemModel](docs/V0AppWebhookResponseItemModel.md)
 - [V0AppWebhookResponseModel](docs/V0AppWebhookResponseModel.md)
 - [V0AppWebhookUpdateParams](docs/V0AppWebhookUpdateParams.md)
 - [V0AppleAPICredentialResponseItem](docs/V0AppleAPICredentialResponseItem.md)
 - [V0AppleAPICredentialsListResponse](docs/V0AppleAPICredentialsListResponse.md)
 - [V0ArchivedBuildArtifact](docs/V0ArchivedBuildArtifact.md)
 - [V0ArchivedBuildListResponseModel](docs/V0ArchivedBuildListResponseModel.md)
 - [V0ArchivedBuildResponseItemModel](docs/V0ArchivedBuildResponseItemModel.md)
 - [V0ArtifactDeleteResponseModel](docs/V0ArtifactDeleteResponseModel.md)
 - [V0ArtifactListElementResponseModel](docs/V0ArtifactListElementResponseModel.md)
 - [V0ArtifactListResponseModel](docs/V0ArtifactListResponseModel.md)
 - [V0ArtifactResponseItemModel](docs/V0ArtifactResponseItemModel.md)
 - [V0ArtifactShowResponseModel](docs/V0ArtifactShowResponseModel.md)
 - [V0ArtifactUpdateParams](docs/V0ArtifactUpdateParams.md)
 - [V0BitriseYMLConfigGetResponse](docs/V0BitriseYMLConfigGetResponse.md)
 - [V0BitriseYMLConfigUpdateParams](docs/V0BitriseYMLConfigUpdateParams.md)
 - [V0BranchListResponseModel](docs/V0BranchListResponseModel.md)
 - [V0BuildAbortParams](docs/V0BuildAbortParams.md)
 - [V0BuildAbortResponseModel](docs/V0BuildAbortResponseModel.md)
 - [V0BuildCertificateListResponseModel](docs/V0BuildCertificateListResponseModel.md)
 - [V0BuildCertificateResponseItemModel](docs/V0BuildCertificateResponseItemModel.md)
 - [V0BuildCertificateResponseModel](docs/V0BuildCertificateResponseModel.md)
 - [V0BuildCertificateUpdateParams](docs/V0BuildCertificateUpdateParams.md)
 - [V0BuildCertificateUploadParams](docs/V0BuildCertificateUploadParams.md)
 - [V0BuildListAllResponseItemModel](docs/V0BuildListAllResponseItemModel.md)
 - [V0BuildListAllResponseModel](docs/V0BuildListAllResponseModel.md)
 - [V0BuildListResponseModel](docs/V0BuildListResponseModel.md)
 - [V0BuildParamsEnvironment](docs/V0BuildParamsEnvironment.md)
 - [V0BuildRequestListResponseModel](docs/V0BuildRequestListResponseModel.md)
 - [V0BuildRequestResponseItemModel](docs/V0BuildRequestResponseItemModel.md)
 - [V0BuildRequestUpdateParams](docs/V0BuildRequestUpdateParams.md)
 - [V0BuildRequestUpdateResponseModel](docs/V0BuildRequestUpdateResponseModel.md)
 - [V0BuildResponseItemModel](docs/V0BuildResponseItemModel.md)
 - [V0BuildShowResponseModel](docs/V0BuildShowResponseModel.md)
 - [V0BuildTriggerParams](docs/V0BuildTriggerParams.md)
 - [V0BuildTriggerParamsBuildParams](docs/V0BuildTriggerParamsBuildParams.md)
 - [V0BuildTriggerParamsHookInfo](docs/V0BuildTriggerParamsHookInfo.md)
 - [V0BuildTriggerRespModel](docs/V0BuildTriggerRespModel.md)
 - [V0BuildWorkflowListResponseModel](docs/V0BuildWorkflowListResponseModel.md)
 - [V0CacheItemDownloadResponseModel](docs/V0CacheItemDownloadResponseModel.md)
 - [V0CacheItemDownloadURLResponseModel](docs/V0CacheItemDownloadURLResponseModel.md)
 - [V0CacheItemListResponseItemModel](docs/V0CacheItemListResponseItemModel.md)
 - [V0CacheItemListResponseModel](docs/V0CacheItemListResponseModel.md)
 - [V0CommitPaths](docs/V0CommitPaths.md)
 - [V0OrganizationDataModel](docs/V0OrganizationDataModel.md)
 - [V0OrganizationListRespModel](docs/V0OrganizationListRespModel.md)
 - [V0OrganizationOwner](docs/V0OrganizationOwner.md)
 - [V0OrganizationRespModel](docs/V0OrganizationRespModel.md)
 - [V0OrganizationUpdateMachineTypeParams](docs/V0OrganizationUpdateMachineTypeParams.md)
 - [V0OrganizationUpdateMachineTypeResponse](docs/V0OrganizationUpdateMachineTypeResponse.md)
 - [V0OrganizationUpdateMachineTypeResponseErrorsInner](docs/V0OrganizationUpdateMachineTypeResponseErrorsInner.md)
 - [V0OwnerAccountResponseModel](docs/V0OwnerAccountResponseModel.md)
 - [V0OwnerAddOnResponseItemModel](docs/V0OwnerAddOnResponseItemModel.md)
 - [V0OwnerAddOnsListResponseModel](docs/V0OwnerAddOnsListResponseModel.md)
 - [V0PagingResponseModel](docs/V0PagingResponseModel.md)
 - [V0PipelineAbortParams](docs/V0PipelineAbortParams.md)
 - [V0PipelineListAllResponseItemModel](docs/V0PipelineListAllResponseItemModel.md)
 - [V0PipelineListAllResponseModel](docs/V0PipelineListAllResponseModel.md)
 - [V0PipelineListResponseItemModel](docs/V0PipelineListResponseItemModel.md)
 - [V0PipelineListResponseModel](docs/V0PipelineListResponseModel.md)
 - [V0PipelineRebuildParams](docs/V0PipelineRebuildParams.md)
 - [V0PipelineShowAppResponseModel](docs/V0PipelineShowAppResponseModel.md)
 - [V0PipelineShowAttemptResponseModel](docs/V0PipelineShowAttemptResponseModel.md)
 - [V0PipelineShowEnvironmentsResponseModel](docs/V0PipelineShowEnvironmentsResponseModel.md)
 - [V0PipelineShowResponseModel](docs/V0PipelineShowResponseModel.md)
 - [V0PipelineShowStageResponseModel](docs/V0PipelineShowStageResponseModel.md)
 - [V0PipelineShowTriggerParamsResponseModel](docs/V0PipelineShowTriggerParamsResponseModel.md)
 - [V0PipelineShowWorkflowResponseModel](docs/V0PipelineShowWorkflowResponseModel.md)
 - [V0PlanDataModel](docs/V0PlanDataModel.md)
 - [V0ProjectFileStorageDocumentUpdateParams](docs/V0ProjectFileStorageDocumentUpdateParams.md)
 - [V0ProjectFileStorageListResponseModel](docs/V0ProjectFileStorageListResponseModel.md)
 - [V0ProjectFileStorageResponseItemModel](docs/V0ProjectFileStorageResponseItemModel.md)
 - [V0ProjectFileStorageResponseModel](docs/V0ProjectFileStorageResponseModel.md)
 - [V0ProjectFileStorageUploadParams](docs/V0ProjectFileStorageUploadParams.md)
 - [V0ProvProfileDocumentUpdateParams](docs/V0ProvProfileDocumentUpdateParams.md)
 - [V0ProvisionProfileListResponseModel](docs/V0ProvisionProfileListResponseModel.md)
 - [V0ProvisionProfileResponseItemModel](docs/V0ProvisionProfileResponseItemModel.md)
 - [V0ProvisionProfileResponseModel](docs/V0ProvisionProfileResponseModel.md)
 - [V0ProvisionProfileUploadParams](docs/V0ProvisionProfileUploadParams.md)
 - [V0ReleaseCreateAppStoreParams](docs/V0ReleaseCreateAppStoreParams.md)
 - [V0ReleaseCreateAppStoreRespModel](docs/V0ReleaseCreateAppStoreRespModel.md)
 - [V0ReleaseCreateGooglePlayParams](docs/V0ReleaseCreateGooglePlayParams.md)
 - [V0ReleaseCreateGooglePlayRespModel](docs/V0ReleaseCreateGooglePlayRespModel.md)
 - [V0SSHKeyRespModel](docs/V0SSHKeyRespModel.md)
 - [V0SSHKeyUploadParams](docs/V0SSHKeyUploadParams.md)
 - [V0TestDeviceListResponseModel](docs/V0TestDeviceListResponseModel.md)
 - [V0TestDeviceResponseItemModel](docs/V0TestDeviceResponseItemModel.md)
 - [V0UserPlanDataModel](docs/V0UserPlanDataModel.md)
 - [V0UserPlanRespModel](docs/V0UserPlanRespModel.md)
 - [V0UserProfileDataModel](docs/V0UserProfileDataModel.md)
 - [V0UserProfileRespModel](docs/V0UserProfileRespModel.md)
 - [V0WebhookDeliveryItemResponseModel](docs/V0WebhookDeliveryItemResponseModel.md)
 - [V0WebhookDeliveryItemShowResponseModel](docs/V0WebhookDeliveryItemShowResponseModel.md)
 - [V0WebhookRespModel](docs/V0WebhookRespModel.md)
 - [WebsiteBitriseYMLLocation](docs/WebsiteBitriseYMLLocation.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### AddonAuthToken

- **Type**: API key
- **API key parameter name**: Bitrise-Addon-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Bitrise-Addon-Auth-Token and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Bitrise-Addon-Auth-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### PersonalAccessToken

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

letsconnect@bitrise.io

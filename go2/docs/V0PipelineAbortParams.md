# V0PipelineAbortParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortReason** | **string** |  | 
**AbortWithSuccess** | **bool** |  | 
**SkipNotifications** | **bool** |  | 

## Methods

### NewV0PipelineAbortParams

`func NewV0PipelineAbortParams(abortReason string, abortWithSuccess bool, skipNotifications bool, ) *V0PipelineAbortParams`

NewV0PipelineAbortParams instantiates a new V0PipelineAbortParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineAbortParamsWithDefaults

`func NewV0PipelineAbortParamsWithDefaults() *V0PipelineAbortParams`

NewV0PipelineAbortParamsWithDefaults instantiates a new V0PipelineAbortParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *V0PipelineAbortParams) GetAbortReason() string`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *V0PipelineAbortParams) GetAbortReasonOk() (*string, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *V0PipelineAbortParams) SetAbortReason(v string)`

SetAbortReason sets AbortReason field to given value.


### GetAbortWithSuccess

`func (o *V0PipelineAbortParams) GetAbortWithSuccess() bool`

GetAbortWithSuccess returns the AbortWithSuccess field if non-nil, zero value otherwise.

### GetAbortWithSuccessOk

`func (o *V0PipelineAbortParams) GetAbortWithSuccessOk() (*bool, bool)`

GetAbortWithSuccessOk returns a tuple with the AbortWithSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortWithSuccess

`func (o *V0PipelineAbortParams) SetAbortWithSuccess(v bool)`

SetAbortWithSuccess sets AbortWithSuccess field to given value.


### GetSkipNotifications

`func (o *V0PipelineAbortParams) GetSkipNotifications() bool`

GetSkipNotifications returns the SkipNotifications field if non-nil, zero value otherwise.

### GetSkipNotificationsOk

`func (o *V0PipelineAbortParams) GetSkipNotificationsOk() (*bool, bool)`

GetSkipNotificationsOk returns a tuple with the SkipNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNotifications

`func (o *V0PipelineAbortParams) SetSkipNotifications(v bool)`

SetSkipNotifications sets SkipNotifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



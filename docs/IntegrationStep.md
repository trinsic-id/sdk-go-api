# IntegrationStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**IntegrationLaunchMethod**](IntegrationLaunchMethod.md) | The launch method to perform | 
**Content** | **string** | Step type-specific content related to the step: a URL for &#x60;LaunchBrowser&#x60;, a deeplink for &#x60;DeeplinkToMobile&#x60; or a string to show to the user for &#x60;ShowContent&#x60;. | 
**Refresh** | Pointer to [**NullableStepRefreshInfo**](StepRefreshInfo.md) | If non-null, contains metadata about how to refresh the value of &#x60;content&#x60;. | [optional] 

## Methods

### NewIntegrationStep

`func NewIntegrationStep(method IntegrationLaunchMethod, content string, ) *IntegrationStep`

NewIntegrationStep instantiates a new IntegrationStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStepWithDefaults

`func NewIntegrationStepWithDefaults() *IntegrationStep`

NewIntegrationStepWithDefaults instantiates a new IntegrationStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *IntegrationStep) GetMethod() IntegrationLaunchMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IntegrationStep) GetMethodOk() (*IntegrationLaunchMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IntegrationStep) SetMethod(v IntegrationLaunchMethod)`

SetMethod sets Method field to given value.


### GetContent

`func (o *IntegrationStep) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IntegrationStep) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IntegrationStep) SetContent(v string)`

SetContent sets Content field to given value.


### GetRefresh

`func (o *IntegrationStep) GetRefresh() StepRefreshInfo`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *IntegrationStep) GetRefreshOk() (*StepRefreshInfo, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *IntegrationStep) SetRefresh(v StepRefreshInfo)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *IntegrationStep) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *IntegrationStep) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *IntegrationStep) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



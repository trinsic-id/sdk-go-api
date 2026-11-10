# ConnectIdInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubProviderId** | Pointer to **NullableString** | The ID of the specific sub-provider to invoke with ConnectID.              If not specified, the user will be prompted to select the sub-provider when &#x60;fallbackToHostedUI&#x60; is enabled. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`ConnectIdInput{FieldName: value}`), or call
`NewConnectIdInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewConnectIdInputWithDefaults

`func NewConnectIdInputWithDefaults() *ConnectIdInput`

NewConnectIdInputWithDefaults instantiates a new ConnectIdInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubProviderId

`func (o *ConnectIdInput) GetSubProviderId() string`

GetSubProviderId returns the SubProviderId field if non-nil, zero value otherwise.

### GetSubProviderIdOk

`func (o *ConnectIdInput) GetSubProviderIdOk() (*string, bool)`

GetSubProviderIdOk returns a tuple with the SubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviderId

`func (o *ConnectIdInput) SetSubProviderId(v string)`

SetSubProviderId sets SubProviderId field to given value.

### HasSubProviderId

`func (o *ConnectIdInput) HasSubProviderId() bool`

HasSubProviderId returns a boolean if a field has been set.

### SetSubProviderIdNil

`func (o *ConnectIdInput) SetSubProviderIdNil(b bool)`

 SetSubProviderIdNil sets the value for SubProviderId to be an explicit nil

### UnsetSubProviderId
`func (o *ConnectIdInput) UnsetSubProviderId()`

UnsetSubProviderId ensures that no value is present for SubProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



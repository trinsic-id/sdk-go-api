# IdinInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubProviderId** | Pointer to **NullableString** | The ID of the specific bank to invoke with IDIN.              If not specified, the user will be prompted to select a bank. | [optional] 

## Methods

### NewIdinInput

`func NewIdinInput() *IdinInput`

NewIdinInput instantiates a new IdinInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdinInputWithDefaults

`func NewIdinInputWithDefaults() *IdinInput`

NewIdinInputWithDefaults instantiates a new IdinInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubProviderId

`func (o *IdinInput) GetSubProviderId() string`

GetSubProviderId returns the SubProviderId field if non-nil, zero value otherwise.

### GetSubProviderIdOk

`func (o *IdinInput) GetSubProviderIdOk() (*string, bool)`

GetSubProviderIdOk returns a tuple with the SubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviderId

`func (o *IdinInput) SetSubProviderId(v string)`

SetSubProviderId sets SubProviderId field to given value.

### HasSubProviderId

`func (o *IdinInput) HasSubProviderId() bool`

HasSubProviderId returns a boolean if a field has been set.

### SetSubProviderIdNil

`func (o *IdinInput) SetSubProviderIdNil(b bool)`

 SetSubProviderIdNil sets the value for SubProviderId to be an explicit nil

### UnsetSubProviderId
`func (o *IdinInput) UnsetSubProviderId()`

UnsetSubProviderId ensures that no value is present for SubProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



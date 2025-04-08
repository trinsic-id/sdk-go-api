# TrinsicTestSubProvidersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubProviderId** | Pointer to **NullableString** | The ID of the specific IDP to invoke within the test federated provider.              Valid options are &#x60;sub-provider-a&#x60; and &#x60;sub-provider-b&#x60;.              If not specified, the user will be prompted to select one. | [optional] 

## Methods

### NewTrinsicTestSubProvidersInput

`func NewTrinsicTestSubProvidersInput() *TrinsicTestSubProvidersInput`

NewTrinsicTestSubProvidersInput instantiates a new TrinsicTestSubProvidersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrinsicTestSubProvidersInputWithDefaults

`func NewTrinsicTestSubProvidersInputWithDefaults() *TrinsicTestSubProvidersInput`

NewTrinsicTestSubProvidersInputWithDefaults instantiates a new TrinsicTestSubProvidersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubProviderId

`func (o *TrinsicTestSubProvidersInput) GetSubProviderId() string`

GetSubProviderId returns the SubProviderId field if non-nil, zero value otherwise.

### GetSubProviderIdOk

`func (o *TrinsicTestSubProvidersInput) GetSubProviderIdOk() (*string, bool)`

GetSubProviderIdOk returns a tuple with the SubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviderId

`func (o *TrinsicTestSubProvidersInput) SetSubProviderId(v string)`

SetSubProviderId sets SubProviderId field to given value.

### HasSubProviderId

`func (o *TrinsicTestSubProvidersInput) HasSubProviderId() bool`

HasSubProviderId returns a boolean if a field has been set.

### SetSubProviderIdNil

`func (o *TrinsicTestSubProvidersInput) SetSubProviderIdNil(b bool)`

 SetSubProviderIdNil sets the value for SubProviderId to be an explicit nil

### UnsetSubProviderId
`func (o *TrinsicTestSubProvidersInput) UnsetSubProviderId()`

UnsetSubProviderId ensures that no value is present for SubProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



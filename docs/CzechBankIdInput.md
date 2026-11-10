# CzechBankIdInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubProviderId** | Pointer to **NullableString** | The ID of the Czech Bank to invoke. If omitted, the user selects a sub-provider. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`CzechBankIdInput{FieldName: value}`), or call
`NewCzechBankIdInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewCzechBankIdInputWithDefaults

`func NewCzechBankIdInputWithDefaults() *CzechBankIdInput`

NewCzechBankIdInputWithDefaults instantiates a new CzechBankIdInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubProviderId

`func (o *CzechBankIdInput) GetSubProviderId() string`

GetSubProviderId returns the SubProviderId field if non-nil, zero value otherwise.

### GetSubProviderIdOk

`func (o *CzechBankIdInput) GetSubProviderIdOk() (*string, bool)`

GetSubProviderIdOk returns a tuple with the SubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProviderId

`func (o *CzechBankIdInput) SetSubProviderId(v string)`

SetSubProviderId sets SubProviderId field to given value.

### HasSubProviderId

`func (o *CzechBankIdInput) HasSubProviderId() bool`

HasSubProviderId returns a boolean if a field has been set.

### SetSubProviderIdNil

`func (o *CzechBankIdInput) SetSubProviderIdNil(b bool)`

 SetSubProviderIdNil sets the value for SubProviderId to be an explicit nil

### UnsetSubProviderId
`func (o *CzechBankIdInput) UnsetSubProviderId()`

UnsetSubProviderId ensures that no value is present for SubProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



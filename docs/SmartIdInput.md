# SmartIdInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmartIdNumber** | Pointer to **NullableString** | The user&#39;s Smart ID ETSI number or document number. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`SmartIdInput{FieldName: value}`), or call
`NewSmartIdInputWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSmartIdInputWithDefaults

`func NewSmartIdInputWithDefaults() *SmartIdInput`

NewSmartIdInputWithDefaults instantiates a new SmartIdInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmartIdNumber

`func (o *SmartIdInput) GetSmartIdNumber() string`

GetSmartIdNumber returns the SmartIdNumber field if non-nil, zero value otherwise.

### GetSmartIdNumberOk

`func (o *SmartIdInput) GetSmartIdNumberOk() (*string, bool)`

GetSmartIdNumberOk returns a tuple with the SmartIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartIdNumber

`func (o *SmartIdInput) SetSmartIdNumber(v string)`

SetSmartIdNumber sets SmartIdNumber field to given value.

### HasSmartIdNumber

`func (o *SmartIdInput) HasSmartIdNumber() bool`

HasSmartIdNumber returns a boolean if a field has been set.

### SetSmartIdNumberNil

`func (o *SmartIdInput) SetSmartIdNumberNil(b bool)`

 SetSmartIdNumberNil sets the value for SmartIdNumber to be an explicit nil

### UnsetSmartIdNumber
`func (o *SmartIdInput) UnsetSmartIdNumber()`

UnsetSmartIdNumber ensures that no value is present for SmartIdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



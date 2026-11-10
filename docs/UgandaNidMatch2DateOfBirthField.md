# UgandaNidMatch2DateOfBirthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The date of birth submitted for this check. | [optional] 
**Outcome** | Pointer to **NullableString** | The outcome of comparing the submitted date of birth with NIRA database.              For Uganda match, partial or transposed outcomes are not used; matching is exact or no match.              Possible values: - Exact Match - No Match - Not Returned | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`UgandaNidMatch2DateOfBirthField{FieldName: value}`), or call
`NewUgandaNidMatch2DateOfBirthFieldWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewUgandaNidMatch2DateOfBirthFieldWithDefaults

`func NewUgandaNidMatch2DateOfBirthFieldWithDefaults() *UgandaNidMatch2DateOfBirthField`

NewUgandaNidMatch2DateOfBirthFieldWithDefaults instantiates a new UgandaNidMatch2DateOfBirthField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *UgandaNidMatch2DateOfBirthField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *UgandaNidMatch2DateOfBirthField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *UgandaNidMatch2DateOfBirthField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *UgandaNidMatch2DateOfBirthField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *UgandaNidMatch2DateOfBirthField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *UgandaNidMatch2DateOfBirthField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *UgandaNidMatch2DateOfBirthField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *UgandaNidMatch2DateOfBirthField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *UgandaNidMatch2DateOfBirthField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *UgandaNidMatch2DateOfBirthField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *UgandaNidMatch2DateOfBirthField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *UgandaNidMatch2DateOfBirthField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



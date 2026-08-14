# NigeriaNinMatch2FullNameField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The full name string submitted for this check (as provided in the match input). | [optional] 
**Outcome** | Pointer to **NullableString** | The outcome of comparing the submitted full name to the issuer record.              Possible values: - Exact Match - Partial Match - Transposed - No Match - Not Returned | [optional] 

## Methods

### NewNigeriaNinMatch2FullNameField

`func NewNigeriaNinMatch2FullNameField() *NigeriaNinMatch2FullNameField`

NewNigeriaNinMatch2FullNameField instantiates a new NigeriaNinMatch2FullNameField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinMatch2FullNameFieldWithDefaults

`func NewNigeriaNinMatch2FullNameFieldWithDefaults() *NigeriaNinMatch2FullNameField`

NewNigeriaNinMatch2FullNameFieldWithDefaults instantiates a new NigeriaNinMatch2FullNameField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *NigeriaNinMatch2FullNameField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *NigeriaNinMatch2FullNameField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *NigeriaNinMatch2FullNameField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *NigeriaNinMatch2FullNameField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *NigeriaNinMatch2FullNameField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *NigeriaNinMatch2FullNameField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *NigeriaNinMatch2FullNameField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *NigeriaNinMatch2FullNameField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *NigeriaNinMatch2FullNameField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *NigeriaNinMatch2FullNameField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *NigeriaNinMatch2FullNameField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *NigeriaNinMatch2FullNameField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



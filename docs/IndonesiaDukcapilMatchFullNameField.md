# IndonesiaDukcapilMatchFullNameField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The full name submitted for this check. | [optional] 
**Outcome** | Pointer to **NullableFloat64** | The provider assessment result for the submitted full name.              Format: numeric score from 0.0 to 1.0, where higher values indicate closer agreement with the Dukcapil record. | [optional] 

## Methods

### NewIndonesiaDukcapilMatchFullNameField

`func NewIndonesiaDukcapilMatchFullNameField() *IndonesiaDukcapilMatchFullNameField`

NewIndonesiaDukcapilMatchFullNameField instantiates a new IndonesiaDukcapilMatchFullNameField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilMatchFullNameFieldWithDefaults

`func NewIndonesiaDukcapilMatchFullNameFieldWithDefaults() *IndonesiaDukcapilMatchFullNameField`

NewIndonesiaDukcapilMatchFullNameFieldWithDefaults instantiates a new IndonesiaDukcapilMatchFullNameField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *IndonesiaDukcapilMatchFullNameField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *IndonesiaDukcapilMatchFullNameField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *IndonesiaDukcapilMatchFullNameField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *IndonesiaDukcapilMatchFullNameField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *IndonesiaDukcapilMatchFullNameField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *IndonesiaDukcapilMatchFullNameField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *IndonesiaDukcapilMatchFullNameField) GetOutcome() float64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *IndonesiaDukcapilMatchFullNameField) GetOutcomeOk() (*float64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *IndonesiaDukcapilMatchFullNameField) SetOutcome(v float64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *IndonesiaDukcapilMatchFullNameField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *IndonesiaDukcapilMatchFullNameField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *IndonesiaDukcapilMatchFullNameField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



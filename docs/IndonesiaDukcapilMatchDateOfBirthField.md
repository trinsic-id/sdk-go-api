# IndonesiaDukcapilMatchDateOfBirthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The date of birth submitted for this check. | [optional] 
**Outcome** | Pointer to **NullableFloat64** | The provider assessment result for the submitted date of birth.              Local integration code expects the provider to return 0 or 1 for this assessment. | [optional] 

## Methods

### NewIndonesiaDukcapilMatchDateOfBirthField

`func NewIndonesiaDukcapilMatchDateOfBirthField() *IndonesiaDukcapilMatchDateOfBirthField`

NewIndonesiaDukcapilMatchDateOfBirthField instantiates a new IndonesiaDukcapilMatchDateOfBirthField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndonesiaDukcapilMatchDateOfBirthFieldWithDefaults

`func NewIndonesiaDukcapilMatchDateOfBirthFieldWithDefaults() *IndonesiaDukcapilMatchDateOfBirthField`

NewIndonesiaDukcapilMatchDateOfBirthFieldWithDefaults instantiates a new IndonesiaDukcapilMatchDateOfBirthField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *IndonesiaDukcapilMatchDateOfBirthField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *IndonesiaDukcapilMatchDateOfBirthField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *IndonesiaDukcapilMatchDateOfBirthField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *IndonesiaDukcapilMatchDateOfBirthField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *IndonesiaDukcapilMatchDateOfBirthField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *IndonesiaDukcapilMatchDateOfBirthField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *IndonesiaDukcapilMatchDateOfBirthField) GetOutcome() float64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *IndonesiaDukcapilMatchDateOfBirthField) GetOutcomeOk() (*float64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *IndonesiaDukcapilMatchDateOfBirthField) SetOutcome(v float64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *IndonesiaDukcapilMatchDateOfBirthField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *IndonesiaDukcapilMatchDateOfBirthField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *IndonesiaDukcapilMatchDateOfBirthField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



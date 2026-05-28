# IndonesiaDukcapilMatchDateOfBirthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The date of birth submitted for this check. | 
**Outcome** | **float64** | The provider assessment result for the submitted date of birth.              Local integration code expects the provider to return 0 or 1 for this assessment. | 

## Methods

### NewIndonesiaDukcapilMatchDateOfBirthField

`func NewIndonesiaDukcapilMatchDateOfBirthField(inputValue string, outcome float64, ) *IndonesiaDukcapilMatchDateOfBirthField`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



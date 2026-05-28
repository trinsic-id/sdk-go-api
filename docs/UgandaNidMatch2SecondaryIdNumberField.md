# UgandaNidMatch2SecondaryIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The card number on the national ID document, distinct from the National ID number.              Format: - 9 numeric digits - There is no publicly documented encoding scheme for encoding personal information in the card number - No check digit or algorithm has been publicly documented by NIRA | 
**Outcome** | **string** | The outcome of comparing the submitted card number with NIRA database.              Possible values: - Exact Match - No Match - Not Returned | 

## Methods

### NewUgandaNidMatch2SecondaryIdNumberField

`func NewUgandaNidMatch2SecondaryIdNumberField(inputValue string, outcome string, ) *UgandaNidMatch2SecondaryIdNumberField`

NewUgandaNidMatch2SecondaryIdNumberField instantiates a new UgandaNidMatch2SecondaryIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUgandaNidMatch2SecondaryIdNumberFieldWithDefaults

`func NewUgandaNidMatch2SecondaryIdNumberFieldWithDefaults() *UgandaNidMatch2SecondaryIdNumberField`

NewUgandaNidMatch2SecondaryIdNumberFieldWithDefaults instantiates a new UgandaNidMatch2SecondaryIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *UgandaNidMatch2SecondaryIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *UgandaNidMatch2SecondaryIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *UgandaNidMatch2SecondaryIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *UgandaNidMatch2SecondaryIdNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *UgandaNidMatch2SecondaryIdNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *UgandaNidMatch2SecondaryIdNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



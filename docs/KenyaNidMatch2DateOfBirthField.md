# KenyaNidMatch2DateOfBirthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The date of birth submitted for this check. | 
**Outcome** | **string** | The outcome of comparing the submitted date of birth against the Integrated Population Registration System (IPRS).              Possible values: - Exact Match - Partial Match - Transposed - No Match - Not Returned | 

## Methods

### NewKenyaNidMatch2DateOfBirthField

`func NewKenyaNidMatch2DateOfBirthField(inputValue string, outcome string, ) *KenyaNidMatch2DateOfBirthField`

NewKenyaNidMatch2DateOfBirthField instantiates a new KenyaNidMatch2DateOfBirthField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidMatch2DateOfBirthFieldWithDefaults

`func NewKenyaNidMatch2DateOfBirthFieldWithDefaults() *KenyaNidMatch2DateOfBirthField`

NewKenyaNidMatch2DateOfBirthFieldWithDefaults instantiates a new KenyaNidMatch2DateOfBirthField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *KenyaNidMatch2DateOfBirthField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *KenyaNidMatch2DateOfBirthField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *KenyaNidMatch2DateOfBirthField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *KenyaNidMatch2DateOfBirthField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *KenyaNidMatch2DateOfBirthField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *KenyaNidMatch2DateOfBirthField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



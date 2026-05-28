# UgandaNidMatch2DateOfBirthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The date of birth submitted for this check. | 
**Outcome** | **string** | The outcome of comparing the submitted date of birth with NIRA database.              For Uganda match, partial or transposed outcomes are not used; matching is exact or no match.              Possible values: - Exact Match - No Match - Not Returned | 

## Methods

### NewUgandaNidMatch2DateOfBirthField

`func NewUgandaNidMatch2DateOfBirthField(inputValue string, outcome string, ) *UgandaNidMatch2DateOfBirthField`

NewUgandaNidMatch2DateOfBirthField instantiates a new UgandaNidMatch2DateOfBirthField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



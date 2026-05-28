# UgandaNidMatch2NationalIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | Uganda National ID number (NID) assigned by the National Identification and Registration Authority (NIRA).              Format: - 14 alphanumeric characters (A-Z, 0-9) - There is no publicly documented encoding scheme for encoding personal information in the NID - No check digit or algorithm has been publicly documented by NIRA | 
**Outcome** | **string** | The outcome of looking up the NID in NIRA database.              Possible values: - Verified - Not Verified - Not Done - Issuer Unavailable - Not Returned | 

## Methods

### NewUgandaNidMatch2NationalIdNumberField

`func NewUgandaNidMatch2NationalIdNumberField(inputValue string, outcome string, ) *UgandaNidMatch2NationalIdNumberField`

NewUgandaNidMatch2NationalIdNumberField instantiates a new UgandaNidMatch2NationalIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUgandaNidMatch2NationalIdNumberFieldWithDefaults

`func NewUgandaNidMatch2NationalIdNumberFieldWithDefaults() *UgandaNidMatch2NationalIdNumberField`

NewUgandaNidMatch2NationalIdNumberFieldWithDefaults instantiates a new UgandaNidMatch2NationalIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *UgandaNidMatch2NationalIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *UgandaNidMatch2NationalIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *UgandaNidMatch2NationalIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *UgandaNidMatch2NationalIdNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *UgandaNidMatch2NationalIdNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *UgandaNidMatch2NationalIdNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



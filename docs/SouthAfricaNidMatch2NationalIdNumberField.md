# SouthAfricaNidMatch2NationalIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The South African National Identity Number (13 digits) is issued for life by the Department of Home Affairs (DHA) and stored in the Home Affairs National Identification System (HANIS) database. The same number is mandatory for banking, employment, taxation, and voting, and is printed on both the legacy green ID book and the Smart ID Card (rolled out from 2013 onward).              Format: - YYMMDD G(4) C A Z - YYMMDD is the date of birth - G(4) is the gender code (below 5000 female, 5000 or above male) - C is the citizenship indicator (0 citizen, 1 permanent resident) - A is reserved (it had a politically sensitive meaning in the past, but is currently   semantically meaningless) - Z is a Luhn check digit | [optional] 
**Outcome** | Pointer to **NullableString** | The outcome of verifying the national ID number.              Possible values: - Verified - Not Verified - Not Done - Issuer Unavailable - Not Returned | [optional] 

## Methods

### NewSouthAfricaNidMatch2NationalIdNumberField

`func NewSouthAfricaNidMatch2NationalIdNumberField() *SouthAfricaNidMatch2NationalIdNumberField`

NewSouthAfricaNidMatch2NationalIdNumberField instantiates a new SouthAfricaNidMatch2NationalIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidMatch2NationalIdNumberFieldWithDefaults

`func NewSouthAfricaNidMatch2NationalIdNumberFieldWithDefaults() *SouthAfricaNidMatch2NationalIdNumberField`

NewSouthAfricaNidMatch2NationalIdNumberFieldWithDefaults instantiates a new SouthAfricaNidMatch2NationalIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *SouthAfricaNidMatch2NationalIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *SouthAfricaNidMatch2NationalIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *SouthAfricaNidMatch2NationalIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *SouthAfricaNidMatch2NationalIdNumberField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *SouthAfricaNidMatch2NationalIdNumberField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *SouthAfricaNidMatch2NationalIdNumberField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *SouthAfricaNidMatch2NationalIdNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *SouthAfricaNidMatch2NationalIdNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *SouthAfricaNidMatch2NationalIdNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *SouthAfricaNidMatch2NationalIdNumberField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *SouthAfricaNidMatch2NationalIdNumberField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *SouthAfricaNidMatch2NationalIdNumberField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



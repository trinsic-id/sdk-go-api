# KenyaNidMatch2NationalIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | Pointer to **NullableString** | The Kenya National ID Number (Nambari ya Kitambulisho) or Unique Personal Identifier (Maisha Namba).              This is the primary unique identifier for Kenyan citizens in all government systems, issued by the National Registration Bureau (NRB). The format is up to 8 digits for National ID or 9 digits for Maisha Namba UPI (the new format since 2023).              Format: - Up to 8 digits for National ID, unique and sequentially generated - 9 digits for Maisha Namba UPI, unique and randomly generated - Neither number encodes any information about the individual | [optional] 
**Outcome** | Pointer to **NullableString** | The outcome of verifying the national ID number against the Integrated Population Registration System (IPRS).              Possible values: - Verified - Not Verified - Not Done - Issuer Unavailable - Not Returned | [optional] 

## Methods

### NewKenyaNidMatch2NationalIdNumberField

`func NewKenyaNidMatch2NationalIdNumberField() *KenyaNidMatch2NationalIdNumberField`

NewKenyaNidMatch2NationalIdNumberField instantiates a new KenyaNidMatch2NationalIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidMatch2NationalIdNumberFieldWithDefaults

`func NewKenyaNidMatch2NationalIdNumberFieldWithDefaults() *KenyaNidMatch2NationalIdNumberField`

NewKenyaNidMatch2NationalIdNumberFieldWithDefaults instantiates a new KenyaNidMatch2NationalIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *KenyaNidMatch2NationalIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *KenyaNidMatch2NationalIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *KenyaNidMatch2NationalIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *KenyaNidMatch2NationalIdNumberField) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *KenyaNidMatch2NationalIdNumberField) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *KenyaNidMatch2NationalIdNumberField) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetOutcome

`func (o *KenyaNidMatch2NationalIdNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *KenyaNidMatch2NationalIdNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *KenyaNidMatch2NationalIdNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *KenyaNidMatch2NationalIdNumberField) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *KenyaNidMatch2NationalIdNumberField) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *KenyaNidMatch2NationalIdNumberField) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



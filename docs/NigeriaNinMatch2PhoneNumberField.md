# NigeriaNinMatch2PhoneNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The phone number as submitted for this check.              Format: - International E.164 | 
**Outcome** | **string** | The outcome of comparing the submitted phone number to the issuer record.              Possible values: - Exact Match - No Match - Not Provided - Not Returned | 

## Methods

### NewNigeriaNinMatch2PhoneNumberField

`func NewNigeriaNinMatch2PhoneNumberField(inputValue string, outcome string, ) *NigeriaNinMatch2PhoneNumberField`

NewNigeriaNinMatch2PhoneNumberField instantiates a new NigeriaNinMatch2PhoneNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinMatch2PhoneNumberFieldWithDefaults

`func NewNigeriaNinMatch2PhoneNumberFieldWithDefaults() *NigeriaNinMatch2PhoneNumberField`

NewNigeriaNinMatch2PhoneNumberFieldWithDefaults instantiates a new NigeriaNinMatch2PhoneNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *NigeriaNinMatch2PhoneNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *NigeriaNinMatch2PhoneNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *NigeriaNinMatch2PhoneNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *NigeriaNinMatch2PhoneNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *NigeriaNinMatch2PhoneNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *NigeriaNinMatch2PhoneNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



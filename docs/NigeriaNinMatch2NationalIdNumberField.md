# NigeriaNinMatch2NationalIdNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | National Identification Number (NIN).              This is a unique, permanent identifier assigned by the National Identity Management Commission (NIMC) upon enrollment.              Format: - 11 numeric digits - No personal information is encoded in the NIN | 
**Outcome** | **string** | The outcome of verifying the NIN against the National Identity Database.              Possible values: - Verified - Not Verified - Not Done - Issuer Unavailable - Not Returned | 

## Methods

### NewNigeriaNinMatch2NationalIdNumberField

`func NewNigeriaNinMatch2NationalIdNumberField(inputValue string, outcome string, ) *NigeriaNinMatch2NationalIdNumberField`

NewNigeriaNinMatch2NationalIdNumberField instantiates a new NigeriaNinMatch2NationalIdNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinMatch2NationalIdNumberFieldWithDefaults

`func NewNigeriaNinMatch2NationalIdNumberFieldWithDefaults() *NigeriaNinMatch2NationalIdNumberField`

NewNigeriaNinMatch2NationalIdNumberFieldWithDefaults instantiates a new NigeriaNinMatch2NationalIdNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *NigeriaNinMatch2NationalIdNumberField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *NigeriaNinMatch2NationalIdNumberField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *NigeriaNinMatch2NationalIdNumberField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *NigeriaNinMatch2NationalIdNumberField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *NigeriaNinMatch2NationalIdNumberField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *NigeriaNinMatch2NationalIdNumberField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



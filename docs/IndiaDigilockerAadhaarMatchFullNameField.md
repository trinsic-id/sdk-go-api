# IndiaDigilockerAadhaarMatchFullNameField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The full name string submitted for this check, as provided in the match input. | 
**Outcome** | **float64** | Full-name similarity score from the DigiLocker match result for this session.              Format: - Between 0.0 to 1.0, where higher values indicate closer agreement with the Aadhaar record. | 

## Methods

### NewIndiaDigilockerAadhaarMatchFullNameField

`func NewIndiaDigilockerAadhaarMatchFullNameField(inputValue string, outcome float64, ) *IndiaDigilockerAadhaarMatchFullNameField`

NewIndiaDigilockerAadhaarMatchFullNameField instantiates a new IndiaDigilockerAadhaarMatchFullNameField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndiaDigilockerAadhaarMatchFullNameFieldWithDefaults

`func NewIndiaDigilockerAadhaarMatchFullNameFieldWithDefaults() *IndiaDigilockerAadhaarMatchFullNameField`

NewIndiaDigilockerAadhaarMatchFullNameFieldWithDefaults instantiates a new IndiaDigilockerAadhaarMatchFullNameField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *IndiaDigilockerAadhaarMatchFullNameField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *IndiaDigilockerAadhaarMatchFullNameField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *IndiaDigilockerAadhaarMatchFullNameField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *IndiaDigilockerAadhaarMatchFullNameField) GetOutcome() float64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *IndiaDigilockerAadhaarMatchFullNameField) GetOutcomeOk() (*float64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *IndiaDigilockerAadhaarMatchFullNameField) SetOutcome(v float64)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



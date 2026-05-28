# KenyaNidMatch2FullNameField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The full name string submitted for this check (as provided in the match input). | 
**Outcome** | **string** | The outcome of comparing the submitted full name against the Integrated Population Registration System (IPRS).              Possible values: - Exact Match - Partial Match - Transposed - No Match - Not Returned | 

## Methods

### NewKenyaNidMatch2FullNameField

`func NewKenyaNidMatch2FullNameField(inputValue string, outcome string, ) *KenyaNidMatch2FullNameField`

NewKenyaNidMatch2FullNameField instantiates a new KenyaNidMatch2FullNameField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidMatch2FullNameFieldWithDefaults

`func NewKenyaNidMatch2FullNameFieldWithDefaults() *KenyaNidMatch2FullNameField`

NewKenyaNidMatch2FullNameFieldWithDefaults instantiates a new KenyaNidMatch2FullNameField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *KenyaNidMatch2FullNameField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *KenyaNidMatch2FullNameField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *KenyaNidMatch2FullNameField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *KenyaNidMatch2FullNameField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *KenyaNidMatch2FullNameField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *KenyaNidMatch2FullNameField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



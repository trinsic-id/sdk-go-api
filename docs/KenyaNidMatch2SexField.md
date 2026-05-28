# KenyaNidMatch2SexField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputValue** | **string** | The sex value submitted for this check.              Possible input values: - Male - Female | 
**Outcome** | **string** | The outcome of comparing the submitted sex against the Integrated Population Registration System (IPRS).              Possible values: - Exact Match - No Match - Not Provided - Not Returned | 

## Methods

### NewKenyaNidMatch2SexField

`func NewKenyaNidMatch2SexField(inputValue string, outcome string, ) *KenyaNidMatch2SexField`

NewKenyaNidMatch2SexField instantiates a new KenyaNidMatch2SexField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidMatch2SexFieldWithDefaults

`func NewKenyaNidMatch2SexFieldWithDefaults() *KenyaNidMatch2SexField`

NewKenyaNidMatch2SexFieldWithDefaults instantiates a new KenyaNidMatch2SexField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputValue

`func (o *KenyaNidMatch2SexField) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *KenyaNidMatch2SexField) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *KenyaNidMatch2SexField) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.


### GetOutcome

`func (o *KenyaNidMatch2SexField) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *KenyaNidMatch2SexField) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *KenyaNidMatch2SexField) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



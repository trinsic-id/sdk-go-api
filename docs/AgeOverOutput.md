# AgeOverOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **int32** | The age threshold from the &#x60;age over X&#x60; claim. | 
**IsOver** | **bool** | Whether the individual is at least the given age. | 

## Methods

### NewAgeOverOutput

`func NewAgeOverOutput(age int32, isOver bool, ) *AgeOverOutput`

NewAgeOverOutput instantiates a new AgeOverOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgeOverOutputWithDefaults

`func NewAgeOverOutputWithDefaults() *AgeOverOutput`

NewAgeOverOutputWithDefaults instantiates a new AgeOverOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *AgeOverOutput) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *AgeOverOutput) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *AgeOverOutput) SetAge(v int32)`

SetAge sets Age field to given value.


### GetIsOver

`func (o *AgeOverOutput) GetIsOver() bool`

GetIsOver returns the IsOver field if non-nil, zero value otherwise.

### GetIsOverOk

`func (o *AgeOverOutput) GetIsOverOk() (*bool, bool)`

GetIsOverOk returns a tuple with the IsOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOver

`func (o *AgeOverOutput) SetIsOver(v bool)`

SetIsOver sets IsOver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



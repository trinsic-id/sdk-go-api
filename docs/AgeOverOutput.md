# AgeOverOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **NullableInt32** | The age threshold from the &#x60;age over X&#x60; claim. | [optional] 
**IsOver** | Pointer to **NullableBool** | Whether the individual is at least the given age. | [optional] 

## Methods

### NewAgeOverOutput

`func NewAgeOverOutput() *AgeOverOutput`

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

### HasAge

`func (o *AgeOverOutput) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *AgeOverOutput) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *AgeOverOutput) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
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

### HasIsOver

`func (o *AgeOverOutput) HasIsOver() bool`

HasIsOver returns a boolean if a field has been set.

### SetIsOverNil

`func (o *AgeOverOutput) SetIsOverNil(b bool)`

 SetIsOverNil sets the value for IsOver to be an explicit nil

### UnsetIsOver
`func (o *AgeOverOutput) UnsetIsOver()`

UnsetIsOver ensures that no value is present for IsOver, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



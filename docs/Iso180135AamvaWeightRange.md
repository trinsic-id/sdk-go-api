# Iso180135AamvaWeightRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinKilograms** | **int32** | Minimum weight in the range, in kilograms. | 
**MaxKilograms** | Pointer to **NullableInt32** | Maximum weight in the range, in kilograms.              If null, the maximum weight has no limit. | [optional] 

## Methods

### NewIso180135AamvaWeightRange

`func NewIso180135AamvaWeightRange(minKilograms int32, ) *Iso180135AamvaWeightRange`

NewIso180135AamvaWeightRange instantiates a new Iso180135AamvaWeightRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180135AamvaWeightRangeWithDefaults

`func NewIso180135AamvaWeightRangeWithDefaults() *Iso180135AamvaWeightRange`

NewIso180135AamvaWeightRangeWithDefaults instantiates a new Iso180135AamvaWeightRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinKilograms

`func (o *Iso180135AamvaWeightRange) GetMinKilograms() int32`

GetMinKilograms returns the MinKilograms field if non-nil, zero value otherwise.

### GetMinKilogramsOk

`func (o *Iso180135AamvaWeightRange) GetMinKilogramsOk() (*int32, bool)`

GetMinKilogramsOk returns a tuple with the MinKilograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinKilograms

`func (o *Iso180135AamvaWeightRange) SetMinKilograms(v int32)`

SetMinKilograms sets MinKilograms field to given value.


### GetMaxKilograms

`func (o *Iso180135AamvaWeightRange) GetMaxKilograms() int32`

GetMaxKilograms returns the MaxKilograms field if non-nil, zero value otherwise.

### GetMaxKilogramsOk

`func (o *Iso180135AamvaWeightRange) GetMaxKilogramsOk() (*int32, bool)`

GetMaxKilogramsOk returns a tuple with the MaxKilograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKilograms

`func (o *Iso180135AamvaWeightRange) SetMaxKilograms(v int32)`

SetMaxKilograms sets MaxKilograms field to given value.

### HasMaxKilograms

`func (o *Iso180135AamvaWeightRange) HasMaxKilograms() bool`

HasMaxKilograms returns a boolean if a field has been set.

### SetMaxKilogramsNil

`func (o *Iso180135AamvaWeightRange) SetMaxKilogramsNil(b bool)`

 SetMaxKilogramsNil sets the value for MaxKilograms to be an explicit nil

### UnsetMaxKilograms
`func (o *Iso180135AamvaWeightRange) UnsetMaxKilograms()`

UnsetMaxKilograms ensures that no value is present for MaxKilograms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



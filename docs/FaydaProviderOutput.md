# FaydaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | A unique eKYC identifying token used to match the original eKYC token received from the provider when the user was initially registered.              Since Fayda does not return identifying data, it is the responsibility of the relying party to keep the unique user token received from Fayda when the user was initially registered to do a comparison of the subs to verify that it is the same person. | [optional] 

## Methods

### NewFaydaProviderOutput

`func NewFaydaProviderOutput() *FaydaProviderOutput`

NewFaydaProviderOutput instantiates a new FaydaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaydaProviderOutputWithDefaults

`func NewFaydaProviderOutputWithDefaults() *FaydaProviderOutput`

NewFaydaProviderOutputWithDefaults instantiates a new FaydaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *FaydaProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *FaydaProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *FaydaProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *FaydaProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *FaydaProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *FaydaProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



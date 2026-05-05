# SingpassMyInfoProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The Singpass subject identifier. | [optional] 
**PersonInfo** | Pointer to [**NullableSingpassPersonalCatalog**](SingpassPersonalCatalog.md) | The MyInfo personal catalog. | [optional] 

## Methods

### NewSingpassMyInfoProviderOutput

`func NewSingpassMyInfoProviderOutput() *SingpassMyInfoProviderOutput`

NewSingpassMyInfoProviderOutput instantiates a new SingpassMyInfoProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingpassMyInfoProviderOutputWithDefaults

`func NewSingpassMyInfoProviderOutputWithDefaults() *SingpassMyInfoProviderOutput`

NewSingpassMyInfoProviderOutputWithDefaults instantiates a new SingpassMyInfoProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *SingpassMyInfoProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *SingpassMyInfoProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *SingpassMyInfoProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *SingpassMyInfoProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *SingpassMyInfoProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *SingpassMyInfoProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetPersonInfo

`func (o *SingpassMyInfoProviderOutput) GetPersonInfo() SingpassPersonalCatalog`

GetPersonInfo returns the PersonInfo field if non-nil, zero value otherwise.

### GetPersonInfoOk

`func (o *SingpassMyInfoProviderOutput) GetPersonInfoOk() (*SingpassPersonalCatalog, bool)`

GetPersonInfoOk returns a tuple with the PersonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonInfo

`func (o *SingpassMyInfoProviderOutput) SetPersonInfo(v SingpassPersonalCatalog)`

SetPersonInfo sets PersonInfo field to given value.

### HasPersonInfo

`func (o *SingpassMyInfoProviderOutput) HasPersonInfo() bool`

HasPersonInfo returns a boolean if a field has been set.

### SetPersonInfoNil

`func (o *SingpassMyInfoProviderOutput) SetPersonInfoNil(b bool)`

 SetPersonInfoNil sets the value for PersonInfo to be an explicit nil

### UnsetPersonInfo
`func (o *SingpassMyInfoProviderOutput) UnsetPersonInfo()`

UnsetPersonInfo ensures that no value is present for PersonInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



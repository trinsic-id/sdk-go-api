# FaydaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | A unique eKYC identifying token used to match the original eKYC token received from the provider when the user was initially registered.              Since Fayda does not return identifying data, it is the responsibility of the relying party to keep the unique user token received from Fayda when the user was initially registered to do a comparison of the subs to verify that it is the same person. | [optional] 
**Name** | Pointer to **NullableString** | The full name of the verified individual.              This may be an English or Arabic name if the individual only has it one language, otherwise this will be null and the other names will be populated. | [optional] 
**EnglishName** | Pointer to **NullableString** | The full English name of the verified individual. | [optional] 
**ArabicName** | Pointer to **NullableString** | The full Arabic name of the verified individual. | [optional] 

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
### GetName

`func (o *FaydaProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FaydaProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FaydaProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FaydaProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FaydaProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FaydaProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnglishName

`func (o *FaydaProviderOutput) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *FaydaProviderOutput) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *FaydaProviderOutput) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.

### HasEnglishName

`func (o *FaydaProviderOutput) HasEnglishName() bool`

HasEnglishName returns a boolean if a field has been set.

### SetEnglishNameNil

`func (o *FaydaProviderOutput) SetEnglishNameNil(b bool)`

 SetEnglishNameNil sets the value for EnglishName to be an explicit nil

### UnsetEnglishName
`func (o *FaydaProviderOutput) UnsetEnglishName()`

UnsetEnglishName ensures that no value is present for EnglishName, not even an explicit nil
### GetArabicName

`func (o *FaydaProviderOutput) GetArabicName() string`

GetArabicName returns the ArabicName field if non-nil, zero value otherwise.

### GetArabicNameOk

`func (o *FaydaProviderOutput) GetArabicNameOk() (*string, bool)`

GetArabicNameOk returns a tuple with the ArabicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArabicName

`func (o *FaydaProviderOutput) SetArabicName(v string)`

SetArabicName sets ArabicName field to given value.

### HasArabicName

`func (o *FaydaProviderOutput) HasArabicName() bool`

HasArabicName returns a boolean if a field has been set.

### SetArabicNameNil

`func (o *FaydaProviderOutput) SetArabicNameNil(b bool)`

 SetArabicNameNil sets the value for ArabicName to be an explicit nil

### UnsetArabicName
`func (o *FaydaProviderOutput) UnsetArabicName()`

UnsetArabicName ensures that no value is present for ArabicName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



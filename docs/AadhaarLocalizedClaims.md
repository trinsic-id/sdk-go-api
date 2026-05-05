# AadhaarLocalizedClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**NullableAadhaarLanguage**](AadhaarLanguage.md) | The language code for the localized claims. | [optional] 
**Name** | Pointer to **NullableString** | The full name. | [optional] 
**Address** | Pointer to [**NullableAadhaarAddress**](AadhaarAddress.md) | The structured address. | [optional] 

## Methods

### NewAadhaarLocalizedClaims

`func NewAadhaarLocalizedClaims() *AadhaarLocalizedClaims`

NewAadhaarLocalizedClaims instantiates a new AadhaarLocalizedClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarLocalizedClaimsWithDefaults

`func NewAadhaarLocalizedClaimsWithDefaults() *AadhaarLocalizedClaims`

NewAadhaarLocalizedClaimsWithDefaults instantiates a new AadhaarLocalizedClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *AadhaarLocalizedClaims) GetLanguage() AadhaarLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AadhaarLocalizedClaims) GetLanguageOk() (*AadhaarLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AadhaarLocalizedClaims) SetLanguage(v AadhaarLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AadhaarLocalizedClaims) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AadhaarLocalizedClaims) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AadhaarLocalizedClaims) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetName

`func (o *AadhaarLocalizedClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AadhaarLocalizedClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AadhaarLocalizedClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AadhaarLocalizedClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AadhaarLocalizedClaims) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AadhaarLocalizedClaims) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *AadhaarLocalizedClaims) GetAddress() AadhaarAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AadhaarLocalizedClaims) GetAddressOk() (*AadhaarAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AadhaarLocalizedClaims) SetAddress(v AadhaarAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AadhaarLocalizedClaims) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *AadhaarLocalizedClaims) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *AadhaarLocalizedClaims) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



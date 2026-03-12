# OneIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The OpenID Connect (OIDC) subject identifier (sub). | [optional] 
**FullName** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family name. | [optional] 
**Birthdate** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**Address** | Pointer to [**NullableOneIdProviderAddress**](OneIdProviderAddress.md) | The address of the individual. | [optional] 

## Methods

### NewOneIdProviderOutput

`func NewOneIdProviderOutput() *OneIdProviderOutput`

NewOneIdProviderOutput instantiates a new OneIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneIdProviderOutputWithDefaults

`func NewOneIdProviderOutputWithDefaults() *OneIdProviderOutput`

NewOneIdProviderOutputWithDefaults instantiates a new OneIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *OneIdProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *OneIdProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *OneIdProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *OneIdProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *OneIdProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *OneIdProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetFullName

`func (o *OneIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OneIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OneIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OneIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *OneIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OneIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *OneIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OneIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OneIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OneIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *OneIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *OneIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *OneIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *OneIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *OneIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *OneIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *OneIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *OneIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetBirthdate

`func (o *OneIdProviderOutput) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *OneIdProviderOutput) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *OneIdProviderOutput) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *OneIdProviderOutput) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### SetBirthdateNil

`func (o *OneIdProviderOutput) SetBirthdateNil(b bool)`

 SetBirthdateNil sets the value for Birthdate to be an explicit nil

### UnsetBirthdate
`func (o *OneIdProviderOutput) UnsetBirthdate()`

UnsetBirthdate ensures that no value is present for Birthdate, not even an explicit nil
### GetAddress

`func (o *OneIdProviderOutput) GetAddress() OneIdProviderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OneIdProviderOutput) GetAddressOk() (*OneIdProviderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OneIdProviderOutput) SetAddress(v OneIdProviderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OneIdProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *OneIdProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *OneIdProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



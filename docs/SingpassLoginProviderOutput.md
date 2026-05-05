# SingpassLoginProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The Singpass subject identifier.              This is a public identifier which is globally unique across the Singpass system. | [optional] 
**AccountType** | Pointer to **NullableString** | The Singpass account type for this user.              Possible values: - foreign (Singpass Foreign Account (SFA) holders) - standard (Singapore Citizens, Permanent Residents, or FIN holders) | [optional] 
**IdentityNumber** | Pointer to **NullableString** | The Unique Identifier Number (UINFIN) for the user.              The number can be either the National Registration Identity Card or Foreign Identification Number (NRIC/FIN). Follows the format @xxxxxxx#: - @ is the status of the holder.     - Singapore citizens and permanent residents born before 1 January 2000 are assigned the letter \&quot;S\&quot;.     - Singapore citizens and permanent residents born on or after 1 January 2000 are assigned the letter \&quot;T\&quot;.     - Foreigners issued with long-term passes before 1 January 2000 are assigned the letter \&quot;F\&quot;.     - Foreigners issued with long-term passes from 1 January 2000 to 31 December 2021 are assigned the letter \&quot;G\&quot;.     - Foreigners issued with long-term passes on or after 1 January 2022 are assigned the letter \&quot;M\&quot;. - xxxxxxx is seven digit serial number. - # is the checksum letter. | [optional] 
**CountryOfIssuance** | Pointer to **NullableString** | The individual&#39;s identity country of issuance. | [optional] 
**Name** | Pointer to **NullableString** | The individual&#39;s principal name. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**MobileNumber** | Pointer to **NullableString** | The individual&#39;s mobile number.              For Singpass Foreign Account (SFA) users, this will always be null. | [optional] 

## Methods

### NewSingpassLoginProviderOutput

`func NewSingpassLoginProviderOutput() *SingpassLoginProviderOutput`

NewSingpassLoginProviderOutput instantiates a new SingpassLoginProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingpassLoginProviderOutputWithDefaults

`func NewSingpassLoginProviderOutputWithDefaults() *SingpassLoginProviderOutput`

NewSingpassLoginProviderOutputWithDefaults instantiates a new SingpassLoginProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *SingpassLoginProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *SingpassLoginProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *SingpassLoginProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *SingpassLoginProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *SingpassLoginProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *SingpassLoginProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetAccountType

`func (o *SingpassLoginProviderOutput) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SingpassLoginProviderOutput) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SingpassLoginProviderOutput) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SingpassLoginProviderOutput) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *SingpassLoginProviderOutput) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *SingpassLoginProviderOutput) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetIdentityNumber

`func (o *SingpassLoginProviderOutput) GetIdentityNumber() string`

GetIdentityNumber returns the IdentityNumber field if non-nil, zero value otherwise.

### GetIdentityNumberOk

`func (o *SingpassLoginProviderOutput) GetIdentityNumberOk() (*string, bool)`

GetIdentityNumberOk returns a tuple with the IdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityNumber

`func (o *SingpassLoginProviderOutput) SetIdentityNumber(v string)`

SetIdentityNumber sets IdentityNumber field to given value.

### HasIdentityNumber

`func (o *SingpassLoginProviderOutput) HasIdentityNumber() bool`

HasIdentityNumber returns a boolean if a field has been set.

### SetIdentityNumberNil

`func (o *SingpassLoginProviderOutput) SetIdentityNumberNil(b bool)`

 SetIdentityNumberNil sets the value for IdentityNumber to be an explicit nil

### UnsetIdentityNumber
`func (o *SingpassLoginProviderOutput) UnsetIdentityNumber()`

UnsetIdentityNumber ensures that no value is present for IdentityNumber, not even an explicit nil
### GetCountryOfIssuance

`func (o *SingpassLoginProviderOutput) GetCountryOfIssuance() string`

GetCountryOfIssuance returns the CountryOfIssuance field if non-nil, zero value otherwise.

### GetCountryOfIssuanceOk

`func (o *SingpassLoginProviderOutput) GetCountryOfIssuanceOk() (*string, bool)`

GetCountryOfIssuanceOk returns a tuple with the CountryOfIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfIssuance

`func (o *SingpassLoginProviderOutput) SetCountryOfIssuance(v string)`

SetCountryOfIssuance sets CountryOfIssuance field to given value.

### HasCountryOfIssuance

`func (o *SingpassLoginProviderOutput) HasCountryOfIssuance() bool`

HasCountryOfIssuance returns a boolean if a field has been set.

### SetCountryOfIssuanceNil

`func (o *SingpassLoginProviderOutput) SetCountryOfIssuanceNil(b bool)`

 SetCountryOfIssuanceNil sets the value for CountryOfIssuance to be an explicit nil

### UnsetCountryOfIssuance
`func (o *SingpassLoginProviderOutput) UnsetCountryOfIssuance()`

UnsetCountryOfIssuance ensures that no value is present for CountryOfIssuance, not even an explicit nil
### GetName

`func (o *SingpassLoginProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingpassLoginProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingpassLoginProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingpassLoginProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SingpassLoginProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SingpassLoginProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *SingpassLoginProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SingpassLoginProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SingpassLoginProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SingpassLoginProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SingpassLoginProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SingpassLoginProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobileNumber

`func (o *SingpassLoginProviderOutput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *SingpassLoginProviderOutput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *SingpassLoginProviderOutput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *SingpassLoginProviderOutput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *SingpassLoginProviderOutput) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *SingpassLoginProviderOutput) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



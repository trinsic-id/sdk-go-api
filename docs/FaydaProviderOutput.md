# FaydaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | A unique eKYC identifying token used to match the original eKYC token received from the provider when the individual was initially registered.              Since Fayda does not return identifying data, it is the responsibility of the relying party to keep the unique individual token received from Fayda when the individual was initially registered to do a comparison of the subs to verify that it is the same person. | [optional] 
**Name** | Pointer to **NullableString** | The full name of the verified individual.              This may be an English or Arabic name if the individual only has it one language, otherwise this will be null and the other names will be populated. | [optional] 
**EnglishName** | Pointer to **NullableString** | The full English name of the verified individual. | [optional] 
**ArabicName** | Pointer to **NullableString** | The full Arabic name of the verified individual. | [optional] 
**Birthdate** | Pointer to **NullableString** | The date of birth of the verified individual.              This attribute is only available if registered directly. | [optional] 
**Gender** | Pointer to **NullableString** | The sex of the verified individual.              Possible values: - Unknown - NotApplicable - Male - Female              This attribute is only available if registered directly. | [optional] 
**Nationality** | Pointer to **NullableString** | The nationality of the verified individual as an ISO 3166-1 alpha-2 country code.              This attribute is only available if registered directly. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the verified individual.              This attribute is only available if registered directly. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the verified individual.              This attribute is only available if registered directly. | [optional] 
**Address** | Pointer to [**NullableFaydaProviderAddress**](FaydaProviderAddress.md) | The address of the verified individual.              This attribute is only available if registered directly. | [optional] 

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
### GetBirthdate

`func (o *FaydaProviderOutput) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *FaydaProviderOutput) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *FaydaProviderOutput) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *FaydaProviderOutput) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### SetBirthdateNil

`func (o *FaydaProviderOutput) SetBirthdateNil(b bool)`

 SetBirthdateNil sets the value for Birthdate to be an explicit nil

### UnsetBirthdate
`func (o *FaydaProviderOutput) UnsetBirthdate()`

UnsetBirthdate ensures that no value is present for Birthdate, not even an explicit nil
### GetGender

`func (o *FaydaProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *FaydaProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *FaydaProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *FaydaProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *FaydaProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *FaydaProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *FaydaProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *FaydaProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *FaydaProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *FaydaProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *FaydaProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *FaydaProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetPhoneNumber

`func (o *FaydaProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FaydaProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FaydaProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FaydaProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *FaydaProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *FaydaProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *FaydaProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FaydaProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FaydaProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FaydaProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *FaydaProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FaydaProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *FaydaProviderOutput) GetAddress() FaydaProviderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FaydaProviderOutput) GetAddressOk() (*FaydaProviderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FaydaProviderOutput) SetAddress(v FaydaProviderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FaydaProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *FaydaProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *FaydaProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



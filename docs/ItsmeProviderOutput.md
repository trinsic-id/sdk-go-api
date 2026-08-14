# ItsmeProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the verified individual.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the verified individual, with a leading + country calling code.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**Sub** | Pointer to **NullableString** | The stable OpenID Connect (OIDC) subject (sub) identifier.              This should be a stable identifier, however, if a user deletes and recreates an account, this identifier will change. | [optional] 
**FullName** | Pointer to **NullableString** | The full name of the verified individual.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the verified individual.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the verified individual.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**DateOfBirthAsString** | Pointer to **NullableString** | The date of birth of the verified individual in itsme&#39;s document-facing string format.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 
**Gender** | Pointer to **NullableString** | The gender claim for the verified individual.              Availability by ID document issuing country: always returned for supported issuing countries except Netherlands, where it is best effort.              Known values: - Female - Male - Unknown | [optional] 
**Locale** | Pointer to **NullableString** | The itsme app language as an uppercase language code.              Availability by ID document issuing country: best effort for all supported issuing countries.              Known values: - NL - FR - DE - EN | [optional] 
**PictureUrl** | Pointer to **NullableString** | The URL of the profile picture resource.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether itsme reports the email address as verified.              Availability by ID document issuing country: returned only if &#x60;email&#x60; is available.              Note: itsme currently documents that this value is usually false because email verification is not implemented in its systems. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether itsme reports the phone number as verified.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**Address** | Pointer to [**NullableItsmeProviderAddress**](ItsmeProviderAddress.md) | The address of the verified individual.              Availability by ID document issuing country: always returned for Belgian-issued ID documents, best effort for Netherlands-issued ID documents, and not returned for other supported issuing countries. | [optional] 
**Citizenship** | Pointer to **NullableString** | The citizenship of the verified individual as an ISO 3166-1 alpha-2 country code.              Availability by ID document issuing country: always returned for supported issuing countries except Belgium, where it is best effort. | [optional] 
**BelgianNationalNumber** | Pointer to [**NullableItsmeBelgianNationalNumber**](ItsmeBelgianNationalNumber.md) | The Belgian National Register Number and related metadata.              Availability by ID document issuing country: returned only for Belgian-issued ID documents. | [optional] 
**BelgianIdentityCard** | Pointer to [**NullableItsmeBelgianIdentityCard**](ItsmeBelgianIdentityCard.md) | The Belgian eID card document number and related metadata.              Availability by ID document issuing country: returned only for Belgian-issued ID documents. | [optional] 
**IdentityDocument** | Pointer to [**NullableItsmeIdentityDocument**](ItsmeIdentityDocument.md) | The identity document and related metadata.              Availability by ID document issuing country: always returned for all supported issuing countries. | [optional] 
**PlaceOfBirth** | Pointer to [**NullableItsmePlaceOfBirth**](ItsmePlaceOfBirth.md) | The place of birth.              Availability by ID document issuing country: best effort for Belgian-issued ID documents; not returned for other supported issuing countries. | [optional] 
**Device** | Pointer to [**NullableItsmeDirectDevice**](ItsmeDirectDevice.md) | The device metadata for the verification.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 

## Methods

### NewItsmeProviderOutput

`func NewItsmeProviderOutput() *ItsmeProviderOutput`

NewItsmeProviderOutput instantiates a new ItsmeProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeProviderOutputWithDefaults

`func NewItsmeProviderOutputWithDefaults() *ItsmeProviderOutput`

NewItsmeProviderOutputWithDefaults instantiates a new ItsmeProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfBirth

`func (o *ItsmeProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ItsmeProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ItsmeProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ItsmeProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ItsmeProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ItsmeProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmail

`func (o *ItsmeProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ItsmeProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ItsmeProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ItsmeProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ItsmeProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ItsmeProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ItsmeProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ItsmeProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ItsmeProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ItsmeProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ItsmeProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ItsmeProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSub

`func (o *ItsmeProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *ItsmeProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *ItsmeProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *ItsmeProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *ItsmeProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *ItsmeProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetFullName

`func (o *ItsmeProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ItsmeProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ItsmeProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ItsmeProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ItsmeProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ItsmeProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *ItsmeProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ItsmeProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ItsmeProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *ItsmeProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *ItsmeProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *ItsmeProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *ItsmeProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ItsmeProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ItsmeProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ItsmeProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *ItsmeProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *ItsmeProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirthAsString

`func (o *ItsmeProviderOutput) GetDateOfBirthAsString() string`

GetDateOfBirthAsString returns the DateOfBirthAsString field if non-nil, zero value otherwise.

### GetDateOfBirthAsStringOk

`func (o *ItsmeProviderOutput) GetDateOfBirthAsStringOk() (*string, bool)`

GetDateOfBirthAsStringOk returns a tuple with the DateOfBirthAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthAsString

`func (o *ItsmeProviderOutput) SetDateOfBirthAsString(v string)`

SetDateOfBirthAsString sets DateOfBirthAsString field to given value.

### HasDateOfBirthAsString

`func (o *ItsmeProviderOutput) HasDateOfBirthAsString() bool`

HasDateOfBirthAsString returns a boolean if a field has been set.

### SetDateOfBirthAsStringNil

`func (o *ItsmeProviderOutput) SetDateOfBirthAsStringNil(b bool)`

 SetDateOfBirthAsStringNil sets the value for DateOfBirthAsString to be an explicit nil

### UnsetDateOfBirthAsString
`func (o *ItsmeProviderOutput) UnsetDateOfBirthAsString()`

UnsetDateOfBirthAsString ensures that no value is present for DateOfBirthAsString, not even an explicit nil
### GetGender

`func (o *ItsmeProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ItsmeProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ItsmeProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ItsmeProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ItsmeProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ItsmeProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetLocale

`func (o *ItsmeProviderOutput) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ItsmeProviderOutput) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ItsmeProviderOutput) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ItsmeProviderOutput) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *ItsmeProviderOutput) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *ItsmeProviderOutput) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetPictureUrl

`func (o *ItsmeProviderOutput) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *ItsmeProviderOutput) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *ItsmeProviderOutput) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.

### HasPictureUrl

`func (o *ItsmeProviderOutput) HasPictureUrl() bool`

HasPictureUrl returns a boolean if a field has been set.

### SetPictureUrlNil

`func (o *ItsmeProviderOutput) SetPictureUrlNil(b bool)`

 SetPictureUrlNil sets the value for PictureUrl to be an explicit nil

### UnsetPictureUrl
`func (o *ItsmeProviderOutput) UnsetPictureUrl()`

UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
### GetEmailVerified

`func (o *ItsmeProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *ItsmeProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *ItsmeProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *ItsmeProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *ItsmeProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *ItsmeProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhoneNumberVerified

`func (o *ItsmeProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *ItsmeProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *ItsmeProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *ItsmeProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *ItsmeProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *ItsmeProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetAddress

`func (o *ItsmeProviderOutput) GetAddress() ItsmeProviderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ItsmeProviderOutput) GetAddressOk() (*ItsmeProviderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ItsmeProviderOutput) SetAddress(v ItsmeProviderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ItsmeProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ItsmeProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ItsmeProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCitizenship

`func (o *ItsmeProviderOutput) GetCitizenship() string`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ItsmeProviderOutput) GetCitizenshipOk() (*string, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ItsmeProviderOutput) SetCitizenship(v string)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *ItsmeProviderOutput) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### SetCitizenshipNil

`func (o *ItsmeProviderOutput) SetCitizenshipNil(b bool)`

 SetCitizenshipNil sets the value for Citizenship to be an explicit nil

### UnsetCitizenship
`func (o *ItsmeProviderOutput) UnsetCitizenship()`

UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
### GetBelgianNationalNumber

`func (o *ItsmeProviderOutput) GetBelgianNationalNumber() ItsmeBelgianNationalNumber`

GetBelgianNationalNumber returns the BelgianNationalNumber field if non-nil, zero value otherwise.

### GetBelgianNationalNumberOk

`func (o *ItsmeProviderOutput) GetBelgianNationalNumberOk() (*ItsmeBelgianNationalNumber, bool)`

GetBelgianNationalNumberOk returns a tuple with the BelgianNationalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelgianNationalNumber

`func (o *ItsmeProviderOutput) SetBelgianNationalNumber(v ItsmeBelgianNationalNumber)`

SetBelgianNationalNumber sets BelgianNationalNumber field to given value.

### HasBelgianNationalNumber

`func (o *ItsmeProviderOutput) HasBelgianNationalNumber() bool`

HasBelgianNationalNumber returns a boolean if a field has been set.

### SetBelgianNationalNumberNil

`func (o *ItsmeProviderOutput) SetBelgianNationalNumberNil(b bool)`

 SetBelgianNationalNumberNil sets the value for BelgianNationalNumber to be an explicit nil

### UnsetBelgianNationalNumber
`func (o *ItsmeProviderOutput) UnsetBelgianNationalNumber()`

UnsetBelgianNationalNumber ensures that no value is present for BelgianNationalNumber, not even an explicit nil
### GetBelgianIdentityCard

`func (o *ItsmeProviderOutput) GetBelgianIdentityCard() ItsmeBelgianIdentityCard`

GetBelgianIdentityCard returns the BelgianIdentityCard field if non-nil, zero value otherwise.

### GetBelgianIdentityCardOk

`func (o *ItsmeProviderOutput) GetBelgianIdentityCardOk() (*ItsmeBelgianIdentityCard, bool)`

GetBelgianIdentityCardOk returns a tuple with the BelgianIdentityCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelgianIdentityCard

`func (o *ItsmeProviderOutput) SetBelgianIdentityCard(v ItsmeBelgianIdentityCard)`

SetBelgianIdentityCard sets BelgianIdentityCard field to given value.

### HasBelgianIdentityCard

`func (o *ItsmeProviderOutput) HasBelgianIdentityCard() bool`

HasBelgianIdentityCard returns a boolean if a field has been set.

### SetBelgianIdentityCardNil

`func (o *ItsmeProviderOutput) SetBelgianIdentityCardNil(b bool)`

 SetBelgianIdentityCardNil sets the value for BelgianIdentityCard to be an explicit nil

### UnsetBelgianIdentityCard
`func (o *ItsmeProviderOutput) UnsetBelgianIdentityCard()`

UnsetBelgianIdentityCard ensures that no value is present for BelgianIdentityCard, not even an explicit nil
### GetIdentityDocument

`func (o *ItsmeProviderOutput) GetIdentityDocument() ItsmeIdentityDocument`

GetIdentityDocument returns the IdentityDocument field if non-nil, zero value otherwise.

### GetIdentityDocumentOk

`func (o *ItsmeProviderOutput) GetIdentityDocumentOk() (*ItsmeIdentityDocument, bool)`

GetIdentityDocumentOk returns a tuple with the IdentityDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocument

`func (o *ItsmeProviderOutput) SetIdentityDocument(v ItsmeIdentityDocument)`

SetIdentityDocument sets IdentityDocument field to given value.

### HasIdentityDocument

`func (o *ItsmeProviderOutput) HasIdentityDocument() bool`

HasIdentityDocument returns a boolean if a field has been set.

### SetIdentityDocumentNil

`func (o *ItsmeProviderOutput) SetIdentityDocumentNil(b bool)`

 SetIdentityDocumentNil sets the value for IdentityDocument to be an explicit nil

### UnsetIdentityDocument
`func (o *ItsmeProviderOutput) UnsetIdentityDocument()`

UnsetIdentityDocument ensures that no value is present for IdentityDocument, not even an explicit nil
### GetPlaceOfBirth

`func (o *ItsmeProviderOutput) GetPlaceOfBirth() ItsmePlaceOfBirth`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ItsmeProviderOutput) GetPlaceOfBirthOk() (*ItsmePlaceOfBirth, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ItsmeProviderOutput) SetPlaceOfBirth(v ItsmePlaceOfBirth)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ItsmeProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ItsmeProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ItsmeProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetDevice

`func (o *ItsmeProviderOutput) GetDevice() ItsmeDirectDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ItsmeProviderOutput) GetDeviceOk() (*ItsmeDirectDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ItsmeProviderOutput) SetDevice(v ItsmeDirectDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ItsmeProviderOutput) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ItsmeProviderOutput) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ItsmeProviderOutput) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



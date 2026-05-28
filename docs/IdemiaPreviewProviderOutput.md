# IdemiaPreviewProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | The OpenID Connect subject identifier for the verified individual. | [optional] 
**Name** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given name. | [optional] 
**MiddleName** | Pointer to **NullableString** | The individual&#39;s middle name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family name. | [optional] 
**Nickname** | Pointer to **NullableString** | The individual&#39;s nickname. | [optional] 
**PreferredUsername** | Pointer to **NullableString** | The individual&#39;s preferred username. | [optional] 
**Profile** | Pointer to **NullableString** | The individual&#39;s profile URL. | [optional] 
**Picture** | Pointer to **NullableString** | URL of the individual&#39;s profile picture. | [optional] 
**Website** | Pointer to **NullableString** | The individual&#39;s website URL. | [optional] 
**Birthdate** | Pointer to **NullableString** | The individual&#39;s date of birth in YYYY-MM-DD format. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth as returned by legacy Idemia claim sets. | [optional] 
**Gender** | Pointer to **NullableString** | The individual&#39;s gender. | [optional] 
**Zoneinfo** | Pointer to **NullableString** | The individual&#39;s time zone. | [optional] 
**Locale** | Pointer to **NullableString** | The individual&#39;s locale. | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The individual&#39;s preferred language. | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | Timestamp when the individual&#39;s profile data was last updated. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the individual&#39;s email address has been verified. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The individual&#39;s phone number. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether the individual&#39;s phone number has been verified. | [optional] 
**MobilePhoneNumber** | Pointer to **NullableString** | The individual&#39;s mobile phone number returned by the mobile_phone scope. | [optional] 
**Address** | Pointer to [**NullableIdemiaPreviewProviderAddress**](IdemiaPreviewProviderAddress.md) | The individual&#39;s address. | [optional] 
**Pcr** | Pointer to **NullableString** | Idemia pairwise relying-party correlation reference. | [optional] 
**MidUid** | Pointer to **NullableString** | Idemia Mobile ID unique identifier. | [optional] 
**UserName** | Pointer to **NullableString** | The user name returned by Idemia. | [optional] 
**Inum** | Pointer to **NullableString** | Idemia client information identifier. | [optional] 
**Role** | Pointer to **NullableString** | Permission role returned by Idemia. | [optional] 
**PersonalIdentityCode** | Pointer to **NullableString** | Personal identity code returned by Idemia when available. | [optional] 
**PersonIdentityCodeSample** | Pointer to **NullableString** | Sample personal identity code returned by Idemia when available. | [optional] 

## Methods

### NewIdemiaPreviewProviderOutput

`func NewIdemiaPreviewProviderOutput() *IdemiaPreviewProviderOutput`

NewIdemiaPreviewProviderOutput instantiates a new IdemiaPreviewProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdemiaPreviewProviderOutputWithDefaults

`func NewIdemiaPreviewProviderOutputWithDefaults() *IdemiaPreviewProviderOutput`

NewIdemiaPreviewProviderOutputWithDefaults instantiates a new IdemiaPreviewProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *IdemiaPreviewProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IdemiaPreviewProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IdemiaPreviewProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IdemiaPreviewProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *IdemiaPreviewProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *IdemiaPreviewProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetName

`func (o *IdemiaPreviewProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdemiaPreviewProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdemiaPreviewProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdemiaPreviewProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdemiaPreviewProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdemiaPreviewProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGivenName

`func (o *IdemiaPreviewProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IdemiaPreviewProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IdemiaPreviewProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *IdemiaPreviewProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *IdemiaPreviewProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *IdemiaPreviewProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *IdemiaPreviewProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IdemiaPreviewProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IdemiaPreviewProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IdemiaPreviewProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *IdemiaPreviewProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *IdemiaPreviewProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *IdemiaPreviewProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IdemiaPreviewProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IdemiaPreviewProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *IdemiaPreviewProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *IdemiaPreviewProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *IdemiaPreviewProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetNickname

`func (o *IdemiaPreviewProviderOutput) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *IdemiaPreviewProviderOutput) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *IdemiaPreviewProviderOutput) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *IdemiaPreviewProviderOutput) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *IdemiaPreviewProviderOutput) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *IdemiaPreviewProviderOutput) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetPreferredUsername

`func (o *IdemiaPreviewProviderOutput) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *IdemiaPreviewProviderOutput) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *IdemiaPreviewProviderOutput) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *IdemiaPreviewProviderOutput) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### SetPreferredUsernameNil

`func (o *IdemiaPreviewProviderOutput) SetPreferredUsernameNil(b bool)`

 SetPreferredUsernameNil sets the value for PreferredUsername to be an explicit nil

### UnsetPreferredUsername
`func (o *IdemiaPreviewProviderOutput) UnsetPreferredUsername()`

UnsetPreferredUsername ensures that no value is present for PreferredUsername, not even an explicit nil
### GetProfile

`func (o *IdemiaPreviewProviderOutput) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IdemiaPreviewProviderOutput) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IdemiaPreviewProviderOutput) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *IdemiaPreviewProviderOutput) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *IdemiaPreviewProviderOutput) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *IdemiaPreviewProviderOutput) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetPicture

`func (o *IdemiaPreviewProviderOutput) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *IdemiaPreviewProviderOutput) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *IdemiaPreviewProviderOutput) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *IdemiaPreviewProviderOutput) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### SetPictureNil

`func (o *IdemiaPreviewProviderOutput) SetPictureNil(b bool)`

 SetPictureNil sets the value for Picture to be an explicit nil

### UnsetPicture
`func (o *IdemiaPreviewProviderOutput) UnsetPicture()`

UnsetPicture ensures that no value is present for Picture, not even an explicit nil
### GetWebsite

`func (o *IdemiaPreviewProviderOutput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *IdemiaPreviewProviderOutput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *IdemiaPreviewProviderOutput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *IdemiaPreviewProviderOutput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *IdemiaPreviewProviderOutput) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *IdemiaPreviewProviderOutput) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetBirthdate

`func (o *IdemiaPreviewProviderOutput) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *IdemiaPreviewProviderOutput) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *IdemiaPreviewProviderOutput) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *IdemiaPreviewProviderOutput) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### SetBirthdateNil

`func (o *IdemiaPreviewProviderOutput) SetBirthdateNil(b bool)`

 SetBirthdateNil sets the value for Birthdate to be an explicit nil

### UnsetBirthdate
`func (o *IdemiaPreviewProviderOutput) UnsetBirthdate()`

UnsetBirthdate ensures that no value is present for Birthdate, not even an explicit nil
### GetDateOfBirth

`func (o *IdemiaPreviewProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IdemiaPreviewProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IdemiaPreviewProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IdemiaPreviewProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IdemiaPreviewProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IdemiaPreviewProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *IdemiaPreviewProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IdemiaPreviewProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IdemiaPreviewProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IdemiaPreviewProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *IdemiaPreviewProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *IdemiaPreviewProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetZoneinfo

`func (o *IdemiaPreviewProviderOutput) GetZoneinfo() string`

GetZoneinfo returns the Zoneinfo field if non-nil, zero value otherwise.

### GetZoneinfoOk

`func (o *IdemiaPreviewProviderOutput) GetZoneinfoOk() (*string, bool)`

GetZoneinfoOk returns a tuple with the Zoneinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneinfo

`func (o *IdemiaPreviewProviderOutput) SetZoneinfo(v string)`

SetZoneinfo sets Zoneinfo field to given value.

### HasZoneinfo

`func (o *IdemiaPreviewProviderOutput) HasZoneinfo() bool`

HasZoneinfo returns a boolean if a field has been set.

### SetZoneinfoNil

`func (o *IdemiaPreviewProviderOutput) SetZoneinfoNil(b bool)`

 SetZoneinfoNil sets the value for Zoneinfo to be an explicit nil

### UnsetZoneinfo
`func (o *IdemiaPreviewProviderOutput) UnsetZoneinfo()`

UnsetZoneinfo ensures that no value is present for Zoneinfo, not even an explicit nil
### GetLocale

`func (o *IdemiaPreviewProviderOutput) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *IdemiaPreviewProviderOutput) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *IdemiaPreviewProviderOutput) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *IdemiaPreviewProviderOutput) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *IdemiaPreviewProviderOutput) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *IdemiaPreviewProviderOutput) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetPreferredLanguage

`func (o *IdemiaPreviewProviderOutput) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *IdemiaPreviewProviderOutput) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *IdemiaPreviewProviderOutput) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *IdemiaPreviewProviderOutput) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *IdemiaPreviewProviderOutput) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *IdemiaPreviewProviderOutput) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetUpdatedAt

`func (o *IdemiaPreviewProviderOutput) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdemiaPreviewProviderOutput) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdemiaPreviewProviderOutput) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdemiaPreviewProviderOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IdemiaPreviewProviderOutput) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IdemiaPreviewProviderOutput) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetEmail

`func (o *IdemiaPreviewProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdemiaPreviewProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdemiaPreviewProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdemiaPreviewProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IdemiaPreviewProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IdemiaPreviewProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *IdemiaPreviewProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *IdemiaPreviewProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *IdemiaPreviewProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *IdemiaPreviewProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *IdemiaPreviewProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *IdemiaPreviewProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhoneNumber

`func (o *IdemiaPreviewProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdemiaPreviewProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdemiaPreviewProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdemiaPreviewProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *IdemiaPreviewProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *IdemiaPreviewProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *IdemiaPreviewProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *IdemiaPreviewProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *IdemiaPreviewProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *IdemiaPreviewProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *IdemiaPreviewProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *IdemiaPreviewProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetMobilePhoneNumber

`func (o *IdemiaPreviewProviderOutput) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *IdemiaPreviewProviderOutput) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *IdemiaPreviewProviderOutput) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *IdemiaPreviewProviderOutput) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### SetMobilePhoneNumberNil

`func (o *IdemiaPreviewProviderOutput) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *IdemiaPreviewProviderOutput) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetAddress

`func (o *IdemiaPreviewProviderOutput) GetAddress() IdemiaPreviewProviderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IdemiaPreviewProviderOutput) GetAddressOk() (*IdemiaPreviewProviderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IdemiaPreviewProviderOutput) SetAddress(v IdemiaPreviewProviderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IdemiaPreviewProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *IdemiaPreviewProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *IdemiaPreviewProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPcr

`func (o *IdemiaPreviewProviderOutput) GetPcr() string`

GetPcr returns the Pcr field if non-nil, zero value otherwise.

### GetPcrOk

`func (o *IdemiaPreviewProviderOutput) GetPcrOk() (*string, bool)`

GetPcrOk returns a tuple with the Pcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcr

`func (o *IdemiaPreviewProviderOutput) SetPcr(v string)`

SetPcr sets Pcr field to given value.

### HasPcr

`func (o *IdemiaPreviewProviderOutput) HasPcr() bool`

HasPcr returns a boolean if a field has been set.

### SetPcrNil

`func (o *IdemiaPreviewProviderOutput) SetPcrNil(b bool)`

 SetPcrNil sets the value for Pcr to be an explicit nil

### UnsetPcr
`func (o *IdemiaPreviewProviderOutput) UnsetPcr()`

UnsetPcr ensures that no value is present for Pcr, not even an explicit nil
### GetMidUid

`func (o *IdemiaPreviewProviderOutput) GetMidUid() string`

GetMidUid returns the MidUid field if non-nil, zero value otherwise.

### GetMidUidOk

`func (o *IdemiaPreviewProviderOutput) GetMidUidOk() (*string, bool)`

GetMidUidOk returns a tuple with the MidUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidUid

`func (o *IdemiaPreviewProviderOutput) SetMidUid(v string)`

SetMidUid sets MidUid field to given value.

### HasMidUid

`func (o *IdemiaPreviewProviderOutput) HasMidUid() bool`

HasMidUid returns a boolean if a field has been set.

### SetMidUidNil

`func (o *IdemiaPreviewProviderOutput) SetMidUidNil(b bool)`

 SetMidUidNil sets the value for MidUid to be an explicit nil

### UnsetMidUid
`func (o *IdemiaPreviewProviderOutput) UnsetMidUid()`

UnsetMidUid ensures that no value is present for MidUid, not even an explicit nil
### GetUserName

`func (o *IdemiaPreviewProviderOutput) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IdemiaPreviewProviderOutput) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IdemiaPreviewProviderOutput) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *IdemiaPreviewProviderOutput) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *IdemiaPreviewProviderOutput) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *IdemiaPreviewProviderOutput) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetInum

`func (o *IdemiaPreviewProviderOutput) GetInum() string`

GetInum returns the Inum field if non-nil, zero value otherwise.

### GetInumOk

`func (o *IdemiaPreviewProviderOutput) GetInumOk() (*string, bool)`

GetInumOk returns a tuple with the Inum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInum

`func (o *IdemiaPreviewProviderOutput) SetInum(v string)`

SetInum sets Inum field to given value.

### HasInum

`func (o *IdemiaPreviewProviderOutput) HasInum() bool`

HasInum returns a boolean if a field has been set.

### SetInumNil

`func (o *IdemiaPreviewProviderOutput) SetInumNil(b bool)`

 SetInumNil sets the value for Inum to be an explicit nil

### UnsetInum
`func (o *IdemiaPreviewProviderOutput) UnsetInum()`

UnsetInum ensures that no value is present for Inum, not even an explicit nil
### GetRole

`func (o *IdemiaPreviewProviderOutput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IdemiaPreviewProviderOutput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IdemiaPreviewProviderOutput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IdemiaPreviewProviderOutput) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IdemiaPreviewProviderOutput) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IdemiaPreviewProviderOutput) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPersonalIdentityCode

`func (o *IdemiaPreviewProviderOutput) GetPersonalIdentityCode() string`

GetPersonalIdentityCode returns the PersonalIdentityCode field if non-nil, zero value otherwise.

### GetPersonalIdentityCodeOk

`func (o *IdemiaPreviewProviderOutput) GetPersonalIdentityCodeOk() (*string, bool)`

GetPersonalIdentityCodeOk returns a tuple with the PersonalIdentityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentityCode

`func (o *IdemiaPreviewProviderOutput) SetPersonalIdentityCode(v string)`

SetPersonalIdentityCode sets PersonalIdentityCode field to given value.

### HasPersonalIdentityCode

`func (o *IdemiaPreviewProviderOutput) HasPersonalIdentityCode() bool`

HasPersonalIdentityCode returns a boolean if a field has been set.

### SetPersonalIdentityCodeNil

`func (o *IdemiaPreviewProviderOutput) SetPersonalIdentityCodeNil(b bool)`

 SetPersonalIdentityCodeNil sets the value for PersonalIdentityCode to be an explicit nil

### UnsetPersonalIdentityCode
`func (o *IdemiaPreviewProviderOutput) UnsetPersonalIdentityCode()`

UnsetPersonalIdentityCode ensures that no value is present for PersonalIdentityCode, not even an explicit nil
### GetPersonIdentityCodeSample

`func (o *IdemiaPreviewProviderOutput) GetPersonIdentityCodeSample() string`

GetPersonIdentityCodeSample returns the PersonIdentityCodeSample field if non-nil, zero value otherwise.

### GetPersonIdentityCodeSampleOk

`func (o *IdemiaPreviewProviderOutput) GetPersonIdentityCodeSampleOk() (*string, bool)`

GetPersonIdentityCodeSampleOk returns a tuple with the PersonIdentityCodeSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIdentityCodeSample

`func (o *IdemiaPreviewProviderOutput) SetPersonIdentityCodeSample(v string)`

SetPersonIdentityCodeSample sets PersonIdentityCodeSample field to given value.

### HasPersonIdentityCodeSample

`func (o *IdemiaPreviewProviderOutput) HasPersonIdentityCodeSample() bool`

HasPersonIdentityCodeSample returns a boolean if a field has been set.

### SetPersonIdentityCodeSampleNil

`func (o *IdemiaPreviewProviderOutput) SetPersonIdentityCodeSampleNil(b bool)`

 SetPersonIdentityCodeSampleNil sets the value for PersonIdentityCodeSample to be an explicit nil

### UnsetPersonIdentityCodeSample
`func (o *IdemiaPreviewProviderOutput) UnsetPersonIdentityCodeSample()`

UnsetPersonIdentityCodeSample ensures that no value is present for PersonIdentityCodeSample, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UAEPassProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | Pointer to **NullableString** | OIDC subject (&#x60;sub&#x60;) returned by UAE Pass.              This is an opaque identifier for the authenticated user in the UAE Pass OIDC transaction. It can be prefixed or unprefixed; do not parse this value or assume it matches the Emirates ID number, UAE Pass UUID, or UAE unified ID. | [optional] 
**FullNameArabic** | Pointer to **NullableString** | Full name in Arabic. | [optional] 
**Gender** | Pointer to **NullableString** | Normalized sex parsed from the UAE Pass gender value.              Possible values: - Male - Female - Unknown - NotApplicable              This is null when the UAE Pass value cannot be parsed. | [optional] 
**Mobile** | Pointer to **NullableString** | Mobile number. | [optional] 
**LastNameEnglish** | Pointer to **NullableString** | Last name in the Latin-script value returned by UAE Pass. | [optional] 
**FullNameEnglish** | Pointer to **NullableString** | Full name in the Latin-script value returned by UAE Pass. | [optional] 
**UaePassUuid** | Pointer to **NullableString** | UAE Pass UUID.              The unique UAE Pass user identifier. It is separate from the OIDC subject identifier, Emirates ID number, and UAE unified ID. | [optional] 
**SmartPassUuid** | Pointer to **NullableString** | SmartPass unique user identifier, if returned.              SmartPass was a previous UAE government single-sign-on identity system. UAE Pass can return this legacy identifier for SOP2 and SOP3 accounts that were verified through SmartPass. | [optional] 
**LastNameArabic** | Pointer to **NullableString** | Last name in Arabic. | [optional] 
**EmiratesIdNumber** | Pointer to **NullableString** |  Verified Emirates ID number, if returned for the user profile type and scope.   The Emirates ID is issued by the Federal Authority for Identity, Citizenship, Customs and Port  Security (ICP). The number is a 15-digit identifier following the format 784-YYYY-NNNNNNN-C:  The number is a 15-digit identifier following the format 784-YYYY-NNNNNNN-C:  - 784: UAE ISO 3166-1 numeric country code.  - YYYY: A year value, which could be birth year, registration year, or another year value. This should      not be relied upon to infer birth or registration details.  - NNNNNNN: 7-digit random serial number for the individual.  - C: Check digit.   This field may not be returned for visitor profiles. | [optional] 
**IdType** | Pointer to **NullableString** | ID type.              Known values: - ID: Emirates ID or national identity document.              Returned for SOP2 and SOP3 citizen or resident profiles and SOP3 visitor profiles. | [optional] 
**NationalityEnglish** | Pointer to **NullableString** | Nationality in the Latin-script value returned by UAE Pass. | [optional] 
**FirstNameEnglish** | Pointer to **NullableString** | First name in the Latin-script value returned by UAE Pass. | [optional] 
**UserType** | Pointer to **NullableString** | UAE Pass user type.              Possible values: - SOP1: Basic account. Unverified account with verified email and mobile number. - SOP2: Advanced account. Verified account with Emirates ID verification through SmartPass,     Dubai ID, or Emirates ID PIN registration. - SOP3: Qualified account. Verified account with Emirates ID verification through finger     biometrics or face biometrics. | [optional] 
**NationalityArabic** | Pointer to **NullableString** | Nationality in Arabic. | [optional] 
**FirstNameArabic** | Pointer to **NullableString** | First name in Arabic. | [optional] 
**Email** | Pointer to **NullableString** | Email address. | [optional] 
**TitleEnglish** | Pointer to **NullableString** | Title in the Latin-script value returned by UAE Pass.              Returned for SOP2 and SOP3 citizen or resident profiles and SOP3 visitor profiles. | [optional] 
**TitleArabic** | Pointer to **NullableString** | Title in Arabic.              Returned for SOP2 and SOP3 citizen or resident profiles and SOP3 visitor profiles. | [optional] 
**ProfileType** | Pointer to **NullableInt32** | UAE Pass profile type, if returned for the requested scope.              Possible values: - 1: Citizen or resident profile. - 2: Visitor profile. | [optional] 
**UnifiedId** | Pointer to **NullableString** | UAE unified ID, if returned for the requested scope.              A unique ID for the user. It is separate from the Emirates ID number, UAE Pass UUID, and OIDC subject identifier. | [optional] 
**AuthenticationAssuranceLevel** | Pointer to **NullableString** | UAE Pass authentication assurance level, if present.              This classifies the authentication policy that was satisfied for the session. Known values include: - urn:safelayer:tws:policies:authentication:level:low - urn:safelayer:tws:policies:authentication:level:high | [optional] 
**AuthenticationMethods** | Pointer to **[]string** | UAE Pass authentication methods, if present.              This classifies the concrete methods used during authentication. Known values include: - urn:safelayer:tws:policies:authentication:adaptive:methods:mobileid - urn:uae:authentication:method:verified - urn:oasis:names:tc:SAML:1:0:am:password | [optional] 

## Methods

### NewUAEPassProviderOutput

`func NewUAEPassProviderOutput() *UAEPassProviderOutput`

NewUAEPassProviderOutput instantiates a new UAEPassProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUAEPassProviderOutputWithDefaults

`func NewUAEPassProviderOutputWithDefaults() *UAEPassProviderOutput`

NewUAEPassProviderOutputWithDefaults instantiates a new UAEPassProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *UAEPassProviderOutput) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UAEPassProviderOutput) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UAEPassProviderOutput) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *UAEPassProviderOutput) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *UAEPassProviderOutput) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *UAEPassProviderOutput) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetFullNameArabic

`func (o *UAEPassProviderOutput) GetFullNameArabic() string`

GetFullNameArabic returns the FullNameArabic field if non-nil, zero value otherwise.

### GetFullNameArabicOk

`func (o *UAEPassProviderOutput) GetFullNameArabicOk() (*string, bool)`

GetFullNameArabicOk returns a tuple with the FullNameArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNameArabic

`func (o *UAEPassProviderOutput) SetFullNameArabic(v string)`

SetFullNameArabic sets FullNameArabic field to given value.

### HasFullNameArabic

`func (o *UAEPassProviderOutput) HasFullNameArabic() bool`

HasFullNameArabic returns a boolean if a field has been set.

### SetFullNameArabicNil

`func (o *UAEPassProviderOutput) SetFullNameArabicNil(b bool)`

 SetFullNameArabicNil sets the value for FullNameArabic to be an explicit nil

### UnsetFullNameArabic
`func (o *UAEPassProviderOutput) UnsetFullNameArabic()`

UnsetFullNameArabic ensures that no value is present for FullNameArabic, not even an explicit nil
### GetGender

`func (o *UAEPassProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UAEPassProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UAEPassProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UAEPassProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *UAEPassProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *UAEPassProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetMobile

`func (o *UAEPassProviderOutput) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UAEPassProviderOutput) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UAEPassProviderOutput) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UAEPassProviderOutput) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### SetMobileNil

`func (o *UAEPassProviderOutput) SetMobileNil(b bool)`

 SetMobileNil sets the value for Mobile to be an explicit nil

### UnsetMobile
`func (o *UAEPassProviderOutput) UnsetMobile()`

UnsetMobile ensures that no value is present for Mobile, not even an explicit nil
### GetLastNameEnglish

`func (o *UAEPassProviderOutput) GetLastNameEnglish() string`

GetLastNameEnglish returns the LastNameEnglish field if non-nil, zero value otherwise.

### GetLastNameEnglishOk

`func (o *UAEPassProviderOutput) GetLastNameEnglishOk() (*string, bool)`

GetLastNameEnglishOk returns a tuple with the LastNameEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameEnglish

`func (o *UAEPassProviderOutput) SetLastNameEnglish(v string)`

SetLastNameEnglish sets LastNameEnglish field to given value.

### HasLastNameEnglish

`func (o *UAEPassProviderOutput) HasLastNameEnglish() bool`

HasLastNameEnglish returns a boolean if a field has been set.

### SetLastNameEnglishNil

`func (o *UAEPassProviderOutput) SetLastNameEnglishNil(b bool)`

 SetLastNameEnglishNil sets the value for LastNameEnglish to be an explicit nil

### UnsetLastNameEnglish
`func (o *UAEPassProviderOutput) UnsetLastNameEnglish()`

UnsetLastNameEnglish ensures that no value is present for LastNameEnglish, not even an explicit nil
### GetFullNameEnglish

`func (o *UAEPassProviderOutput) GetFullNameEnglish() string`

GetFullNameEnglish returns the FullNameEnglish field if non-nil, zero value otherwise.

### GetFullNameEnglishOk

`func (o *UAEPassProviderOutput) GetFullNameEnglishOk() (*string, bool)`

GetFullNameEnglishOk returns a tuple with the FullNameEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNameEnglish

`func (o *UAEPassProviderOutput) SetFullNameEnglish(v string)`

SetFullNameEnglish sets FullNameEnglish field to given value.

### HasFullNameEnglish

`func (o *UAEPassProviderOutput) HasFullNameEnglish() bool`

HasFullNameEnglish returns a boolean if a field has been set.

### SetFullNameEnglishNil

`func (o *UAEPassProviderOutput) SetFullNameEnglishNil(b bool)`

 SetFullNameEnglishNil sets the value for FullNameEnglish to be an explicit nil

### UnsetFullNameEnglish
`func (o *UAEPassProviderOutput) UnsetFullNameEnglish()`

UnsetFullNameEnglish ensures that no value is present for FullNameEnglish, not even an explicit nil
### GetUaePassUuid

`func (o *UAEPassProviderOutput) GetUaePassUuid() string`

GetUaePassUuid returns the UaePassUuid field if non-nil, zero value otherwise.

### GetUaePassUuidOk

`func (o *UAEPassProviderOutput) GetUaePassUuidOk() (*string, bool)`

GetUaePassUuidOk returns a tuple with the UaePassUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaePassUuid

`func (o *UAEPassProviderOutput) SetUaePassUuid(v string)`

SetUaePassUuid sets UaePassUuid field to given value.

### HasUaePassUuid

`func (o *UAEPassProviderOutput) HasUaePassUuid() bool`

HasUaePassUuid returns a boolean if a field has been set.

### SetUaePassUuidNil

`func (o *UAEPassProviderOutput) SetUaePassUuidNil(b bool)`

 SetUaePassUuidNil sets the value for UaePassUuid to be an explicit nil

### UnsetUaePassUuid
`func (o *UAEPassProviderOutput) UnsetUaePassUuid()`

UnsetUaePassUuid ensures that no value is present for UaePassUuid, not even an explicit nil
### GetSmartPassUuid

`func (o *UAEPassProviderOutput) GetSmartPassUuid() string`

GetSmartPassUuid returns the SmartPassUuid field if non-nil, zero value otherwise.

### GetSmartPassUuidOk

`func (o *UAEPassProviderOutput) GetSmartPassUuidOk() (*string, bool)`

GetSmartPassUuidOk returns a tuple with the SmartPassUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartPassUuid

`func (o *UAEPassProviderOutput) SetSmartPassUuid(v string)`

SetSmartPassUuid sets SmartPassUuid field to given value.

### HasSmartPassUuid

`func (o *UAEPassProviderOutput) HasSmartPassUuid() bool`

HasSmartPassUuid returns a boolean if a field has been set.

### SetSmartPassUuidNil

`func (o *UAEPassProviderOutput) SetSmartPassUuidNil(b bool)`

 SetSmartPassUuidNil sets the value for SmartPassUuid to be an explicit nil

### UnsetSmartPassUuid
`func (o *UAEPassProviderOutput) UnsetSmartPassUuid()`

UnsetSmartPassUuid ensures that no value is present for SmartPassUuid, not even an explicit nil
### GetLastNameArabic

`func (o *UAEPassProviderOutput) GetLastNameArabic() string`

GetLastNameArabic returns the LastNameArabic field if non-nil, zero value otherwise.

### GetLastNameArabicOk

`func (o *UAEPassProviderOutput) GetLastNameArabicOk() (*string, bool)`

GetLastNameArabicOk returns a tuple with the LastNameArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameArabic

`func (o *UAEPassProviderOutput) SetLastNameArabic(v string)`

SetLastNameArabic sets LastNameArabic field to given value.

### HasLastNameArabic

`func (o *UAEPassProviderOutput) HasLastNameArabic() bool`

HasLastNameArabic returns a boolean if a field has been set.

### SetLastNameArabicNil

`func (o *UAEPassProviderOutput) SetLastNameArabicNil(b bool)`

 SetLastNameArabicNil sets the value for LastNameArabic to be an explicit nil

### UnsetLastNameArabic
`func (o *UAEPassProviderOutput) UnsetLastNameArabic()`

UnsetLastNameArabic ensures that no value is present for LastNameArabic, not even an explicit nil
### GetEmiratesIdNumber

`func (o *UAEPassProviderOutput) GetEmiratesIdNumber() string`

GetEmiratesIdNumber returns the EmiratesIdNumber field if non-nil, zero value otherwise.

### GetEmiratesIdNumberOk

`func (o *UAEPassProviderOutput) GetEmiratesIdNumberOk() (*string, bool)`

GetEmiratesIdNumberOk returns a tuple with the EmiratesIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiratesIdNumber

`func (o *UAEPassProviderOutput) SetEmiratesIdNumber(v string)`

SetEmiratesIdNumber sets EmiratesIdNumber field to given value.

### HasEmiratesIdNumber

`func (o *UAEPassProviderOutput) HasEmiratesIdNumber() bool`

HasEmiratesIdNumber returns a boolean if a field has been set.

### SetEmiratesIdNumberNil

`func (o *UAEPassProviderOutput) SetEmiratesIdNumberNil(b bool)`

 SetEmiratesIdNumberNil sets the value for EmiratesIdNumber to be an explicit nil

### UnsetEmiratesIdNumber
`func (o *UAEPassProviderOutput) UnsetEmiratesIdNumber()`

UnsetEmiratesIdNumber ensures that no value is present for EmiratesIdNumber, not even an explicit nil
### GetIdType

`func (o *UAEPassProviderOutput) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *UAEPassProviderOutput) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *UAEPassProviderOutput) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *UAEPassProviderOutput) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### SetIdTypeNil

`func (o *UAEPassProviderOutput) SetIdTypeNil(b bool)`

 SetIdTypeNil sets the value for IdType to be an explicit nil

### UnsetIdType
`func (o *UAEPassProviderOutput) UnsetIdType()`

UnsetIdType ensures that no value is present for IdType, not even an explicit nil
### GetNationalityEnglish

`func (o *UAEPassProviderOutput) GetNationalityEnglish() string`

GetNationalityEnglish returns the NationalityEnglish field if non-nil, zero value otherwise.

### GetNationalityEnglishOk

`func (o *UAEPassProviderOutput) GetNationalityEnglishOk() (*string, bool)`

GetNationalityEnglishOk returns a tuple with the NationalityEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityEnglish

`func (o *UAEPassProviderOutput) SetNationalityEnglish(v string)`

SetNationalityEnglish sets NationalityEnglish field to given value.

### HasNationalityEnglish

`func (o *UAEPassProviderOutput) HasNationalityEnglish() bool`

HasNationalityEnglish returns a boolean if a field has been set.

### SetNationalityEnglishNil

`func (o *UAEPassProviderOutput) SetNationalityEnglishNil(b bool)`

 SetNationalityEnglishNil sets the value for NationalityEnglish to be an explicit nil

### UnsetNationalityEnglish
`func (o *UAEPassProviderOutput) UnsetNationalityEnglish()`

UnsetNationalityEnglish ensures that no value is present for NationalityEnglish, not even an explicit nil
### GetFirstNameEnglish

`func (o *UAEPassProviderOutput) GetFirstNameEnglish() string`

GetFirstNameEnglish returns the FirstNameEnglish field if non-nil, zero value otherwise.

### GetFirstNameEnglishOk

`func (o *UAEPassProviderOutput) GetFirstNameEnglishOk() (*string, bool)`

GetFirstNameEnglishOk returns a tuple with the FirstNameEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameEnglish

`func (o *UAEPassProviderOutput) SetFirstNameEnglish(v string)`

SetFirstNameEnglish sets FirstNameEnglish field to given value.

### HasFirstNameEnglish

`func (o *UAEPassProviderOutput) HasFirstNameEnglish() bool`

HasFirstNameEnglish returns a boolean if a field has been set.

### SetFirstNameEnglishNil

`func (o *UAEPassProviderOutput) SetFirstNameEnglishNil(b bool)`

 SetFirstNameEnglishNil sets the value for FirstNameEnglish to be an explicit nil

### UnsetFirstNameEnglish
`func (o *UAEPassProviderOutput) UnsetFirstNameEnglish()`

UnsetFirstNameEnglish ensures that no value is present for FirstNameEnglish, not even an explicit nil
### GetUserType

`func (o *UAEPassProviderOutput) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UAEPassProviderOutput) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UAEPassProviderOutput) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UAEPassProviderOutput) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *UAEPassProviderOutput) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UAEPassProviderOutput) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetNationalityArabic

`func (o *UAEPassProviderOutput) GetNationalityArabic() string`

GetNationalityArabic returns the NationalityArabic field if non-nil, zero value otherwise.

### GetNationalityArabicOk

`func (o *UAEPassProviderOutput) GetNationalityArabicOk() (*string, bool)`

GetNationalityArabicOk returns a tuple with the NationalityArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityArabic

`func (o *UAEPassProviderOutput) SetNationalityArabic(v string)`

SetNationalityArabic sets NationalityArabic field to given value.

### HasNationalityArabic

`func (o *UAEPassProviderOutput) HasNationalityArabic() bool`

HasNationalityArabic returns a boolean if a field has been set.

### SetNationalityArabicNil

`func (o *UAEPassProviderOutput) SetNationalityArabicNil(b bool)`

 SetNationalityArabicNil sets the value for NationalityArabic to be an explicit nil

### UnsetNationalityArabic
`func (o *UAEPassProviderOutput) UnsetNationalityArabic()`

UnsetNationalityArabic ensures that no value is present for NationalityArabic, not even an explicit nil
### GetFirstNameArabic

`func (o *UAEPassProviderOutput) GetFirstNameArabic() string`

GetFirstNameArabic returns the FirstNameArabic field if non-nil, zero value otherwise.

### GetFirstNameArabicOk

`func (o *UAEPassProviderOutput) GetFirstNameArabicOk() (*string, bool)`

GetFirstNameArabicOk returns a tuple with the FirstNameArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameArabic

`func (o *UAEPassProviderOutput) SetFirstNameArabic(v string)`

SetFirstNameArabic sets FirstNameArabic field to given value.

### HasFirstNameArabic

`func (o *UAEPassProviderOutput) HasFirstNameArabic() bool`

HasFirstNameArabic returns a boolean if a field has been set.

### SetFirstNameArabicNil

`func (o *UAEPassProviderOutput) SetFirstNameArabicNil(b bool)`

 SetFirstNameArabicNil sets the value for FirstNameArabic to be an explicit nil

### UnsetFirstNameArabic
`func (o *UAEPassProviderOutput) UnsetFirstNameArabic()`

UnsetFirstNameArabic ensures that no value is present for FirstNameArabic, not even an explicit nil
### GetEmail

`func (o *UAEPassProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UAEPassProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UAEPassProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UAEPassProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UAEPassProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UAEPassProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTitleEnglish

`func (o *UAEPassProviderOutput) GetTitleEnglish() string`

GetTitleEnglish returns the TitleEnglish field if non-nil, zero value otherwise.

### GetTitleEnglishOk

`func (o *UAEPassProviderOutput) GetTitleEnglishOk() (*string, bool)`

GetTitleEnglishOk returns a tuple with the TitleEnglish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEnglish

`func (o *UAEPassProviderOutput) SetTitleEnglish(v string)`

SetTitleEnglish sets TitleEnglish field to given value.

### HasTitleEnglish

`func (o *UAEPassProviderOutput) HasTitleEnglish() bool`

HasTitleEnglish returns a boolean if a field has been set.

### SetTitleEnglishNil

`func (o *UAEPassProviderOutput) SetTitleEnglishNil(b bool)`

 SetTitleEnglishNil sets the value for TitleEnglish to be an explicit nil

### UnsetTitleEnglish
`func (o *UAEPassProviderOutput) UnsetTitleEnglish()`

UnsetTitleEnglish ensures that no value is present for TitleEnglish, not even an explicit nil
### GetTitleArabic

`func (o *UAEPassProviderOutput) GetTitleArabic() string`

GetTitleArabic returns the TitleArabic field if non-nil, zero value otherwise.

### GetTitleArabicOk

`func (o *UAEPassProviderOutput) GetTitleArabicOk() (*string, bool)`

GetTitleArabicOk returns a tuple with the TitleArabic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleArabic

`func (o *UAEPassProviderOutput) SetTitleArabic(v string)`

SetTitleArabic sets TitleArabic field to given value.

### HasTitleArabic

`func (o *UAEPassProviderOutput) HasTitleArabic() bool`

HasTitleArabic returns a boolean if a field has been set.

### SetTitleArabicNil

`func (o *UAEPassProviderOutput) SetTitleArabicNil(b bool)`

 SetTitleArabicNil sets the value for TitleArabic to be an explicit nil

### UnsetTitleArabic
`func (o *UAEPassProviderOutput) UnsetTitleArabic()`

UnsetTitleArabic ensures that no value is present for TitleArabic, not even an explicit nil
### GetProfileType

`func (o *UAEPassProviderOutput) GetProfileType() int32`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *UAEPassProviderOutput) GetProfileTypeOk() (*int32, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *UAEPassProviderOutput) SetProfileType(v int32)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *UAEPassProviderOutput) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### SetProfileTypeNil

`func (o *UAEPassProviderOutput) SetProfileTypeNil(b bool)`

 SetProfileTypeNil sets the value for ProfileType to be an explicit nil

### UnsetProfileType
`func (o *UAEPassProviderOutput) UnsetProfileType()`

UnsetProfileType ensures that no value is present for ProfileType, not even an explicit nil
### GetUnifiedId

`func (o *UAEPassProviderOutput) GetUnifiedId() string`

GetUnifiedId returns the UnifiedId field if non-nil, zero value otherwise.

### GetUnifiedIdOk

`func (o *UAEPassProviderOutput) GetUnifiedIdOk() (*string, bool)`

GetUnifiedIdOk returns a tuple with the UnifiedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedId

`func (o *UAEPassProviderOutput) SetUnifiedId(v string)`

SetUnifiedId sets UnifiedId field to given value.

### HasUnifiedId

`func (o *UAEPassProviderOutput) HasUnifiedId() bool`

HasUnifiedId returns a boolean if a field has been set.

### SetUnifiedIdNil

`func (o *UAEPassProviderOutput) SetUnifiedIdNil(b bool)`

 SetUnifiedIdNil sets the value for UnifiedId to be an explicit nil

### UnsetUnifiedId
`func (o *UAEPassProviderOutput) UnsetUnifiedId()`

UnsetUnifiedId ensures that no value is present for UnifiedId, not even an explicit nil
### GetAuthenticationAssuranceLevel

`func (o *UAEPassProviderOutput) GetAuthenticationAssuranceLevel() string`

GetAuthenticationAssuranceLevel returns the AuthenticationAssuranceLevel field if non-nil, zero value otherwise.

### GetAuthenticationAssuranceLevelOk

`func (o *UAEPassProviderOutput) GetAuthenticationAssuranceLevelOk() (*string, bool)`

GetAuthenticationAssuranceLevelOk returns a tuple with the AuthenticationAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAssuranceLevel

`func (o *UAEPassProviderOutput) SetAuthenticationAssuranceLevel(v string)`

SetAuthenticationAssuranceLevel sets AuthenticationAssuranceLevel field to given value.

### HasAuthenticationAssuranceLevel

`func (o *UAEPassProviderOutput) HasAuthenticationAssuranceLevel() bool`

HasAuthenticationAssuranceLevel returns a boolean if a field has been set.

### SetAuthenticationAssuranceLevelNil

`func (o *UAEPassProviderOutput) SetAuthenticationAssuranceLevelNil(b bool)`

 SetAuthenticationAssuranceLevelNil sets the value for AuthenticationAssuranceLevel to be an explicit nil

### UnsetAuthenticationAssuranceLevel
`func (o *UAEPassProviderOutput) UnsetAuthenticationAssuranceLevel()`

UnsetAuthenticationAssuranceLevel ensures that no value is present for AuthenticationAssuranceLevel, not even an explicit nil
### GetAuthenticationMethods

`func (o *UAEPassProviderOutput) GetAuthenticationMethods() []string`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *UAEPassProviderOutput) GetAuthenticationMethodsOk() (*[]string, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *UAEPassProviderOutput) SetAuthenticationMethods(v []string)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *UAEPassProviderOutput) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### SetAuthenticationMethodsNil

`func (o *UAEPassProviderOutput) SetAuthenticationMethodsNil(b bool)`

 SetAuthenticationMethodsNil sets the value for AuthenticationMethods to be an explicit nil

### UnsetAuthenticationMethods
`func (o *UAEPassProviderOutput) UnsetAuthenticationMethods()`

UnsetAuthenticationMethods ensures that no value is present for AuthenticationMethods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



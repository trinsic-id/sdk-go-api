# CzechMojeIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectIdentifier** | Pointer to **NullableString** | The subject identifier (sub) of the verified individual&#39;s MojeID account.              This is a unique identifier that represents the user within the MojeID system. | [optional] 
**Name** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given (first) name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family (last) name. | [optional] 
**Nickname** | Pointer to **NullableString** | The individual&#39;s nickname. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the individual&#39;s email address has been verified by MojeID. (Verification email and link clicked) | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The individual&#39;s phone number. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | An individual&#39;s phone number has been verified by MojeID. (SMS verification) | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth.              Formatted as an ISO 8601 Date. | [optional] 
**Gender** | Pointer to **NullableString** | The individual&#39;s gender as reported by MojeID. | [optional] 
**IdCardNumber** | Pointer to **NullableString** | The individual&#39;s Czech ID card number. | [optional] 
**PassportNumber** | Pointer to **NullableString** | The individual&#39;s Czech passport number. | [optional] 
**SocialSecurityNumber** | Pointer to **NullableString** | The individual&#39;s Ministry of Labour and Social Affairs identifier (Czech social security equivalent). | [optional] 
**IsAdult** | Pointer to **NullableBool** | The individual is an adult (18 years or older). | [optional] 
**AccountValidated** | Pointer to **NullableBool** | The MojeID account has been validated.              A validated account indicates the individual&#39;s identity has been verified to a higher level of assurance within the MojeID system. | [optional] 
**CompanyRegistrationNumber** | Pointer to **NullableString** | The individual&#39;s or organization&#39;s Czech Registration ID (ICO), the Czech company registration number. | [optional] 
**TaxIdentificationNumber** | Pointer to **NullableString** | The individual&#39;s or organization&#39;s Danove Identifikacni Cislo (DIC), the Czech tax identification number. | [optional] 
**NiaVerified** | Pointer to **NullableBool** | The individual&#39;s identity has been verified through the Czech National Identity Authority (NIA).              NIA verification indicates a higher level of identity assurance, typically involving government-issued credentials verified through the Czech eGovernment infrastructure. | [optional] 
**TransactionId** | Pointer to **NullableString** | Unique login transaction identifier | [optional] 
**HomeAddress** | Pointer to [**NullableCzechMojeIdAddressOutput**](CzechMojeIdAddressOutput.md) | The individual&#39;s home (default) address, from the &#x60;mojeid_address_def&#x60; claim. | [optional] 
**BillingAddress** | Pointer to [**NullableCzechMojeIdAddressOutput**](CzechMojeIdAddressOutput.md) | The individual&#39;s billing address, from the &#x60;mojeid_address_bill&#x60; claim. | [optional] 
**ShippingAddress** | Pointer to [**NullableCzechMojeIdAddressOutput**](CzechMojeIdAddressOutput.md) | The individual&#39;s shipping address, from the &#x60;mojeid_address_ship&#x60; claim. | [optional] 

## Methods

### NewCzechMojeIdProviderOutput

`func NewCzechMojeIdProviderOutput() *CzechMojeIdProviderOutput`

NewCzechMojeIdProviderOutput instantiates a new CzechMojeIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechMojeIdProviderOutputWithDefaults

`func NewCzechMojeIdProviderOutputWithDefaults() *CzechMojeIdProviderOutput`

NewCzechMojeIdProviderOutputWithDefaults instantiates a new CzechMojeIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectIdentifier

`func (o *CzechMojeIdProviderOutput) GetSubjectIdentifier() string`

GetSubjectIdentifier returns the SubjectIdentifier field if non-nil, zero value otherwise.

### GetSubjectIdentifierOk

`func (o *CzechMojeIdProviderOutput) GetSubjectIdentifierOk() (*string, bool)`

GetSubjectIdentifierOk returns a tuple with the SubjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIdentifier

`func (o *CzechMojeIdProviderOutput) SetSubjectIdentifier(v string)`

SetSubjectIdentifier sets SubjectIdentifier field to given value.

### HasSubjectIdentifier

`func (o *CzechMojeIdProviderOutput) HasSubjectIdentifier() bool`

HasSubjectIdentifier returns a boolean if a field has been set.

### SetSubjectIdentifierNil

`func (o *CzechMojeIdProviderOutput) SetSubjectIdentifierNil(b bool)`

 SetSubjectIdentifierNil sets the value for SubjectIdentifier to be an explicit nil

### UnsetSubjectIdentifier
`func (o *CzechMojeIdProviderOutput) UnsetSubjectIdentifier()`

UnsetSubjectIdentifier ensures that no value is present for SubjectIdentifier, not even an explicit nil
### GetName

`func (o *CzechMojeIdProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CzechMojeIdProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CzechMojeIdProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CzechMojeIdProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CzechMojeIdProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CzechMojeIdProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGivenName

`func (o *CzechMojeIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *CzechMojeIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *CzechMojeIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *CzechMojeIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *CzechMojeIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *CzechMojeIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *CzechMojeIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *CzechMojeIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *CzechMojeIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *CzechMojeIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *CzechMojeIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *CzechMojeIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetNickname

`func (o *CzechMojeIdProviderOutput) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CzechMojeIdProviderOutput) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CzechMojeIdProviderOutput) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CzechMojeIdProviderOutput) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *CzechMojeIdProviderOutput) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *CzechMojeIdProviderOutput) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetEmail

`func (o *CzechMojeIdProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CzechMojeIdProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CzechMojeIdProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CzechMojeIdProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CzechMojeIdProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CzechMojeIdProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *CzechMojeIdProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CzechMojeIdProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CzechMojeIdProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CzechMojeIdProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *CzechMojeIdProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *CzechMojeIdProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhoneNumber

`func (o *CzechMojeIdProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CzechMojeIdProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CzechMojeIdProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CzechMojeIdProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *CzechMojeIdProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *CzechMojeIdProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *CzechMojeIdProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *CzechMojeIdProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *CzechMojeIdProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *CzechMojeIdProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *CzechMojeIdProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *CzechMojeIdProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetDateOfBirth

`func (o *CzechMojeIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CzechMojeIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CzechMojeIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CzechMojeIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *CzechMojeIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *CzechMojeIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGender

`func (o *CzechMojeIdProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CzechMojeIdProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CzechMojeIdProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CzechMojeIdProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CzechMojeIdProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CzechMojeIdProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetIdCardNumber

`func (o *CzechMojeIdProviderOutput) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *CzechMojeIdProviderOutput) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *CzechMojeIdProviderOutput) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.

### HasIdCardNumber

`func (o *CzechMojeIdProviderOutput) HasIdCardNumber() bool`

HasIdCardNumber returns a boolean if a field has been set.

### SetIdCardNumberNil

`func (o *CzechMojeIdProviderOutput) SetIdCardNumberNil(b bool)`

 SetIdCardNumberNil sets the value for IdCardNumber to be an explicit nil

### UnsetIdCardNumber
`func (o *CzechMojeIdProviderOutput) UnsetIdCardNumber()`

UnsetIdCardNumber ensures that no value is present for IdCardNumber, not even an explicit nil
### GetPassportNumber

`func (o *CzechMojeIdProviderOutput) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *CzechMojeIdProviderOutput) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *CzechMojeIdProviderOutput) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *CzechMojeIdProviderOutput) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### SetPassportNumberNil

`func (o *CzechMojeIdProviderOutput) SetPassportNumberNil(b bool)`

 SetPassportNumberNil sets the value for PassportNumber to be an explicit nil

### UnsetPassportNumber
`func (o *CzechMojeIdProviderOutput) UnsetPassportNumber()`

UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
### GetSocialSecurityNumber

`func (o *CzechMojeIdProviderOutput) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *CzechMojeIdProviderOutput) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *CzechMojeIdProviderOutput) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *CzechMojeIdProviderOutput) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### SetSocialSecurityNumberNil

`func (o *CzechMojeIdProviderOutput) SetSocialSecurityNumberNil(b bool)`

 SetSocialSecurityNumberNil sets the value for SocialSecurityNumber to be an explicit nil

### UnsetSocialSecurityNumber
`func (o *CzechMojeIdProviderOutput) UnsetSocialSecurityNumber()`

UnsetSocialSecurityNumber ensures that no value is present for SocialSecurityNumber, not even an explicit nil
### GetIsAdult

`func (o *CzechMojeIdProviderOutput) GetIsAdult() bool`

GetIsAdult returns the IsAdult field if non-nil, zero value otherwise.

### GetIsAdultOk

`func (o *CzechMojeIdProviderOutput) GetIsAdultOk() (*bool, bool)`

GetIsAdultOk returns a tuple with the IsAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdult

`func (o *CzechMojeIdProviderOutput) SetIsAdult(v bool)`

SetIsAdult sets IsAdult field to given value.

### HasIsAdult

`func (o *CzechMojeIdProviderOutput) HasIsAdult() bool`

HasIsAdult returns a boolean if a field has been set.

### SetIsAdultNil

`func (o *CzechMojeIdProviderOutput) SetIsAdultNil(b bool)`

 SetIsAdultNil sets the value for IsAdult to be an explicit nil

### UnsetIsAdult
`func (o *CzechMojeIdProviderOutput) UnsetIsAdult()`

UnsetIsAdult ensures that no value is present for IsAdult, not even an explicit nil
### GetAccountValidated

`func (o *CzechMojeIdProviderOutput) GetAccountValidated() bool`

GetAccountValidated returns the AccountValidated field if non-nil, zero value otherwise.

### GetAccountValidatedOk

`func (o *CzechMojeIdProviderOutput) GetAccountValidatedOk() (*bool, bool)`

GetAccountValidatedOk returns a tuple with the AccountValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountValidated

`func (o *CzechMojeIdProviderOutput) SetAccountValidated(v bool)`

SetAccountValidated sets AccountValidated field to given value.

### HasAccountValidated

`func (o *CzechMojeIdProviderOutput) HasAccountValidated() bool`

HasAccountValidated returns a boolean if a field has been set.

### SetAccountValidatedNil

`func (o *CzechMojeIdProviderOutput) SetAccountValidatedNil(b bool)`

 SetAccountValidatedNil sets the value for AccountValidated to be an explicit nil

### UnsetAccountValidated
`func (o *CzechMojeIdProviderOutput) UnsetAccountValidated()`

UnsetAccountValidated ensures that no value is present for AccountValidated, not even an explicit nil
### GetCompanyRegistrationNumber

`func (o *CzechMojeIdProviderOutput) GetCompanyRegistrationNumber() string`

GetCompanyRegistrationNumber returns the CompanyRegistrationNumber field if non-nil, zero value otherwise.

### GetCompanyRegistrationNumberOk

`func (o *CzechMojeIdProviderOutput) GetCompanyRegistrationNumberOk() (*string, bool)`

GetCompanyRegistrationNumberOk returns a tuple with the CompanyRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRegistrationNumber

`func (o *CzechMojeIdProviderOutput) SetCompanyRegistrationNumber(v string)`

SetCompanyRegistrationNumber sets CompanyRegistrationNumber field to given value.

### HasCompanyRegistrationNumber

`func (o *CzechMojeIdProviderOutput) HasCompanyRegistrationNumber() bool`

HasCompanyRegistrationNumber returns a boolean if a field has been set.

### SetCompanyRegistrationNumberNil

`func (o *CzechMojeIdProviderOutput) SetCompanyRegistrationNumberNil(b bool)`

 SetCompanyRegistrationNumberNil sets the value for CompanyRegistrationNumber to be an explicit nil

### UnsetCompanyRegistrationNumber
`func (o *CzechMojeIdProviderOutput) UnsetCompanyRegistrationNumber()`

UnsetCompanyRegistrationNumber ensures that no value is present for CompanyRegistrationNumber, not even an explicit nil
### GetTaxIdentificationNumber

`func (o *CzechMojeIdProviderOutput) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *CzechMojeIdProviderOutput) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *CzechMojeIdProviderOutput) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *CzechMojeIdProviderOutput) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### SetTaxIdentificationNumberNil

`func (o *CzechMojeIdProviderOutput) SetTaxIdentificationNumberNil(b bool)`

 SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil

### UnsetTaxIdentificationNumber
`func (o *CzechMojeIdProviderOutput) UnsetTaxIdentificationNumber()`

UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
### GetNiaVerified

`func (o *CzechMojeIdProviderOutput) GetNiaVerified() bool`

GetNiaVerified returns the NiaVerified field if non-nil, zero value otherwise.

### GetNiaVerifiedOk

`func (o *CzechMojeIdProviderOutput) GetNiaVerifiedOk() (*bool, bool)`

GetNiaVerifiedOk returns a tuple with the NiaVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiaVerified

`func (o *CzechMojeIdProviderOutput) SetNiaVerified(v bool)`

SetNiaVerified sets NiaVerified field to given value.

### HasNiaVerified

`func (o *CzechMojeIdProviderOutput) HasNiaVerified() bool`

HasNiaVerified returns a boolean if a field has been set.

### SetNiaVerifiedNil

`func (o *CzechMojeIdProviderOutput) SetNiaVerifiedNil(b bool)`

 SetNiaVerifiedNil sets the value for NiaVerified to be an explicit nil

### UnsetNiaVerified
`func (o *CzechMojeIdProviderOutput) UnsetNiaVerified()`

UnsetNiaVerified ensures that no value is present for NiaVerified, not even an explicit nil
### GetTransactionId

`func (o *CzechMojeIdProviderOutput) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CzechMojeIdProviderOutput) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CzechMojeIdProviderOutput) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CzechMojeIdProviderOutput) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *CzechMojeIdProviderOutput) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CzechMojeIdProviderOutput) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetHomeAddress

`func (o *CzechMojeIdProviderOutput) GetHomeAddress() CzechMojeIdAddressOutput`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *CzechMojeIdProviderOutput) GetHomeAddressOk() (*CzechMojeIdAddressOutput, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeAddress

`func (o *CzechMojeIdProviderOutput) SetHomeAddress(v CzechMojeIdAddressOutput)`

SetHomeAddress sets HomeAddress field to given value.

### HasHomeAddress

`func (o *CzechMojeIdProviderOutput) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddressNil

`func (o *CzechMojeIdProviderOutput) SetHomeAddressNil(b bool)`

 SetHomeAddressNil sets the value for HomeAddress to be an explicit nil

### UnsetHomeAddress
`func (o *CzechMojeIdProviderOutput) UnsetHomeAddress()`

UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
### GetBillingAddress

`func (o *CzechMojeIdProviderOutput) GetBillingAddress() CzechMojeIdAddressOutput`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CzechMojeIdProviderOutput) GetBillingAddressOk() (*CzechMojeIdAddressOutput, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CzechMojeIdProviderOutput) SetBillingAddress(v CzechMojeIdAddressOutput)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CzechMojeIdProviderOutput) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *CzechMojeIdProviderOutput) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *CzechMojeIdProviderOutput) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetShippingAddress

`func (o *CzechMojeIdProviderOutput) GetShippingAddress() CzechMojeIdAddressOutput`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CzechMojeIdProviderOutput) GetShippingAddressOk() (*CzechMojeIdAddressOutput, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CzechMojeIdProviderOutput) SetShippingAddress(v CzechMojeIdAddressOutput)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CzechMojeIdProviderOutput) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### SetShippingAddressNil

`func (o *CzechMojeIdProviderOutput) SetShippingAddressNil(b bool)`

 SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil

### UnsetShippingAddress
`func (o *CzechMojeIdProviderOutput) UnsetShippingAddress()`

UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



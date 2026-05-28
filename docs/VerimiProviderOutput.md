# VerimiProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the individual. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the email address has been verified by Verimi. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the individual in E.164 format. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether the phone number has been verified by Verimi. | [optional] 
**Gender** | Pointer to **NullableString** | The gender of the individual. | [optional] 
**ZoneInformation** | Pointer to **NullableString** | The individual&#39;s time zone as an IANA Time Zone Database zoneinfo identifier. | [optional] 
**Locale** | Pointer to **NullableString** | The selected language or locale of the individual as a BCP 47 language tag. | [optional] 
**Citizenship** | Pointer to **NullableString** | The citizenship of the individual.              This is an ISO-3166-1 alpha-2 country code. | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** | The place of birth of the individual. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number of the identity document used by the individual. | [optional] 
**DocumentType** | Pointer to **NullableString** | The identity document type.              Known values: - I: ID card. - P: Passport. | [optional] 
**DocumentExpirationDate** | Pointer to **NullableString** | The expiration date of the identity document. | [optional] 
**DocumentIssueDate** | Pointer to **NullableString** | The issue date of the identity document. | [optional] 
**DocumentIssuingAuthority** | Pointer to **NullableString** | The authority that issued the identity document. | [optional] 
**VerificationMethod** | Pointer to **NullableString** | The verification method used to prove the individual&#39;s identity. | [optional] 
**VerificationDate** | Pointer to **NullableString** | The date when the individual proved their identity with Verimi. | [optional] 
**LevelOfAssurance** | Pointer to **NullableString** | The level of assurance for the verification.              Known values: - Low: The individual has self-asserted their identity. - Substantial: The individual has completed identity proofing and strong authentication. - High: The individual has completed identity proofing with stronger cryptographic authentication requirements. | [optional] 
**AuthenticationMethod** | Pointer to **NullableString** | The authentication method for the completed identification.              Known values: - email - loa.dipp.default - loa.dipp.2fa - idcard | [optional] 
**Pseudonym** | Pointer to **NullableString** | The pseudonymous identifier.              For German eID, this is the restricted identifier (rID), scoped to the relying-party sector. | [optional] 
**Address** | Pointer to [**NullableVerimiAddressOutput**](VerimiAddressOutput.md) | The individual&#39;s structured address. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code for the country that issued the identity document. | [optional] 

## Methods

### NewVerimiProviderOutput

`func NewVerimiProviderOutput() *VerimiProviderOutput`

NewVerimiProviderOutput instantiates a new VerimiProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerimiProviderOutputWithDefaults

`func NewVerimiProviderOutputWithDefaults() *VerimiProviderOutput`

NewVerimiProviderOutputWithDefaults instantiates a new VerimiProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *VerimiProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VerimiProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VerimiProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *VerimiProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *VerimiProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *VerimiProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *VerimiProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *VerimiProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *VerimiProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *VerimiProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *VerimiProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *VerimiProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetGivenName

`func (o *VerimiProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *VerimiProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *VerimiProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *VerimiProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *VerimiProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *VerimiProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *VerimiProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *VerimiProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *VerimiProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *VerimiProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *VerimiProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *VerimiProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *VerimiProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *VerimiProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *VerimiProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *VerimiProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *VerimiProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *VerimiProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetEmail

`func (o *VerimiProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VerimiProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VerimiProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VerimiProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *VerimiProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VerimiProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *VerimiProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *VerimiProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *VerimiProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *VerimiProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *VerimiProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *VerimiProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhoneNumber

`func (o *VerimiProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *VerimiProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *VerimiProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *VerimiProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *VerimiProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *VerimiProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *VerimiProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *VerimiProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *VerimiProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *VerimiProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *VerimiProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *VerimiProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetGender

`func (o *VerimiProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *VerimiProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *VerimiProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *VerimiProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *VerimiProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *VerimiProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetZoneInformation

`func (o *VerimiProviderOutput) GetZoneInformation() string`

GetZoneInformation returns the ZoneInformation field if non-nil, zero value otherwise.

### GetZoneInformationOk

`func (o *VerimiProviderOutput) GetZoneInformationOk() (*string, bool)`

GetZoneInformationOk returns a tuple with the ZoneInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInformation

`func (o *VerimiProviderOutput) SetZoneInformation(v string)`

SetZoneInformation sets ZoneInformation field to given value.

### HasZoneInformation

`func (o *VerimiProviderOutput) HasZoneInformation() bool`

HasZoneInformation returns a boolean if a field has been set.

### SetZoneInformationNil

`func (o *VerimiProviderOutput) SetZoneInformationNil(b bool)`

 SetZoneInformationNil sets the value for ZoneInformation to be an explicit nil

### UnsetZoneInformation
`func (o *VerimiProviderOutput) UnsetZoneInformation()`

UnsetZoneInformation ensures that no value is present for ZoneInformation, not even an explicit nil
### GetLocale

`func (o *VerimiProviderOutput) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *VerimiProviderOutput) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *VerimiProviderOutput) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *VerimiProviderOutput) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *VerimiProviderOutput) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *VerimiProviderOutput) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetCitizenship

`func (o *VerimiProviderOutput) GetCitizenship() string`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *VerimiProviderOutput) GetCitizenshipOk() (*string, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *VerimiProviderOutput) SetCitizenship(v string)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *VerimiProviderOutput) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### SetCitizenshipNil

`func (o *VerimiProviderOutput) SetCitizenshipNil(b bool)`

 SetCitizenshipNil sets the value for Citizenship to be an explicit nil

### UnsetCitizenship
`func (o *VerimiProviderOutput) UnsetCitizenship()`

UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
### GetPlaceOfBirth

`func (o *VerimiProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *VerimiProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *VerimiProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *VerimiProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *VerimiProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *VerimiProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *VerimiProviderOutput) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *VerimiProviderOutput) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *VerimiProviderOutput) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *VerimiProviderOutput) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *VerimiProviderOutput) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *VerimiProviderOutput) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDocumentType

`func (o *VerimiProviderOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *VerimiProviderOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *VerimiProviderOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *VerimiProviderOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *VerimiProviderOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *VerimiProviderOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetDocumentExpirationDate

`func (o *VerimiProviderOutput) GetDocumentExpirationDate() string`

GetDocumentExpirationDate returns the DocumentExpirationDate field if non-nil, zero value otherwise.

### GetDocumentExpirationDateOk

`func (o *VerimiProviderOutput) GetDocumentExpirationDateOk() (*string, bool)`

GetDocumentExpirationDateOk returns a tuple with the DocumentExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentExpirationDate

`func (o *VerimiProviderOutput) SetDocumentExpirationDate(v string)`

SetDocumentExpirationDate sets DocumentExpirationDate field to given value.

### HasDocumentExpirationDate

`func (o *VerimiProviderOutput) HasDocumentExpirationDate() bool`

HasDocumentExpirationDate returns a boolean if a field has been set.

### SetDocumentExpirationDateNil

`func (o *VerimiProviderOutput) SetDocumentExpirationDateNil(b bool)`

 SetDocumentExpirationDateNil sets the value for DocumentExpirationDate to be an explicit nil

### UnsetDocumentExpirationDate
`func (o *VerimiProviderOutput) UnsetDocumentExpirationDate()`

UnsetDocumentExpirationDate ensures that no value is present for DocumentExpirationDate, not even an explicit nil
### GetDocumentIssueDate

`func (o *VerimiProviderOutput) GetDocumentIssueDate() string`

GetDocumentIssueDate returns the DocumentIssueDate field if non-nil, zero value otherwise.

### GetDocumentIssueDateOk

`func (o *VerimiProviderOutput) GetDocumentIssueDateOk() (*string, bool)`

GetDocumentIssueDateOk returns a tuple with the DocumentIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssueDate

`func (o *VerimiProviderOutput) SetDocumentIssueDate(v string)`

SetDocumentIssueDate sets DocumentIssueDate field to given value.

### HasDocumentIssueDate

`func (o *VerimiProviderOutput) HasDocumentIssueDate() bool`

HasDocumentIssueDate returns a boolean if a field has been set.

### SetDocumentIssueDateNil

`func (o *VerimiProviderOutput) SetDocumentIssueDateNil(b bool)`

 SetDocumentIssueDateNil sets the value for DocumentIssueDate to be an explicit nil

### UnsetDocumentIssueDate
`func (o *VerimiProviderOutput) UnsetDocumentIssueDate()`

UnsetDocumentIssueDate ensures that no value is present for DocumentIssueDate, not even an explicit nil
### GetDocumentIssuingAuthority

`func (o *VerimiProviderOutput) GetDocumentIssuingAuthority() string`

GetDocumentIssuingAuthority returns the DocumentIssuingAuthority field if non-nil, zero value otherwise.

### GetDocumentIssuingAuthorityOk

`func (o *VerimiProviderOutput) GetDocumentIssuingAuthorityOk() (*string, bool)`

GetDocumentIssuingAuthorityOk returns a tuple with the DocumentIssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssuingAuthority

`func (o *VerimiProviderOutput) SetDocumentIssuingAuthority(v string)`

SetDocumentIssuingAuthority sets DocumentIssuingAuthority field to given value.

### HasDocumentIssuingAuthority

`func (o *VerimiProviderOutput) HasDocumentIssuingAuthority() bool`

HasDocumentIssuingAuthority returns a boolean if a field has been set.

### SetDocumentIssuingAuthorityNil

`func (o *VerimiProviderOutput) SetDocumentIssuingAuthorityNil(b bool)`

 SetDocumentIssuingAuthorityNil sets the value for DocumentIssuingAuthority to be an explicit nil

### UnsetDocumentIssuingAuthority
`func (o *VerimiProviderOutput) UnsetDocumentIssuingAuthority()`

UnsetDocumentIssuingAuthority ensures that no value is present for DocumentIssuingAuthority, not even an explicit nil
### GetVerificationMethod

`func (o *VerimiProviderOutput) GetVerificationMethod() string`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *VerimiProviderOutput) GetVerificationMethodOk() (*string, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *VerimiProviderOutput) SetVerificationMethod(v string)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *VerimiProviderOutput) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.

### SetVerificationMethodNil

`func (o *VerimiProviderOutput) SetVerificationMethodNil(b bool)`

 SetVerificationMethodNil sets the value for VerificationMethod to be an explicit nil

### UnsetVerificationMethod
`func (o *VerimiProviderOutput) UnsetVerificationMethod()`

UnsetVerificationMethod ensures that no value is present for VerificationMethod, not even an explicit nil
### GetVerificationDate

`func (o *VerimiProviderOutput) GetVerificationDate() string`

GetVerificationDate returns the VerificationDate field if non-nil, zero value otherwise.

### GetVerificationDateOk

`func (o *VerimiProviderOutput) GetVerificationDateOk() (*string, bool)`

GetVerificationDateOk returns a tuple with the VerificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDate

`func (o *VerimiProviderOutput) SetVerificationDate(v string)`

SetVerificationDate sets VerificationDate field to given value.

### HasVerificationDate

`func (o *VerimiProviderOutput) HasVerificationDate() bool`

HasVerificationDate returns a boolean if a field has been set.

### SetVerificationDateNil

`func (o *VerimiProviderOutput) SetVerificationDateNil(b bool)`

 SetVerificationDateNil sets the value for VerificationDate to be an explicit nil

### UnsetVerificationDate
`func (o *VerimiProviderOutput) UnsetVerificationDate()`

UnsetVerificationDate ensures that no value is present for VerificationDate, not even an explicit nil
### GetLevelOfAssurance

`func (o *VerimiProviderOutput) GetLevelOfAssurance() string`

GetLevelOfAssurance returns the LevelOfAssurance field if non-nil, zero value otherwise.

### GetLevelOfAssuranceOk

`func (o *VerimiProviderOutput) GetLevelOfAssuranceOk() (*string, bool)`

GetLevelOfAssuranceOk returns a tuple with the LevelOfAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOfAssurance

`func (o *VerimiProviderOutput) SetLevelOfAssurance(v string)`

SetLevelOfAssurance sets LevelOfAssurance field to given value.

### HasLevelOfAssurance

`func (o *VerimiProviderOutput) HasLevelOfAssurance() bool`

HasLevelOfAssurance returns a boolean if a field has been set.

### SetLevelOfAssuranceNil

`func (o *VerimiProviderOutput) SetLevelOfAssuranceNil(b bool)`

 SetLevelOfAssuranceNil sets the value for LevelOfAssurance to be an explicit nil

### UnsetLevelOfAssurance
`func (o *VerimiProviderOutput) UnsetLevelOfAssurance()`

UnsetLevelOfAssurance ensures that no value is present for LevelOfAssurance, not even an explicit nil
### GetAuthenticationMethod

`func (o *VerimiProviderOutput) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VerimiProviderOutput) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VerimiProviderOutput) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *VerimiProviderOutput) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### SetAuthenticationMethodNil

`func (o *VerimiProviderOutput) SetAuthenticationMethodNil(b bool)`

 SetAuthenticationMethodNil sets the value for AuthenticationMethod to be an explicit nil

### UnsetAuthenticationMethod
`func (o *VerimiProviderOutput) UnsetAuthenticationMethod()`

UnsetAuthenticationMethod ensures that no value is present for AuthenticationMethod, not even an explicit nil
### GetPseudonym

`func (o *VerimiProviderOutput) GetPseudonym() string`

GetPseudonym returns the Pseudonym field if non-nil, zero value otherwise.

### GetPseudonymOk

`func (o *VerimiProviderOutput) GetPseudonymOk() (*string, bool)`

GetPseudonymOk returns a tuple with the Pseudonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonym

`func (o *VerimiProviderOutput) SetPseudonym(v string)`

SetPseudonym sets Pseudonym field to given value.

### HasPseudonym

`func (o *VerimiProviderOutput) HasPseudonym() bool`

HasPseudonym returns a boolean if a field has been set.

### SetPseudonymNil

`func (o *VerimiProviderOutput) SetPseudonymNil(b bool)`

 SetPseudonymNil sets the value for Pseudonym to be an explicit nil

### UnsetPseudonym
`func (o *VerimiProviderOutput) UnsetPseudonym()`

UnsetPseudonym ensures that no value is present for Pseudonym, not even an explicit nil
### GetAddress

`func (o *VerimiProviderOutput) GetAddress() VerimiAddressOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VerimiProviderOutput) GetAddressOk() (*VerimiAddressOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VerimiProviderOutput) SetAddress(v VerimiAddressOutput)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VerimiProviderOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VerimiProviderOutput) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VerimiProviderOutput) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetIssuingCountry

`func (o *VerimiProviderOutput) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *VerimiProviderOutput) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *VerimiProviderOutput) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *VerimiProviderOutput) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *VerimiProviderOutput) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *VerimiProviderOutput) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



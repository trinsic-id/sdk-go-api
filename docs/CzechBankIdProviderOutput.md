# CzechBankIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectIdentifier** | **string** | The subject identifier for the verified individual. | 
**FullName** | Pointer to **NullableString** | The individual&#39;s full name. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given or first name. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family or last name. | [optional] 
**MiddleName** | Pointer to **NullableString** | The individual&#39;s middle name. | [optional] 
**Nickname** | Pointer to **NullableString** | Nickname used by a physical person              A casual name, not necessarily the legal given name. | [optional] 
**PreferredUsername** | Pointer to **NullableString** | The individual&#39;s preferred username.              The user has chosen this as the preferred name for logins. | [optional] 
**Gender** | Pointer to **NullableString** | The individual&#39;s gender.              Possible values: - Male - Female - Other              Only returned for Identify verifications. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth.              This is a string because the value may have reduced precision, such as only a year. | [optional] 
**BirthNumber** | Pointer to **NullableString** | The individual&#39;s Czech birth number.              This value is written without the traditional forward slash. Czech birth numbers are Czech public-authority identifiers administered through the Ministry of Interior&#39;s birth-number register. The Ministry can assign one to Czech citizens and to foreign nationals with temporary or permanent residence in the Czech Republic. For foreign nationals, this is still a Czech birth number in the Czech format, not an identity number from another country.              The common current format is YYXXDDSSSC: - YY: Two digit birth year - XX: Encoded birth month, see below for encoding logic - DD: Birth day - SSS: Sequence number - C: Modulo-11 checksum digit. Not present for people born before 1954. The encoded month is MM for men, MM + 50 for women, and can be MM + 20 for men or MM + 70 for women when additional number ranges are needed. Older numbers for people born before 1954 can use YYXXDDSSS and do not use the modulo-11 checksum.              Only returned for Identify verifications. | [optional] 
**Age** | Pointer to **NullableInt32** | The individual&#39;s age. | [optional] 
**IsAdult** | Pointer to **NullableBool** | Whether the individual is an adult.              In the Czech Republic, legal majority (adulthood) is reached at 18 years old. Only returned for Identify Plus verifications. | [optional] 
**DateOfDeath** | Pointer to **NullableString** | The individual&#39;s date of death, if available. | [optional] 
**BirthPlace** | Pointer to **NullableString** | The individual&#39;s place of birth.              Only returned for Identify Plus verifications. | [optional] 
**BirthCountry** | Pointer to **NullableString** | The individual&#39;s country of birth.              Only returned for Identify Plus verifications. | [optional] 
**PrimaryNationality** | Pointer to **NullableString** | The individual&#39;s primary nationality.              The value is an ISO 3166-1 alpha-2 country code. Only returned for Identify Plus verifications. | [optional] 
**Nationalities** | Pointer to **[]string** | The individual&#39;s nationalities. | [optional] 
**MaritalStatus** | Pointer to **NullableString** | The individual&#39;s marital status.              Possible values: - COHABITATION - MARRIED - DIVORCED - REGISTERED_PARTNERSHIP - REGISTERED_PARTNERSHIP_CANCELED - WIDOWED - SINGLE - PARTNERSHIP - PARTNERSHIP_CANCELED - SEPARATED - REGISTERED_PARTNERSHIP_WIDOWED - LAPSED_MARRIAGE - UNKNOWN              Only returned for Identify Plus verifications. | [optional] 
**Email** | Pointer to **NullableString** | The individual&#39;s email address. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the individual&#39;s email address has been verified. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The individual&#39;s phone number.              This value is normalized to E.164 format. | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether the individual&#39;s phone number has been verified. | [optional] 
**IsPoliticallyExposedPerson** | Pointer to **NullableBool** | Whether the individual is a politically exposed person.              This follows Czech anti-money-laundering rules, indicating a person that is at risk for corruption, bribery or money laundering activities. Mandatory screening of these individuals is required for compliance. Typically set for people such as the Prime Minister, ministers, etc. | [optional] 
**LimitedLegalCapacity** | Pointer to **NullableBool** | Whether the individual is a person with limited legal capacity. | [optional] 
**Verification** | Pointer to [**NullableCzechBankIdVerification**](CzechBankIdVerification.md) | Metadata about how the identity information was verified for Anti Money Laundering scope.              Only returned for Identify AML verifications. | [optional] 
**Addresses** | Pointer to [**[]CzechBankIdAddress**](CzechBankIdAddress.md) | The individual&#39;s addresses.              This can include permanent residence and contact addresses. Only returned for Identify verifications. | [optional] 
**IdCards** | Pointer to [**[]CzechBankIdCard**](CzechBankIdCard.md) | The individual&#39;s identity documents.              This can include document type, country, number, issuer, and validity dates. Only returned for Identify Plus verifications. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The UTC date and time when the identity data was last updated. | [optional] 
**ZoneInfo** | Pointer to **NullableString** | The individual&#39;s time zone.              This is represented as a zoneinfo/IANA Time Zone Database value. | [optional] 
**Locale** | Pointer to **NullableString** | The individual&#39;s locale.              This is represented as a BCP 47/RFC 5646 language tag. Some values may use an underscore separator, such as cs_CZ. | [optional] 
**Titles** | Pointer to [**[]CzechBankIdTitle**](CzechBankIdTitle.md) | The individual&#39;s title prefixes and suffixes.              Only returned for Identify verifications. | [optional] 

## Methods

### NewCzechBankIdProviderOutput

`func NewCzechBankIdProviderOutput(subjectIdentifier string, ) *CzechBankIdProviderOutput`

NewCzechBankIdProviderOutput instantiates a new CzechBankIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechBankIdProviderOutputWithDefaults

`func NewCzechBankIdProviderOutputWithDefaults() *CzechBankIdProviderOutput`

NewCzechBankIdProviderOutputWithDefaults instantiates a new CzechBankIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectIdentifier

`func (o *CzechBankIdProviderOutput) GetSubjectIdentifier() string`

GetSubjectIdentifier returns the SubjectIdentifier field if non-nil, zero value otherwise.

### GetSubjectIdentifierOk

`func (o *CzechBankIdProviderOutput) GetSubjectIdentifierOk() (*string, bool)`

GetSubjectIdentifierOk returns a tuple with the SubjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIdentifier

`func (o *CzechBankIdProviderOutput) SetSubjectIdentifier(v string)`

SetSubjectIdentifier sets SubjectIdentifier field to given value.


### GetFullName

`func (o *CzechBankIdProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CzechBankIdProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CzechBankIdProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CzechBankIdProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *CzechBankIdProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *CzechBankIdProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *CzechBankIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *CzechBankIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *CzechBankIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *CzechBankIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *CzechBankIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *CzechBankIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *CzechBankIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *CzechBankIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *CzechBankIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *CzechBankIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *CzechBankIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *CzechBankIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetMiddleName

`func (o *CzechBankIdProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CzechBankIdProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CzechBankIdProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CzechBankIdProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *CzechBankIdProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *CzechBankIdProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetNickname

`func (o *CzechBankIdProviderOutput) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CzechBankIdProviderOutput) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CzechBankIdProviderOutput) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CzechBankIdProviderOutput) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *CzechBankIdProviderOutput) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *CzechBankIdProviderOutput) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetPreferredUsername

`func (o *CzechBankIdProviderOutput) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *CzechBankIdProviderOutput) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *CzechBankIdProviderOutput) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *CzechBankIdProviderOutput) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### SetPreferredUsernameNil

`func (o *CzechBankIdProviderOutput) SetPreferredUsernameNil(b bool)`

 SetPreferredUsernameNil sets the value for PreferredUsername to be an explicit nil

### UnsetPreferredUsername
`func (o *CzechBankIdProviderOutput) UnsetPreferredUsername()`

UnsetPreferredUsername ensures that no value is present for PreferredUsername, not even an explicit nil
### GetGender

`func (o *CzechBankIdProviderOutput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CzechBankIdProviderOutput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CzechBankIdProviderOutput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CzechBankIdProviderOutput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CzechBankIdProviderOutput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CzechBankIdProviderOutput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDateOfBirth

`func (o *CzechBankIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CzechBankIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CzechBankIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CzechBankIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *CzechBankIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *CzechBankIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetBirthNumber

`func (o *CzechBankIdProviderOutput) GetBirthNumber() string`

GetBirthNumber returns the BirthNumber field if non-nil, zero value otherwise.

### GetBirthNumberOk

`func (o *CzechBankIdProviderOutput) GetBirthNumberOk() (*string, bool)`

GetBirthNumberOk returns a tuple with the BirthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthNumber

`func (o *CzechBankIdProviderOutput) SetBirthNumber(v string)`

SetBirthNumber sets BirthNumber field to given value.

### HasBirthNumber

`func (o *CzechBankIdProviderOutput) HasBirthNumber() bool`

HasBirthNumber returns a boolean if a field has been set.

### SetBirthNumberNil

`func (o *CzechBankIdProviderOutput) SetBirthNumberNil(b bool)`

 SetBirthNumberNil sets the value for BirthNumber to be an explicit nil

### UnsetBirthNumber
`func (o *CzechBankIdProviderOutput) UnsetBirthNumber()`

UnsetBirthNumber ensures that no value is present for BirthNumber, not even an explicit nil
### GetAge

`func (o *CzechBankIdProviderOutput) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *CzechBankIdProviderOutput) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *CzechBankIdProviderOutput) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *CzechBankIdProviderOutput) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *CzechBankIdProviderOutput) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *CzechBankIdProviderOutput) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetIsAdult

`func (o *CzechBankIdProviderOutput) GetIsAdult() bool`

GetIsAdult returns the IsAdult field if non-nil, zero value otherwise.

### GetIsAdultOk

`func (o *CzechBankIdProviderOutput) GetIsAdultOk() (*bool, bool)`

GetIsAdultOk returns a tuple with the IsAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdult

`func (o *CzechBankIdProviderOutput) SetIsAdult(v bool)`

SetIsAdult sets IsAdult field to given value.

### HasIsAdult

`func (o *CzechBankIdProviderOutput) HasIsAdult() bool`

HasIsAdult returns a boolean if a field has been set.

### SetIsAdultNil

`func (o *CzechBankIdProviderOutput) SetIsAdultNil(b bool)`

 SetIsAdultNil sets the value for IsAdult to be an explicit nil

### UnsetIsAdult
`func (o *CzechBankIdProviderOutput) UnsetIsAdult()`

UnsetIsAdult ensures that no value is present for IsAdult, not even an explicit nil
### GetDateOfDeath

`func (o *CzechBankIdProviderOutput) GetDateOfDeath() string`

GetDateOfDeath returns the DateOfDeath field if non-nil, zero value otherwise.

### GetDateOfDeathOk

`func (o *CzechBankIdProviderOutput) GetDateOfDeathOk() (*string, bool)`

GetDateOfDeathOk returns a tuple with the DateOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDeath

`func (o *CzechBankIdProviderOutput) SetDateOfDeath(v string)`

SetDateOfDeath sets DateOfDeath field to given value.

### HasDateOfDeath

`func (o *CzechBankIdProviderOutput) HasDateOfDeath() bool`

HasDateOfDeath returns a boolean if a field has been set.

### SetDateOfDeathNil

`func (o *CzechBankIdProviderOutput) SetDateOfDeathNil(b bool)`

 SetDateOfDeathNil sets the value for DateOfDeath to be an explicit nil

### UnsetDateOfDeath
`func (o *CzechBankIdProviderOutput) UnsetDateOfDeath()`

UnsetDateOfDeath ensures that no value is present for DateOfDeath, not even an explicit nil
### GetBirthPlace

`func (o *CzechBankIdProviderOutput) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *CzechBankIdProviderOutput) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *CzechBankIdProviderOutput) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *CzechBankIdProviderOutput) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### SetBirthPlaceNil

`func (o *CzechBankIdProviderOutput) SetBirthPlaceNil(b bool)`

 SetBirthPlaceNil sets the value for BirthPlace to be an explicit nil

### UnsetBirthPlace
`func (o *CzechBankIdProviderOutput) UnsetBirthPlace()`

UnsetBirthPlace ensures that no value is present for BirthPlace, not even an explicit nil
### GetBirthCountry

`func (o *CzechBankIdProviderOutput) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *CzechBankIdProviderOutput) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *CzechBankIdProviderOutput) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *CzechBankIdProviderOutput) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *CzechBankIdProviderOutput) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *CzechBankIdProviderOutput) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetPrimaryNationality

`func (o *CzechBankIdProviderOutput) GetPrimaryNationality() string`

GetPrimaryNationality returns the PrimaryNationality field if non-nil, zero value otherwise.

### GetPrimaryNationalityOk

`func (o *CzechBankIdProviderOutput) GetPrimaryNationalityOk() (*string, bool)`

GetPrimaryNationalityOk returns a tuple with the PrimaryNationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNationality

`func (o *CzechBankIdProviderOutput) SetPrimaryNationality(v string)`

SetPrimaryNationality sets PrimaryNationality field to given value.

### HasPrimaryNationality

`func (o *CzechBankIdProviderOutput) HasPrimaryNationality() bool`

HasPrimaryNationality returns a boolean if a field has been set.

### SetPrimaryNationalityNil

`func (o *CzechBankIdProviderOutput) SetPrimaryNationalityNil(b bool)`

 SetPrimaryNationalityNil sets the value for PrimaryNationality to be an explicit nil

### UnsetPrimaryNationality
`func (o *CzechBankIdProviderOutput) UnsetPrimaryNationality()`

UnsetPrimaryNationality ensures that no value is present for PrimaryNationality, not even an explicit nil
### GetNationalities

`func (o *CzechBankIdProviderOutput) GetNationalities() []string`

GetNationalities returns the Nationalities field if non-nil, zero value otherwise.

### GetNationalitiesOk

`func (o *CzechBankIdProviderOutput) GetNationalitiesOk() (*[]string, bool)`

GetNationalitiesOk returns a tuple with the Nationalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalities

`func (o *CzechBankIdProviderOutput) SetNationalities(v []string)`

SetNationalities sets Nationalities field to given value.

### HasNationalities

`func (o *CzechBankIdProviderOutput) HasNationalities() bool`

HasNationalities returns a boolean if a field has been set.

### SetNationalitiesNil

`func (o *CzechBankIdProviderOutput) SetNationalitiesNil(b bool)`

 SetNationalitiesNil sets the value for Nationalities to be an explicit nil

### UnsetNationalities
`func (o *CzechBankIdProviderOutput) UnsetNationalities()`

UnsetNationalities ensures that no value is present for Nationalities, not even an explicit nil
### GetMaritalStatus

`func (o *CzechBankIdProviderOutput) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *CzechBankIdProviderOutput) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *CzechBankIdProviderOutput) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *CzechBankIdProviderOutput) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### SetMaritalStatusNil

`func (o *CzechBankIdProviderOutput) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *CzechBankIdProviderOutput) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetEmail

`func (o *CzechBankIdProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CzechBankIdProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CzechBankIdProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CzechBankIdProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CzechBankIdProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CzechBankIdProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *CzechBankIdProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CzechBankIdProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CzechBankIdProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CzechBankIdProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *CzechBankIdProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *CzechBankIdProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPhoneNumber

`func (o *CzechBankIdProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CzechBankIdProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CzechBankIdProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CzechBankIdProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *CzechBankIdProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *CzechBankIdProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *CzechBankIdProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *CzechBankIdProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *CzechBankIdProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *CzechBankIdProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *CzechBankIdProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *CzechBankIdProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetIsPoliticallyExposedPerson

`func (o *CzechBankIdProviderOutput) GetIsPoliticallyExposedPerson() bool`

GetIsPoliticallyExposedPerson returns the IsPoliticallyExposedPerson field if non-nil, zero value otherwise.

### GetIsPoliticallyExposedPersonOk

`func (o *CzechBankIdProviderOutput) GetIsPoliticallyExposedPersonOk() (*bool, bool)`

GetIsPoliticallyExposedPersonOk returns a tuple with the IsPoliticallyExposedPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPoliticallyExposedPerson

`func (o *CzechBankIdProviderOutput) SetIsPoliticallyExposedPerson(v bool)`

SetIsPoliticallyExposedPerson sets IsPoliticallyExposedPerson field to given value.

### HasIsPoliticallyExposedPerson

`func (o *CzechBankIdProviderOutput) HasIsPoliticallyExposedPerson() bool`

HasIsPoliticallyExposedPerson returns a boolean if a field has been set.

### SetIsPoliticallyExposedPersonNil

`func (o *CzechBankIdProviderOutput) SetIsPoliticallyExposedPersonNil(b bool)`

 SetIsPoliticallyExposedPersonNil sets the value for IsPoliticallyExposedPerson to be an explicit nil

### UnsetIsPoliticallyExposedPerson
`func (o *CzechBankIdProviderOutput) UnsetIsPoliticallyExposedPerson()`

UnsetIsPoliticallyExposedPerson ensures that no value is present for IsPoliticallyExposedPerson, not even an explicit nil
### GetLimitedLegalCapacity

`func (o *CzechBankIdProviderOutput) GetLimitedLegalCapacity() bool`

GetLimitedLegalCapacity returns the LimitedLegalCapacity field if non-nil, zero value otherwise.

### GetLimitedLegalCapacityOk

`func (o *CzechBankIdProviderOutput) GetLimitedLegalCapacityOk() (*bool, bool)`

GetLimitedLegalCapacityOk returns a tuple with the LimitedLegalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedLegalCapacity

`func (o *CzechBankIdProviderOutput) SetLimitedLegalCapacity(v bool)`

SetLimitedLegalCapacity sets LimitedLegalCapacity field to given value.

### HasLimitedLegalCapacity

`func (o *CzechBankIdProviderOutput) HasLimitedLegalCapacity() bool`

HasLimitedLegalCapacity returns a boolean if a field has been set.

### SetLimitedLegalCapacityNil

`func (o *CzechBankIdProviderOutput) SetLimitedLegalCapacityNil(b bool)`

 SetLimitedLegalCapacityNil sets the value for LimitedLegalCapacity to be an explicit nil

### UnsetLimitedLegalCapacity
`func (o *CzechBankIdProviderOutput) UnsetLimitedLegalCapacity()`

UnsetLimitedLegalCapacity ensures that no value is present for LimitedLegalCapacity, not even an explicit nil
### GetVerification

`func (o *CzechBankIdProviderOutput) GetVerification() CzechBankIdVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *CzechBankIdProviderOutput) GetVerificationOk() (*CzechBankIdVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *CzechBankIdProviderOutput) SetVerification(v CzechBankIdVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *CzechBankIdProviderOutput) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *CzechBankIdProviderOutput) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *CzechBankIdProviderOutput) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil
### GetAddresses

`func (o *CzechBankIdProviderOutput) GetAddresses() []CzechBankIdAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CzechBankIdProviderOutput) GetAddressesOk() (*[]CzechBankIdAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CzechBankIdProviderOutput) SetAddresses(v []CzechBankIdAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CzechBankIdProviderOutput) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *CzechBankIdProviderOutput) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *CzechBankIdProviderOutput) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIdCards

`func (o *CzechBankIdProviderOutput) GetIdCards() []CzechBankIdCard`

GetIdCards returns the IdCards field if non-nil, zero value otherwise.

### GetIdCardsOk

`func (o *CzechBankIdProviderOutput) GetIdCardsOk() (*[]CzechBankIdCard, bool)`

GetIdCardsOk returns a tuple with the IdCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCards

`func (o *CzechBankIdProviderOutput) SetIdCards(v []CzechBankIdCard)`

SetIdCards sets IdCards field to given value.

### HasIdCards

`func (o *CzechBankIdProviderOutput) HasIdCards() bool`

HasIdCards returns a boolean if a field has been set.

### SetIdCardsNil

`func (o *CzechBankIdProviderOutput) SetIdCardsNil(b bool)`

 SetIdCardsNil sets the value for IdCards to be an explicit nil

### UnsetIdCards
`func (o *CzechBankIdProviderOutput) UnsetIdCards()`

UnsetIdCards ensures that no value is present for IdCards, not even an explicit nil
### GetUpdatedAt

`func (o *CzechBankIdProviderOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CzechBankIdProviderOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CzechBankIdProviderOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CzechBankIdProviderOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CzechBankIdProviderOutput) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CzechBankIdProviderOutput) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetZoneInfo

`func (o *CzechBankIdProviderOutput) GetZoneInfo() string`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CzechBankIdProviderOutput) GetZoneInfoOk() (*string, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CzechBankIdProviderOutput) SetZoneInfo(v string)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CzechBankIdProviderOutput) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CzechBankIdProviderOutput) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CzechBankIdProviderOutput) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil
### GetLocale

`func (o *CzechBankIdProviderOutput) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CzechBankIdProviderOutput) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CzechBankIdProviderOutput) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CzechBankIdProviderOutput) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *CzechBankIdProviderOutput) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *CzechBankIdProviderOutput) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetTitles

`func (o *CzechBankIdProviderOutput) GetTitles() []CzechBankIdTitle`

GetTitles returns the Titles field if non-nil, zero value otherwise.

### GetTitlesOk

`func (o *CzechBankIdProviderOutput) GetTitlesOk() (*[]CzechBankIdTitle, bool)`

GetTitlesOk returns a tuple with the Titles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitles

`func (o *CzechBankIdProviderOutput) SetTitles(v []CzechBankIdTitle)`

SetTitles sets Titles field to given value.

### HasTitles

`func (o *CzechBankIdProviderOutput) HasTitles() bool`

HasTitles returns a boolean if a field has been set.

### SetTitlesNil

`func (o *CzechBankIdProviderOutput) SetTitlesNil(b bool)`

 SetTitlesNil sets the value for Titles to be an explicit nil

### UnsetTitles
`func (o *CzechBankIdProviderOutput) UnsetTitles()`

UnsetTitles ensures that no value is present for Titles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



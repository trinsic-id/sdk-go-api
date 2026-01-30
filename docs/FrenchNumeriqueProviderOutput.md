# FrenchNumeriqueProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to [**NullableFrenchNumeriqueGivenName**](FrenchNumeriqueGivenName.md) | Given name information including the full given name and its structured components (first name and middle name(s)). | [optional] 
**FamilyName** | Pointer to **NullableString** | Birth name (nom de naissance) as recorded on the user&#39;s French identity document. In France, this is the family name given at birth and may differ from the usage name. | [optional] 
**PreferredSurname** | Pointer to **NullableString** | Usage surname (\&quot;nom d&#39;usage\&quot;) - according to French law, this is the surname used in daily life. This is typically a married name or other preferred surname, as opposed to FamilyName which is the birth surname (\&quot;nom de naissance\&quot;). | [optional] 
**Birthdate** | Pointer to **NullableString** | Date of birth in YYYY-MM-DD format, as recorded on the user&#39;s French identity document. | [optional] 
**Nationality** | Pointer to [**NullableFrenchNumeriqueNationality**](FrenchNumeriqueNationality.md) | Nationality information with ISO 3166 alpha-3 code and French label. The label is provided in French (e.g., \&quot;Française\&quot; for French nationality). | [optional] 
**Sex** | Pointer to **NullableString** | Sex - \&quot;Male\&quot; or \&quot;Female\&quot;. | [optional] 
**Majority** | Pointer to **NullableBool** | Majority status (majorité) indicating whether the user has reached the French age of legal majority (18 years old). This value is computed by the provider from the user&#39;s birthdate. | [optional] 
**PhoneNumber** | Pointer to [**NullableFrenchNumeriquePhone**](FrenchNumeriquePhone.md) | Phone number information including the full phone number and its structured components (country prefix and national number). | [optional] 
**PhoneNumberVerified** | Pointer to **NullableBool** | Whether the phone number has been verified by the provider. | [optional] 
**Email** | Pointer to **NullableString** | Email address registered in the user&#39;s digital identity account. | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the email address has been verified by the provider. | [optional] 
**Birthplace** | Pointer to **NullableString** | National Institute of Statistics and Economic Studies (INSEE) official geographic code (COG) of the birthplace. This is a 5-digit French administrative code identifying the commune (municipality) of birth. For births in metropolitan France, the first two digits represent the department.              See: https://www.insee.fr/fr/information/2560452 | [optional] 
**BirthplaceLabel** | Pointer to **NullableString** | Name of the birthplace commune (municipality). If applicable, this includes the full name of the arrondissement (borough). | [optional] 
**BirthCountry** | Pointer to **NullableString** | National Institute of Statistics and Economic Studies (INSEE) official geographic code (COG) for the user&#39;s country of birth.              LaPoste returns this as a 5-character numeric string. This is a country/territory COG, so it always starts with \&quot;99\&quot; (99xxx). France is coded as 99100.              See full list here: https://www.insee.fr/fr/information/2560452 | [optional] 
**BirthCountryIso** | Pointer to **NullableString** | Birth country as an ISO 3166 alpha-3 code. | [optional] 
**BirthCountryLabel** | Pointer to **NullableString** | Label of the birth country, in English. | [optional] 
**BirthDepartment** | Pointer to **NullableString** | Number of the department of birth (e.g., \&quot;75\&quot; for Paris, \&quot;2A\&quot; for Corse-du-Sud).              See full list here: https://en.wikipedia.org/wiki/Departments_of_France | [optional] 
**DigitalIdentityCreationDate** | Pointer to **NullableString** | Date when the user&#39;s digital identity was created. | [optional] 
**DigitalIdentityExpirationDate** | Pointer to **NullableString** | Date when the digital identity expires. It is valid for 5 years from the date of identity verification. Users are notified ~1 month before and must re-verify identity via the app or in-person at La Poste. After expiration, the identity is deactivated but can be renewed within 1 year; otherwise deleted. | [optional] 
**IdentityDocumentType** | Pointer to **NullableString** | Type of identity document used for verification. Can be \&quot;ID_CARD\&quot;, \&quot;PASSPORT\&quot;, or \&quot;RESIDENCE_PERMIT\&quot;. | [optional] 
**IdentityDocumentNumber** | Pointer to **NullableString** | Identity document number as it appears on the ID Card, Passport or Residence Permit. | [optional] 
**IdentityDocumentEmittingDate** | Pointer to **NullableString** | Date when the identity document was issued. | [optional] 
**IdentityDocumentExpirationDate** | Pointer to **NullableString** | Date when the identity document expires. | [optional] 
**IdentityDocumentEmittingCountry** | Pointer to **NullableString** | Country that issued the identity document (ISO 3166 alpha-3 code). | [optional] 
**IdentityDocumentMrz** | Pointer to **NullableString** | Machine Readable Zone (MRZ) data from the ID Card, Passport or Residence Permit. | [optional] 

## Methods

### NewFrenchNumeriqueProviderOutput

`func NewFrenchNumeriqueProviderOutput() *FrenchNumeriqueProviderOutput`

NewFrenchNumeriqueProviderOutput instantiates a new FrenchNumeriqueProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrenchNumeriqueProviderOutputWithDefaults

`func NewFrenchNumeriqueProviderOutputWithDefaults() *FrenchNumeriqueProviderOutput`

NewFrenchNumeriqueProviderOutputWithDefaults instantiates a new FrenchNumeriqueProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *FrenchNumeriqueProviderOutput) GetGivenName() FrenchNumeriqueGivenName`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *FrenchNumeriqueProviderOutput) GetGivenNameOk() (*FrenchNumeriqueGivenName, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *FrenchNumeriqueProviderOutput) SetGivenName(v FrenchNumeriqueGivenName)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *FrenchNumeriqueProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *FrenchNumeriqueProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *FrenchNumeriqueProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *FrenchNumeriqueProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *FrenchNumeriqueProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *FrenchNumeriqueProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *FrenchNumeriqueProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *FrenchNumeriqueProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *FrenchNumeriqueProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetPreferredSurname

`func (o *FrenchNumeriqueProviderOutput) GetPreferredSurname() string`

GetPreferredSurname returns the PreferredSurname field if non-nil, zero value otherwise.

### GetPreferredSurnameOk

`func (o *FrenchNumeriqueProviderOutput) GetPreferredSurnameOk() (*string, bool)`

GetPreferredSurnameOk returns a tuple with the PreferredSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSurname

`func (o *FrenchNumeriqueProviderOutput) SetPreferredSurname(v string)`

SetPreferredSurname sets PreferredSurname field to given value.

### HasPreferredSurname

`func (o *FrenchNumeriqueProviderOutput) HasPreferredSurname() bool`

HasPreferredSurname returns a boolean if a field has been set.

### SetPreferredSurnameNil

`func (o *FrenchNumeriqueProviderOutput) SetPreferredSurnameNil(b bool)`

 SetPreferredSurnameNil sets the value for PreferredSurname to be an explicit nil

### UnsetPreferredSurname
`func (o *FrenchNumeriqueProviderOutput) UnsetPreferredSurname()`

UnsetPreferredSurname ensures that no value is present for PreferredSurname, not even an explicit nil
### GetBirthdate

`func (o *FrenchNumeriqueProviderOutput) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *FrenchNumeriqueProviderOutput) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *FrenchNumeriqueProviderOutput) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### SetBirthdateNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthdateNil(b bool)`

 SetBirthdateNil sets the value for Birthdate to be an explicit nil

### UnsetBirthdate
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthdate()`

UnsetBirthdate ensures that no value is present for Birthdate, not even an explicit nil
### GetNationality

`func (o *FrenchNumeriqueProviderOutput) GetNationality() FrenchNumeriqueNationality`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *FrenchNumeriqueProviderOutput) GetNationalityOk() (*FrenchNumeriqueNationality, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *FrenchNumeriqueProviderOutput) SetNationality(v FrenchNumeriqueNationality)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *FrenchNumeriqueProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *FrenchNumeriqueProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *FrenchNumeriqueProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetSex

`func (o *FrenchNumeriqueProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *FrenchNumeriqueProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *FrenchNumeriqueProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *FrenchNumeriqueProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *FrenchNumeriqueProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *FrenchNumeriqueProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetMajority

`func (o *FrenchNumeriqueProviderOutput) GetMajority() bool`

GetMajority returns the Majority field if non-nil, zero value otherwise.

### GetMajorityOk

`func (o *FrenchNumeriqueProviderOutput) GetMajorityOk() (*bool, bool)`

GetMajorityOk returns a tuple with the Majority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajority

`func (o *FrenchNumeriqueProviderOutput) SetMajority(v bool)`

SetMajority sets Majority field to given value.

### HasMajority

`func (o *FrenchNumeriqueProviderOutput) HasMajority() bool`

HasMajority returns a boolean if a field has been set.

### SetMajorityNil

`func (o *FrenchNumeriqueProviderOutput) SetMajorityNil(b bool)`

 SetMajorityNil sets the value for Majority to be an explicit nil

### UnsetMajority
`func (o *FrenchNumeriqueProviderOutput) UnsetMajority()`

UnsetMajority ensures that no value is present for Majority, not even an explicit nil
### GetPhoneNumber

`func (o *FrenchNumeriqueProviderOutput) GetPhoneNumber() FrenchNumeriquePhone`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FrenchNumeriqueProviderOutput) GetPhoneNumberOk() (*FrenchNumeriquePhone, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FrenchNumeriqueProviderOutput) SetPhoneNumber(v FrenchNumeriquePhone)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FrenchNumeriqueProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *FrenchNumeriqueProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *FrenchNumeriqueProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberVerified

`func (o *FrenchNumeriqueProviderOutput) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *FrenchNumeriqueProviderOutput) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *FrenchNumeriqueProviderOutput) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *FrenchNumeriqueProviderOutput) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### SetPhoneNumberVerifiedNil

`func (o *FrenchNumeriqueProviderOutput) SetPhoneNumberVerifiedNil(b bool)`

 SetPhoneNumberVerifiedNil sets the value for PhoneNumberVerified to be an explicit nil

### UnsetPhoneNumberVerified
`func (o *FrenchNumeriqueProviderOutput) UnsetPhoneNumberVerified()`

UnsetPhoneNumberVerified ensures that no value is present for PhoneNumberVerified, not even an explicit nil
### GetEmail

`func (o *FrenchNumeriqueProviderOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FrenchNumeriqueProviderOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FrenchNumeriqueProviderOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FrenchNumeriqueProviderOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *FrenchNumeriqueProviderOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FrenchNumeriqueProviderOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *FrenchNumeriqueProviderOutput) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *FrenchNumeriqueProviderOutput) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *FrenchNumeriqueProviderOutput) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *FrenchNumeriqueProviderOutput) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *FrenchNumeriqueProviderOutput) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *FrenchNumeriqueProviderOutput) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetBirthplace

`func (o *FrenchNumeriqueProviderOutput) GetBirthplace() string`

GetBirthplace returns the Birthplace field if non-nil, zero value otherwise.

### GetBirthplaceOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthplaceOk() (*string, bool)`

GetBirthplaceOk returns a tuple with the Birthplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplace

`func (o *FrenchNumeriqueProviderOutput) SetBirthplace(v string)`

SetBirthplace sets Birthplace field to given value.

### HasBirthplace

`func (o *FrenchNumeriqueProviderOutput) HasBirthplace() bool`

HasBirthplace returns a boolean if a field has been set.

### SetBirthplaceNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthplaceNil(b bool)`

 SetBirthplaceNil sets the value for Birthplace to be an explicit nil

### UnsetBirthplace
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthplace()`

UnsetBirthplace ensures that no value is present for Birthplace, not even an explicit nil
### GetBirthplaceLabel

`func (o *FrenchNumeriqueProviderOutput) GetBirthplaceLabel() string`

GetBirthplaceLabel returns the BirthplaceLabel field if non-nil, zero value otherwise.

### GetBirthplaceLabelOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthplaceLabelOk() (*string, bool)`

GetBirthplaceLabelOk returns a tuple with the BirthplaceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplaceLabel

`func (o *FrenchNumeriqueProviderOutput) SetBirthplaceLabel(v string)`

SetBirthplaceLabel sets BirthplaceLabel field to given value.

### HasBirthplaceLabel

`func (o *FrenchNumeriqueProviderOutput) HasBirthplaceLabel() bool`

HasBirthplaceLabel returns a boolean if a field has been set.

### SetBirthplaceLabelNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthplaceLabelNil(b bool)`

 SetBirthplaceLabelNil sets the value for BirthplaceLabel to be an explicit nil

### UnsetBirthplaceLabel
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthplaceLabel()`

UnsetBirthplaceLabel ensures that no value is present for BirthplaceLabel, not even an explicit nil
### GetBirthCountry

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *FrenchNumeriqueProviderOutput) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetBirthCountryIso

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountryIso() string`

GetBirthCountryIso returns the BirthCountryIso field if non-nil, zero value otherwise.

### GetBirthCountryIsoOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountryIsoOk() (*string, bool)`

GetBirthCountryIsoOk returns a tuple with the BirthCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountryIso

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountryIso(v string)`

SetBirthCountryIso sets BirthCountryIso field to given value.

### HasBirthCountryIso

`func (o *FrenchNumeriqueProviderOutput) HasBirthCountryIso() bool`

HasBirthCountryIso returns a boolean if a field has been set.

### SetBirthCountryIsoNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountryIsoNil(b bool)`

 SetBirthCountryIsoNil sets the value for BirthCountryIso to be an explicit nil

### UnsetBirthCountryIso
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthCountryIso()`

UnsetBirthCountryIso ensures that no value is present for BirthCountryIso, not even an explicit nil
### GetBirthCountryLabel

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountryLabel() string`

GetBirthCountryLabel returns the BirthCountryLabel field if non-nil, zero value otherwise.

### GetBirthCountryLabelOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthCountryLabelOk() (*string, bool)`

GetBirthCountryLabelOk returns a tuple with the BirthCountryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountryLabel

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountryLabel(v string)`

SetBirthCountryLabel sets BirthCountryLabel field to given value.

### HasBirthCountryLabel

`func (o *FrenchNumeriqueProviderOutput) HasBirthCountryLabel() bool`

HasBirthCountryLabel returns a boolean if a field has been set.

### SetBirthCountryLabelNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthCountryLabelNil(b bool)`

 SetBirthCountryLabelNil sets the value for BirthCountryLabel to be an explicit nil

### UnsetBirthCountryLabel
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthCountryLabel()`

UnsetBirthCountryLabel ensures that no value is present for BirthCountryLabel, not even an explicit nil
### GetBirthDepartment

`func (o *FrenchNumeriqueProviderOutput) GetBirthDepartment() string`

GetBirthDepartment returns the BirthDepartment field if non-nil, zero value otherwise.

### GetBirthDepartmentOk

`func (o *FrenchNumeriqueProviderOutput) GetBirthDepartmentOk() (*string, bool)`

GetBirthDepartmentOk returns a tuple with the BirthDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDepartment

`func (o *FrenchNumeriqueProviderOutput) SetBirthDepartment(v string)`

SetBirthDepartment sets BirthDepartment field to given value.

### HasBirthDepartment

`func (o *FrenchNumeriqueProviderOutput) HasBirthDepartment() bool`

HasBirthDepartment returns a boolean if a field has been set.

### SetBirthDepartmentNil

`func (o *FrenchNumeriqueProviderOutput) SetBirthDepartmentNil(b bool)`

 SetBirthDepartmentNil sets the value for BirthDepartment to be an explicit nil

### UnsetBirthDepartment
`func (o *FrenchNumeriqueProviderOutput) UnsetBirthDepartment()`

UnsetBirthDepartment ensures that no value is present for BirthDepartment, not even an explicit nil
### GetDigitalIdentityCreationDate

`func (o *FrenchNumeriqueProviderOutput) GetDigitalIdentityCreationDate() string`

GetDigitalIdentityCreationDate returns the DigitalIdentityCreationDate field if non-nil, zero value otherwise.

### GetDigitalIdentityCreationDateOk

`func (o *FrenchNumeriqueProviderOutput) GetDigitalIdentityCreationDateOk() (*string, bool)`

GetDigitalIdentityCreationDateOk returns a tuple with the DigitalIdentityCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalIdentityCreationDate

`func (o *FrenchNumeriqueProviderOutput) SetDigitalIdentityCreationDate(v string)`

SetDigitalIdentityCreationDate sets DigitalIdentityCreationDate field to given value.

### HasDigitalIdentityCreationDate

`func (o *FrenchNumeriqueProviderOutput) HasDigitalIdentityCreationDate() bool`

HasDigitalIdentityCreationDate returns a boolean if a field has been set.

### SetDigitalIdentityCreationDateNil

`func (o *FrenchNumeriqueProviderOutput) SetDigitalIdentityCreationDateNil(b bool)`

 SetDigitalIdentityCreationDateNil sets the value for DigitalIdentityCreationDate to be an explicit nil

### UnsetDigitalIdentityCreationDate
`func (o *FrenchNumeriqueProviderOutput) UnsetDigitalIdentityCreationDate()`

UnsetDigitalIdentityCreationDate ensures that no value is present for DigitalIdentityCreationDate, not even an explicit nil
### GetDigitalIdentityExpirationDate

`func (o *FrenchNumeriqueProviderOutput) GetDigitalIdentityExpirationDate() string`

GetDigitalIdentityExpirationDate returns the DigitalIdentityExpirationDate field if non-nil, zero value otherwise.

### GetDigitalIdentityExpirationDateOk

`func (o *FrenchNumeriqueProviderOutput) GetDigitalIdentityExpirationDateOk() (*string, bool)`

GetDigitalIdentityExpirationDateOk returns a tuple with the DigitalIdentityExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalIdentityExpirationDate

`func (o *FrenchNumeriqueProviderOutput) SetDigitalIdentityExpirationDate(v string)`

SetDigitalIdentityExpirationDate sets DigitalIdentityExpirationDate field to given value.

### HasDigitalIdentityExpirationDate

`func (o *FrenchNumeriqueProviderOutput) HasDigitalIdentityExpirationDate() bool`

HasDigitalIdentityExpirationDate returns a boolean if a field has been set.

### SetDigitalIdentityExpirationDateNil

`func (o *FrenchNumeriqueProviderOutput) SetDigitalIdentityExpirationDateNil(b bool)`

 SetDigitalIdentityExpirationDateNil sets the value for DigitalIdentityExpirationDate to be an explicit nil

### UnsetDigitalIdentityExpirationDate
`func (o *FrenchNumeriqueProviderOutput) UnsetDigitalIdentityExpirationDate()`

UnsetDigitalIdentityExpirationDate ensures that no value is present for DigitalIdentityExpirationDate, not even an explicit nil
### GetIdentityDocumentType

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentType() string`

GetIdentityDocumentType returns the IdentityDocumentType field if non-nil, zero value otherwise.

### GetIdentityDocumentTypeOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentTypeOk() (*string, bool)`

GetIdentityDocumentTypeOk returns a tuple with the IdentityDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentType

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentType(v string)`

SetIdentityDocumentType sets IdentityDocumentType field to given value.

### HasIdentityDocumentType

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentType() bool`

HasIdentityDocumentType returns a boolean if a field has been set.

### SetIdentityDocumentTypeNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentTypeNil(b bool)`

 SetIdentityDocumentTypeNil sets the value for IdentityDocumentType to be an explicit nil

### UnsetIdentityDocumentType
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentType()`

UnsetIdentityDocumentType ensures that no value is present for IdentityDocumentType, not even an explicit nil
### GetIdentityDocumentNumber

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentNumber() string`

GetIdentityDocumentNumber returns the IdentityDocumentNumber field if non-nil, zero value otherwise.

### GetIdentityDocumentNumberOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentNumberOk() (*string, bool)`

GetIdentityDocumentNumberOk returns a tuple with the IdentityDocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentNumber

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentNumber(v string)`

SetIdentityDocumentNumber sets IdentityDocumentNumber field to given value.

### HasIdentityDocumentNumber

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentNumber() bool`

HasIdentityDocumentNumber returns a boolean if a field has been set.

### SetIdentityDocumentNumberNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentNumberNil(b bool)`

 SetIdentityDocumentNumberNil sets the value for IdentityDocumentNumber to be an explicit nil

### UnsetIdentityDocumentNumber
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentNumber()`

UnsetIdentityDocumentNumber ensures that no value is present for IdentityDocumentNumber, not even an explicit nil
### GetIdentityDocumentEmittingDate

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentEmittingDate() string`

GetIdentityDocumentEmittingDate returns the IdentityDocumentEmittingDate field if non-nil, zero value otherwise.

### GetIdentityDocumentEmittingDateOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentEmittingDateOk() (*string, bool)`

GetIdentityDocumentEmittingDateOk returns a tuple with the IdentityDocumentEmittingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentEmittingDate

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentEmittingDate(v string)`

SetIdentityDocumentEmittingDate sets IdentityDocumentEmittingDate field to given value.

### HasIdentityDocumentEmittingDate

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentEmittingDate() bool`

HasIdentityDocumentEmittingDate returns a boolean if a field has been set.

### SetIdentityDocumentEmittingDateNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentEmittingDateNil(b bool)`

 SetIdentityDocumentEmittingDateNil sets the value for IdentityDocumentEmittingDate to be an explicit nil

### UnsetIdentityDocumentEmittingDate
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentEmittingDate()`

UnsetIdentityDocumentEmittingDate ensures that no value is present for IdentityDocumentEmittingDate, not even an explicit nil
### GetIdentityDocumentExpirationDate

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentExpirationDate() string`

GetIdentityDocumentExpirationDate returns the IdentityDocumentExpirationDate field if non-nil, zero value otherwise.

### GetIdentityDocumentExpirationDateOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentExpirationDateOk() (*string, bool)`

GetIdentityDocumentExpirationDateOk returns a tuple with the IdentityDocumentExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentExpirationDate

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentExpirationDate(v string)`

SetIdentityDocumentExpirationDate sets IdentityDocumentExpirationDate field to given value.

### HasIdentityDocumentExpirationDate

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentExpirationDate() bool`

HasIdentityDocumentExpirationDate returns a boolean if a field has been set.

### SetIdentityDocumentExpirationDateNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentExpirationDateNil(b bool)`

 SetIdentityDocumentExpirationDateNil sets the value for IdentityDocumentExpirationDate to be an explicit nil

### UnsetIdentityDocumentExpirationDate
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentExpirationDate()`

UnsetIdentityDocumentExpirationDate ensures that no value is present for IdentityDocumentExpirationDate, not even an explicit nil
### GetIdentityDocumentEmittingCountry

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentEmittingCountry() string`

GetIdentityDocumentEmittingCountry returns the IdentityDocumentEmittingCountry field if non-nil, zero value otherwise.

### GetIdentityDocumentEmittingCountryOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentEmittingCountryOk() (*string, bool)`

GetIdentityDocumentEmittingCountryOk returns a tuple with the IdentityDocumentEmittingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentEmittingCountry

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentEmittingCountry(v string)`

SetIdentityDocumentEmittingCountry sets IdentityDocumentEmittingCountry field to given value.

### HasIdentityDocumentEmittingCountry

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentEmittingCountry() bool`

HasIdentityDocumentEmittingCountry returns a boolean if a field has been set.

### SetIdentityDocumentEmittingCountryNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentEmittingCountryNil(b bool)`

 SetIdentityDocumentEmittingCountryNil sets the value for IdentityDocumentEmittingCountry to be an explicit nil

### UnsetIdentityDocumentEmittingCountry
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentEmittingCountry()`

UnsetIdentityDocumentEmittingCountry ensures that no value is present for IdentityDocumentEmittingCountry, not even an explicit nil
### GetIdentityDocumentMrz

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentMrz() string`

GetIdentityDocumentMrz returns the IdentityDocumentMrz field if non-nil, zero value otherwise.

### GetIdentityDocumentMrzOk

`func (o *FrenchNumeriqueProviderOutput) GetIdentityDocumentMrzOk() (*string, bool)`

GetIdentityDocumentMrzOk returns a tuple with the IdentityDocumentMrz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentMrz

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentMrz(v string)`

SetIdentityDocumentMrz sets IdentityDocumentMrz field to given value.

### HasIdentityDocumentMrz

`func (o *FrenchNumeriqueProviderOutput) HasIdentityDocumentMrz() bool`

HasIdentityDocumentMrz returns a boolean if a field has been set.

### SetIdentityDocumentMrzNil

`func (o *FrenchNumeriqueProviderOutput) SetIdentityDocumentMrzNil(b bool)`

 SetIdentityDocumentMrzNil sets the value for IdentityDocumentMrz to be an explicit nil

### UnsetIdentityDocumentMrz
`func (o *FrenchNumeriqueProviderOutput) UnsetIdentityDocumentMrz()`

UnsetIdentityDocumentMrz ensures that no value is present for IdentityDocumentMrz, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MobileIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | The given name of the individual, extracted from the Mobile ID authentication certificate&#39;s Subject Distinguished Name \&quot;G\&quot; (givenName) attribute. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name (surname) of the individual, extracted from the Mobile ID authentication certificate&#39;s Subject Distinguished Name \&quot;SN\&quot; (surname) attribute. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The individual&#39;s date of birth, derived from the personal code. | [optional] 
**Sex** | Pointer to **NullableString** | The individual&#39;s sex, derived from the first digit of the personal code.              Possible values: - Male - Female | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code extracted from the Mobile ID authentication certificate&#39;s Subject Distinguished Name \&quot;C\&quot; (country) attribute.              This represents the country of the certificate issuer, not the person&#39;s nationality or citizenship. For Mobile ID, this will be \&quot;EE\&quot; (Estonia) or \&quot;LT\&quot; (Lithuania). | [optional] 
**IdentityType** | Pointer to **NullableString** | The identity document type, extracted from the first 3 characters of the SERIALNUMBER field (before the country code).              Possible values: - PNO: Personal Number (national civic registration number) - PAS: Passport - IDC: National identity card | [optional] 
**PersonalCode** | Pointer to **NullableString** | The personal code (Estonian: isikukood, Lithuanian: asmens kodas) of the individual, extracted from the SERIALNUMBER field of the Mobile ID authentication certificate. This is the portion after the identity type and country prefix (e.g., \&quot;48501010014\&quot; from \&quot;PNOEE-48501010014\&quot;).              Both Estonian and Lithuanian personal codes are 11 digits in the format GYYMMDDSSSC where: - G &#x3D; century/gender (3-4 &#x3D; 1900s, 5-6 &#x3D; 2000s; odd &#x3D; male, even &#x3D; female) - YYMMDD &#x3D; date of birth - SSS &#x3D; sequence number - C &#x3D; checksum digit              See: - https://en.wikipedia.org/wiki/National_identification_number#Estonia - https://en.wikipedia.org/wiki/National_identification_number#Lithuania | [optional] 
**SerialNumber** | Pointer to **NullableString** | The SERIALNUMBER attribute from the Mobile ID authentication certificate&#39;s Subject Distinguished Name. Format: \&quot;{identity-type}{country-code}-{identifier}\&quot;              Components: - identity-type (3 chars): PNO (Personal Number), PAS (Passport), IDC (ID Card) - country-code (2 chars): ISO 3166-1 alpha-2 (EE, LT) - identifier: The personal code | [optional] 
**CertificateSubject** | Pointer to **NullableString** | The full Subject Distinguished Name (Subject DN) from the Mobile ID authentication certificate. Contains comma-separated RDN (Relative Distinguished Name) components including C (Country), CN (Common Name), SN (Surname), G (Given Name), and SERIALNUMBER (Personal identifier). | [optional] 

## Methods

### NewMobileIdProviderOutput

`func NewMobileIdProviderOutput() *MobileIdProviderOutput`

NewMobileIdProviderOutput instantiates a new MobileIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileIdProviderOutputWithDefaults

`func NewMobileIdProviderOutputWithDefaults() *MobileIdProviderOutput`

NewMobileIdProviderOutputWithDefaults instantiates a new MobileIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *MobileIdProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MobileIdProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MobileIdProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MobileIdProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MobileIdProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MobileIdProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *MobileIdProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MobileIdProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MobileIdProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MobileIdProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *MobileIdProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *MobileIdProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *MobileIdProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MobileIdProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MobileIdProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MobileIdProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MobileIdProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MobileIdProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *MobileIdProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *MobileIdProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *MobileIdProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *MobileIdProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *MobileIdProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *MobileIdProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetCountry

`func (o *MobileIdProviderOutput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MobileIdProviderOutput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MobileIdProviderOutput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MobileIdProviderOutput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *MobileIdProviderOutput) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *MobileIdProviderOutput) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetIdentityType

`func (o *MobileIdProviderOutput) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *MobileIdProviderOutput) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *MobileIdProviderOutput) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *MobileIdProviderOutput) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *MobileIdProviderOutput) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *MobileIdProviderOutput) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetPersonalCode

`func (o *MobileIdProviderOutput) GetPersonalCode() string`

GetPersonalCode returns the PersonalCode field if non-nil, zero value otherwise.

### GetPersonalCodeOk

`func (o *MobileIdProviderOutput) GetPersonalCodeOk() (*string, bool)`

GetPersonalCodeOk returns a tuple with the PersonalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCode

`func (o *MobileIdProviderOutput) SetPersonalCode(v string)`

SetPersonalCode sets PersonalCode field to given value.

### HasPersonalCode

`func (o *MobileIdProviderOutput) HasPersonalCode() bool`

HasPersonalCode returns a boolean if a field has been set.

### SetPersonalCodeNil

`func (o *MobileIdProviderOutput) SetPersonalCodeNil(b bool)`

 SetPersonalCodeNil sets the value for PersonalCode to be an explicit nil

### UnsetPersonalCode
`func (o *MobileIdProviderOutput) UnsetPersonalCode()`

UnsetPersonalCode ensures that no value is present for PersonalCode, not even an explicit nil
### GetSerialNumber

`func (o *MobileIdProviderOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileIdProviderOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileIdProviderOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileIdProviderOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *MobileIdProviderOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *MobileIdProviderOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetCertificateSubject

`func (o *MobileIdProviderOutput) GetCertificateSubject() string`

GetCertificateSubject returns the CertificateSubject field if non-nil, zero value otherwise.

### GetCertificateSubjectOk

`func (o *MobileIdProviderOutput) GetCertificateSubjectOk() (*string, bool)`

GetCertificateSubjectOk returns a tuple with the CertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSubject

`func (o *MobileIdProviderOutput) SetCertificateSubject(v string)`

SetCertificateSubject sets CertificateSubject field to given value.

### HasCertificateSubject

`func (o *MobileIdProviderOutput) HasCertificateSubject() bool`

HasCertificateSubject returns a boolean if a field has been set.

### SetCertificateSubjectNil

`func (o *MobileIdProviderOutput) SetCertificateSubjectNil(b bool)`

 SetCertificateSubjectNil sets the value for CertificateSubject to be an explicit nil

### UnsetCertificateSubject
`func (o *MobileIdProviderOutput) UnsetCertificateSubject()`

UnsetCertificateSubject ensures that no value is present for CertificateSubject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



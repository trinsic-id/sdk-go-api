# AppleWalletDigitalIdCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenNameUnicode** | Pointer to **NullableString** | First name(s), other name(s), or secondary identifier of the individual.              This field uses the full UTF-8 character set and can represent any name. | [optional] 
**GivenNameLatin1** | Pointer to **NullableString** | First name(s), other name(s), or secondary identifier of the individual.              This field uses the Latin-1 character set and is limited to names represented by the English alphabet. represented using the English alphabet. | [optional] 
**FamilyNameUnicode** | Pointer to **NullableString** | Last name, surname, or primary identifier of the individual.              This field uses the full UTF-8 character set and can represent any name. | [optional] 
**FamilyNameLatin1** | Pointer to **NullableString** | Last name, surname, or primary identifier of the individual.              This field uses only the Latin-1 character set and is largely limited to names which can be represented using the English alphabet. | [optional] 
**Sex** | Pointer to **NullableInt32** | The individual&#39;s sex as an ISO/IEC 5218 code.              Possible values: - 0: Unknown - 1: Male - 2: Female - 9: Not Applicable | [optional] 
**DateOfBirth** | Pointer to [**NullableAppleWalletDigitalIdBirthDate**](AppleWalletDigitalIdBirthDate.md) | The date of birth of the individual, possibly with an \&quot;approximate mask\&quot; indicating uncertain digits. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number of the underlying passport used to create the Digital ID. | [optional] 
**IssuingAuthorityUnicode** | Pointer to **NullableString** | Name or identifier of the issuing authority of the credential, using the full UTF-8 character set.              This field&#39;s contents are arbitrary; it has no guaranteed format. | [optional] 
**IssuingSubdivision** | Pointer to **NullableString** | ISO 3166-2 country-subdivision code for the jurisdiction that issued the credential. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the issuing authority&#39;s country or territory. | [optional] 
**IssueDate** | Pointer to **NullableString** | Date when the underlying passport used to create the Digital ID was issued. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Date when the underlying passport used to create the Digital ID expires. | [optional] 
**AgeInYears** | Pointer to **NullableInt32** | The age of the individual, as of the issuance date of the credential. | [optional] 
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the credential. | [optional] 
**ResidentAddressUnicode** | Pointer to **NullableString** | The permanent address of the individual, using the full UTF-8 character set.              This is only present if the underlying passport contained an address, which is uncommon. | [optional] 
**ResidentCityUnicode** | Pointer to **NullableString** | City of the individual&#39;s permanent address, using the full UTF-8 character set.              This is only present if the underlying passport contained an address, which is uncommon. | [optional] 
**ResidentCityLatin1** | Pointer to **NullableString** | City of the individual&#39;s permanent address, using the Latin-1 character set.              This is only present if the underlying passport contained an address, which is uncommon. | [optional] 
**ResidentPostalCode** | Pointer to **NullableString** | Postal code of the individual&#39;s permanent address.              This is only present if the underlying passport contained an address, which is uncommon. | [optional] 
**ResidentCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the individual&#39;s permanent address.              This is only present if the underlying passport contained an address, which is uncommon. | [optional] 

## Methods

### NewAppleWalletDigitalIdCredential

`func NewAppleWalletDigitalIdCredential() *AppleWalletDigitalIdCredential`

NewAppleWalletDigitalIdCredential instantiates a new AppleWalletDigitalIdCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleWalletDigitalIdCredentialWithDefaults

`func NewAppleWalletDigitalIdCredentialWithDefaults() *AppleWalletDigitalIdCredential`

NewAppleWalletDigitalIdCredentialWithDefaults instantiates a new AppleWalletDigitalIdCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenNameUnicode

`func (o *AppleWalletDigitalIdCredential) GetGivenNameUnicode() string`

GetGivenNameUnicode returns the GivenNameUnicode field if non-nil, zero value otherwise.

### GetGivenNameUnicodeOk

`func (o *AppleWalletDigitalIdCredential) GetGivenNameUnicodeOk() (*string, bool)`

GetGivenNameUnicodeOk returns a tuple with the GivenNameUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameUnicode

`func (o *AppleWalletDigitalIdCredential) SetGivenNameUnicode(v string)`

SetGivenNameUnicode sets GivenNameUnicode field to given value.

### HasGivenNameUnicode

`func (o *AppleWalletDigitalIdCredential) HasGivenNameUnicode() bool`

HasGivenNameUnicode returns a boolean if a field has been set.

### SetGivenNameUnicodeNil

`func (o *AppleWalletDigitalIdCredential) SetGivenNameUnicodeNil(b bool)`

 SetGivenNameUnicodeNil sets the value for GivenNameUnicode to be an explicit nil

### UnsetGivenNameUnicode
`func (o *AppleWalletDigitalIdCredential) UnsetGivenNameUnicode()`

UnsetGivenNameUnicode ensures that no value is present for GivenNameUnicode, not even an explicit nil
### GetGivenNameLatin1

`func (o *AppleWalletDigitalIdCredential) GetGivenNameLatin1() string`

GetGivenNameLatin1 returns the GivenNameLatin1 field if non-nil, zero value otherwise.

### GetGivenNameLatin1Ok

`func (o *AppleWalletDigitalIdCredential) GetGivenNameLatin1Ok() (*string, bool)`

GetGivenNameLatin1Ok returns a tuple with the GivenNameLatin1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameLatin1

`func (o *AppleWalletDigitalIdCredential) SetGivenNameLatin1(v string)`

SetGivenNameLatin1 sets GivenNameLatin1 field to given value.

### HasGivenNameLatin1

`func (o *AppleWalletDigitalIdCredential) HasGivenNameLatin1() bool`

HasGivenNameLatin1 returns a boolean if a field has been set.

### SetGivenNameLatin1Nil

`func (o *AppleWalletDigitalIdCredential) SetGivenNameLatin1Nil(b bool)`

 SetGivenNameLatin1Nil sets the value for GivenNameLatin1 to be an explicit nil

### UnsetGivenNameLatin1
`func (o *AppleWalletDigitalIdCredential) UnsetGivenNameLatin1()`

UnsetGivenNameLatin1 ensures that no value is present for GivenNameLatin1, not even an explicit nil
### GetFamilyNameUnicode

`func (o *AppleWalletDigitalIdCredential) GetFamilyNameUnicode() string`

GetFamilyNameUnicode returns the FamilyNameUnicode field if non-nil, zero value otherwise.

### GetFamilyNameUnicodeOk

`func (o *AppleWalletDigitalIdCredential) GetFamilyNameUnicodeOk() (*string, bool)`

GetFamilyNameUnicodeOk returns a tuple with the FamilyNameUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNameUnicode

`func (o *AppleWalletDigitalIdCredential) SetFamilyNameUnicode(v string)`

SetFamilyNameUnicode sets FamilyNameUnicode field to given value.

### HasFamilyNameUnicode

`func (o *AppleWalletDigitalIdCredential) HasFamilyNameUnicode() bool`

HasFamilyNameUnicode returns a boolean if a field has been set.

### SetFamilyNameUnicodeNil

`func (o *AppleWalletDigitalIdCredential) SetFamilyNameUnicodeNil(b bool)`

 SetFamilyNameUnicodeNil sets the value for FamilyNameUnicode to be an explicit nil

### UnsetFamilyNameUnicode
`func (o *AppleWalletDigitalIdCredential) UnsetFamilyNameUnicode()`

UnsetFamilyNameUnicode ensures that no value is present for FamilyNameUnicode, not even an explicit nil
### GetFamilyNameLatin1

`func (o *AppleWalletDigitalIdCredential) GetFamilyNameLatin1() string`

GetFamilyNameLatin1 returns the FamilyNameLatin1 field if non-nil, zero value otherwise.

### GetFamilyNameLatin1Ok

`func (o *AppleWalletDigitalIdCredential) GetFamilyNameLatin1Ok() (*string, bool)`

GetFamilyNameLatin1Ok returns a tuple with the FamilyNameLatin1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNameLatin1

`func (o *AppleWalletDigitalIdCredential) SetFamilyNameLatin1(v string)`

SetFamilyNameLatin1 sets FamilyNameLatin1 field to given value.

### HasFamilyNameLatin1

`func (o *AppleWalletDigitalIdCredential) HasFamilyNameLatin1() bool`

HasFamilyNameLatin1 returns a boolean if a field has been set.

### SetFamilyNameLatin1Nil

`func (o *AppleWalletDigitalIdCredential) SetFamilyNameLatin1Nil(b bool)`

 SetFamilyNameLatin1Nil sets the value for FamilyNameLatin1 to be an explicit nil

### UnsetFamilyNameLatin1
`func (o *AppleWalletDigitalIdCredential) UnsetFamilyNameLatin1()`

UnsetFamilyNameLatin1 ensures that no value is present for FamilyNameLatin1, not even an explicit nil
### GetSex

`func (o *AppleWalletDigitalIdCredential) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *AppleWalletDigitalIdCredential) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *AppleWalletDigitalIdCredential) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *AppleWalletDigitalIdCredential) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *AppleWalletDigitalIdCredential) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *AppleWalletDigitalIdCredential) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *AppleWalletDigitalIdCredential) GetDateOfBirth() AppleWalletDigitalIdBirthDate`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AppleWalletDigitalIdCredential) GetDateOfBirthOk() (*AppleWalletDigitalIdBirthDate, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AppleWalletDigitalIdCredential) SetDateOfBirth(v AppleWalletDigitalIdBirthDate)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AppleWalletDigitalIdCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AppleWalletDigitalIdCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AppleWalletDigitalIdCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDocumentNumber

`func (o *AppleWalletDigitalIdCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *AppleWalletDigitalIdCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *AppleWalletDigitalIdCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *AppleWalletDigitalIdCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *AppleWalletDigitalIdCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *AppleWalletDigitalIdCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssuingAuthorityUnicode

`func (o *AppleWalletDigitalIdCredential) GetIssuingAuthorityUnicode() string`

GetIssuingAuthorityUnicode returns the IssuingAuthorityUnicode field if non-nil, zero value otherwise.

### GetIssuingAuthorityUnicodeOk

`func (o *AppleWalletDigitalIdCredential) GetIssuingAuthorityUnicodeOk() (*string, bool)`

GetIssuingAuthorityUnicodeOk returns a tuple with the IssuingAuthorityUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthorityUnicode

`func (o *AppleWalletDigitalIdCredential) SetIssuingAuthorityUnicode(v string)`

SetIssuingAuthorityUnicode sets IssuingAuthorityUnicode field to given value.

### HasIssuingAuthorityUnicode

`func (o *AppleWalletDigitalIdCredential) HasIssuingAuthorityUnicode() bool`

HasIssuingAuthorityUnicode returns a boolean if a field has been set.

### SetIssuingAuthorityUnicodeNil

`func (o *AppleWalletDigitalIdCredential) SetIssuingAuthorityUnicodeNil(b bool)`

 SetIssuingAuthorityUnicodeNil sets the value for IssuingAuthorityUnicode to be an explicit nil

### UnsetIssuingAuthorityUnicode
`func (o *AppleWalletDigitalIdCredential) UnsetIssuingAuthorityUnicode()`

UnsetIssuingAuthorityUnicode ensures that no value is present for IssuingAuthorityUnicode, not even an explicit nil
### GetIssuingSubdivision

`func (o *AppleWalletDigitalIdCredential) GetIssuingSubdivision() string`

GetIssuingSubdivision returns the IssuingSubdivision field if non-nil, zero value otherwise.

### GetIssuingSubdivisionOk

`func (o *AppleWalletDigitalIdCredential) GetIssuingSubdivisionOk() (*string, bool)`

GetIssuingSubdivisionOk returns a tuple with the IssuingSubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingSubdivision

`func (o *AppleWalletDigitalIdCredential) SetIssuingSubdivision(v string)`

SetIssuingSubdivision sets IssuingSubdivision field to given value.

### HasIssuingSubdivision

`func (o *AppleWalletDigitalIdCredential) HasIssuingSubdivision() bool`

HasIssuingSubdivision returns a boolean if a field has been set.

### SetIssuingSubdivisionNil

`func (o *AppleWalletDigitalIdCredential) SetIssuingSubdivisionNil(b bool)`

 SetIssuingSubdivisionNil sets the value for IssuingSubdivision to be an explicit nil

### UnsetIssuingSubdivision
`func (o *AppleWalletDigitalIdCredential) UnsetIssuingSubdivision()`

UnsetIssuingSubdivision ensures that no value is present for IssuingSubdivision, not even an explicit nil
### GetIssuingCountry

`func (o *AppleWalletDigitalIdCredential) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *AppleWalletDigitalIdCredential) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *AppleWalletDigitalIdCredential) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *AppleWalletDigitalIdCredential) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *AppleWalletDigitalIdCredential) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *AppleWalletDigitalIdCredential) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssueDate

`func (o *AppleWalletDigitalIdCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *AppleWalletDigitalIdCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *AppleWalletDigitalIdCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *AppleWalletDigitalIdCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *AppleWalletDigitalIdCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *AppleWalletDigitalIdCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *AppleWalletDigitalIdCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AppleWalletDigitalIdCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *AppleWalletDigitalIdCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *AppleWalletDigitalIdCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *AppleWalletDigitalIdCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *AppleWalletDigitalIdCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetAgeInYears

`func (o *AppleWalletDigitalIdCredential) GetAgeInYears() int32`

GetAgeInYears returns the AgeInYears field if non-nil, zero value otherwise.

### GetAgeInYearsOk

`func (o *AppleWalletDigitalIdCredential) GetAgeInYearsOk() (*int32, bool)`

GetAgeInYearsOk returns a tuple with the AgeInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeInYears

`func (o *AppleWalletDigitalIdCredential) SetAgeInYears(v int32)`

SetAgeInYears sets AgeInYears field to given value.

### HasAgeInYears

`func (o *AppleWalletDigitalIdCredential) HasAgeInYears() bool`

HasAgeInYears returns a boolean if a field has been set.

### SetAgeInYearsNil

`func (o *AppleWalletDigitalIdCredential) SetAgeInYearsNil(b bool)`

 SetAgeInYearsNil sets the value for AgeInYears to be an explicit nil

### UnsetAgeInYears
`func (o *AppleWalletDigitalIdCredential) UnsetAgeInYears()`

UnsetAgeInYears ensures that no value is present for AgeInYears, not even an explicit nil
### GetAgeOver

`func (o *AppleWalletDigitalIdCredential) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *AppleWalletDigitalIdCredential) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *AppleWalletDigitalIdCredential) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *AppleWalletDigitalIdCredential) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *AppleWalletDigitalIdCredential) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *AppleWalletDigitalIdCredential) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil
### GetResidentAddressUnicode

`func (o *AppleWalletDigitalIdCredential) GetResidentAddressUnicode() string`

GetResidentAddressUnicode returns the ResidentAddressUnicode field if non-nil, zero value otherwise.

### GetResidentAddressUnicodeOk

`func (o *AppleWalletDigitalIdCredential) GetResidentAddressUnicodeOk() (*string, bool)`

GetResidentAddressUnicodeOk returns a tuple with the ResidentAddressUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentAddressUnicode

`func (o *AppleWalletDigitalIdCredential) SetResidentAddressUnicode(v string)`

SetResidentAddressUnicode sets ResidentAddressUnicode field to given value.

### HasResidentAddressUnicode

`func (o *AppleWalletDigitalIdCredential) HasResidentAddressUnicode() bool`

HasResidentAddressUnicode returns a boolean if a field has been set.

### SetResidentAddressUnicodeNil

`func (o *AppleWalletDigitalIdCredential) SetResidentAddressUnicodeNil(b bool)`

 SetResidentAddressUnicodeNil sets the value for ResidentAddressUnicode to be an explicit nil

### UnsetResidentAddressUnicode
`func (o *AppleWalletDigitalIdCredential) UnsetResidentAddressUnicode()`

UnsetResidentAddressUnicode ensures that no value is present for ResidentAddressUnicode, not even an explicit nil
### GetResidentCityUnicode

`func (o *AppleWalletDigitalIdCredential) GetResidentCityUnicode() string`

GetResidentCityUnicode returns the ResidentCityUnicode field if non-nil, zero value otherwise.

### GetResidentCityUnicodeOk

`func (o *AppleWalletDigitalIdCredential) GetResidentCityUnicodeOk() (*string, bool)`

GetResidentCityUnicodeOk returns a tuple with the ResidentCityUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCityUnicode

`func (o *AppleWalletDigitalIdCredential) SetResidentCityUnicode(v string)`

SetResidentCityUnicode sets ResidentCityUnicode field to given value.

### HasResidentCityUnicode

`func (o *AppleWalletDigitalIdCredential) HasResidentCityUnicode() bool`

HasResidentCityUnicode returns a boolean if a field has been set.

### SetResidentCityUnicodeNil

`func (o *AppleWalletDigitalIdCredential) SetResidentCityUnicodeNil(b bool)`

 SetResidentCityUnicodeNil sets the value for ResidentCityUnicode to be an explicit nil

### UnsetResidentCityUnicode
`func (o *AppleWalletDigitalIdCredential) UnsetResidentCityUnicode()`

UnsetResidentCityUnicode ensures that no value is present for ResidentCityUnicode, not even an explicit nil
### GetResidentCityLatin1

`func (o *AppleWalletDigitalIdCredential) GetResidentCityLatin1() string`

GetResidentCityLatin1 returns the ResidentCityLatin1 field if non-nil, zero value otherwise.

### GetResidentCityLatin1Ok

`func (o *AppleWalletDigitalIdCredential) GetResidentCityLatin1Ok() (*string, bool)`

GetResidentCityLatin1Ok returns a tuple with the ResidentCityLatin1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCityLatin1

`func (o *AppleWalletDigitalIdCredential) SetResidentCityLatin1(v string)`

SetResidentCityLatin1 sets ResidentCityLatin1 field to given value.

### HasResidentCityLatin1

`func (o *AppleWalletDigitalIdCredential) HasResidentCityLatin1() bool`

HasResidentCityLatin1 returns a boolean if a field has been set.

### SetResidentCityLatin1Nil

`func (o *AppleWalletDigitalIdCredential) SetResidentCityLatin1Nil(b bool)`

 SetResidentCityLatin1Nil sets the value for ResidentCityLatin1 to be an explicit nil

### UnsetResidentCityLatin1
`func (o *AppleWalletDigitalIdCredential) UnsetResidentCityLatin1()`

UnsetResidentCityLatin1 ensures that no value is present for ResidentCityLatin1, not even an explicit nil
### GetResidentPostalCode

`func (o *AppleWalletDigitalIdCredential) GetResidentPostalCode() string`

GetResidentPostalCode returns the ResidentPostalCode field if non-nil, zero value otherwise.

### GetResidentPostalCodeOk

`func (o *AppleWalletDigitalIdCredential) GetResidentPostalCodeOk() (*string, bool)`

GetResidentPostalCodeOk returns a tuple with the ResidentPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentPostalCode

`func (o *AppleWalletDigitalIdCredential) SetResidentPostalCode(v string)`

SetResidentPostalCode sets ResidentPostalCode field to given value.

### HasResidentPostalCode

`func (o *AppleWalletDigitalIdCredential) HasResidentPostalCode() bool`

HasResidentPostalCode returns a boolean if a field has been set.

### SetResidentPostalCodeNil

`func (o *AppleWalletDigitalIdCredential) SetResidentPostalCodeNil(b bool)`

 SetResidentPostalCodeNil sets the value for ResidentPostalCode to be an explicit nil

### UnsetResidentPostalCode
`func (o *AppleWalletDigitalIdCredential) UnsetResidentPostalCode()`

UnsetResidentPostalCode ensures that no value is present for ResidentPostalCode, not even an explicit nil
### GetResidentCountry

`func (o *AppleWalletDigitalIdCredential) GetResidentCountry() string`

GetResidentCountry returns the ResidentCountry field if non-nil, zero value otherwise.

### GetResidentCountryOk

`func (o *AppleWalletDigitalIdCredential) GetResidentCountryOk() (*string, bool)`

GetResidentCountryOk returns a tuple with the ResidentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountry

`func (o *AppleWalletDigitalIdCredential) SetResidentCountry(v string)`

SetResidentCountry sets ResidentCountry field to given value.

### HasResidentCountry

`func (o *AppleWalletDigitalIdCredential) HasResidentCountry() bool`

HasResidentCountry returns a boolean if a field has been set.

### SetResidentCountryNil

`func (o *AppleWalletDigitalIdCredential) SetResidentCountryNil(b bool)`

 SetResidentCountryNil sets the value for ResidentCountry to be an explicit nil

### UnsetResidentCountry
`func (o *AppleWalletDigitalIdCredential) UnsetResidentCountry()`

UnsetResidentCountry ensures that no value is present for ResidentCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



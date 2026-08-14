# MoldovaIdentityCardCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idnp** | Pointer to **NullableString** | The individual&#39;s IDNP (Numărul de Identificare Personal), the Moldovan state personal identification number.              This is a 13-digit number which uniquely identifies all natural persons in the Republic of Moldova. It has the format &#x60;2YYYOOOSSSSSX&#x60;, where: - &#x60;2&#x60; is the literal number &#x60;2&#x60; to indicate that the identifier belongs to a natural person - &#x60;YYY&#x60; is the last 3 digits of the year in which the identifier was issued - &#x60;OOO&#x60; is the code of the registrar office which issued the identifier - &#x60;SSSSS&#x60; is the sequential birth number for that year - &#x60;X&#x60; is a check digit              Note that the year encoded in the identifier is not necessarily the same as the year of birth of the individual. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given name(s), as recorded on their Moldovan identity card. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family name (surname), as recorded on their Moldovan identity card. | [optional] 
**Sex** | Pointer to **NullableInt32** | The individual&#39;s sex, as an ISO/IEC 5218 numeric code.              Possible values: - 0: Not known - 1: Male - 2: Female - 9: Not applicable | [optional] 
**Nationality** | Pointer to **NullableString** | The individual&#39;s nationality, as an ISO 3166-1 alpha-2 country code.              Moldovan nationals are &#x60;MD&#x60;. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**AgeOver18** | Pointer to **NullableBool** | Whether the individual is at least 18 years of age at the time of issuance of the digital credential. | [optional] 
**AgeOver21** | Pointer to **NullableBool** | Whether the individual is at least 21 years of age at the time of issuance of the digital credential. | [optional] 
**ResidentAddress** | Pointer to **NullableString** | The individual&#39;s full registered residential address as a single formatted string. | [optional] 
**ResidentCountry** | Pointer to **NullableString** | The country of the individual&#39;s registered residence, as an ISO 3166-1 alpha-2 country code. | [optional] 
**ResidentRegion** | Pointer to **NullableString** | The region of the individual&#39;s registered residence. | [optional] 
**ResidentCity** | Pointer to **NullableString** | The locality (municipality, city, town, or village) of the individual&#39;s registered residence. | [optional] 
**ResidentStreet** | Pointer to **NullableString** | The street name of the individual&#39;s registered residence. | [optional] 
**ResidentHouseNumber** | Pointer to **NullableString** | The house or building number of the individual&#39;s registered residence. | [optional] 
**ResidentBlock** | Pointer to **NullableString** | The building block (bloc) of the individual&#39;s registered residence, where applicable. | [optional] 
**ResidentFlat** | Pointer to **NullableString** | The flat (apartment) number of the individual&#39;s registered residence, where applicable. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the physical identity card, used to create this digital credential, was issued. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date the physical identity card, used to create this digital credential, expires. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The name of the authority that issued the identity card. | [optional] 
**DocumentType** | Pointer to **NullableString** | The document type              TODO: Get example values | [optional] 
**DocumentSeries** | Pointer to **NullableString** | The series of the underlying physical identity card.              This is the prefix to the full Moldovan identity card number, which is in the format &#x60;SNNN...&#x60;, where:              - &#x60;S&#x60; is a one-letter document series - &#x60;NNN...&#x60; is a document number (typically 8 digits) | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The number of the underlying physical identity card, excluding the series.              This is the suffix to the full Moldovan identity card number, which is in the format &#x60;SNNN...&#x60;, where:              - &#x60;S&#x60; is a one-letter document series - &#x60;NNN...&#x60; is a document number (typically 8 digits) | [optional] 

## Methods

### NewMoldovaIdentityCardCredential

`func NewMoldovaIdentityCardCredential() *MoldovaIdentityCardCredential`

NewMoldovaIdentityCardCredential instantiates a new MoldovaIdentityCardCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoldovaIdentityCardCredentialWithDefaults

`func NewMoldovaIdentityCardCredentialWithDefaults() *MoldovaIdentityCardCredential`

NewMoldovaIdentityCardCredentialWithDefaults instantiates a new MoldovaIdentityCardCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdnp

`func (o *MoldovaIdentityCardCredential) GetIdnp() string`

GetIdnp returns the Idnp field if non-nil, zero value otherwise.

### GetIdnpOk

`func (o *MoldovaIdentityCardCredential) GetIdnpOk() (*string, bool)`

GetIdnpOk returns a tuple with the Idnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnp

`func (o *MoldovaIdentityCardCredential) SetIdnp(v string)`

SetIdnp sets Idnp field to given value.

### HasIdnp

`func (o *MoldovaIdentityCardCredential) HasIdnp() bool`

HasIdnp returns a boolean if a field has been set.

### SetIdnpNil

`func (o *MoldovaIdentityCardCredential) SetIdnpNil(b bool)`

 SetIdnpNil sets the value for Idnp to be an explicit nil

### UnsetIdnp
`func (o *MoldovaIdentityCardCredential) UnsetIdnp()`

UnsetIdnp ensures that no value is present for Idnp, not even an explicit nil
### GetGivenName

`func (o *MoldovaIdentityCardCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MoldovaIdentityCardCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MoldovaIdentityCardCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MoldovaIdentityCardCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MoldovaIdentityCardCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MoldovaIdentityCardCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *MoldovaIdentityCardCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MoldovaIdentityCardCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MoldovaIdentityCardCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MoldovaIdentityCardCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *MoldovaIdentityCardCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *MoldovaIdentityCardCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetSex

`func (o *MoldovaIdentityCardCredential) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *MoldovaIdentityCardCredential) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *MoldovaIdentityCardCredential) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *MoldovaIdentityCardCredential) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *MoldovaIdentityCardCredential) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *MoldovaIdentityCardCredential) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetNationality

`func (o *MoldovaIdentityCardCredential) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *MoldovaIdentityCardCredential) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *MoldovaIdentityCardCredential) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *MoldovaIdentityCardCredential) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *MoldovaIdentityCardCredential) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *MoldovaIdentityCardCredential) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDateOfBirth

`func (o *MoldovaIdentityCardCredential) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MoldovaIdentityCardCredential) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MoldovaIdentityCardCredential) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MoldovaIdentityCardCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MoldovaIdentityCardCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MoldovaIdentityCardCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetAgeOver18

`func (o *MoldovaIdentityCardCredential) GetAgeOver18() bool`

GetAgeOver18 returns the AgeOver18 field if non-nil, zero value otherwise.

### GetAgeOver18Ok

`func (o *MoldovaIdentityCardCredential) GetAgeOver18Ok() (*bool, bool)`

GetAgeOver18Ok returns a tuple with the AgeOver18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver18

`func (o *MoldovaIdentityCardCredential) SetAgeOver18(v bool)`

SetAgeOver18 sets AgeOver18 field to given value.

### HasAgeOver18

`func (o *MoldovaIdentityCardCredential) HasAgeOver18() bool`

HasAgeOver18 returns a boolean if a field has been set.

### SetAgeOver18Nil

`func (o *MoldovaIdentityCardCredential) SetAgeOver18Nil(b bool)`

 SetAgeOver18Nil sets the value for AgeOver18 to be an explicit nil

### UnsetAgeOver18
`func (o *MoldovaIdentityCardCredential) UnsetAgeOver18()`

UnsetAgeOver18 ensures that no value is present for AgeOver18, not even an explicit nil
### GetAgeOver21

`func (o *MoldovaIdentityCardCredential) GetAgeOver21() bool`

GetAgeOver21 returns the AgeOver21 field if non-nil, zero value otherwise.

### GetAgeOver21Ok

`func (o *MoldovaIdentityCardCredential) GetAgeOver21Ok() (*bool, bool)`

GetAgeOver21Ok returns a tuple with the AgeOver21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver21

`func (o *MoldovaIdentityCardCredential) SetAgeOver21(v bool)`

SetAgeOver21 sets AgeOver21 field to given value.

### HasAgeOver21

`func (o *MoldovaIdentityCardCredential) HasAgeOver21() bool`

HasAgeOver21 returns a boolean if a field has been set.

### SetAgeOver21Nil

`func (o *MoldovaIdentityCardCredential) SetAgeOver21Nil(b bool)`

 SetAgeOver21Nil sets the value for AgeOver21 to be an explicit nil

### UnsetAgeOver21
`func (o *MoldovaIdentityCardCredential) UnsetAgeOver21()`

UnsetAgeOver21 ensures that no value is present for AgeOver21, not even an explicit nil
### GetResidentAddress

`func (o *MoldovaIdentityCardCredential) GetResidentAddress() string`

GetResidentAddress returns the ResidentAddress field if non-nil, zero value otherwise.

### GetResidentAddressOk

`func (o *MoldovaIdentityCardCredential) GetResidentAddressOk() (*string, bool)`

GetResidentAddressOk returns a tuple with the ResidentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentAddress

`func (o *MoldovaIdentityCardCredential) SetResidentAddress(v string)`

SetResidentAddress sets ResidentAddress field to given value.

### HasResidentAddress

`func (o *MoldovaIdentityCardCredential) HasResidentAddress() bool`

HasResidentAddress returns a boolean if a field has been set.

### SetResidentAddressNil

`func (o *MoldovaIdentityCardCredential) SetResidentAddressNil(b bool)`

 SetResidentAddressNil sets the value for ResidentAddress to be an explicit nil

### UnsetResidentAddress
`func (o *MoldovaIdentityCardCredential) UnsetResidentAddress()`

UnsetResidentAddress ensures that no value is present for ResidentAddress, not even an explicit nil
### GetResidentCountry

`func (o *MoldovaIdentityCardCredential) GetResidentCountry() string`

GetResidentCountry returns the ResidentCountry field if non-nil, zero value otherwise.

### GetResidentCountryOk

`func (o *MoldovaIdentityCardCredential) GetResidentCountryOk() (*string, bool)`

GetResidentCountryOk returns a tuple with the ResidentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCountry

`func (o *MoldovaIdentityCardCredential) SetResidentCountry(v string)`

SetResidentCountry sets ResidentCountry field to given value.

### HasResidentCountry

`func (o *MoldovaIdentityCardCredential) HasResidentCountry() bool`

HasResidentCountry returns a boolean if a field has been set.

### SetResidentCountryNil

`func (o *MoldovaIdentityCardCredential) SetResidentCountryNil(b bool)`

 SetResidentCountryNil sets the value for ResidentCountry to be an explicit nil

### UnsetResidentCountry
`func (o *MoldovaIdentityCardCredential) UnsetResidentCountry()`

UnsetResidentCountry ensures that no value is present for ResidentCountry, not even an explicit nil
### GetResidentRegion

`func (o *MoldovaIdentityCardCredential) GetResidentRegion() string`

GetResidentRegion returns the ResidentRegion field if non-nil, zero value otherwise.

### GetResidentRegionOk

`func (o *MoldovaIdentityCardCredential) GetResidentRegionOk() (*string, bool)`

GetResidentRegionOk returns a tuple with the ResidentRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentRegion

`func (o *MoldovaIdentityCardCredential) SetResidentRegion(v string)`

SetResidentRegion sets ResidentRegion field to given value.

### HasResidentRegion

`func (o *MoldovaIdentityCardCredential) HasResidentRegion() bool`

HasResidentRegion returns a boolean if a field has been set.

### SetResidentRegionNil

`func (o *MoldovaIdentityCardCredential) SetResidentRegionNil(b bool)`

 SetResidentRegionNil sets the value for ResidentRegion to be an explicit nil

### UnsetResidentRegion
`func (o *MoldovaIdentityCardCredential) UnsetResidentRegion()`

UnsetResidentRegion ensures that no value is present for ResidentRegion, not even an explicit nil
### GetResidentCity

`func (o *MoldovaIdentityCardCredential) GetResidentCity() string`

GetResidentCity returns the ResidentCity field if non-nil, zero value otherwise.

### GetResidentCityOk

`func (o *MoldovaIdentityCardCredential) GetResidentCityOk() (*string, bool)`

GetResidentCityOk returns a tuple with the ResidentCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentCity

`func (o *MoldovaIdentityCardCredential) SetResidentCity(v string)`

SetResidentCity sets ResidentCity field to given value.

### HasResidentCity

`func (o *MoldovaIdentityCardCredential) HasResidentCity() bool`

HasResidentCity returns a boolean if a field has been set.

### SetResidentCityNil

`func (o *MoldovaIdentityCardCredential) SetResidentCityNil(b bool)`

 SetResidentCityNil sets the value for ResidentCity to be an explicit nil

### UnsetResidentCity
`func (o *MoldovaIdentityCardCredential) UnsetResidentCity()`

UnsetResidentCity ensures that no value is present for ResidentCity, not even an explicit nil
### GetResidentStreet

`func (o *MoldovaIdentityCardCredential) GetResidentStreet() string`

GetResidentStreet returns the ResidentStreet field if non-nil, zero value otherwise.

### GetResidentStreetOk

`func (o *MoldovaIdentityCardCredential) GetResidentStreetOk() (*string, bool)`

GetResidentStreetOk returns a tuple with the ResidentStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentStreet

`func (o *MoldovaIdentityCardCredential) SetResidentStreet(v string)`

SetResidentStreet sets ResidentStreet field to given value.

### HasResidentStreet

`func (o *MoldovaIdentityCardCredential) HasResidentStreet() bool`

HasResidentStreet returns a boolean if a field has been set.

### SetResidentStreetNil

`func (o *MoldovaIdentityCardCredential) SetResidentStreetNil(b bool)`

 SetResidentStreetNil sets the value for ResidentStreet to be an explicit nil

### UnsetResidentStreet
`func (o *MoldovaIdentityCardCredential) UnsetResidentStreet()`

UnsetResidentStreet ensures that no value is present for ResidentStreet, not even an explicit nil
### GetResidentHouseNumber

`func (o *MoldovaIdentityCardCredential) GetResidentHouseNumber() string`

GetResidentHouseNumber returns the ResidentHouseNumber field if non-nil, zero value otherwise.

### GetResidentHouseNumberOk

`func (o *MoldovaIdentityCardCredential) GetResidentHouseNumberOk() (*string, bool)`

GetResidentHouseNumberOk returns a tuple with the ResidentHouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentHouseNumber

`func (o *MoldovaIdentityCardCredential) SetResidentHouseNumber(v string)`

SetResidentHouseNumber sets ResidentHouseNumber field to given value.

### HasResidentHouseNumber

`func (o *MoldovaIdentityCardCredential) HasResidentHouseNumber() bool`

HasResidentHouseNumber returns a boolean if a field has been set.

### SetResidentHouseNumberNil

`func (o *MoldovaIdentityCardCredential) SetResidentHouseNumberNil(b bool)`

 SetResidentHouseNumberNil sets the value for ResidentHouseNumber to be an explicit nil

### UnsetResidentHouseNumber
`func (o *MoldovaIdentityCardCredential) UnsetResidentHouseNumber()`

UnsetResidentHouseNumber ensures that no value is present for ResidentHouseNumber, not even an explicit nil
### GetResidentBlock

`func (o *MoldovaIdentityCardCredential) GetResidentBlock() string`

GetResidentBlock returns the ResidentBlock field if non-nil, zero value otherwise.

### GetResidentBlockOk

`func (o *MoldovaIdentityCardCredential) GetResidentBlockOk() (*string, bool)`

GetResidentBlockOk returns a tuple with the ResidentBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentBlock

`func (o *MoldovaIdentityCardCredential) SetResidentBlock(v string)`

SetResidentBlock sets ResidentBlock field to given value.

### HasResidentBlock

`func (o *MoldovaIdentityCardCredential) HasResidentBlock() bool`

HasResidentBlock returns a boolean if a field has been set.

### SetResidentBlockNil

`func (o *MoldovaIdentityCardCredential) SetResidentBlockNil(b bool)`

 SetResidentBlockNil sets the value for ResidentBlock to be an explicit nil

### UnsetResidentBlock
`func (o *MoldovaIdentityCardCredential) UnsetResidentBlock()`

UnsetResidentBlock ensures that no value is present for ResidentBlock, not even an explicit nil
### GetResidentFlat

`func (o *MoldovaIdentityCardCredential) GetResidentFlat() string`

GetResidentFlat returns the ResidentFlat field if non-nil, zero value otherwise.

### GetResidentFlatOk

`func (o *MoldovaIdentityCardCredential) GetResidentFlatOk() (*string, bool)`

GetResidentFlatOk returns a tuple with the ResidentFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentFlat

`func (o *MoldovaIdentityCardCredential) SetResidentFlat(v string)`

SetResidentFlat sets ResidentFlat field to given value.

### HasResidentFlat

`func (o *MoldovaIdentityCardCredential) HasResidentFlat() bool`

HasResidentFlat returns a boolean if a field has been set.

### SetResidentFlatNil

`func (o *MoldovaIdentityCardCredential) SetResidentFlatNil(b bool)`

 SetResidentFlatNil sets the value for ResidentFlat to be an explicit nil

### UnsetResidentFlat
`func (o *MoldovaIdentityCardCredential) UnsetResidentFlat()`

UnsetResidentFlat ensures that no value is present for ResidentFlat, not even an explicit nil
### GetIssueDate

`func (o *MoldovaIdentityCardCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *MoldovaIdentityCardCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *MoldovaIdentityCardCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *MoldovaIdentityCardCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *MoldovaIdentityCardCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *MoldovaIdentityCardCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *MoldovaIdentityCardCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MoldovaIdentityCardCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MoldovaIdentityCardCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MoldovaIdentityCardCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *MoldovaIdentityCardCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *MoldovaIdentityCardCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingAuthority

`func (o *MoldovaIdentityCardCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *MoldovaIdentityCardCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *MoldovaIdentityCardCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *MoldovaIdentityCardCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *MoldovaIdentityCardCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *MoldovaIdentityCardCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetDocumentType

`func (o *MoldovaIdentityCardCredential) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *MoldovaIdentityCardCredential) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *MoldovaIdentityCardCredential) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *MoldovaIdentityCardCredential) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *MoldovaIdentityCardCredential) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *MoldovaIdentityCardCredential) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetDocumentSeries

`func (o *MoldovaIdentityCardCredential) GetDocumentSeries() string`

GetDocumentSeries returns the DocumentSeries field if non-nil, zero value otherwise.

### GetDocumentSeriesOk

`func (o *MoldovaIdentityCardCredential) GetDocumentSeriesOk() (*string, bool)`

GetDocumentSeriesOk returns a tuple with the DocumentSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSeries

`func (o *MoldovaIdentityCardCredential) SetDocumentSeries(v string)`

SetDocumentSeries sets DocumentSeries field to given value.

### HasDocumentSeries

`func (o *MoldovaIdentityCardCredential) HasDocumentSeries() bool`

HasDocumentSeries returns a boolean if a field has been set.

### SetDocumentSeriesNil

`func (o *MoldovaIdentityCardCredential) SetDocumentSeriesNil(b bool)`

 SetDocumentSeriesNil sets the value for DocumentSeries to be an explicit nil

### UnsetDocumentSeries
`func (o *MoldovaIdentityCardCredential) UnsetDocumentSeries()`

UnsetDocumentSeries ensures that no value is present for DocumentSeries, not even an explicit nil
### GetDocumentNumber

`func (o *MoldovaIdentityCardCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *MoldovaIdentityCardCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *MoldovaIdentityCardCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *MoldovaIdentityCardCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *MoldovaIdentityCardCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *MoldovaIdentityCardCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



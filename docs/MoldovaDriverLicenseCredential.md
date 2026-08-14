# MoldovaDriverLicenseCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idnp** | Pointer to **NullableString** | The individual&#39;s IDNP (Numărul de Identificare Personal), the Moldovan state personal identification number.              This is a 13-digit number which uniquely identifies all natural persons in the Republic of Moldova. It has the format &#x60;2YYYOOOSSSSSX&#x60;, where: - &#x60;2&#x60; is the literal number &#x60;2&#x60; to indicate that the identifier belongs to a natural person - &#x60;YYY&#x60; is the last 3 digits of the year in which the identifier was issued - &#x60;OOO&#x60; is the code of the registrar office which issued the identifier - &#x60;SSSSS&#x60; is the sequential birth number for that year - &#x60;X&#x60; is a check digit              Note that the year encoded in the identifier is not necessarily the same as the year of birth of the individual. | [optional] 
**GivenName** | Pointer to **NullableString** | The individual&#39;s given name(s), as recorded on their Moldovan driver license. | [optional] 
**FamilyName** | Pointer to **NullableString** | The individual&#39;s family name (surname), as recorded on their Moldovan driver license. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**AgeOver18** | Pointer to **NullableBool** | Whether the individual is at least 18 years of age at the time of issuance of the digital credential. | [optional] 
**AgeOver21** | Pointer to **NullableBool** | Whether the individual is at least 21 years of age at the time of issuance of the digital credential. | [optional] 
**BirthCountry** | Pointer to **NullableString** | The country where the individual was born, as an ISO 3166-1 alpha-2 country code. | [optional] 
**BirthCity** | Pointer to **NullableString** | The locality (municipality, city, town, or village) where the individual was born. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the driver license was issued. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date the driver license expires. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | The name of the authority that issued the driver license. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The driver license number. | [optional] 
**DrivingPrivileges** | Pointer to [**[]Iso180135DrivingPrivilege**](Iso180135DrivingPrivilege.md) | The categories of vehicle the individual is licensed to drive, parsed per ISO 18013-5.              Each entry has a vehicle category code with its own issue / expiry dates and possible restrictions. | [optional] 

## Methods

### NewMoldovaDriverLicenseCredential

`func NewMoldovaDriverLicenseCredential() *MoldovaDriverLicenseCredential`

NewMoldovaDriverLicenseCredential instantiates a new MoldovaDriverLicenseCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoldovaDriverLicenseCredentialWithDefaults

`func NewMoldovaDriverLicenseCredentialWithDefaults() *MoldovaDriverLicenseCredential`

NewMoldovaDriverLicenseCredentialWithDefaults instantiates a new MoldovaDriverLicenseCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdnp

`func (o *MoldovaDriverLicenseCredential) GetIdnp() string`

GetIdnp returns the Idnp field if non-nil, zero value otherwise.

### GetIdnpOk

`func (o *MoldovaDriverLicenseCredential) GetIdnpOk() (*string, bool)`

GetIdnpOk returns a tuple with the Idnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnp

`func (o *MoldovaDriverLicenseCredential) SetIdnp(v string)`

SetIdnp sets Idnp field to given value.

### HasIdnp

`func (o *MoldovaDriverLicenseCredential) HasIdnp() bool`

HasIdnp returns a boolean if a field has been set.

### SetIdnpNil

`func (o *MoldovaDriverLicenseCredential) SetIdnpNil(b bool)`

 SetIdnpNil sets the value for Idnp to be an explicit nil

### UnsetIdnp
`func (o *MoldovaDriverLicenseCredential) UnsetIdnp()`

UnsetIdnp ensures that no value is present for Idnp, not even an explicit nil
### GetGivenName

`func (o *MoldovaDriverLicenseCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MoldovaDriverLicenseCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MoldovaDriverLicenseCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MoldovaDriverLicenseCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MoldovaDriverLicenseCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MoldovaDriverLicenseCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *MoldovaDriverLicenseCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MoldovaDriverLicenseCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MoldovaDriverLicenseCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MoldovaDriverLicenseCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *MoldovaDriverLicenseCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *MoldovaDriverLicenseCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *MoldovaDriverLicenseCredential) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MoldovaDriverLicenseCredential) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MoldovaDriverLicenseCredential) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MoldovaDriverLicenseCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MoldovaDriverLicenseCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MoldovaDriverLicenseCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetAgeOver18

`func (o *MoldovaDriverLicenseCredential) GetAgeOver18() bool`

GetAgeOver18 returns the AgeOver18 field if non-nil, zero value otherwise.

### GetAgeOver18Ok

`func (o *MoldovaDriverLicenseCredential) GetAgeOver18Ok() (*bool, bool)`

GetAgeOver18Ok returns a tuple with the AgeOver18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver18

`func (o *MoldovaDriverLicenseCredential) SetAgeOver18(v bool)`

SetAgeOver18 sets AgeOver18 field to given value.

### HasAgeOver18

`func (o *MoldovaDriverLicenseCredential) HasAgeOver18() bool`

HasAgeOver18 returns a boolean if a field has been set.

### SetAgeOver18Nil

`func (o *MoldovaDriverLicenseCredential) SetAgeOver18Nil(b bool)`

 SetAgeOver18Nil sets the value for AgeOver18 to be an explicit nil

### UnsetAgeOver18
`func (o *MoldovaDriverLicenseCredential) UnsetAgeOver18()`

UnsetAgeOver18 ensures that no value is present for AgeOver18, not even an explicit nil
### GetAgeOver21

`func (o *MoldovaDriverLicenseCredential) GetAgeOver21() bool`

GetAgeOver21 returns the AgeOver21 field if non-nil, zero value otherwise.

### GetAgeOver21Ok

`func (o *MoldovaDriverLicenseCredential) GetAgeOver21Ok() (*bool, bool)`

GetAgeOver21Ok returns a tuple with the AgeOver21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver21

`func (o *MoldovaDriverLicenseCredential) SetAgeOver21(v bool)`

SetAgeOver21 sets AgeOver21 field to given value.

### HasAgeOver21

`func (o *MoldovaDriverLicenseCredential) HasAgeOver21() bool`

HasAgeOver21 returns a boolean if a field has been set.

### SetAgeOver21Nil

`func (o *MoldovaDriverLicenseCredential) SetAgeOver21Nil(b bool)`

 SetAgeOver21Nil sets the value for AgeOver21 to be an explicit nil

### UnsetAgeOver21
`func (o *MoldovaDriverLicenseCredential) UnsetAgeOver21()`

UnsetAgeOver21 ensures that no value is present for AgeOver21, not even an explicit nil
### GetBirthCountry

`func (o *MoldovaDriverLicenseCredential) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *MoldovaDriverLicenseCredential) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *MoldovaDriverLicenseCredential) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *MoldovaDriverLicenseCredential) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *MoldovaDriverLicenseCredential) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *MoldovaDriverLicenseCredential) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetBirthCity

`func (o *MoldovaDriverLicenseCredential) GetBirthCity() string`

GetBirthCity returns the BirthCity field if non-nil, zero value otherwise.

### GetBirthCityOk

`func (o *MoldovaDriverLicenseCredential) GetBirthCityOk() (*string, bool)`

GetBirthCityOk returns a tuple with the BirthCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCity

`func (o *MoldovaDriverLicenseCredential) SetBirthCity(v string)`

SetBirthCity sets BirthCity field to given value.

### HasBirthCity

`func (o *MoldovaDriverLicenseCredential) HasBirthCity() bool`

HasBirthCity returns a boolean if a field has been set.

### SetBirthCityNil

`func (o *MoldovaDriverLicenseCredential) SetBirthCityNil(b bool)`

 SetBirthCityNil sets the value for BirthCity to be an explicit nil

### UnsetBirthCity
`func (o *MoldovaDriverLicenseCredential) UnsetBirthCity()`

UnsetBirthCity ensures that no value is present for BirthCity, not even an explicit nil
### GetIssueDate

`func (o *MoldovaDriverLicenseCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *MoldovaDriverLicenseCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *MoldovaDriverLicenseCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *MoldovaDriverLicenseCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *MoldovaDriverLicenseCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *MoldovaDriverLicenseCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *MoldovaDriverLicenseCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MoldovaDriverLicenseCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MoldovaDriverLicenseCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MoldovaDriverLicenseCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *MoldovaDriverLicenseCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *MoldovaDriverLicenseCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingAuthority

`func (o *MoldovaDriverLicenseCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *MoldovaDriverLicenseCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *MoldovaDriverLicenseCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *MoldovaDriverLicenseCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *MoldovaDriverLicenseCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *MoldovaDriverLicenseCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetDocumentNumber

`func (o *MoldovaDriverLicenseCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *MoldovaDriverLicenseCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *MoldovaDriverLicenseCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *MoldovaDriverLicenseCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *MoldovaDriverLicenseCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *MoldovaDriverLicenseCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetDrivingPrivileges

`func (o *MoldovaDriverLicenseCredential) GetDrivingPrivileges() []Iso180135DrivingPrivilege`

GetDrivingPrivileges returns the DrivingPrivileges field if non-nil, zero value otherwise.

### GetDrivingPrivilegesOk

`func (o *MoldovaDriverLicenseCredential) GetDrivingPrivilegesOk() (*[]Iso180135DrivingPrivilege, bool)`

GetDrivingPrivilegesOk returns a tuple with the DrivingPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingPrivileges

`func (o *MoldovaDriverLicenseCredential) SetDrivingPrivileges(v []Iso180135DrivingPrivilege)`

SetDrivingPrivileges sets DrivingPrivileges field to given value.

### HasDrivingPrivileges

`func (o *MoldovaDriverLicenseCredential) HasDrivingPrivileges() bool`

HasDrivingPrivileges returns a boolean if a field has been set.

### SetDrivingPrivilegesNil

`func (o *MoldovaDriverLicenseCredential) SetDrivingPrivilegesNil(b bool)`

 SetDrivingPrivilegesNil sets the value for DrivingPrivileges to be an explicit nil

### UnsetDrivingPrivileges
`func (o *MoldovaDriverLicenseCredential) UnsetDrivingPrivileges()`

UnsetDrivingPrivileges ensures that no value is present for DrivingPrivileges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



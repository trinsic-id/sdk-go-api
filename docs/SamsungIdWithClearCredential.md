# SamsungIdWithClearCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | First name(s), other name(s), or secondary identifier of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | Last name, surname, or primary identifier of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date the Samsung ID was issued.              This is not the issue date of the passport used to create the Samsung ID; see &#x60;originalDocumentIssueDate&#x60;. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date the Samsung ID expires.              This is not necessarily the expiry date of the passport used to create the Samsung ID; see &#x60;originalDocumentExpiryDate&#x60;. | [optional] 
**OriginalDocumentIssueDate** | Pointer to **NullableString** | Issue date of the underlying passport used to create the Samsung ID. | [optional] 
**OriginalDocumentExpiryDate** | Pointer to **NullableString** | Expiry date of the underlying passport used to create the Samsung ID. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The document number of the underlying passport used to create the Samsung ID. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | Identifier of Samsung as the issuing authority of the Samsung ID credential.              This always begins with \&quot;XS-\&quot;. | [optional] 
**Sex** | Pointer to **NullableInt32** | The individual&#39;s sex as an ISO/IEC 5218 code.              Possible values: - 0: Unknown - 1: Male - 2: Female - 9: Not Applicable | [optional] 
**DhsCompliance** | Pointer to **NullableBool** | Whether the credential is compliant with REAL ID.              This is always &#x60;true&#x60; for Samsung ID with CLEAR credentials. | [optional] 
**AgeInYears** | Pointer to **NullableInt32** | Age of the individual in years. | [optional] 
**AgeBirthYear** | Pointer to **NullableInt32** | Year of birth of the individual. | [optional] 
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the credential. | [optional] 

## Methods

### NewSamsungIdWithClearCredential

`func NewSamsungIdWithClearCredential() *SamsungIdWithClearCredential`

NewSamsungIdWithClearCredential instantiates a new SamsungIdWithClearCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamsungIdWithClearCredentialWithDefaults

`func NewSamsungIdWithClearCredentialWithDefaults() *SamsungIdWithClearCredential`

NewSamsungIdWithClearCredentialWithDefaults instantiates a new SamsungIdWithClearCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *SamsungIdWithClearCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *SamsungIdWithClearCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *SamsungIdWithClearCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *SamsungIdWithClearCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *SamsungIdWithClearCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *SamsungIdWithClearCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *SamsungIdWithClearCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *SamsungIdWithClearCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *SamsungIdWithClearCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *SamsungIdWithClearCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *SamsungIdWithClearCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *SamsungIdWithClearCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *SamsungIdWithClearCredential) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SamsungIdWithClearCredential) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SamsungIdWithClearCredential) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SamsungIdWithClearCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SamsungIdWithClearCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SamsungIdWithClearCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetIssueDate

`func (o *SamsungIdWithClearCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *SamsungIdWithClearCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *SamsungIdWithClearCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *SamsungIdWithClearCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *SamsungIdWithClearCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *SamsungIdWithClearCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *SamsungIdWithClearCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *SamsungIdWithClearCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *SamsungIdWithClearCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *SamsungIdWithClearCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *SamsungIdWithClearCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *SamsungIdWithClearCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetOriginalDocumentIssueDate

`func (o *SamsungIdWithClearCredential) GetOriginalDocumentIssueDate() string`

GetOriginalDocumentIssueDate returns the OriginalDocumentIssueDate field if non-nil, zero value otherwise.

### GetOriginalDocumentIssueDateOk

`func (o *SamsungIdWithClearCredential) GetOriginalDocumentIssueDateOk() (*string, bool)`

GetOriginalDocumentIssueDateOk returns a tuple with the OriginalDocumentIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDocumentIssueDate

`func (o *SamsungIdWithClearCredential) SetOriginalDocumentIssueDate(v string)`

SetOriginalDocumentIssueDate sets OriginalDocumentIssueDate field to given value.

### HasOriginalDocumentIssueDate

`func (o *SamsungIdWithClearCredential) HasOriginalDocumentIssueDate() bool`

HasOriginalDocumentIssueDate returns a boolean if a field has been set.

### SetOriginalDocumentIssueDateNil

`func (o *SamsungIdWithClearCredential) SetOriginalDocumentIssueDateNil(b bool)`

 SetOriginalDocumentIssueDateNil sets the value for OriginalDocumentIssueDate to be an explicit nil

### UnsetOriginalDocumentIssueDate
`func (o *SamsungIdWithClearCredential) UnsetOriginalDocumentIssueDate()`

UnsetOriginalDocumentIssueDate ensures that no value is present for OriginalDocumentIssueDate, not even an explicit nil
### GetOriginalDocumentExpiryDate

`func (o *SamsungIdWithClearCredential) GetOriginalDocumentExpiryDate() string`

GetOriginalDocumentExpiryDate returns the OriginalDocumentExpiryDate field if non-nil, zero value otherwise.

### GetOriginalDocumentExpiryDateOk

`func (o *SamsungIdWithClearCredential) GetOriginalDocumentExpiryDateOk() (*string, bool)`

GetOriginalDocumentExpiryDateOk returns a tuple with the OriginalDocumentExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDocumentExpiryDate

`func (o *SamsungIdWithClearCredential) SetOriginalDocumentExpiryDate(v string)`

SetOriginalDocumentExpiryDate sets OriginalDocumentExpiryDate field to given value.

### HasOriginalDocumentExpiryDate

`func (o *SamsungIdWithClearCredential) HasOriginalDocumentExpiryDate() bool`

HasOriginalDocumentExpiryDate returns a boolean if a field has been set.

### SetOriginalDocumentExpiryDateNil

`func (o *SamsungIdWithClearCredential) SetOriginalDocumentExpiryDateNil(b bool)`

 SetOriginalDocumentExpiryDateNil sets the value for OriginalDocumentExpiryDate to be an explicit nil

### UnsetOriginalDocumentExpiryDate
`func (o *SamsungIdWithClearCredential) UnsetOriginalDocumentExpiryDate()`

UnsetOriginalDocumentExpiryDate ensures that no value is present for OriginalDocumentExpiryDate, not even an explicit nil
### GetDocumentNumber

`func (o *SamsungIdWithClearCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SamsungIdWithClearCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SamsungIdWithClearCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SamsungIdWithClearCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SamsungIdWithClearCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SamsungIdWithClearCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetIssuingAuthority

`func (o *SamsungIdWithClearCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *SamsungIdWithClearCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *SamsungIdWithClearCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *SamsungIdWithClearCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *SamsungIdWithClearCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *SamsungIdWithClearCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetSex

`func (o *SamsungIdWithClearCredential) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SamsungIdWithClearCredential) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SamsungIdWithClearCredential) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SamsungIdWithClearCredential) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SamsungIdWithClearCredential) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SamsungIdWithClearCredential) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDhsCompliance

`func (o *SamsungIdWithClearCredential) GetDhsCompliance() bool`

GetDhsCompliance returns the DhsCompliance field if non-nil, zero value otherwise.

### GetDhsComplianceOk

`func (o *SamsungIdWithClearCredential) GetDhsComplianceOk() (*bool, bool)`

GetDhsComplianceOk returns a tuple with the DhsCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhsCompliance

`func (o *SamsungIdWithClearCredential) SetDhsCompliance(v bool)`

SetDhsCompliance sets DhsCompliance field to given value.

### HasDhsCompliance

`func (o *SamsungIdWithClearCredential) HasDhsCompliance() bool`

HasDhsCompliance returns a boolean if a field has been set.

### SetDhsComplianceNil

`func (o *SamsungIdWithClearCredential) SetDhsComplianceNil(b bool)`

 SetDhsComplianceNil sets the value for DhsCompliance to be an explicit nil

### UnsetDhsCompliance
`func (o *SamsungIdWithClearCredential) UnsetDhsCompliance()`

UnsetDhsCompliance ensures that no value is present for DhsCompliance, not even an explicit nil
### GetAgeInYears

`func (o *SamsungIdWithClearCredential) GetAgeInYears() int32`

GetAgeInYears returns the AgeInYears field if non-nil, zero value otherwise.

### GetAgeInYearsOk

`func (o *SamsungIdWithClearCredential) GetAgeInYearsOk() (*int32, bool)`

GetAgeInYearsOk returns a tuple with the AgeInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeInYears

`func (o *SamsungIdWithClearCredential) SetAgeInYears(v int32)`

SetAgeInYears sets AgeInYears field to given value.

### HasAgeInYears

`func (o *SamsungIdWithClearCredential) HasAgeInYears() bool`

HasAgeInYears returns a boolean if a field has been set.

### SetAgeInYearsNil

`func (o *SamsungIdWithClearCredential) SetAgeInYearsNil(b bool)`

 SetAgeInYearsNil sets the value for AgeInYears to be an explicit nil

### UnsetAgeInYears
`func (o *SamsungIdWithClearCredential) UnsetAgeInYears()`

UnsetAgeInYears ensures that no value is present for AgeInYears, not even an explicit nil
### GetAgeBirthYear

`func (o *SamsungIdWithClearCredential) GetAgeBirthYear() int32`

GetAgeBirthYear returns the AgeBirthYear field if non-nil, zero value otherwise.

### GetAgeBirthYearOk

`func (o *SamsungIdWithClearCredential) GetAgeBirthYearOk() (*int32, bool)`

GetAgeBirthYearOk returns a tuple with the AgeBirthYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeBirthYear

`func (o *SamsungIdWithClearCredential) SetAgeBirthYear(v int32)`

SetAgeBirthYear sets AgeBirthYear field to given value.

### HasAgeBirthYear

`func (o *SamsungIdWithClearCredential) HasAgeBirthYear() bool`

HasAgeBirthYear returns a boolean if a field has been set.

### SetAgeBirthYearNil

`func (o *SamsungIdWithClearCredential) SetAgeBirthYearNil(b bool)`

 SetAgeBirthYearNil sets the value for AgeBirthYear to be an explicit nil

### UnsetAgeBirthYear
`func (o *SamsungIdWithClearCredential) UnsetAgeBirthYear()`

UnsetAgeBirthYear ensures that no value is present for AgeBirthYear, not even an explicit nil
### GetAgeOver

`func (o *SamsungIdWithClearCredential) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *SamsungIdWithClearCredential) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *SamsungIdWithClearCredential) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *SamsungIdWithClearCredential) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *SamsungIdWithClearCredential) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *SamsungIdWithClearCredential) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



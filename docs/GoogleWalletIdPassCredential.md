# GoogleWalletIdPassCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | First name(s), other name(s), or secondary identifier of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | Last name, surname, or primary identifier of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**IssueDate** | Pointer to **NullableString** | The date when the ID Pass was issued.              This is not the same as the issue date of the underlying physical passport used to create the ID Pass. | [optional] 
**ExpiryDate** | Pointer to **NullableString** | The date when the ID Pass expires.              This is not necessarily the same as the expiration date of the underlying physical passport used to create the ID Pass. | [optional] 
**IssuingCountry** | Pointer to **NullableString** | Alpha-2 country code of the issuing authority&#39;s country or territory. | [optional] 
**IssuingAuthority** | Pointer to **NullableString** | Name or identifier of the mDL issuing authority.              This field&#39;s contents are arbitrary; it has no guaranteed format. | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality of the individual as an ISO 3166-1 alpha-2 country code. | [optional] 
**DocumentNumber** | Pointer to **NullableString** | The number of the underlying passport used to create the ID Pass. | [optional] 
**Sex** | Pointer to **NullableInt32** | The individual&#39;s sex as an ISO/IEC 5218 code.              Possible values: - 0: Unknown - 1: Male - 2: Female - 9: Not Applicable | [optional] 
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the ID Pass. | [optional] 
**IssueDateOfUnderlyingDocument** | Pointer to **NullableString** | Date when the underlying passport backing this digital credential was originally issued. | [optional] 

## Methods

### NewGoogleWalletIdPassCredential

`func NewGoogleWalletIdPassCredential() *GoogleWalletIdPassCredential`

NewGoogleWalletIdPassCredential instantiates a new GoogleWalletIdPassCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWalletIdPassCredentialWithDefaults

`func NewGoogleWalletIdPassCredentialWithDefaults() *GoogleWalletIdPassCredential`

NewGoogleWalletIdPassCredentialWithDefaults instantiates a new GoogleWalletIdPassCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *GoogleWalletIdPassCredential) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *GoogleWalletIdPassCredential) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *GoogleWalletIdPassCredential) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *GoogleWalletIdPassCredential) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *GoogleWalletIdPassCredential) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *GoogleWalletIdPassCredential) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *GoogleWalletIdPassCredential) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *GoogleWalletIdPassCredential) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *GoogleWalletIdPassCredential) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *GoogleWalletIdPassCredential) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *GoogleWalletIdPassCredential) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *GoogleWalletIdPassCredential) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *GoogleWalletIdPassCredential) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *GoogleWalletIdPassCredential) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *GoogleWalletIdPassCredential) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *GoogleWalletIdPassCredential) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *GoogleWalletIdPassCredential) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *GoogleWalletIdPassCredential) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetIssueDate

`func (o *GoogleWalletIdPassCredential) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *GoogleWalletIdPassCredential) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *GoogleWalletIdPassCredential) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *GoogleWalletIdPassCredential) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *GoogleWalletIdPassCredential) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *GoogleWalletIdPassCredential) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetExpiryDate

`func (o *GoogleWalletIdPassCredential) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GoogleWalletIdPassCredential) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GoogleWalletIdPassCredential) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GoogleWalletIdPassCredential) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *GoogleWalletIdPassCredential) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *GoogleWalletIdPassCredential) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetIssuingCountry

`func (o *GoogleWalletIdPassCredential) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *GoogleWalletIdPassCredential) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *GoogleWalletIdPassCredential) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *GoogleWalletIdPassCredential) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### SetIssuingCountryNil

`func (o *GoogleWalletIdPassCredential) SetIssuingCountryNil(b bool)`

 SetIssuingCountryNil sets the value for IssuingCountry to be an explicit nil

### UnsetIssuingCountry
`func (o *GoogleWalletIdPassCredential) UnsetIssuingCountry()`

UnsetIssuingCountry ensures that no value is present for IssuingCountry, not even an explicit nil
### GetIssuingAuthority

`func (o *GoogleWalletIdPassCredential) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *GoogleWalletIdPassCredential) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *GoogleWalletIdPassCredential) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *GoogleWalletIdPassCredential) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### SetIssuingAuthorityNil

`func (o *GoogleWalletIdPassCredential) SetIssuingAuthorityNil(b bool)`

 SetIssuingAuthorityNil sets the value for IssuingAuthority to be an explicit nil

### UnsetIssuingAuthority
`func (o *GoogleWalletIdPassCredential) UnsetIssuingAuthority()`

UnsetIssuingAuthority ensures that no value is present for IssuingAuthority, not even an explicit nil
### GetNationality

`func (o *GoogleWalletIdPassCredential) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *GoogleWalletIdPassCredential) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *GoogleWalletIdPassCredential) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *GoogleWalletIdPassCredential) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *GoogleWalletIdPassCredential) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *GoogleWalletIdPassCredential) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDocumentNumber

`func (o *GoogleWalletIdPassCredential) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GoogleWalletIdPassCredential) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GoogleWalletIdPassCredential) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GoogleWalletIdPassCredential) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *GoogleWalletIdPassCredential) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *GoogleWalletIdPassCredential) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetSex

`func (o *GoogleWalletIdPassCredential) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *GoogleWalletIdPassCredential) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *GoogleWalletIdPassCredential) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *GoogleWalletIdPassCredential) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *GoogleWalletIdPassCredential) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *GoogleWalletIdPassCredential) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetAgeOver

`func (o *GoogleWalletIdPassCredential) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *GoogleWalletIdPassCredential) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *GoogleWalletIdPassCredential) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *GoogleWalletIdPassCredential) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *GoogleWalletIdPassCredential) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *GoogleWalletIdPassCredential) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil
### GetIssueDateOfUnderlyingDocument

`func (o *GoogleWalletIdPassCredential) GetIssueDateOfUnderlyingDocument() string`

GetIssueDateOfUnderlyingDocument returns the IssueDateOfUnderlyingDocument field if non-nil, zero value otherwise.

### GetIssueDateOfUnderlyingDocumentOk

`func (o *GoogleWalletIdPassCredential) GetIssueDateOfUnderlyingDocumentOk() (*string, bool)`

GetIssueDateOfUnderlyingDocumentOk returns a tuple with the IssueDateOfUnderlyingDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDateOfUnderlyingDocument

`func (o *GoogleWalletIdPassCredential) SetIssueDateOfUnderlyingDocument(v string)`

SetIssueDateOfUnderlyingDocument sets IssueDateOfUnderlyingDocument field to given value.

### HasIssueDateOfUnderlyingDocument

`func (o *GoogleWalletIdPassCredential) HasIssueDateOfUnderlyingDocument() bool`

HasIssueDateOfUnderlyingDocument returns a boolean if a field has been set.

### SetIssueDateOfUnderlyingDocumentNil

`func (o *GoogleWalletIdPassCredential) SetIssueDateOfUnderlyingDocumentNil(b bool)`

 SetIssueDateOfUnderlyingDocumentNil sets the value for IssueDateOfUnderlyingDocument to be an explicit nil

### UnsetIssueDateOfUnderlyingDocument
`func (o *GoogleWalletIdPassCredential) UnsetIssueDateOfUnderlyingDocument()`

UnsetIssueDateOfUnderlyingDocument ensures that no value is present for IssueDateOfUnderlyingDocument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



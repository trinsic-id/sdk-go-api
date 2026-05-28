# SouthAfricaNidMatch2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The South African National Identity Number (13 digits).              Issued for life by the Department of Home Affairs (DHA) and stored in the HANIS (Home Affairs National Identification System) database. The same number is mandatory for banking, employment, taxation, and voting, and is printed on both the legacy green ID book and the Smart ID Card (rolled out from 2013 onward).              Format: - YYMMDD G(4) C A Z - YYMMDD is the date of birth - G(4) is the gender code (below 5000 female, 5000 or above male) - C is the citizenship indicator (0 citizen, 1 permanent resident) - A is reserved (it had a politically sensitive meaning in the past, but is currently   semantically meaningless) - Z is a Luhn check digit | [optional] 
**GivenName** | Pointer to **NullableString** | The user&#39;s first name as it appears in their National ID | [optional] 
**FamilyName** | Pointer to **NullableString** | The user&#39;s last name as it appears in their National ID | [optional] 
**MiddleName** | Pointer to **NullableString** | The user&#39;s middle name as it appears in their National ID (optional) | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | [optional] 
**Sex** | Pointer to [**NullableSouthAfricaNidMatch2InputSex**](SouthAfricaNidMatch2InputSex.md) | The user&#39;s sex as it appears in their National ID | [optional] 

## Methods

### NewSouthAfricaNidMatch2Input

`func NewSouthAfricaNidMatch2Input() *SouthAfricaNidMatch2Input`

NewSouthAfricaNidMatch2Input instantiates a new SouthAfricaNidMatch2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidMatch2InputWithDefaults

`func NewSouthAfricaNidMatch2InputWithDefaults() *SouthAfricaNidMatch2Input`

NewSouthAfricaNidMatch2InputWithDefaults instantiates a new SouthAfricaNidMatch2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *SouthAfricaNidMatch2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *SouthAfricaNidMatch2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *SouthAfricaNidMatch2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *SouthAfricaNidMatch2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *SouthAfricaNidMatch2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *SouthAfricaNidMatch2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetGivenName

`func (o *SouthAfricaNidMatch2Input) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *SouthAfricaNidMatch2Input) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *SouthAfricaNidMatch2Input) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *SouthAfricaNidMatch2Input) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *SouthAfricaNidMatch2Input) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *SouthAfricaNidMatch2Input) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *SouthAfricaNidMatch2Input) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *SouthAfricaNidMatch2Input) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *SouthAfricaNidMatch2Input) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *SouthAfricaNidMatch2Input) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *SouthAfricaNidMatch2Input) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *SouthAfricaNidMatch2Input) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetMiddleName

`func (o *SouthAfricaNidMatch2Input) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *SouthAfricaNidMatch2Input) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *SouthAfricaNidMatch2Input) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *SouthAfricaNidMatch2Input) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *SouthAfricaNidMatch2Input) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *SouthAfricaNidMatch2Input) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetDateOfBirth

`func (o *SouthAfricaNidMatch2Input) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SouthAfricaNidMatch2Input) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SouthAfricaNidMatch2Input) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SouthAfricaNidMatch2Input) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SouthAfricaNidMatch2Input) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SouthAfricaNidMatch2Input) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *SouthAfricaNidMatch2Input) GetSex() SouthAfricaNidMatch2InputSex`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SouthAfricaNidMatch2Input) GetSexOk() (*SouthAfricaNidMatch2InputSex, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SouthAfricaNidMatch2Input) SetSex(v SouthAfricaNidMatch2InputSex)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SouthAfricaNidMatch2Input) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SouthAfricaNidMatch2Input) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SouthAfricaNidMatch2Input) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ChinaResidentIdMatchOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The fullname of the individual in Chinese. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The birthdate of the individual. | [optional] 
**ResidentIdNumber** | Pointer to **NullableString** | The Resident ID number from the People&#39;s Republic of China (PRC) ID card.              The Resident ID number is an 18-digit number follows the format (AAAAAAYYYYMMDDXXX@): - Six digits (AAAAAA) are the area code of where the person was born or the area of first issuance for those who were born before the resident system was created. - Eight digits are the birthdate of the individual in YYYYMMDD format - Three digits are the order code assigned to individual. Males are assigned odd numbers and females even numbers. - Final digit is the checksum confirming the validity of the ID number&#39;s first 17 digits using ISO 7064:1983, MOD 11-2. | [optional] 

## Methods

### NewChinaResidentIdMatchOutput

`func NewChinaResidentIdMatchOutput() *ChinaResidentIdMatchOutput`

NewChinaResidentIdMatchOutput instantiates a new ChinaResidentIdMatchOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChinaResidentIdMatchOutputWithDefaults

`func NewChinaResidentIdMatchOutputWithDefaults() *ChinaResidentIdMatchOutput`

NewChinaResidentIdMatchOutputWithDefaults instantiates a new ChinaResidentIdMatchOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ChinaResidentIdMatchOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ChinaResidentIdMatchOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ChinaResidentIdMatchOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ChinaResidentIdMatchOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ChinaResidentIdMatchOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ChinaResidentIdMatchOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateOfBirth

`func (o *ChinaResidentIdMatchOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ChinaResidentIdMatchOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ChinaResidentIdMatchOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ChinaResidentIdMatchOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ChinaResidentIdMatchOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ChinaResidentIdMatchOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetResidentIdNumber

`func (o *ChinaResidentIdMatchOutput) GetResidentIdNumber() string`

GetResidentIdNumber returns the ResidentIdNumber field if non-nil, zero value otherwise.

### GetResidentIdNumberOk

`func (o *ChinaResidentIdMatchOutput) GetResidentIdNumberOk() (*string, bool)`

GetResidentIdNumberOk returns a tuple with the ResidentIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentIdNumber

`func (o *ChinaResidentIdMatchOutput) SetResidentIdNumber(v string)`

SetResidentIdNumber sets ResidentIdNumber field to given value.

### HasResidentIdNumber

`func (o *ChinaResidentIdMatchOutput) HasResidentIdNumber() bool`

HasResidentIdNumber returns a boolean if a field has been set.

### SetResidentIdNumberNil

`func (o *ChinaResidentIdMatchOutput) SetResidentIdNumberNil(b bool)`

 SetResidentIdNumberNil sets the value for ResidentIdNumber to be an explicit nil

### UnsetResidentIdNumber
`func (o *ChinaResidentIdMatchOutput) UnsetResidentIdNumber()`

UnsetResidentIdNumber ensures that no value is present for ResidentIdNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



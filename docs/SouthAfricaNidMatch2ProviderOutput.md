# SouthAfricaNidMatch2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | [**SouthAfricaNidMatch2NationalIdNumberField**](SouthAfricaNidMatch2NationalIdNumberField.md) | Outcome of the verification of the national ID number. | 
**FullName** | [**SouthAfricaNidMatch2FullNameField**](SouthAfricaNidMatch2FullNameField.md) | Outcome of the verification of the full name. | 
**DateOfBirth** | Pointer to [**NullableSouthAfricaNidMatch2DateOfBirthField**](SouthAfricaNidMatch2DateOfBirthField.md) | Outcome of the verification of the date of birth. | [optional] 
**Sex** | Pointer to [**NullableSouthAfricaNidMatch2SexField**](SouthAfricaNidMatch2SexField.md) | Outcome of the verification of the sex. | [optional] 
**CitizenshipStatus** | Pointer to **NullableString** | Citizenship status of the individual, if the ID number is valid.              Possible values: - Citizen - PermanentResident | [optional] 

## Methods

### NewSouthAfricaNidMatch2ProviderOutput

`func NewSouthAfricaNidMatch2ProviderOutput(nationalIdNumber SouthAfricaNidMatch2NationalIdNumberField, fullName SouthAfricaNidMatch2FullNameField, ) *SouthAfricaNidMatch2ProviderOutput`

NewSouthAfricaNidMatch2ProviderOutput instantiates a new SouthAfricaNidMatch2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidMatch2ProviderOutputWithDefaults

`func NewSouthAfricaNidMatch2ProviderOutputWithDefaults() *SouthAfricaNidMatch2ProviderOutput`

NewSouthAfricaNidMatch2ProviderOutputWithDefaults instantiates a new SouthAfricaNidMatch2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *SouthAfricaNidMatch2ProviderOutput) GetNationalIdNumber() SouthAfricaNidMatch2NationalIdNumberField`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *SouthAfricaNidMatch2ProviderOutput) GetNationalIdNumberOk() (*SouthAfricaNidMatch2NationalIdNumberField, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *SouthAfricaNidMatch2ProviderOutput) SetNationalIdNumber(v SouthAfricaNidMatch2NationalIdNumberField)`

SetNationalIdNumber sets NationalIdNumber field to given value.


### GetFullName

`func (o *SouthAfricaNidMatch2ProviderOutput) GetFullName() SouthAfricaNidMatch2FullNameField`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SouthAfricaNidMatch2ProviderOutput) GetFullNameOk() (*SouthAfricaNidMatch2FullNameField, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SouthAfricaNidMatch2ProviderOutput) SetFullName(v SouthAfricaNidMatch2FullNameField)`

SetFullName sets FullName field to given value.


### GetDateOfBirth

`func (o *SouthAfricaNidMatch2ProviderOutput) GetDateOfBirth() SouthAfricaNidMatch2DateOfBirthField`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SouthAfricaNidMatch2ProviderOutput) GetDateOfBirthOk() (*SouthAfricaNidMatch2DateOfBirthField, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SouthAfricaNidMatch2ProviderOutput) SetDateOfBirth(v SouthAfricaNidMatch2DateOfBirthField)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SouthAfricaNidMatch2ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SouthAfricaNidMatch2ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SouthAfricaNidMatch2ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetSex

`func (o *SouthAfricaNidMatch2ProviderOutput) GetSex() SouthAfricaNidMatch2SexField`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SouthAfricaNidMatch2ProviderOutput) GetSexOk() (*SouthAfricaNidMatch2SexField, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SouthAfricaNidMatch2ProviderOutput) SetSex(v SouthAfricaNidMatch2SexField)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SouthAfricaNidMatch2ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SouthAfricaNidMatch2ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SouthAfricaNidMatch2ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetCitizenshipStatus

`func (o *SouthAfricaNidMatch2ProviderOutput) GetCitizenshipStatus() string`

GetCitizenshipStatus returns the CitizenshipStatus field if non-nil, zero value otherwise.

### GetCitizenshipStatusOk

`func (o *SouthAfricaNidMatch2ProviderOutput) GetCitizenshipStatusOk() (*string, bool)`

GetCitizenshipStatusOk returns a tuple with the CitizenshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipStatus

`func (o *SouthAfricaNidMatch2ProviderOutput) SetCitizenshipStatus(v string)`

SetCitizenshipStatus sets CitizenshipStatus field to given value.

### HasCitizenshipStatus

`func (o *SouthAfricaNidMatch2ProviderOutput) HasCitizenshipStatus() bool`

HasCitizenshipStatus returns a boolean if a field has been set.

### SetCitizenshipStatusNil

`func (o *SouthAfricaNidMatch2ProviderOutput) SetCitizenshipStatusNil(b bool)`

 SetCitizenshipStatusNil sets the value for CitizenshipStatus to be an explicit nil

### UnsetCitizenshipStatus
`func (o *SouthAfricaNidMatch2ProviderOutput) UnsetCitizenshipStatus()`

UnsetCitizenshipStatus ensures that no value is present for CitizenshipStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



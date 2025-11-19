# KenyaNidInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The user&#39;s first name as it appears in their National ID | 
**MiddleName** | Pointer to **NullableString** | The user&#39;s middle name as it appears in their National ID (optional) | [optional] 
**LastName** | **string** | The user&#39;s middle name as it appears in their National ID (optional) | 
**DateOfBirth** | **string** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | 
**Gender** | **string** | The user&#39;s gender as it appears in their National ID (e.g., male, female) | 
**NationalIdNumber** | **string** | The user&#39;s Kenya National ID number | 

## Methods

### NewKenyaNidInput

`func NewKenyaNidInput(firstName string, lastName string, dateOfBirth string, gender string, nationalIdNumber string, ) *KenyaNidInput`

NewKenyaNidInput instantiates a new KenyaNidInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidInputWithDefaults

`func NewKenyaNidInputWithDefaults() *KenyaNidInput`

NewKenyaNidInputWithDefaults instantiates a new KenyaNidInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *KenyaNidInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *KenyaNidInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *KenyaNidInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *KenyaNidInput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *KenyaNidInput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *KenyaNidInput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *KenyaNidInput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *KenyaNidInput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *KenyaNidInput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *KenyaNidInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *KenyaNidInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *KenyaNidInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *KenyaNidInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *KenyaNidInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *KenyaNidInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetGender

`func (o *KenyaNidInput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *KenyaNidInput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *KenyaNidInput) SetGender(v string)`

SetGender sets Gender field to given value.


### GetNationalIdNumber

`func (o *KenyaNidInput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *KenyaNidInput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *KenyaNidInput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



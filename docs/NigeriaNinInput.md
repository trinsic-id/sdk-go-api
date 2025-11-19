# NigeriaNinInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The user&#39;s first name as it appears in their National ID | 
**MiddleName** | Pointer to **NullableString** | The user&#39;s middle name as it appears in their National ID (optional) | [optional] 
**LastName** | **string** | The user&#39;s last name as it appears in their National ID (optional) | 
**PhoneNumber** | Pointer to **NullableString** | The user&#39;s phone number (optional) | [optional] 
**DateOfBirth** | **string** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | 
**Gender** | Pointer to **NullableString** | The user&#39;s gender as it appears in their National ID (e.g., male, female) | [optional] 
**NationalIdNumber** | **string** | The user&#39;s National ID number | 

## Methods

### NewNigeriaNinInput

`func NewNigeriaNinInput(firstName string, lastName string, dateOfBirth string, nationalIdNumber string, ) *NigeriaNinInput`

NewNigeriaNinInput instantiates a new NigeriaNinInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinInputWithDefaults

`func NewNigeriaNinInputWithDefaults() *NigeriaNinInput`

NewNigeriaNinInputWithDefaults instantiates a new NigeriaNinInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *NigeriaNinInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NigeriaNinInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NigeriaNinInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *NigeriaNinInput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *NigeriaNinInput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *NigeriaNinInput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *NigeriaNinInput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *NigeriaNinInput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *NigeriaNinInput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *NigeriaNinInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NigeriaNinInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NigeriaNinInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhoneNumber

`func (o *NigeriaNinInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NigeriaNinInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NigeriaNinInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NigeriaNinInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *NigeriaNinInput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *NigeriaNinInput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetDateOfBirth

`func (o *NigeriaNinInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *NigeriaNinInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *NigeriaNinInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetGender

`func (o *NigeriaNinInput) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *NigeriaNinInput) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *NigeriaNinInput) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *NigeriaNinInput) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *NigeriaNinInput) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *NigeriaNinInput) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationalIdNumber

`func (o *NigeriaNinInput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *NigeriaNinInput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *NigeriaNinInput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UgandaNidMatch2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The user&#39;s Uganda National ID number (NIN) - 14 alphanumeric characters | [optional] 
**SecondaryIdNumber** | Pointer to **NullableString** | The card number on the document (secondary ID number). Required for Uganda Basic KYC instead of first name and last name. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The user&#39;s date of birth, in &#x60;YYYY-MM-DD&#x60; format | [optional] 

## Methods

### NewUgandaNidMatch2Input

`func NewUgandaNidMatch2Input() *UgandaNidMatch2Input`

NewUgandaNidMatch2Input instantiates a new UgandaNidMatch2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUgandaNidMatch2InputWithDefaults

`func NewUgandaNidMatch2InputWithDefaults() *UgandaNidMatch2Input`

NewUgandaNidMatch2InputWithDefaults instantiates a new UgandaNidMatch2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *UgandaNidMatch2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *UgandaNidMatch2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *UgandaNidMatch2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *UgandaNidMatch2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *UgandaNidMatch2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *UgandaNidMatch2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetSecondaryIdNumber

`func (o *UgandaNidMatch2Input) GetSecondaryIdNumber() string`

GetSecondaryIdNumber returns the SecondaryIdNumber field if non-nil, zero value otherwise.

### GetSecondaryIdNumberOk

`func (o *UgandaNidMatch2Input) GetSecondaryIdNumberOk() (*string, bool)`

GetSecondaryIdNumberOk returns a tuple with the SecondaryIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIdNumber

`func (o *UgandaNidMatch2Input) SetSecondaryIdNumber(v string)`

SetSecondaryIdNumber sets SecondaryIdNumber field to given value.

### HasSecondaryIdNumber

`func (o *UgandaNidMatch2Input) HasSecondaryIdNumber() bool`

HasSecondaryIdNumber returns a boolean if a field has been set.

### SetSecondaryIdNumberNil

`func (o *UgandaNidMatch2Input) SetSecondaryIdNumberNil(b bool)`

 SetSecondaryIdNumberNil sets the value for SecondaryIdNumber to be an explicit nil

### UnsetSecondaryIdNumber
`func (o *UgandaNidMatch2Input) UnsetSecondaryIdNumber()`

UnsetSecondaryIdNumber ensures that no value is present for SecondaryIdNumber, not even an explicit nil
### GetDateOfBirth

`func (o *UgandaNidMatch2Input) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UgandaNidMatch2Input) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UgandaNidMatch2Input) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UgandaNidMatch2Input) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *UgandaNidMatch2Input) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *UgandaNidMatch2Input) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



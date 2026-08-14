# AudkenniProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalRegisterId** | Pointer to **NullableString** | National Register ID number (kennitala) of the individual.              Often written with a hyphen following the 6th digit (123456-1234). ID numbers are composed of ten digits. The first six of these are the individual’s date of birth in the format DDMMYY. The seventh and eighth digits are randomly chosen when the ID number is allocated, the ninth digit used to be a check digit (modulus (11) but this has been removed. The tenth indicates the century of the individual’s birth: ‘9’ for 1900–1999, ‘0’ for 2000 and beyond. | [optional] 
**SubjectId** | Pointer to **NullableString** | Unique identifier for the individual across Audkenni.              It is recommended to use this identifier instead of using the national register ID directly. | [optional] 
**Name** | Pointer to **NullableString** | The full name of the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual as parsed from the National Register ID number. Is null if the date of birth could not be parsed, but should always be present | [optional] 

## Methods

### NewAudkenniProviderOutput

`func NewAudkenniProviderOutput() *AudkenniProviderOutput`

NewAudkenniProviderOutput instantiates a new AudkenniProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudkenniProviderOutputWithDefaults

`func NewAudkenniProviderOutputWithDefaults() *AudkenniProviderOutput`

NewAudkenniProviderOutputWithDefaults instantiates a new AudkenniProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalRegisterId

`func (o *AudkenniProviderOutput) GetNationalRegisterId() string`

GetNationalRegisterId returns the NationalRegisterId field if non-nil, zero value otherwise.

### GetNationalRegisterIdOk

`func (o *AudkenniProviderOutput) GetNationalRegisterIdOk() (*string, bool)`

GetNationalRegisterIdOk returns a tuple with the NationalRegisterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegisterId

`func (o *AudkenniProviderOutput) SetNationalRegisterId(v string)`

SetNationalRegisterId sets NationalRegisterId field to given value.

### HasNationalRegisterId

`func (o *AudkenniProviderOutput) HasNationalRegisterId() bool`

HasNationalRegisterId returns a boolean if a field has been set.

### SetNationalRegisterIdNil

`func (o *AudkenniProviderOutput) SetNationalRegisterIdNil(b bool)`

 SetNationalRegisterIdNil sets the value for NationalRegisterId to be an explicit nil

### UnsetNationalRegisterId
`func (o *AudkenniProviderOutput) UnsetNationalRegisterId()`

UnsetNationalRegisterId ensures that no value is present for NationalRegisterId, not even an explicit nil
### GetSubjectId

`func (o *AudkenniProviderOutput) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *AudkenniProviderOutput) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *AudkenniProviderOutput) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *AudkenniProviderOutput) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### SetSubjectIdNil

`func (o *AudkenniProviderOutput) SetSubjectIdNil(b bool)`

 SetSubjectIdNil sets the value for SubjectId to be an explicit nil

### UnsetSubjectId
`func (o *AudkenniProviderOutput) UnsetSubjectId()`

UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
### GetName

`func (o *AudkenniProviderOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudkenniProviderOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudkenniProviderOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AudkenniProviderOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AudkenniProviderOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AudkenniProviderOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDateOfBirth

`func (o *AudkenniProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AudkenniProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AudkenniProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AudkenniProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AudkenniProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AudkenniProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



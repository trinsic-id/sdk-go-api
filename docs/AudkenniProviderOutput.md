# AudkenniProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalRegisterId** | **string** | National Register ID number (kennitala) of the individual.              Often written with a hyphen following the 6th digit (123456-1234). ID numbers are composed of ten digits. The first six of these are the individual’s date of birth in the format DDMMYY. The seventh and eighth digits are randomly chosen when the ID number is allocated, the ninth digit used to be a check digit (modulus (11) but this has been removed. The tenth indicates the century of the individual’s birth: ‘9’ for 1900–1999, ‘0’ for 2000 and beyond. | 
**SubjectId** | **string** | Unique identifier for the individual across Audkenni.              It is recommended to use this identifier instead of using the national register ID directly. | 
**Name** | **string** | The full name of the individual. | 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual as parsed from the National Register ID number. Is null if the date of birth could not be parsed, but should always be present | [optional] 

## Methods

### NewAudkenniProviderOutput

`func NewAudkenniProviderOutput(nationalRegisterId string, subjectId string, name string, ) *AudkenniProviderOutput`

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



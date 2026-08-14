# UgandaNidMatch2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | Pointer to [**NullableUgandaNidMatch2NationalIdNumberField**](UgandaNidMatch2NationalIdNumberField.md) | Outcome of verifying the national identification number in the NIRA database. | [optional] 
**SecondaryIdNumber** | Pointer to [**NullableUgandaNidMatch2SecondaryIdNumberField**](UgandaNidMatch2SecondaryIdNumberField.md) | Outcome of comparing the submitted card number with the NIRA database. | [optional] 
**DateOfBirth** | Pointer to [**NullableUgandaNidMatch2DateOfBirthField**](UgandaNidMatch2DateOfBirthField.md) | Outcome of comparing the submitted date of birth with the NIRA database. | [optional] 

## Methods

### NewUgandaNidMatch2ProviderOutput

`func NewUgandaNidMatch2ProviderOutput() *UgandaNidMatch2ProviderOutput`

NewUgandaNidMatch2ProviderOutput instantiates a new UgandaNidMatch2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUgandaNidMatch2ProviderOutputWithDefaults

`func NewUgandaNidMatch2ProviderOutputWithDefaults() *UgandaNidMatch2ProviderOutput`

NewUgandaNidMatch2ProviderOutputWithDefaults instantiates a new UgandaNidMatch2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *UgandaNidMatch2ProviderOutput) GetNationalIdNumber() UgandaNidMatch2NationalIdNumberField`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *UgandaNidMatch2ProviderOutput) GetNationalIdNumberOk() (*UgandaNidMatch2NationalIdNumberField, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *UgandaNidMatch2ProviderOutput) SetNationalIdNumber(v UgandaNidMatch2NationalIdNumberField)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *UgandaNidMatch2ProviderOutput) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *UgandaNidMatch2ProviderOutput) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *UgandaNidMatch2ProviderOutput) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetSecondaryIdNumber

`func (o *UgandaNidMatch2ProviderOutput) GetSecondaryIdNumber() UgandaNidMatch2SecondaryIdNumberField`

GetSecondaryIdNumber returns the SecondaryIdNumber field if non-nil, zero value otherwise.

### GetSecondaryIdNumberOk

`func (o *UgandaNidMatch2ProviderOutput) GetSecondaryIdNumberOk() (*UgandaNidMatch2SecondaryIdNumberField, bool)`

GetSecondaryIdNumberOk returns a tuple with the SecondaryIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIdNumber

`func (o *UgandaNidMatch2ProviderOutput) SetSecondaryIdNumber(v UgandaNidMatch2SecondaryIdNumberField)`

SetSecondaryIdNumber sets SecondaryIdNumber field to given value.

### HasSecondaryIdNumber

`func (o *UgandaNidMatch2ProviderOutput) HasSecondaryIdNumber() bool`

HasSecondaryIdNumber returns a boolean if a field has been set.

### SetSecondaryIdNumberNil

`func (o *UgandaNidMatch2ProviderOutput) SetSecondaryIdNumberNil(b bool)`

 SetSecondaryIdNumberNil sets the value for SecondaryIdNumber to be an explicit nil

### UnsetSecondaryIdNumber
`func (o *UgandaNidMatch2ProviderOutput) UnsetSecondaryIdNumber()`

UnsetSecondaryIdNumber ensures that no value is present for SecondaryIdNumber, not even an explicit nil
### GetDateOfBirth

`func (o *UgandaNidMatch2ProviderOutput) GetDateOfBirth() UgandaNidMatch2DateOfBirthField`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UgandaNidMatch2ProviderOutput) GetDateOfBirthOk() (*UgandaNidMatch2DateOfBirthField, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UgandaNidMatch2ProviderOutput) SetDateOfBirth(v UgandaNidMatch2DateOfBirthField)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UgandaNidMatch2ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *UgandaNidMatch2ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *UgandaNidMatch2ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



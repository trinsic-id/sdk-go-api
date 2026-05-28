# UgandaNidMatch2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | [**UgandaNidMatch2NationalIdNumberField**](UgandaNidMatch2NationalIdNumberField.md) | Outcome of verifying the national identification number in the NIRA database. | 
**SecondaryIdNumber** | [**UgandaNidMatch2SecondaryIdNumberField**](UgandaNidMatch2SecondaryIdNumberField.md) | Outcome of comparing the submitted card number with the NIRA database. | 
**DateOfBirth** | [**UgandaNidMatch2DateOfBirthField**](UgandaNidMatch2DateOfBirthField.md) | Outcome of comparing the submitted date of birth with the NIRA database. | 

## Methods

### NewUgandaNidMatch2ProviderOutput

`func NewUgandaNidMatch2ProviderOutput(nationalIdNumber UgandaNidMatch2NationalIdNumberField, secondaryIdNumber UgandaNidMatch2SecondaryIdNumberField, dateOfBirth UgandaNidMatch2DateOfBirthField, ) *UgandaNidMatch2ProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



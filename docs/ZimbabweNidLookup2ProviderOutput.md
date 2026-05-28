# ZimbabweNidLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | **string** | Zimbabwe National ID number (NID) issued by the Zimbabwean government.              Format: - 8-9 digits followed by 1 letter and 2 digits - Regex: /^[0-9]{8,9}[A-Za-z]\\d{2}$/ - There is no publicly documented encoding scheme for encoding personal information in the NID - No check digit or algorithm has been publicly documented by the Zimbabwean government | 
**GivenName** | **string** | Given name(s) of the individual. | 
**FamilyName** | **string** | Family name of the individual. | 
**FullName** | **string** | Full name of the individual. | 
**Sex** | **string** | Sex of the individual.              Possible values: - Male - Female - Not Applicable - Unknown | 
**DateOfBirth** | **string** | Date of birth of the individual. | 
**PlaceOfBirth** | Pointer to **NullableString** | Place of birth, typically a municipality or city within Zimbabwe. | [optional] 
**IsAlive** | Pointer to **NullableBool** | Whether individual is reported as alive by the Zimbabwean government. | [optional] 
**DateOfDeath** | Pointer to **NullableString** | Date of death of the individual, when available. | [optional] 

## Methods

### NewZimbabweNidLookup2ProviderOutput

`func NewZimbabweNidLookup2ProviderOutput(nationalIdNumber string, givenName string, familyName string, fullName string, sex string, dateOfBirth string, ) *ZimbabweNidLookup2ProviderOutput`

NewZimbabweNidLookup2ProviderOutput instantiates a new ZimbabweNidLookup2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZimbabweNidLookup2ProviderOutputWithDefaults

`func NewZimbabweNidLookup2ProviderOutputWithDefaults() *ZimbabweNidLookup2ProviderOutput`

NewZimbabweNidLookup2ProviderOutputWithDefaults instantiates a new ZimbabweNidLookup2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *ZimbabweNidLookup2ProviderOutput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *ZimbabweNidLookup2ProviderOutput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.


### GetGivenName

`func (o *ZimbabweNidLookup2ProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ZimbabweNidLookup2ProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *ZimbabweNidLookup2ProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ZimbabweNidLookup2ProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetFullName

`func (o *ZimbabweNidLookup2ProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ZimbabweNidLookup2ProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetSex

`func (o *ZimbabweNidLookup2ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *ZimbabweNidLookup2ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.


### GetDateOfBirth

`func (o *ZimbabweNidLookup2ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ZimbabweNidLookup2ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetPlaceOfBirth

`func (o *ZimbabweNidLookup2ProviderOutput) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ZimbabweNidLookup2ProviderOutput) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ZimbabweNidLookup2ProviderOutput) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ZimbabweNidLookup2ProviderOutput) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ZimbabweNidLookup2ProviderOutput) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetIsAlive

`func (o *ZimbabweNidLookup2ProviderOutput) GetIsAlive() bool`

GetIsAlive returns the IsAlive field if non-nil, zero value otherwise.

### GetIsAliveOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetIsAliveOk() (*bool, bool)`

GetIsAliveOk returns a tuple with the IsAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlive

`func (o *ZimbabweNidLookup2ProviderOutput) SetIsAlive(v bool)`

SetIsAlive sets IsAlive field to given value.

### HasIsAlive

`func (o *ZimbabweNidLookup2ProviderOutput) HasIsAlive() bool`

HasIsAlive returns a boolean if a field has been set.

### SetIsAliveNil

`func (o *ZimbabweNidLookup2ProviderOutput) SetIsAliveNil(b bool)`

 SetIsAliveNil sets the value for IsAlive to be an explicit nil

### UnsetIsAlive
`func (o *ZimbabweNidLookup2ProviderOutput) UnsetIsAlive()`

UnsetIsAlive ensures that no value is present for IsAlive, not even an explicit nil
### GetDateOfDeath

`func (o *ZimbabweNidLookup2ProviderOutput) GetDateOfDeath() string`

GetDateOfDeath returns the DateOfDeath field if non-nil, zero value otherwise.

### GetDateOfDeathOk

`func (o *ZimbabweNidLookup2ProviderOutput) GetDateOfDeathOk() (*string, bool)`

GetDateOfDeathOk returns a tuple with the DateOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDeath

`func (o *ZimbabweNidLookup2ProviderOutput) SetDateOfDeath(v string)`

SetDateOfDeath sets DateOfDeath field to given value.

### HasDateOfDeath

`func (o *ZimbabweNidLookup2ProviderOutput) HasDateOfDeath() bool`

HasDateOfDeath returns a boolean if a field has been set.

### SetDateOfDeathNil

`func (o *ZimbabweNidLookup2ProviderOutput) SetDateOfDeathNil(b bool)`

 SetDateOfDeathNil sets the value for DateOfDeath to be an explicit nil

### UnsetDateOfDeath
`func (o *ZimbabweNidLookup2ProviderOutput) UnsetDateOfDeath()`

UnsetDateOfDeath ensures that no value is present for DateOfDeath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



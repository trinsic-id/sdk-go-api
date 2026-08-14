# PolandEdoAppProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**NationalIdentificationNumber** | Pointer to **NullableString** | The 11-digit Polish national identification number (PESEL) of the verified individual.              This is in the format YYMMDDZZZGQ, where: - YYMMDD is the date of birth - ZZZ is a unique identifier - G is sex (even for females, odd for males) - Q is a checksum digit              The year of birth encoded in this identifier assumes a default year of birth in the 20th century. If the year of birth is in the range [1800, 1899], the month portion is incremented by 80. If the year of birth is in the range [2000, 2099] the month portion is incremented by 20. If the year of birth is in the range [2100, 2199], the month portion is incremented by 40. If the year of birth is in the range [2200, 2299], the month portion is incremented by 60. | [optional] 

## Methods

### NewPolandEdoAppProviderOutput

`func NewPolandEdoAppProviderOutput() *PolandEdoAppProviderOutput`

NewPolandEdoAppProviderOutput instantiates a new PolandEdoAppProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolandEdoAppProviderOutputWithDefaults

`func NewPolandEdoAppProviderOutputWithDefaults() *PolandEdoAppProviderOutput`

NewPolandEdoAppProviderOutputWithDefaults instantiates a new PolandEdoAppProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PolandEdoAppProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PolandEdoAppProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PolandEdoAppProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PolandEdoAppProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PolandEdoAppProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PolandEdoAppProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PolandEdoAppProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PolandEdoAppProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PolandEdoAppProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PolandEdoAppProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PolandEdoAppProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PolandEdoAppProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *PolandEdoAppProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PolandEdoAppProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PolandEdoAppProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PolandEdoAppProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PolandEdoAppProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PolandEdoAppProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationalIdentificationNumber

`func (o *PolandEdoAppProviderOutput) GetNationalIdentificationNumber() string`

GetNationalIdentificationNumber returns the NationalIdentificationNumber field if non-nil, zero value otherwise.

### GetNationalIdentificationNumberOk

`func (o *PolandEdoAppProviderOutput) GetNationalIdentificationNumberOk() (*string, bool)`

GetNationalIdentificationNumberOk returns a tuple with the NationalIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentificationNumber

`func (o *PolandEdoAppProviderOutput) SetNationalIdentificationNumber(v string)`

SetNationalIdentificationNumber sets NationalIdentificationNumber field to given value.

### HasNationalIdentificationNumber

`func (o *PolandEdoAppProviderOutput) HasNationalIdentificationNumber() bool`

HasNationalIdentificationNumber returns a boolean if a field has been set.

### SetNationalIdentificationNumberNil

`func (o *PolandEdoAppProviderOutput) SetNationalIdentificationNumberNil(b bool)`

 SetNationalIdentificationNumberNil sets the value for NationalIdentificationNumber to be an explicit nil

### UnsetNationalIdentificationNumber
`func (o *PolandEdoAppProviderOutput) UnsetNationalIdentificationNumber()`

UnsetNationalIdentificationNumber ensures that no value is present for NationalIdentificationNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



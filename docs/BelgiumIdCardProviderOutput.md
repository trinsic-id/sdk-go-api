# BelgiumIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**NationalRegisterNumber** | Pointer to **NullableString** | The Belgian National Register Number (\&quot;Rijksregisternummer\&quot;) of the verified individual.              This is an 11-digit number in the format YYMMDDXXXCC, where: - YYMMDD represents the individual&#39;s date of birth (year, month, day). - XXX is a sequential birth number, odd for males and even for females. - CC is a checksum, calculated with the equation: 97 - (YYMMDDXXX mod 97)              For births in the year 2000 or later, the digit &#39;2&#39; is prepended to the first 9 digits during checksum calculation. | [optional] 

## Methods

### NewBelgiumIdCardProviderOutput

`func NewBelgiumIdCardProviderOutput() *BelgiumIdCardProviderOutput`

NewBelgiumIdCardProviderOutput instantiates a new BelgiumIdCardProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBelgiumIdCardProviderOutputWithDefaults

`func NewBelgiumIdCardProviderOutputWithDefaults() *BelgiumIdCardProviderOutput`

NewBelgiumIdCardProviderOutputWithDefaults instantiates a new BelgiumIdCardProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BelgiumIdCardProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BelgiumIdCardProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BelgiumIdCardProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BelgiumIdCardProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *BelgiumIdCardProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *BelgiumIdCardProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *BelgiumIdCardProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BelgiumIdCardProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BelgiumIdCardProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BelgiumIdCardProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *BelgiumIdCardProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *BelgiumIdCardProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *BelgiumIdCardProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BelgiumIdCardProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BelgiumIdCardProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BelgiumIdCardProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *BelgiumIdCardProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *BelgiumIdCardProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationalRegisterNumber

`func (o *BelgiumIdCardProviderOutput) GetNationalRegisterNumber() string`

GetNationalRegisterNumber returns the NationalRegisterNumber field if non-nil, zero value otherwise.

### GetNationalRegisterNumberOk

`func (o *BelgiumIdCardProviderOutput) GetNationalRegisterNumberOk() (*string, bool)`

GetNationalRegisterNumberOk returns a tuple with the NationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegisterNumber

`func (o *BelgiumIdCardProviderOutput) SetNationalRegisterNumber(v string)`

SetNationalRegisterNumber sets NationalRegisterNumber field to given value.

### HasNationalRegisterNumber

`func (o *BelgiumIdCardProviderOutput) HasNationalRegisterNumber() bool`

HasNationalRegisterNumber returns a boolean if a field has been set.

### SetNationalRegisterNumberNil

`func (o *BelgiumIdCardProviderOutput) SetNationalRegisterNumberNil(b bool)`

 SetNationalRegisterNumberNil sets the value for NationalRegisterNumber to be an explicit nil

### UnsetNationalRegisterNumber
`func (o *BelgiumIdCardProviderOutput) UnsetNationalRegisterNumber()`

UnsetNationalRegisterNumber ensures that no value is present for NationalRegisterNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



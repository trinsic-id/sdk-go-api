# EstoniaIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**PersonalIdentificationCode** | Pointer to **NullableString** | The 11-digit Estonian personal identification code (\&quot;isikukood\&quot;) of the verified individual.              This is in the format GYYMMDDSSSC, where: - G combines the century and sex of birth - YYMMDD is the date of birth within the given century - SSS is a serial number distinguishing people born on the same date - C is a checksum digit              If G is odd, the individual is male. If G is even, the individual is female.              If G is 1 or 2, the individual was born in the 19th century (1800-1899). If G is 3 or 4, the individual was born in the 20th century (1900-1999). If G is 5 or 6, the individual was born in the 21st century (2000-2099). If G is 7 or 8, the individual was born in the 22nd century (2100-2199). | [optional] 

## Methods

### NewEstoniaIdCardProviderOutput

`func NewEstoniaIdCardProviderOutput() *EstoniaIdCardProviderOutput`

NewEstoniaIdCardProviderOutput instantiates a new EstoniaIdCardProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstoniaIdCardProviderOutputWithDefaults

`func NewEstoniaIdCardProviderOutputWithDefaults() *EstoniaIdCardProviderOutput`

NewEstoniaIdCardProviderOutputWithDefaults instantiates a new EstoniaIdCardProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *EstoniaIdCardProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EstoniaIdCardProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EstoniaIdCardProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EstoniaIdCardProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *EstoniaIdCardProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *EstoniaIdCardProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *EstoniaIdCardProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EstoniaIdCardProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EstoniaIdCardProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EstoniaIdCardProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *EstoniaIdCardProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *EstoniaIdCardProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *EstoniaIdCardProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *EstoniaIdCardProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *EstoniaIdCardProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *EstoniaIdCardProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *EstoniaIdCardProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *EstoniaIdCardProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalIdentificationCode

`func (o *EstoniaIdCardProviderOutput) GetPersonalIdentificationCode() string`

GetPersonalIdentificationCode returns the PersonalIdentificationCode field if non-nil, zero value otherwise.

### GetPersonalIdentificationCodeOk

`func (o *EstoniaIdCardProviderOutput) GetPersonalIdentificationCodeOk() (*string, bool)`

GetPersonalIdentificationCodeOk returns a tuple with the PersonalIdentificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentificationCode

`func (o *EstoniaIdCardProviderOutput) SetPersonalIdentificationCode(v string)`

SetPersonalIdentificationCode sets PersonalIdentificationCode field to given value.

### HasPersonalIdentificationCode

`func (o *EstoniaIdCardProviderOutput) HasPersonalIdentificationCode() bool`

HasPersonalIdentificationCode returns a boolean if a field has been set.

### SetPersonalIdentificationCodeNil

`func (o *EstoniaIdCardProviderOutput) SetPersonalIdentificationCodeNil(b bool)`

 SetPersonalIdentificationCodeNil sets the value for PersonalIdentificationCode to be an explicit nil

### UnsetPersonalIdentificationCode
`func (o *EstoniaIdCardProviderOutput) UnsetPersonalIdentificationCode()`

UnsetPersonalIdentificationCode ensures that no value is present for PersonalIdentificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



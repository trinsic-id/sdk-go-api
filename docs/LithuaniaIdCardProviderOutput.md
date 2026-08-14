# LithuaniaIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**PersonalCode** | Pointer to **NullableString** | The 11-digit Lithuanian Personal Code (asmens kodas) of the verified individual.              If the first digit is \&quot;9\&quot; (rare), the rest of the identifier is random and has no structure.              Otherwise, this is in the format GYYMMDDSSSC, where:              - G is both gender and century of birth - YYMMDD is the date of birth - SSS is a sequential serial number - C is a checksum digit              If G is odd, the gender is male. If G is even, the gender is female.              If G is 1 or 2, the individual was born in the 19th century (1800-1899). If G is 3 or 4, the individual was born in the 20th century (1900-1999). If G is 5 or 6, the individual was born in the 21st century (2000-2099).              Rarely, the date of birth may be encoded as 000000 if it is not known. | [optional] 

## Methods

### NewLithuaniaIdCardProviderOutput

`func NewLithuaniaIdCardProviderOutput() *LithuaniaIdCardProviderOutput`

NewLithuaniaIdCardProviderOutput instantiates a new LithuaniaIdCardProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLithuaniaIdCardProviderOutputWithDefaults

`func NewLithuaniaIdCardProviderOutputWithDefaults() *LithuaniaIdCardProviderOutput`

NewLithuaniaIdCardProviderOutputWithDefaults instantiates a new LithuaniaIdCardProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *LithuaniaIdCardProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LithuaniaIdCardProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LithuaniaIdCardProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LithuaniaIdCardProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *LithuaniaIdCardProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *LithuaniaIdCardProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *LithuaniaIdCardProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LithuaniaIdCardProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LithuaniaIdCardProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LithuaniaIdCardProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *LithuaniaIdCardProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *LithuaniaIdCardProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *LithuaniaIdCardProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *LithuaniaIdCardProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *LithuaniaIdCardProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *LithuaniaIdCardProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *LithuaniaIdCardProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *LithuaniaIdCardProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalCode

`func (o *LithuaniaIdCardProviderOutput) GetPersonalCode() string`

GetPersonalCode returns the PersonalCode field if non-nil, zero value otherwise.

### GetPersonalCodeOk

`func (o *LithuaniaIdCardProviderOutput) GetPersonalCodeOk() (*string, bool)`

GetPersonalCodeOk returns a tuple with the PersonalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCode

`func (o *LithuaniaIdCardProviderOutput) SetPersonalCode(v string)`

SetPersonalCode sets PersonalCode field to given value.

### HasPersonalCode

`func (o *LithuaniaIdCardProviderOutput) HasPersonalCode() bool`

HasPersonalCode returns a boolean if a field has been set.

### SetPersonalCodeNil

`func (o *LithuaniaIdCardProviderOutput) SetPersonalCodeNil(b bool)`

 SetPersonalCodeNil sets the value for PersonalCode to be an explicit nil

### UnsetPersonalCode
`func (o *LithuaniaIdCardProviderOutput) UnsetPersonalCode()`

UnsetPersonalCode ensures that no value is present for PersonalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



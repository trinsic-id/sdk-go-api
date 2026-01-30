# ItsmeProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**HashedNationalRegisterNumber** | Pointer to **NullableString** | The hashed version of the Belgian National Register Number of the verified individual.              By default, itsme does not return the raw National Register Number of the individual; instead, only a hashed version is returned.              Your account must be approved by itsme to receive the raw, unhashed National Register Number. | [optional] 
**NationalRegisterNumber** | Pointer to **NullableString** | The raw (not hashed) Belgian National Register Number (\&quot;Rijksregisternummer\&quot;) of the verified individual.              Only returned if your account has been explicitly authorized to receive it by itsme; by law, this data is considered sensitive personal data.              This is an 11-digit number in the format YYMMDDXXXCC, where: - YYMMDD represents the individual&#39;s date of birth (year, month, day). - XXX is a sequential birth number, odd for females and even for males. - CC is a checksum, calculated with the equation: 97 - (YYMMDDXXX mod 97)              For births in the year 2000 or later, the digit &#39;2&#39; is prepended to the first 9 digits during checksum calculation. | [optional] 

## Methods

### NewItsmeProviderOutput

`func NewItsmeProviderOutput(firstName string, lastName string, dateOfBirth string, ) *ItsmeProviderOutput`

NewItsmeProviderOutput instantiates a new ItsmeProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeProviderOutputWithDefaults

`func NewItsmeProviderOutputWithDefaults() *ItsmeProviderOutput`

NewItsmeProviderOutputWithDefaults instantiates a new ItsmeProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ItsmeProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ItsmeProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ItsmeProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ItsmeProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ItsmeProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ItsmeProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *ItsmeProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ItsmeProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ItsmeProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) GetHashedNationalRegisterNumber() string`

GetHashedNationalRegisterNumber returns the HashedNationalRegisterNumber field if non-nil, zero value otherwise.

### GetHashedNationalRegisterNumberOk

`func (o *ItsmeProviderOutput) GetHashedNationalRegisterNumberOk() (*string, bool)`

GetHashedNationalRegisterNumberOk returns a tuple with the HashedNationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) SetHashedNationalRegisterNumber(v string)`

SetHashedNationalRegisterNumber sets HashedNationalRegisterNumber field to given value.

### HasHashedNationalRegisterNumber

`func (o *ItsmeProviderOutput) HasHashedNationalRegisterNumber() bool`

HasHashedNationalRegisterNumber returns a boolean if a field has been set.

### SetHashedNationalRegisterNumberNil

`func (o *ItsmeProviderOutput) SetHashedNationalRegisterNumberNil(b bool)`

 SetHashedNationalRegisterNumberNil sets the value for HashedNationalRegisterNumber to be an explicit nil

### UnsetHashedNationalRegisterNumber
`func (o *ItsmeProviderOutput) UnsetHashedNationalRegisterNumber()`

UnsetHashedNationalRegisterNumber ensures that no value is present for HashedNationalRegisterNumber, not even an explicit nil
### GetNationalRegisterNumber

`func (o *ItsmeProviderOutput) GetNationalRegisterNumber() string`

GetNationalRegisterNumber returns the NationalRegisterNumber field if non-nil, zero value otherwise.

### GetNationalRegisterNumberOk

`func (o *ItsmeProviderOutput) GetNationalRegisterNumberOk() (*string, bool)`

GetNationalRegisterNumberOk returns a tuple with the NationalRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRegisterNumber

`func (o *ItsmeProviderOutput) SetNationalRegisterNumber(v string)`

SetNationalRegisterNumber sets NationalRegisterNumber field to given value.

### HasNationalRegisterNumber

`func (o *ItsmeProviderOutput) HasNationalRegisterNumber() bool`

HasNationalRegisterNumber returns a boolean if a field has been set.

### SetNationalRegisterNumberNil

`func (o *ItsmeProviderOutput) SetNationalRegisterNumberNil(b bool)`

 SetNationalRegisterNumberNil sets the value for NationalRegisterNumber to be an explicit nil

### UnsetNationalRegisterNumber
`func (o *ItsmeProviderOutput) UnsetNationalRegisterNumber()`

UnsetNationalRegisterNumber ensures that no value is present for NationalRegisterNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



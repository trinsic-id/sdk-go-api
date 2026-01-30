# EstoniaIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**PersonalIdentificationCode** | **string** | The 11-digit Estonian personal identification code (\&quot;isikukood\&quot;) of the verified individual.              This is in the format GYYMMDDSSSC, where: - G combines the century and sex of birth - YYMMDD is the date of birth within the given century - SSS is a serial number distinguishing people born on the same date - C is a checksum digit              If G is odd, the individual is male. If G is even, the individual is female.              If G is 1 or 2, the individual was born in the 19th century (1800-1899). If G is 3 or 4, the individual was born in the 20th century (1900-1999). If G is 5 or 6, the individual was born in the 21st century (2000-2099). If G is 7 or 8, the individual was born in the 22nd century (2100-2199). | 

## Methods

### NewEstoniaIdCardProviderOutput

`func NewEstoniaIdCardProviderOutput(firstName string, lastName string, dateOfBirth string, personalIdentificationCode string, ) *EstoniaIdCardProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



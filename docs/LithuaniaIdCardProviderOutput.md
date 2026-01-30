# LithuaniaIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**PersonalCode** | **string** | The 11-digit Lithuanian Personal Code (asmens kodas) of the verified individual.              If the first digit is \&quot;9\&quot; (rare), the rest of the identifier is random and has no structure.              Otherwise, this is in the format GYYMMDDSSSC, where:              - G is both gender and century of birth - YYMMDD is the date of birth - SSS is a sequential serial number - C is a checksum digit              If G is odd, the gender is male. If G is even, the gender is female.              If G is 1 or 2, the individual was born in the 19th century (1800-1899). If G is 3 or 4, the individual was born in the 20th century (1900-1999). If G is 5 or 6, the individual was born in the 21st century (2000-2099).              Rarely, the date of birth may be encoded as 000000 if it is not known. | 

## Methods

### NewLithuaniaIdCardProviderOutput

`func NewLithuaniaIdCardProviderOutput(firstName string, lastName string, dateOfBirth string, personalCode string, ) *LithuaniaIdCardProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



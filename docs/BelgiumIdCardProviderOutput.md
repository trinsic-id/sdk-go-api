# BelgiumIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**NationalRegisterNumber** | **string** | The Belgian National Register Number (\&quot;Rijksregisternummer\&quot;) of the verified individual.              This is an 11-digit number in the format YYMMDDXXXCC, where: - YYMMDD represents the individual&#39;s date of birth (year, month, day). - XXX is a sequential birth number, odd for males and even for females. - CC is a checksum, calculated with the equation: 97 - (YYMMDDXXX mod 97)              For births in the year 2000 or later, the digit &#39;2&#39; is prepended to the first 9 digits during checksum calculation. | 

## Methods

### NewBelgiumIdCardProviderOutput

`func NewBelgiumIdCardProviderOutput(firstName string, lastName string, dateOfBirth string, nationalRegisterNumber string, ) *BelgiumIdCardProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



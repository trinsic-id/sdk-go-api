# SerbiaIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**UniqueMasterCitizenNumber** | **string** | The 13-digit Serbian Unique Master Citizen Number (\&quot;JMBG\&quot; / \&quot;Jedinstveni Matični Broj Građana\&quot;) of the verified individual.              This is in the format DDMMYYYRRSSSC, where: - DDMM is the day and month of birth - YYY is the last three digits of the year of birth - RR is the political region code of the region of birth (if born after 1976) or of first registration (if born before 1976) - SSS is a unique sex-specific serial number for individuals born on the same date in the same region - C is a checksum digit              If YYY is between 000 and 099, the millennium digit of the year is \&quot;2\&quot;; the individual was born after the year 2000. If YYY is between 800 and 999, the millennium digit of the year is \&quot;1\&quot;; the individual was born before the year 2000.              If SSS is between 000 and 499, the individual is male. If SSS is between 500 and 999, the individual is female. | 

## Methods

### NewSerbiaIdCardProviderOutput

`func NewSerbiaIdCardProviderOutput(firstName string, lastName string, dateOfBirth string, uniqueMasterCitizenNumber string, ) *SerbiaIdCardProviderOutput`

NewSerbiaIdCardProviderOutput instantiates a new SerbiaIdCardProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSerbiaIdCardProviderOutputWithDefaults

`func NewSerbiaIdCardProviderOutputWithDefaults() *SerbiaIdCardProviderOutput`

NewSerbiaIdCardProviderOutputWithDefaults instantiates a new SerbiaIdCardProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *SerbiaIdCardProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SerbiaIdCardProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SerbiaIdCardProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SerbiaIdCardProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SerbiaIdCardProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SerbiaIdCardProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDateOfBirth

`func (o *SerbiaIdCardProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SerbiaIdCardProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SerbiaIdCardProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetUniqueMasterCitizenNumber

`func (o *SerbiaIdCardProviderOutput) GetUniqueMasterCitizenNumber() string`

GetUniqueMasterCitizenNumber returns the UniqueMasterCitizenNumber field if non-nil, zero value otherwise.

### GetUniqueMasterCitizenNumberOk

`func (o *SerbiaIdCardProviderOutput) GetUniqueMasterCitizenNumberOk() (*string, bool)`

GetUniqueMasterCitizenNumberOk returns a tuple with the UniqueMasterCitizenNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueMasterCitizenNumber

`func (o *SerbiaIdCardProviderOutput) SetUniqueMasterCitizenNumber(v string)`

SetUniqueMasterCitizenNumber sets UniqueMasterCitizenNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



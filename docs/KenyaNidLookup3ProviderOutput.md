# KenyaNidLookup3ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The Kenya National ID Number (Nambari ya Kitambulisho) or Unique Personal Identifier (Maisha Namba).              This is the primary unique identifier for Kenyan citizens in all government systems, issued by the National Registration Bureau (NRB). The format is either 8 digits for National ID or 9 digits for Maisha Namba UPI (the new format since 2023). | [optional] 
**FirstName** | Pointer to **NullableString** | The first name of the individual. | [optional] 
**Surname** | Pointer to **NullableString** | The surname of the individual. | [optional] 
**OtherName** | Pointer to **NullableString** | Other name (middle name) of the individual. | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the individual.              Possible values: - Male - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual.              Format - YYYY-MM-DD | [optional] 
**Citizenship** | Pointer to **NullableString** | Citizenship of the individual.              Format - ISO 3166-1 alpha-2 country code | [optional] 
**SerialNumber** | Pointer to **NullableString** | The physical card serial number printed on the Kenya National ID card.              This is distinct from the ID Number and serves as a card issuance tracking identifier maintained by IPRS. This value changes each time a new physical card is issued (loss, damage, renewal). | [optional] 

## Methods

### NewKenyaNidLookup3ProviderOutput

`func NewKenyaNidLookup3ProviderOutput() *KenyaNidLookup3ProviderOutput`

NewKenyaNidLookup3ProviderOutput instantiates a new KenyaNidLookup3ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKenyaNidLookup3ProviderOutputWithDefaults

`func NewKenyaNidLookup3ProviderOutputWithDefaults() *KenyaNidLookup3ProviderOutput`

NewKenyaNidLookup3ProviderOutputWithDefaults instantiates a new KenyaNidLookup3ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *KenyaNidLookup3ProviderOutput) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *KenyaNidLookup3ProviderOutput) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *KenyaNidLookup3ProviderOutput) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *KenyaNidLookup3ProviderOutput) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *KenyaNidLookup3ProviderOutput) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *KenyaNidLookup3ProviderOutput) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetFirstName

`func (o *KenyaNidLookup3ProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *KenyaNidLookup3ProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *KenyaNidLookup3ProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *KenyaNidLookup3ProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *KenyaNidLookup3ProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *KenyaNidLookup3ProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetSurname

`func (o *KenyaNidLookup3ProviderOutput) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *KenyaNidLookup3ProviderOutput) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *KenyaNidLookup3ProviderOutput) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *KenyaNidLookup3ProviderOutput) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *KenyaNidLookup3ProviderOutput) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *KenyaNidLookup3ProviderOutput) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetOtherName

`func (o *KenyaNidLookup3ProviderOutput) GetOtherName() string`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *KenyaNidLookup3ProviderOutput) GetOtherNameOk() (*string, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *KenyaNidLookup3ProviderOutput) SetOtherName(v string)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *KenyaNidLookup3ProviderOutput) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### SetOtherNameNil

`func (o *KenyaNidLookup3ProviderOutput) SetOtherNameNil(b bool)`

 SetOtherNameNil sets the value for OtherName to be an explicit nil

### UnsetOtherName
`func (o *KenyaNidLookup3ProviderOutput) UnsetOtherName()`

UnsetOtherName ensures that no value is present for OtherName, not even an explicit nil
### GetSex

`func (o *KenyaNidLookup3ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *KenyaNidLookup3ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *KenyaNidLookup3ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *KenyaNidLookup3ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *KenyaNidLookup3ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *KenyaNidLookup3ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *KenyaNidLookup3ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *KenyaNidLookup3ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *KenyaNidLookup3ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *KenyaNidLookup3ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *KenyaNidLookup3ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *KenyaNidLookup3ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetCitizenship

`func (o *KenyaNidLookup3ProviderOutput) GetCitizenship() string`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *KenyaNidLookup3ProviderOutput) GetCitizenshipOk() (*string, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *KenyaNidLookup3ProviderOutput) SetCitizenship(v string)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *KenyaNidLookup3ProviderOutput) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### SetCitizenshipNil

`func (o *KenyaNidLookup3ProviderOutput) SetCitizenshipNil(b bool)`

 SetCitizenshipNil sets the value for Citizenship to be an explicit nil

### UnsetCitizenship
`func (o *KenyaNidLookup3ProviderOutput) UnsetCitizenship()`

UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
### GetSerialNumber

`func (o *KenyaNidLookup3ProviderOutput) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KenyaNidLookup3ProviderOutput) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KenyaNidLookup3ProviderOutput) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *KenyaNidLookup3ProviderOutput) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *KenyaNidLookup3ProviderOutput) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *KenyaNidLookup3ProviderOutput) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



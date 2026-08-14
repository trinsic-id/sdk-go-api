# PolandMojeIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**NationalIdentificationNumber** | Pointer to **NullableString** | The 11-digit Polish national identification number (PESEL) of the verified individual.              This is in the format YYMMDDZZZGQ, where: - YYMMDD is the date of birth - ZZZ is a unique identifier - G is sex (even for females, odd for males) - Q is a checksum digit              The year of birth encoded in this identifier assumes a default year of birth in the 20th century. If the year of birth is in the range [1800, 1899], the month portion is incremented by 80. If the year of birth is in the range [2000, 2099] the month portion is incremented by 20. If the year of birth is in the range [2100, 2199], the month portion is incremented by 40. If the year of birth is in the range [2200, 2299], the month portion is incremented by 60. | [optional] 

## Methods

### NewPolandMojeIdProviderOutput

`func NewPolandMojeIdProviderOutput() *PolandMojeIdProviderOutput`

NewPolandMojeIdProviderOutput instantiates a new PolandMojeIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolandMojeIdProviderOutputWithDefaults

`func NewPolandMojeIdProviderOutputWithDefaults() *PolandMojeIdProviderOutput`

NewPolandMojeIdProviderOutputWithDefaults instantiates a new PolandMojeIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PolandMojeIdProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PolandMojeIdProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PolandMojeIdProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PolandMojeIdProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PolandMojeIdProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PolandMojeIdProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PolandMojeIdProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PolandMojeIdProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PolandMojeIdProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PolandMojeIdProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PolandMojeIdProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PolandMojeIdProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetNationalIdentificationNumber

`func (o *PolandMojeIdProviderOutput) GetNationalIdentificationNumber() string`

GetNationalIdentificationNumber returns the NationalIdentificationNumber field if non-nil, zero value otherwise.

### GetNationalIdentificationNumberOk

`func (o *PolandMojeIdProviderOutput) GetNationalIdentificationNumberOk() (*string, bool)`

GetNationalIdentificationNumberOk returns a tuple with the NationalIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentificationNumber

`func (o *PolandMojeIdProviderOutput) SetNationalIdentificationNumber(v string)`

SetNationalIdentificationNumber sets NationalIdentificationNumber field to given value.

### HasNationalIdentificationNumber

`func (o *PolandMojeIdProviderOutput) HasNationalIdentificationNumber() bool`

HasNationalIdentificationNumber returns a boolean if a field has been set.

### SetNationalIdentificationNumberNil

`func (o *PolandMojeIdProviderOutput) SetNationalIdentificationNumberNil(b bool)`

 SetNationalIdentificationNumberNil sets the value for NationalIdentificationNumber to be an explicit nil

### UnsetNationalIdentificationNumber
`func (o *PolandMojeIdProviderOutput) UnsetNationalIdentificationNumber()`

UnsetNationalIdentificationNumber ensures that no value is present for NationalIdentificationNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



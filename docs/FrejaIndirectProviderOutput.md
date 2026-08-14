# FrejaIndirectProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**PersonalNumber** | Pointer to **NullableString** | The value returned by Freja in the \&quot;ssn\&quot; field.              The actual value of this field depends on the country of origin used to create the Freja credential. It is typically a Social Security Number, National Identification Number, or equivalent personal identifier. | [optional] 
**PersonalNumberCountry** | Pointer to **NullableString** | The 2-digit ISO country code of the country which issued the personal number. | [optional] 

## Methods

### NewFrejaIndirectProviderOutput

`func NewFrejaIndirectProviderOutput() *FrejaIndirectProviderOutput`

NewFrejaIndirectProviderOutput instantiates a new FrejaIndirectProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrejaIndirectProviderOutputWithDefaults

`func NewFrejaIndirectProviderOutputWithDefaults() *FrejaIndirectProviderOutput`

NewFrejaIndirectProviderOutputWithDefaults instantiates a new FrejaIndirectProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *FrejaIndirectProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FrejaIndirectProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FrejaIndirectProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FrejaIndirectProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *FrejaIndirectProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *FrejaIndirectProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *FrejaIndirectProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FrejaIndirectProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FrejaIndirectProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FrejaIndirectProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *FrejaIndirectProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *FrejaIndirectProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPersonalNumber

`func (o *FrejaIndirectProviderOutput) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *FrejaIndirectProviderOutput) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *FrejaIndirectProviderOutput) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *FrejaIndirectProviderOutput) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.

### SetPersonalNumberNil

`func (o *FrejaIndirectProviderOutput) SetPersonalNumberNil(b bool)`

 SetPersonalNumberNil sets the value for PersonalNumber to be an explicit nil

### UnsetPersonalNumber
`func (o *FrejaIndirectProviderOutput) UnsetPersonalNumber()`

UnsetPersonalNumber ensures that no value is present for PersonalNumber, not even an explicit nil
### GetPersonalNumberCountry

`func (o *FrejaIndirectProviderOutput) GetPersonalNumberCountry() string`

GetPersonalNumberCountry returns the PersonalNumberCountry field if non-nil, zero value otherwise.

### GetPersonalNumberCountryOk

`func (o *FrejaIndirectProviderOutput) GetPersonalNumberCountryOk() (*string, bool)`

GetPersonalNumberCountryOk returns a tuple with the PersonalNumberCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumberCountry

`func (o *FrejaIndirectProviderOutput) SetPersonalNumberCountry(v string)`

SetPersonalNumberCountry sets PersonalNumberCountry field to given value.

### HasPersonalNumberCountry

`func (o *FrejaIndirectProviderOutput) HasPersonalNumberCountry() bool`

HasPersonalNumberCountry returns a boolean if a field has been set.

### SetPersonalNumberCountryNil

`func (o *FrejaIndirectProviderOutput) SetPersonalNumberCountryNil(b bool)`

 SetPersonalNumberCountryNil sets the value for PersonalNumberCountry to be an explicit nil

### UnsetPersonalNumberCountry
`func (o *FrejaIndirectProviderOutput) UnsetPersonalNumberCountry()`

UnsetPersonalNumberCountry ensures that no value is present for PersonalNumberCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



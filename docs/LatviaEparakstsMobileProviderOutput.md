# LatviaEparakstsMobileProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**PersonalCode** | Pointer to **NullableString** | The 11-digit Latvian personal code (\&quot;personas kods\&quot;) of the verified individual.              This has two possible formats, depending on whether the personal code was issued after July 1, 2017.              For personal codes issued before July 1, 2017, the format is DDMMYY-CZZZQ, where: - DDMMYY is the date of birth, followed by an optional hyphen - C represents the century of birth (&#39;0&#39; for 1800-1899, &#39;1&#39; for 1900-1999, &#39;2&#39; for 2000-2099) - ZZZ is a serial number - Q is a checksum digit              For personal codes issued on or after July 1, 2017, the format is 32ZZZZZZZZQ, where: - 32 is a fixed prefix - ZZZZZZZQ are 8 random digits - Q is a checksum digit              NOTE: Individuals born before July 1, 2017 can elect to be issued a new personal code which does not contain their birthdate. Therefore, no concrete assumptions may be made about an individual&#39;s date of birth based solely on the format of their personal code. | [optional] 

## Methods

### NewLatviaEparakstsMobileProviderOutput

`func NewLatviaEparakstsMobileProviderOutput() *LatviaEparakstsMobileProviderOutput`

NewLatviaEparakstsMobileProviderOutput instantiates a new LatviaEparakstsMobileProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatviaEparakstsMobileProviderOutputWithDefaults

`func NewLatviaEparakstsMobileProviderOutputWithDefaults() *LatviaEparakstsMobileProviderOutput`

NewLatviaEparakstsMobileProviderOutputWithDefaults instantiates a new LatviaEparakstsMobileProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *LatviaEparakstsMobileProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LatviaEparakstsMobileProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LatviaEparakstsMobileProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LatviaEparakstsMobileProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *LatviaEparakstsMobileProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *LatviaEparakstsMobileProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *LatviaEparakstsMobileProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LatviaEparakstsMobileProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LatviaEparakstsMobileProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LatviaEparakstsMobileProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *LatviaEparakstsMobileProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *LatviaEparakstsMobileProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPersonalCode

`func (o *LatviaEparakstsMobileProviderOutput) GetPersonalCode() string`

GetPersonalCode returns the PersonalCode field if non-nil, zero value otherwise.

### GetPersonalCodeOk

`func (o *LatviaEparakstsMobileProviderOutput) GetPersonalCodeOk() (*string, bool)`

GetPersonalCodeOk returns a tuple with the PersonalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCode

`func (o *LatviaEparakstsMobileProviderOutput) SetPersonalCode(v string)`

SetPersonalCode sets PersonalCode field to given value.

### HasPersonalCode

`func (o *LatviaEparakstsMobileProviderOutput) HasPersonalCode() bool`

HasPersonalCode returns a boolean if a field has been set.

### SetPersonalCodeNil

`func (o *LatviaEparakstsMobileProviderOutput) SetPersonalCodeNil(b bool)`

 SetPersonalCodeNil sets the value for PersonalCode to be an explicit nil

### UnsetPersonalCode
`func (o *LatviaEparakstsMobileProviderOutput) UnsetPersonalCode()`

UnsetPersonalCode ensures that no value is present for PersonalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



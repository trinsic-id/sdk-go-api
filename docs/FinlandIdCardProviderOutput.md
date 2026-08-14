# FinlandIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the verified individual | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the verified individual | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the verified individual | [optional] 
**PersonalIdentificationCode** | Pointer to **NullableString** | The 11-digit Finnish Personal Identification Code (Henkilötunnus) of the verified individual.               This is in the format DDMMYYCZZZQ, where: - DDMMYY is the date of birth - C is a symbol which determines the century of birth - ZZZ is an individual number, indicating gender - Q is a checksum character              If ZZZ is even, the individual is female. If ZZZ is odd, the individual is male.              If C is &#39;+&#39;, the individual was born in the 19th century (1800-1899). If C is &#39;-&#39;, &#39;U&#39;, &#39;V&#39;, &#39;W&#39;, &#39;X&#39;, or &#39;Y&#39;, the individual was born in the 20th century (1900-1999). If C is &#39;A&#39;, &#39;B&#39;, &#39;C&#39;, &#39;D&#39;, &#39;E&#39;, or &#39;F&#39;, the individual was born in the 21st century (2000-2099). | [optional] 

## Methods

### NewFinlandIdCardProviderOutput

`func NewFinlandIdCardProviderOutput() *FinlandIdCardProviderOutput`

NewFinlandIdCardProviderOutput instantiates a new FinlandIdCardProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinlandIdCardProviderOutputWithDefaults

`func NewFinlandIdCardProviderOutputWithDefaults() *FinlandIdCardProviderOutput`

NewFinlandIdCardProviderOutputWithDefaults instantiates a new FinlandIdCardProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *FinlandIdCardProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FinlandIdCardProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FinlandIdCardProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FinlandIdCardProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *FinlandIdCardProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *FinlandIdCardProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *FinlandIdCardProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FinlandIdCardProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FinlandIdCardProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FinlandIdCardProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *FinlandIdCardProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *FinlandIdCardProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *FinlandIdCardProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *FinlandIdCardProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *FinlandIdCardProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *FinlandIdCardProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *FinlandIdCardProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *FinlandIdCardProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalIdentificationCode

`func (o *FinlandIdCardProviderOutput) GetPersonalIdentificationCode() string`

GetPersonalIdentificationCode returns the PersonalIdentificationCode field if non-nil, zero value otherwise.

### GetPersonalIdentificationCodeOk

`func (o *FinlandIdCardProviderOutput) GetPersonalIdentificationCodeOk() (*string, bool)`

GetPersonalIdentificationCodeOk returns a tuple with the PersonalIdentificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdentificationCode

`func (o *FinlandIdCardProviderOutput) SetPersonalIdentificationCode(v string)`

SetPersonalIdentificationCode sets PersonalIdentificationCode field to given value.

### HasPersonalIdentificationCode

`func (o *FinlandIdCardProviderOutput) HasPersonalIdentificationCode() bool`

HasPersonalIdentificationCode returns a boolean if a field has been set.

### SetPersonalIdentificationCodeNil

`func (o *FinlandIdCardProviderOutput) SetPersonalIdentificationCodeNil(b bool)`

 SetPersonalIdentificationCodeNil sets the value for PersonalIdentificationCode to be an explicit nil

### UnsetPersonalIdentificationCode
`func (o *FinlandIdCardProviderOutput) UnsetPersonalIdentificationCode()`

UnsetPersonalIdentificationCode ensures that no value is present for PersonalIdentificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



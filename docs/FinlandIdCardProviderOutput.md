# FinlandIdCardProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**DateOfBirth** | **string** | The date of birth of the verified individual | 
**PersonalIdentificationCode** | **string** | The 11-digit Finnish Personal Identification Code (Henkilötunnus) of the verified individual.               This is in the format DDMMYYCZZZQ, where: - DDMMYY is the date of birth - C is a symbol which determines the century of birth - ZZZ is an individual number, indicating gender - Q is a checksum character              If ZZZ is even, the individual is female. If ZZZ is odd, the individual is male.              If C is &#39;+&#39;, the individual was born in the 19th century (1800-1899). If C is &#39;-&#39;, &#39;U&#39;, &#39;V&#39;, &#39;W&#39;, &#39;X&#39;, or &#39;Y&#39;, the individual was born in the 20th century (1900-1999). If C is &#39;A&#39;, &#39;B&#39;, &#39;C&#39;, &#39;D&#39;, &#39;E&#39;, or &#39;F&#39;, the individual was born in the 21st century (2000-2099). | 

## Methods

### NewFinlandIdCardProviderOutput

`func NewFinlandIdCardProviderOutput(firstName string, lastName string, dateOfBirth string, personalIdentificationCode string, ) *FinlandIdCardProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



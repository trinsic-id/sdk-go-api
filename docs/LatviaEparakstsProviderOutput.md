# LatviaEparakstsProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**PersonalCode** | **string** | The 11-digit Latvian personal code (\&quot;personas kods\&quot;) of the verified individual.              This has two possible formats, depending on whether the personal code was issued after July 1, 2017.              For personal codes issued before July 1, 2017, the format is DDMMYY-CZZZQ, where: - DDMMYY is the date of birth, followed by an optional hyphen - C represents the century of birth (&#39;0&#39; for 1800-1899, &#39;1&#39; for 1900-1999, &#39;2&#39; for 2000-2099) - ZZZ is a serial number - Q is a checksum digit              For personal codes issued on or after July 1, 2017, the format is 32ZZZZZZZZQ, where: - 32 is a fixed prefix - ZZZZZZZQ are 8 random digits - Q is a checksum digit              NOTE: Individuals born before July 1, 2017 can elect to be issued a new personal code which does not contain their birthdate. Therefore, no concrete assumptions may be made about an individual&#39;s date of birth based solely on the format of their personal code. | 

## Methods

### NewLatviaEparakstsProviderOutput

`func NewLatviaEparakstsProviderOutput(firstName string, lastName string, personalCode string, ) *LatviaEparakstsProviderOutput`

NewLatviaEparakstsProviderOutput instantiates a new LatviaEparakstsProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatviaEparakstsProviderOutputWithDefaults

`func NewLatviaEparakstsProviderOutputWithDefaults() *LatviaEparakstsProviderOutput`

NewLatviaEparakstsProviderOutputWithDefaults instantiates a new LatviaEparakstsProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *LatviaEparakstsProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LatviaEparakstsProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LatviaEparakstsProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *LatviaEparakstsProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LatviaEparakstsProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LatviaEparakstsProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPersonalCode

`func (o *LatviaEparakstsProviderOutput) GetPersonalCode() string`

GetPersonalCode returns the PersonalCode field if non-nil, zero value otherwise.

### GetPersonalCodeOk

`func (o *LatviaEparakstsProviderOutput) GetPersonalCodeOk() (*string, bool)`

GetPersonalCodeOk returns a tuple with the PersonalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCode

`func (o *LatviaEparakstsProviderOutput) SetPersonalCode(v string)`

SetPersonalCode sets PersonalCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



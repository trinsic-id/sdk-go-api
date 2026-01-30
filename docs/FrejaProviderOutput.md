# FrejaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the verified individual | 
**LastName** | **string** | The last name of the verified individual | 
**PersonalNumber** | **string** | The value returned by Freja in the \&quot;ssn\&quot; field.              The actual value of this field depends on the country of origin used to create the Freja credential. It is typically a Social Security Number, National Identification Number, or equivalent personal identifier. | 
**PersonalNumberCountry** | **string** | The 2-digit ISO country code of the country which issued the personal number. | 

## Methods

### NewFrejaProviderOutput

`func NewFrejaProviderOutput(firstName string, lastName string, personalNumber string, personalNumberCountry string, ) *FrejaProviderOutput`

NewFrejaProviderOutput instantiates a new FrejaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrejaProviderOutputWithDefaults

`func NewFrejaProviderOutputWithDefaults() *FrejaProviderOutput`

NewFrejaProviderOutputWithDefaults instantiates a new FrejaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *FrejaProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FrejaProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FrejaProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *FrejaProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FrejaProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FrejaProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPersonalNumber

`func (o *FrejaProviderOutput) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *FrejaProviderOutput) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *FrejaProviderOutput) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.


### GetPersonalNumberCountry

`func (o *FrejaProviderOutput) GetPersonalNumberCountry() string`

GetPersonalNumberCountry returns the PersonalNumberCountry field if non-nil, zero value otherwise.

### GetPersonalNumberCountryOk

`func (o *FrejaProviderOutput) GetPersonalNumberCountryOk() (*string, bool)`

GetPersonalNumberCountryOk returns a tuple with the PersonalNumberCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumberCountry

`func (o *FrejaProviderOutput) SetPersonalNumberCountry(v string)`

SetPersonalNumberCountry sets PersonalNumberCountry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



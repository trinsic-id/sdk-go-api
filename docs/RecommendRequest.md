# RecommendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **NullableString** | The phone number of the user you wish to generate a recommendation for. Will be used to look up the user&#39;s identity in the network | [optional] 
**Countries** | Pointer to **[]string** | A list of countries, in alpha-2 ISO 3166 format, you wish to specify for the recommendation, this can include the user&#39;s country of residence, nationality, etc. | [optional] 
**Subdivisions** | Pointer to **[]string** | If one of the countries has subdivisions, for example the US states, you can specify a list of these to further refine the recommendation (e.g., CA, UT, NY) | [optional] 
**IpAddresses** | Pointer to **[]string** | Provide the IP addresses of the user you wish to generate a recommendation for. Will be used to look up the user&#39;s geographic location. | [optional] 
**IncludeDisabledProviders** | Pointer to **NullableBool** | If true, the recommendation will include providers that are disabled for the app | [optional] 

## Methods

### NewRecommendRequest

`func NewRecommendRequest() *RecommendRequest`

NewRecommendRequest instantiates a new RecommendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendRequestWithDefaults

`func NewRecommendRequestWithDefaults() *RecommendRequest`

NewRecommendRequestWithDefaults instantiates a new RecommendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *RecommendRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RecommendRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RecommendRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RecommendRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *RecommendRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *RecommendRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountries

`func (o *RecommendRequest) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *RecommendRequest) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *RecommendRequest) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *RecommendRequest) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### SetCountriesNil

`func (o *RecommendRequest) SetCountriesNil(b bool)`

 SetCountriesNil sets the value for Countries to be an explicit nil

### UnsetCountries
`func (o *RecommendRequest) UnsetCountries()`

UnsetCountries ensures that no value is present for Countries, not even an explicit nil
### GetSubdivisions

`func (o *RecommendRequest) GetSubdivisions() []string`

GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

### GetSubdivisionsOk

`func (o *RecommendRequest) GetSubdivisionsOk() (*[]string, bool)`

GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisions

`func (o *RecommendRequest) SetSubdivisions(v []string)`

SetSubdivisions sets Subdivisions field to given value.

### HasSubdivisions

`func (o *RecommendRequest) HasSubdivisions() bool`

HasSubdivisions returns a boolean if a field has been set.

### SetSubdivisionsNil

`func (o *RecommendRequest) SetSubdivisionsNil(b bool)`

 SetSubdivisionsNil sets the value for Subdivisions to be an explicit nil

### UnsetSubdivisions
`func (o *RecommendRequest) UnsetSubdivisions()`

UnsetSubdivisions ensures that no value is present for Subdivisions, not even an explicit nil
### GetIpAddresses

`func (o *RecommendRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *RecommendRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *RecommendRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *RecommendRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *RecommendRequest) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *RecommendRequest) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetIncludeDisabledProviders

`func (o *RecommendRequest) GetIncludeDisabledProviders() bool`

GetIncludeDisabledProviders returns the IncludeDisabledProviders field if non-nil, zero value otherwise.

### GetIncludeDisabledProvidersOk

`func (o *RecommendRequest) GetIncludeDisabledProvidersOk() (*bool, bool)`

GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabledProviders

`func (o *RecommendRequest) SetIncludeDisabledProviders(v bool)`

SetIncludeDisabledProviders sets IncludeDisabledProviders field to given value.

### HasIncludeDisabledProviders

`func (o *RecommendRequest) HasIncludeDisabledProviders() bool`

HasIncludeDisabledProviders returns a boolean if a field has been set.

### SetIncludeDisabledProvidersNil

`func (o *RecommendRequest) SetIncludeDisabledProvidersNil(b bool)`

 SetIncludeDisabledProvidersNil sets the value for IncludeDisabledProviders to be an explicit nil

### UnsetIncludeDisabledProviders
`func (o *RecommendRequest) UnsetIncludeDisabledProviders()`

UnsetIncludeDisabledProviders ensures that no value is present for IncludeDisabledProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



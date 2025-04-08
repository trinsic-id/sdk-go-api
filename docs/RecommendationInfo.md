# RecommendationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **NullableString** | The phone number of the user you wish to generate a recommendation for.              Will be used to look up the user&#39;s identity in the network, as well as to determine the user&#39;s geographic location. | [optional] 
**Countries** | Pointer to **[]string** | A list of countries, in alpha-2 ISO 3166 format, which the user is related to.              This can include the user&#39;s country of residence, nationality, etc. | [optional] 
**Subdivisions** | Pointer to **[]string** | If one of the countries has subdivisions (for example: US states), specify those related to the user here (e.g., CA, UT, NY) | [optional] 
**IpAddresses** | Pointer to **[]string** | Any IP addresses related to the user.              Will be used to determine the user&#39;s geographic location. | [optional] 

## Methods

### NewRecommendationInfo

`func NewRecommendationInfo() *RecommendationInfo`

NewRecommendationInfo instantiates a new RecommendationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationInfoWithDefaults

`func NewRecommendationInfoWithDefaults() *RecommendationInfo`

NewRecommendationInfoWithDefaults instantiates a new RecommendationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *RecommendationInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RecommendationInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RecommendationInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RecommendationInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *RecommendationInfo) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *RecommendationInfo) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountries

`func (o *RecommendationInfo) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *RecommendationInfo) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *RecommendationInfo) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *RecommendationInfo) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### SetCountriesNil

`func (o *RecommendationInfo) SetCountriesNil(b bool)`

 SetCountriesNil sets the value for Countries to be an explicit nil

### UnsetCountries
`func (o *RecommendationInfo) UnsetCountries()`

UnsetCountries ensures that no value is present for Countries, not even an explicit nil
### GetSubdivisions

`func (o *RecommendationInfo) GetSubdivisions() []string`

GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

### GetSubdivisionsOk

`func (o *RecommendationInfo) GetSubdivisionsOk() (*[]string, bool)`

GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisions

`func (o *RecommendationInfo) SetSubdivisions(v []string)`

SetSubdivisions sets Subdivisions field to given value.

### HasSubdivisions

`func (o *RecommendationInfo) HasSubdivisions() bool`

HasSubdivisions returns a boolean if a field has been set.

### SetSubdivisionsNil

`func (o *RecommendationInfo) SetSubdivisionsNil(b bool)`

 SetSubdivisionsNil sets the value for Subdivisions to be an explicit nil

### UnsetSubdivisions
`func (o *RecommendationInfo) UnsetSubdivisions()`

UnsetSubdivisions ensures that no value is present for Subdivisions, not even an explicit nil
### GetIpAddresses

`func (o *RecommendationInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *RecommendationInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *RecommendationInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *RecommendationInfo) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *RecommendationInfo) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *RecommendationInfo) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



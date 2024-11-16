# RecommendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number of the user you whish to generate a recommendation for. Will be used to lookup the user&#39;s identity in the network | 
**Countries** | Pointer to **[]string** | A list of countries, in alpha-2 ISO 3166 format, you wish to specify for the recommendation, this can include the user&#39;s country of residence, nationality, etc. | [optional] 
**States** | Pointer to **[]string** | If one of the countries is US, you can specify a list of US states to further refine the recommendation (e.g., CA, UT, NY) | [optional] 

## Methods

### NewRecommendRequest

`func NewRecommendRequest(phoneNumber string, ) *RecommendRequest`

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

### GetStates

`func (o *RecommendRequest) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *RecommendRequest) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *RecommendRequest) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *RecommendRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



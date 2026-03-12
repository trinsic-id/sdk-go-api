# RecommendProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendedProviders** | [**[]RecommendProviderInformation**](RecommendProviderInformation.md) |  | 
**Remainder** | [**[]RecommendProviderInformation**](RecommendProviderInformation.md) | All Providers available to your Verification Profile which are not in &#x60;recommendedProviders&#x60; | 

## Methods

### NewRecommendProvidersResponse

`func NewRecommendProvidersResponse(recommendedProviders []RecommendProviderInformation, remainder []RecommendProviderInformation, ) *RecommendProvidersResponse`

NewRecommendProvidersResponse instantiates a new RecommendProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendProvidersResponseWithDefaults

`func NewRecommendProvidersResponseWithDefaults() *RecommendProvidersResponse`

NewRecommendProvidersResponseWithDefaults instantiates a new RecommendProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendedProviders

`func (o *RecommendProvidersResponse) GetRecommendedProviders() []RecommendProviderInformation`

GetRecommendedProviders returns the RecommendedProviders field if non-nil, zero value otherwise.

### GetRecommendedProvidersOk

`func (o *RecommendProvidersResponse) GetRecommendedProvidersOk() (*[]RecommendProviderInformation, bool)`

GetRecommendedProvidersOk returns a tuple with the RecommendedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedProviders

`func (o *RecommendProvidersResponse) SetRecommendedProviders(v []RecommendProviderInformation)`

SetRecommendedProviders sets RecommendedProviders field to given value.


### GetRemainder

`func (o *RecommendProvidersResponse) GetRemainder() []RecommendProviderInformation`

GetRemainder returns the Remainder field if non-nil, zero value otherwise.

### GetRemainderOk

`func (o *RecommendProvidersResponse) GetRemainderOk() (*[]RecommendProviderInformation, bool)`

GetRemainderOk returns a tuple with the Remainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainder

`func (o *RecommendProvidersResponse) SetRemainder(v []RecommendProviderInformation)`

SetRemainder sets Remainder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



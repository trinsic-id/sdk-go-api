# RecommendProvidersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationProfileId** | **string** | The ID of the VerificationProfile to use for this recommendation. | 
**RecommendationInfo** | Pointer to [**NullableRecommendationInfo**](RecommendationInfo.md) | Information about the user you wish to generate a recommendation for. | [optional] 
**Health** | Pointer to [**NullableRecommendProviderHealthOption**](RecommendProviderHealthOption.md) | Filter providers by health status. | [optional] 

## Methods

### NewRecommendProvidersRequest

`func NewRecommendProvidersRequest(verificationProfileId string, ) *RecommendProvidersRequest`

NewRecommendProvidersRequest instantiates a new RecommendProvidersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendProvidersRequestWithDefaults

`func NewRecommendProvidersRequestWithDefaults() *RecommendProvidersRequest`

NewRecommendProvidersRequestWithDefaults instantiates a new RecommendProvidersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationProfileId

`func (o *RecommendProvidersRequest) GetVerificationProfileId() string`

GetVerificationProfileId returns the VerificationProfileId field if non-nil, zero value otherwise.

### GetVerificationProfileIdOk

`func (o *RecommendProvidersRequest) GetVerificationProfileIdOk() (*string, bool)`

GetVerificationProfileIdOk returns a tuple with the VerificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfileId

`func (o *RecommendProvidersRequest) SetVerificationProfileId(v string)`

SetVerificationProfileId sets VerificationProfileId field to given value.


### GetRecommendationInfo

`func (o *RecommendProvidersRequest) GetRecommendationInfo() RecommendationInfo`

GetRecommendationInfo returns the RecommendationInfo field if non-nil, zero value otherwise.

### GetRecommendationInfoOk

`func (o *RecommendProvidersRequest) GetRecommendationInfoOk() (*RecommendationInfo, bool)`

GetRecommendationInfoOk returns a tuple with the RecommendationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationInfo

`func (o *RecommendProvidersRequest) SetRecommendationInfo(v RecommendationInfo)`

SetRecommendationInfo sets RecommendationInfo field to given value.

### HasRecommendationInfo

`func (o *RecommendProvidersRequest) HasRecommendationInfo() bool`

HasRecommendationInfo returns a boolean if a field has been set.

### SetRecommendationInfoNil

`func (o *RecommendProvidersRequest) SetRecommendationInfoNil(b bool)`

 SetRecommendationInfoNil sets the value for RecommendationInfo to be an explicit nil

### UnsetRecommendationInfo
`func (o *RecommendProvidersRequest) UnsetRecommendationInfo()`

UnsetRecommendationInfo ensures that no value is present for RecommendationInfo, not even an explicit nil
### GetHealth

`func (o *RecommendProvidersRequest) GetHealth() RecommendProviderHealthOption`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *RecommendProvidersRequest) GetHealthOk() (*RecommendProviderHealthOption, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *RecommendProvidersRequest) SetHealth(v RecommendProviderHealthOption)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *RecommendProvidersRequest) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### SetHealthNil

`func (o *RecommendProvidersRequest) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *RecommendProvidersRequest) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



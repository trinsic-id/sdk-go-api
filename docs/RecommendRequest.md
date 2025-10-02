# RecommendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationProfileId** | **string** | The ID of the VerificationProfile to use for this recommendation. | 
**RecommendationInfo** | Pointer to [**NullableRecommendationInfo**](RecommendationInfo.md) | Information about the user you wish to generate a recommendation for. | [optional] 
**Health** | Pointer to **NullableString** | Filter providers by health status. Valid values: \&quot;online\&quot;, \&quot;offline\&quot;, \&quot;all\&quot;. Defaults to \&quot;online\&quot;. | [optional] 

## Methods

### NewRecommendRequest

`func NewRecommendRequest(verificationProfileId string, ) *RecommendRequest`

NewRecommendRequest instantiates a new RecommendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendRequestWithDefaults

`func NewRecommendRequestWithDefaults() *RecommendRequest`

NewRecommendRequestWithDefaults instantiates a new RecommendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationProfileId

`func (o *RecommendRequest) GetVerificationProfileId() string`

GetVerificationProfileId returns the VerificationProfileId field if non-nil, zero value otherwise.

### GetVerificationProfileIdOk

`func (o *RecommendRequest) GetVerificationProfileIdOk() (*string, bool)`

GetVerificationProfileIdOk returns a tuple with the VerificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfileId

`func (o *RecommendRequest) SetVerificationProfileId(v string)`

SetVerificationProfileId sets VerificationProfileId field to given value.


### GetRecommendationInfo

`func (o *RecommendRequest) GetRecommendationInfo() RecommendationInfo`

GetRecommendationInfo returns the RecommendationInfo field if non-nil, zero value otherwise.

### GetRecommendationInfoOk

`func (o *RecommendRequest) GetRecommendationInfoOk() (*RecommendationInfo, bool)`

GetRecommendationInfoOk returns a tuple with the RecommendationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationInfo

`func (o *RecommendRequest) SetRecommendationInfo(v RecommendationInfo)`

SetRecommendationInfo sets RecommendationInfo field to given value.

### HasRecommendationInfo

`func (o *RecommendRequest) HasRecommendationInfo() bool`

HasRecommendationInfo returns a boolean if a field has been set.

### SetRecommendationInfoNil

`func (o *RecommendRequest) SetRecommendationInfoNil(b bool)`

 SetRecommendationInfoNil sets the value for RecommendationInfo to be an explicit nil

### UnsetRecommendationInfo
`func (o *RecommendRequest) UnsetRecommendationInfo()`

UnsetRecommendationInfo ensures that no value is present for RecommendationInfo, not even an explicit nil
### GetHealth

`func (o *RecommendRequest) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *RecommendRequest) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *RecommendRequest) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *RecommendRequest) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### SetHealthNil

`func (o *RecommendRequest) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *RecommendRequest) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



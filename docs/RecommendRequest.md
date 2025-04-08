# RecommendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendationInfo** | Pointer to [**NullableRecommendationInfo**](RecommendationInfo.md) | Information about the user you wish to generate a recommendation for. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



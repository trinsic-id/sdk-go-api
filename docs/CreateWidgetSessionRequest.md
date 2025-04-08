# CreateWidgetSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **NullableString** | The URL to redirect the user to after the widget session is complete.              *Note*: this should NOT be set if you intend to use Trinsic&#39;s Web UI SDK to launch the Widget as an embedded iFrame or popup; in that case, session resolution is handled by our SDK, not via redirect. | [optional] 
**Providers** | Pointer to **[]string** | The list of allowed identity providers. If not specified, all available providers will be allowed. | [optional] 
**RecommendationInfo** | Pointer to [**NullableRecommendationInfo**](RecommendationInfo.md) | Data that you already know about the user being verified.   This data is used to improve the user experience during provider selection, by surfacing the most relevant providers first. | [optional] 

## Methods

### NewCreateWidgetSessionRequest

`func NewCreateWidgetSessionRequest() *CreateWidgetSessionRequest`

NewCreateWidgetSessionRequest instantiates a new CreateWidgetSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWidgetSessionRequestWithDefaults

`func NewCreateWidgetSessionRequestWithDefaults() *CreateWidgetSessionRequest`

NewCreateWidgetSessionRequestWithDefaults instantiates a new CreateWidgetSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *CreateWidgetSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateWidgetSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateWidgetSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateWidgetSessionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateWidgetSessionRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateWidgetSessionRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetProviders

`func (o *CreateWidgetSessionRequest) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CreateWidgetSessionRequest) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CreateWidgetSessionRequest) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CreateWidgetSessionRequest) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *CreateWidgetSessionRequest) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *CreateWidgetSessionRequest) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetRecommendationInfo

`func (o *CreateWidgetSessionRequest) GetRecommendationInfo() RecommendationInfo`

GetRecommendationInfo returns the RecommendationInfo field if non-nil, zero value otherwise.

### GetRecommendationInfoOk

`func (o *CreateWidgetSessionRequest) GetRecommendationInfoOk() (*RecommendationInfo, bool)`

GetRecommendationInfoOk returns a tuple with the RecommendationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationInfo

`func (o *CreateWidgetSessionRequest) SetRecommendationInfo(v RecommendationInfo)`

SetRecommendationInfo sets RecommendationInfo field to given value.

### HasRecommendationInfo

`func (o *CreateWidgetSessionRequest) HasRecommendationInfo() bool`

HasRecommendationInfo returns a boolean if a field has been set.

### SetRecommendationInfoNil

`func (o *CreateWidgetSessionRequest) SetRecommendationInfoNil(b bool)`

 SetRecommendationInfoNil sets the value for RecommendationInfo to be an explicit nil

### UnsetRecommendationInfo
`func (o *CreateWidgetSessionRequest) UnsetRecommendationInfo()`

UnsetRecommendationInfo ensures that no value is present for RecommendationInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



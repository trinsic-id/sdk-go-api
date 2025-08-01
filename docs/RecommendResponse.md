# RecommendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recognized** | [**[]ProviderInformation**](ProviderInformation.md) | The providers which the user is known to have a credential with. | 
**Relevant** | [**[]ProviderInformation**](ProviderInformation.md) | The providers which the user is deemed sufficiently likely to have a credential with. | 
**Remainder** | [**[]ProviderInformation**](ProviderInformation.md) | The providers which the user is unlikely to have a credential with.              This includes all providers available to your app which were not included in the Recognized or Relevant lists. | 

## Methods

### NewRecommendResponse

`func NewRecommendResponse(recognized []ProviderInformation, relevant []ProviderInformation, remainder []ProviderInformation, ) *RecommendResponse`

NewRecommendResponse instantiates a new RecommendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendResponseWithDefaults

`func NewRecommendResponseWithDefaults() *RecommendResponse`

NewRecommendResponseWithDefaults instantiates a new RecommendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecognized

`func (o *RecommendResponse) GetRecognized() []ProviderInformation`

GetRecognized returns the Recognized field if non-nil, zero value otherwise.

### GetRecognizedOk

`func (o *RecommendResponse) GetRecognizedOk() (*[]ProviderInformation, bool)`

GetRecognizedOk returns a tuple with the Recognized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecognized

`func (o *RecommendResponse) SetRecognized(v []ProviderInformation)`

SetRecognized sets Recognized field to given value.


### GetRelevant

`func (o *RecommendResponse) GetRelevant() []ProviderInformation`

GetRelevant returns the Relevant field if non-nil, zero value otherwise.

### GetRelevantOk

`func (o *RecommendResponse) GetRelevantOk() (*[]ProviderInformation, bool)`

GetRelevantOk returns a tuple with the Relevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevant

`func (o *RecommendResponse) SetRelevant(v []ProviderInformation)`

SetRelevant sets Relevant field to given value.


### GetRemainder

`func (o *RecommendResponse) GetRemainder() []ProviderInformation`

GetRemainder returns the Remainder field if non-nil, zero value otherwise.

### GetRemainderOk

`func (o *RecommendResponse) GetRemainderOk() (*[]ProviderInformation, bool)`

GetRemainderOk returns a tuple with the Remainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainder

`func (o *RecommendResponse) SetRemainder(v []ProviderInformation)`

SetRemainder sets Remainder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



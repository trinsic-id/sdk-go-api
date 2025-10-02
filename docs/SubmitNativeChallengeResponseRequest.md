# SubmitNativeChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultsAccessKey** | **string** | The Results Access Key for the Session in question | 
**ResponseToken** | **string** | The response token retrieved from a Trinsic UI SDK&#39;s performMdlExchange() call. | 

## Methods

### NewSubmitNativeChallengeResponseRequest

`func NewSubmitNativeChallengeResponseRequest(resultsAccessKey string, responseToken string, ) *SubmitNativeChallengeResponseRequest`

NewSubmitNativeChallengeResponseRequest instantiates a new SubmitNativeChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitNativeChallengeResponseRequestWithDefaults

`func NewSubmitNativeChallengeResponseRequestWithDefaults() *SubmitNativeChallengeResponseRequest`

NewSubmitNativeChallengeResponseRequestWithDefaults instantiates a new SubmitNativeChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultsAccessKey

`func (o *SubmitNativeChallengeResponseRequest) GetResultsAccessKey() string`

GetResultsAccessKey returns the ResultsAccessKey field if non-nil, zero value otherwise.

### GetResultsAccessKeyOk

`func (o *SubmitNativeChallengeResponseRequest) GetResultsAccessKeyOk() (*string, bool)`

GetResultsAccessKeyOk returns a tuple with the ResultsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAccessKey

`func (o *SubmitNativeChallengeResponseRequest) SetResultsAccessKey(v string)`

SetResultsAccessKey sets ResultsAccessKey field to given value.


### GetResponseToken

`func (o *SubmitNativeChallengeResponseRequest) GetResponseToken() string`

GetResponseToken returns the ResponseToken field if non-nil, zero value otherwise.

### GetResponseTokenOk

`func (o *SubmitNativeChallengeResponseRequest) GetResponseTokenOk() (*string, bool)`

GetResponseTokenOk returns a tuple with the ResponseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseToken

`func (o *SubmitNativeChallengeResponseRequest) SetResponseToken(v string)`

SetResponseToken sets ResponseToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



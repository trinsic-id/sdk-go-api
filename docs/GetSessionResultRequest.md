# GetSessionResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultsAccessKey** | **string** | The Results Access Key to exchange | 
**RedirectToken** | Pointer to **NullableString** | The &#x60;redirectToken&#x60; which Trinsic appended to your Redirect URL before sending the user back to your application.              This is only relevant for Sessions which involve a redirect flow.              If a &#x60;redirectToken&#x60; is required for a Session, it is only required for the first call to &#x60;GetSessionResult&#x60;. Future calls do not require a &#x60;redirectToken&#x60;. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`GetSessionResultRequest{FieldName: value}`), or call
`NewGetSessionResultRequestWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewGetSessionResultRequestWithDefaults

`func NewGetSessionResultRequestWithDefaults() *GetSessionResultRequest`

NewGetSessionResultRequestWithDefaults instantiates a new GetSessionResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultsAccessKey

`func (o *GetSessionResultRequest) GetResultsAccessKey() string`

GetResultsAccessKey returns the ResultsAccessKey field if non-nil, zero value otherwise.

### GetResultsAccessKeyOk

`func (o *GetSessionResultRequest) GetResultsAccessKeyOk() (*string, bool)`

GetResultsAccessKeyOk returns a tuple with the ResultsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAccessKey

`func (o *GetSessionResultRequest) SetResultsAccessKey(v string)`

SetResultsAccessKey sets ResultsAccessKey field to given value.


### GetRedirectToken

`func (o *GetSessionResultRequest) GetRedirectToken() string`

GetRedirectToken returns the RedirectToken field if non-nil, zero value otherwise.

### GetRedirectTokenOk

`func (o *GetSessionResultRequest) GetRedirectTokenOk() (*string, bool)`

GetRedirectTokenOk returns a tuple with the RedirectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToken

`func (o *GetSessionResultRequest) SetRedirectToken(v string)`

SetRedirectToken sets RedirectToken field to given value.

### HasRedirectToken

`func (o *GetSessionResultRequest) HasRedirectToken() bool`

HasRedirectToken returns a boolean if a field has been set.

### SetRedirectTokenNil

`func (o *GetSessionResultRequest) SetRedirectTokenNil(b bool)`

 SetRedirectTokenNil sets the value for RedirectToken to be an explicit nil

### UnsetRedirectToken
`func (o *GetSessionResultRequest) UnsetRedirectToken()`

UnsetRedirectToken ensures that no value is present for RedirectToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateHostedProviderSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | The ID of the newly-created Acceptance Session | 
**LaunchUrl** | **string** | The URL that should be used to launch the Hosted Provider Session on your user&#39;s device.              There are two recommended ways to do so: - Redirect their browser to the &#x60;LaunchUrl&#x60; as a top-level navigation - Open a popup window and navigate it to the &#x60;LaunchUrl&#x60;              This URL is sensitive and as such can only be obtained once. If you need to obtain it again, you will need to create a new Acceptance Session. | 
**ResultsAccessKey** | **string** | The &#x60;resultsAccessKey&#x60; for the Acceptance Session.              This is an encrypted payload which contains the decryption key necessary to access the Session&#39;s Data Vault.              Save this securely in your systems; it must be passed back with any API call which requires access to the Session&#39;s Data Vault.              Trinsic cannot access a Session&#39;s Data Vault without this key. | 

## Methods

### NewCreateHostedProviderSessionResponse

`func NewCreateHostedProviderSessionResponse(sessionId string, launchUrl string, resultsAccessKey string, ) *CreateHostedProviderSessionResponse`

NewCreateHostedProviderSessionResponse instantiates a new CreateHostedProviderSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHostedProviderSessionResponseWithDefaults

`func NewCreateHostedProviderSessionResponseWithDefaults() *CreateHostedProviderSessionResponse`

NewCreateHostedProviderSessionResponseWithDefaults instantiates a new CreateHostedProviderSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *CreateHostedProviderSessionResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CreateHostedProviderSessionResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CreateHostedProviderSessionResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetLaunchUrl

`func (o *CreateHostedProviderSessionResponse) GetLaunchUrl() string`

GetLaunchUrl returns the LaunchUrl field if non-nil, zero value otherwise.

### GetLaunchUrlOk

`func (o *CreateHostedProviderSessionResponse) GetLaunchUrlOk() (*string, bool)`

GetLaunchUrlOk returns a tuple with the LaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchUrl

`func (o *CreateHostedProviderSessionResponse) SetLaunchUrl(v string)`

SetLaunchUrl sets LaunchUrl field to given value.


### GetResultsAccessKey

`func (o *CreateHostedProviderSessionResponse) GetResultsAccessKey() string`

GetResultsAccessKey returns the ResultsAccessKey field if non-nil, zero value otherwise.

### GetResultsAccessKeyOk

`func (o *CreateHostedProviderSessionResponse) GetResultsAccessKeyOk() (*string, bool)`

GetResultsAccessKeyOk returns a tuple with the ResultsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAccessKey

`func (o *CreateHostedProviderSessionResponse) SetResultsAccessKey(v string)`

SetResultsAccessKey sets ResultsAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateWidgetSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | The ID of the newly-created Acceptance Session | 
**LaunchUrl** | Pointer to **NullableString** | The URL that should be used to invoke the Acceptance Session on your user&#39;s device.                You can use our frontend SDKs to launch the user into the Acceptance Session, or you can redirect the user&#39;s browser to this URL.                This URL is sensitive and as such can only be obtained once. If you need to obtain it again, you will need to create a new Acceptance Session. | [optional] 

## Methods

### NewCreateWidgetSessionResponse

`func NewCreateWidgetSessionResponse(sessionId string, ) *CreateWidgetSessionResponse`

NewCreateWidgetSessionResponse instantiates a new CreateWidgetSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWidgetSessionResponseWithDefaults

`func NewCreateWidgetSessionResponseWithDefaults() *CreateWidgetSessionResponse`

NewCreateWidgetSessionResponseWithDefaults instantiates a new CreateWidgetSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *CreateWidgetSessionResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CreateWidgetSessionResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CreateWidgetSessionResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetLaunchUrl

`func (o *CreateWidgetSessionResponse) GetLaunchUrl() string`

GetLaunchUrl returns the LaunchUrl field if non-nil, zero value otherwise.

### GetLaunchUrlOk

`func (o *CreateWidgetSessionResponse) GetLaunchUrlOk() (*string, bool)`

GetLaunchUrlOk returns a tuple with the LaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchUrl

`func (o *CreateWidgetSessionResponse) SetLaunchUrl(v string)`

SetLaunchUrl sets LaunchUrl field to given value.

### HasLaunchUrl

`func (o *CreateWidgetSessionResponse) HasLaunchUrl() bool`

HasLaunchUrl returns a boolean if a field has been set.

### SetLaunchUrlNil

`func (o *CreateWidgetSessionResponse) SetLaunchUrlNil(b bool)`

 SetLaunchUrlNil sets the value for LaunchUrl to be an explicit nil

### UnsetLaunchUrl
`func (o *CreateWidgetSessionResponse) UnsetLaunchUrl()`

UnsetLaunchUrl ensures that no value is present for LaunchUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



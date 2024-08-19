# CreateSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) | The created Acceptance Session | 
**LaunchUrl** | Pointer to **string** | The URL that should be used to invoke the Acceptance Session on your user&#39;s device.                If the Session was created with &#x60;LaunchMethodDirectly&#x60; set to &#x60;true&#x60;, you should redirect your user&#39;s browser to this URL. The frontend SDK cannot presently be used to  invoke these Sessions.                Otherwise, you should pass this URL to your user&#39;s frontend and use the frontend SDK to invoke the Session.                This URL is sensitive and as such can only be obtained once. If you need to obtain it again, you will need to create a new Acceptance Session. | [optional] 

## Methods

### NewCreateSessionResponse

`func NewCreateSessionResponse(session Session, ) *CreateSessionResponse`

NewCreateSessionResponse instantiates a new CreateSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionResponseWithDefaults

`func NewCreateSessionResponseWithDefaults() *CreateSessionResponse`

NewCreateSessionResponseWithDefaults instantiates a new CreateSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *CreateSessionResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CreateSessionResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CreateSessionResponse) SetSession(v Session)`

SetSession sets Session field to given value.


### GetLaunchUrl

`func (o *CreateSessionResponse) GetLaunchUrl() string`

GetLaunchUrl returns the LaunchUrl field if non-nil, zero value otherwise.

### GetLaunchUrlOk

`func (o *CreateSessionResponse) GetLaunchUrlOk() (*string, bool)`

GetLaunchUrlOk returns a tuple with the LaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchUrl

`func (o *CreateSessionResponse) SetLaunchUrl(v string)`

SetLaunchUrl sets LaunchUrl field to given value.

### HasLaunchUrl

`func (o *CreateSessionResponse) HasLaunchUrl() bool`

HasLaunchUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



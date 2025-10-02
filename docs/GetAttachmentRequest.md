# GetAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentAccessKey** | **string** | The Attachment Access Key to exchange for the raw file contents of the related Attachment | 
**SessionId** | **string** | The ID of the Acceptance Session for which the Attachment is being requested. | 

## Methods

### NewGetAttachmentRequest

`func NewGetAttachmentRequest(attachmentAccessKey string, sessionId string, ) *GetAttachmentRequest`

NewGetAttachmentRequest instantiates a new GetAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttachmentRequestWithDefaults

`func NewGetAttachmentRequestWithDefaults() *GetAttachmentRequest`

NewGetAttachmentRequestWithDefaults instantiates a new GetAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentAccessKey

`func (o *GetAttachmentRequest) GetAttachmentAccessKey() string`

GetAttachmentAccessKey returns the AttachmentAccessKey field if non-nil, zero value otherwise.

### GetAttachmentAccessKeyOk

`func (o *GetAttachmentRequest) GetAttachmentAccessKeyOk() (*string, bool)`

GetAttachmentAccessKeyOk returns a tuple with the AttachmentAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentAccessKey

`func (o *GetAttachmentRequest) SetAttachmentAccessKey(v string)`

SetAttachmentAccessKey sets AttachmentAccessKey field to given value.


### GetSessionId

`func (o *GetAttachmentRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GetAttachmentRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GetAttachmentRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



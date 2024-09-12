# GetAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The raw file contents of the Attachment | 
**ContentType** | **string** | The MIME type of the Attachment data | 

## Methods

### NewGetAttachmentResponse

`func NewGetAttachmentResponse(content string, contentType string, ) *GetAttachmentResponse`

NewGetAttachmentResponse instantiates a new GetAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttachmentResponseWithDefaults

`func NewGetAttachmentResponseWithDefaults() *GetAttachmentResponse`

NewGetAttachmentResponseWithDefaults instantiates a new GetAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GetAttachmentResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetAttachmentResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetAttachmentResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *GetAttachmentResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetAttachmentResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetAttachmentResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



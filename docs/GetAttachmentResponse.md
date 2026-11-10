# GetAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The file contents of the Attachment. When using an SDK, this can be a raw byte array. In raw JSON form this is Base64. Combine this with the Content type to reconstruct the file.              This is empty when &#x60;encryptedContent&#x60; is present. | 
**EncryptedContent** | Pointer to [**NullableTrinsicHPKEBundle**](TrinsicHPKEBundle.md) | The HPKE Bundle containing the encrypted contents of the Attachment, if secured via HPKE.              Decrypt it using the HPKE info string &#x60;trinsic-hpke:1:attachment-contents:{sessionId}:{attachmentId}&#x60;. | [optional] 
**ContentType** | **string** | The MIME type of the Attachment data | 

## Methods

### Named construction

Use a keyed struct literal (`GetAttachmentResponse{FieldName: value}`), or call
`NewGetAttachmentResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

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


### GetEncryptedContent

`func (o *GetAttachmentResponse) GetEncryptedContent() TrinsicHPKEBundle`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *GetAttachmentResponse) GetEncryptedContentOk() (*TrinsicHPKEBundle, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *GetAttachmentResponse) SetEncryptedContent(v TrinsicHPKEBundle)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *GetAttachmentResponse) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *GetAttachmentResponse) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *GetAttachmentResponse) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil
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



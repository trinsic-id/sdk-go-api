# AttachmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the Attachment.              Pass this to the &#x60;GetAttachment&#x60; endpoint, along with the Session ID and the &#x60;resultsAccessKey&#x60; for said Session. | 
**Type** | **string** | The type of the Attachment.              Possible values: - \&quot;selfie\&quot; - \&quot;document_front\&quot; - \&quot;document_back\&quot; - \&quot;document_portrait\&quot; - \&quot;document_signature\&quot; - \&quot;document_scan_report\&quot;              Additional attachment types may be defined for specific Providers. | 
**ContentType** | **string** | The MIME type of the attachment&#39;s contents; eg, \&quot;image/jpeg\&quot; or \&quot;application/pdf\&quot;. | 
**SizeBytes** | **int32** | The size in bytes of the attachment. | 
**HpkeEncrypted** | Pointer to **NullableBool** | Whether the attachment contents are encrypted via HPKE and must be decrypted using your private key. | [optional] 

## Methods

### NewAttachmentInfo

`func NewAttachmentInfo(id string, type_ string, contentType string, sizeBytes int32, ) *AttachmentInfo`

NewAttachmentInfo instantiates a new AttachmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentInfoWithDefaults

`func NewAttachmentInfoWithDefaults() *AttachmentInfo`

NewAttachmentInfoWithDefaults instantiates a new AttachmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentInfo) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AttachmentInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentInfo) SetType(v string)`

SetType sets Type field to given value.


### GetContentType

`func (o *AttachmentInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AttachmentInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AttachmentInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSizeBytes

`func (o *AttachmentInfo) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *AttachmentInfo) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *AttachmentInfo) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetHpkeEncrypted

`func (o *AttachmentInfo) GetHpkeEncrypted() bool`

GetHpkeEncrypted returns the HpkeEncrypted field if non-nil, zero value otherwise.

### GetHpkeEncryptedOk

`func (o *AttachmentInfo) GetHpkeEncryptedOk() (*bool, bool)`

GetHpkeEncryptedOk returns a tuple with the HpkeEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpkeEncrypted

`func (o *AttachmentInfo) SetHpkeEncrypted(v bool)`

SetHpkeEncrypted sets HpkeEncrypted field to given value.

### HasHpkeEncrypted

`func (o *AttachmentInfo) HasHpkeEncrypted() bool`

HasHpkeEncrypted returns a boolean if a field has been set.

### SetHpkeEncryptedNil

`func (o *AttachmentInfo) SetHpkeEncryptedNil(b bool)`

 SetHpkeEncryptedNil sets the value for HpkeEncrypted to be an explicit nil

### UnsetHpkeEncrypted
`func (o *AttachmentInfo) UnsetHpkeEncrypted()`

UnsetHpkeEncrypted ensures that no value is present for HpkeEncrypted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AttachmentAccessKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selfie** | Pointer to **string** | Key to access the selfie image (if relevant) for this verification | [optional] 
**DocumentFront** | Pointer to **string** | Key to access the document front image (if relevant) for this verification | [optional] 
**DocumentBack** | Pointer to **string** | Key to access the document back image (if relevant) for this verification | [optional] 
**DocumentPortrait** | Pointer to **string** | Key to access the document portrait image (if relevant and available) for this verification.                Specifically, this is a cropped version of the document front image which includes only the portrait on the document. | [optional] 

## Methods

### NewAttachmentAccessKeys

`func NewAttachmentAccessKeys() *AttachmentAccessKeys`

NewAttachmentAccessKeys instantiates a new AttachmentAccessKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentAccessKeysWithDefaults

`func NewAttachmentAccessKeysWithDefaults() *AttachmentAccessKeys`

NewAttachmentAccessKeysWithDefaults instantiates a new AttachmentAccessKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelfie

`func (o *AttachmentAccessKeys) GetSelfie() string`

GetSelfie returns the Selfie field if non-nil, zero value otherwise.

### GetSelfieOk

`func (o *AttachmentAccessKeys) GetSelfieOk() (*string, bool)`

GetSelfieOk returns a tuple with the Selfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfie

`func (o *AttachmentAccessKeys) SetSelfie(v string)`

SetSelfie sets Selfie field to given value.

### HasSelfie

`func (o *AttachmentAccessKeys) HasSelfie() bool`

HasSelfie returns a boolean if a field has been set.

### GetDocumentFront

`func (o *AttachmentAccessKeys) GetDocumentFront() string`

GetDocumentFront returns the DocumentFront field if non-nil, zero value otherwise.

### GetDocumentFrontOk

`func (o *AttachmentAccessKeys) GetDocumentFrontOk() (*string, bool)`

GetDocumentFrontOk returns a tuple with the DocumentFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFront

`func (o *AttachmentAccessKeys) SetDocumentFront(v string)`

SetDocumentFront sets DocumentFront field to given value.

### HasDocumentFront

`func (o *AttachmentAccessKeys) HasDocumentFront() bool`

HasDocumentFront returns a boolean if a field has been set.

### GetDocumentBack

`func (o *AttachmentAccessKeys) GetDocumentBack() string`

GetDocumentBack returns the DocumentBack field if non-nil, zero value otherwise.

### GetDocumentBackOk

`func (o *AttachmentAccessKeys) GetDocumentBackOk() (*string, bool)`

GetDocumentBackOk returns a tuple with the DocumentBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBack

`func (o *AttachmentAccessKeys) SetDocumentBack(v string)`

SetDocumentBack sets DocumentBack field to given value.

### HasDocumentBack

`func (o *AttachmentAccessKeys) HasDocumentBack() bool`

HasDocumentBack returns a boolean if a field has been set.

### GetDocumentPortrait

`func (o *AttachmentAccessKeys) GetDocumentPortrait() string`

GetDocumentPortrait returns the DocumentPortrait field if non-nil, zero value otherwise.

### GetDocumentPortraitOk

`func (o *AttachmentAccessKeys) GetDocumentPortraitOk() (*string, bool)`

GetDocumentPortraitOk returns a tuple with the DocumentPortrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPortrait

`func (o *AttachmentAccessKeys) SetDocumentPortrait(v string)`

SetDocumentPortrait sets DocumentPortrait field to given value.

### HasDocumentPortrait

`func (o *AttachmentAccessKeys) HasDocumentPortrait() bool`

HasDocumentPortrait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



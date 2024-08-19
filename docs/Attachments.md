# Attachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selfie** | Pointer to **string** | Key to access the selfie image (if relevant) for this verification | [optional] 
**DocumentFront** | Pointer to **string** | Key to access the document front image (if relevant) for this verification | [optional] 
**DocumentBack** | Pointer to **string** | Key to access the document back image (if relevant) for this verification | [optional] 
**DocumentPortrait** | Pointer to **string** | Key to access the document portrait image (if relevant and available) for this verification.                Specifically, this is a cropped version of the document front image which includes only the portrait on the document. | [optional] 

## Methods

### NewAttachments

`func NewAttachments() *Attachments`

NewAttachments instantiates a new Attachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsWithDefaults

`func NewAttachmentsWithDefaults() *Attachments`

NewAttachmentsWithDefaults instantiates a new Attachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelfie

`func (o *Attachments) GetSelfie() string`

GetSelfie returns the Selfie field if non-nil, zero value otherwise.

### GetSelfieOk

`func (o *Attachments) GetSelfieOk() (*string, bool)`

GetSelfieOk returns a tuple with the Selfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfie

`func (o *Attachments) SetSelfie(v string)`

SetSelfie sets Selfie field to given value.

### HasSelfie

`func (o *Attachments) HasSelfie() bool`

HasSelfie returns a boolean if a field has been set.

### GetDocumentFront

`func (o *Attachments) GetDocumentFront() string`

GetDocumentFront returns the DocumentFront field if non-nil, zero value otherwise.

### GetDocumentFrontOk

`func (o *Attachments) GetDocumentFrontOk() (*string, bool)`

GetDocumentFrontOk returns a tuple with the DocumentFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFront

`func (o *Attachments) SetDocumentFront(v string)`

SetDocumentFront sets DocumentFront field to given value.

### HasDocumentFront

`func (o *Attachments) HasDocumentFront() bool`

HasDocumentFront returns a boolean if a field has been set.

### GetDocumentBack

`func (o *Attachments) GetDocumentBack() string`

GetDocumentBack returns the DocumentBack field if non-nil, zero value otherwise.

### GetDocumentBackOk

`func (o *Attachments) GetDocumentBackOk() (*string, bool)`

GetDocumentBackOk returns a tuple with the DocumentBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBack

`func (o *Attachments) SetDocumentBack(v string)`

SetDocumentBack sets DocumentBack field to given value.

### HasDocumentBack

`func (o *Attachments) HasDocumentBack() bool`

HasDocumentBack returns a boolean if a field has been set.

### GetDocumentPortrait

`func (o *Attachments) GetDocumentPortrait() string`

GetDocumentPortrait returns the DocumentPortrait field if non-nil, zero value otherwise.

### GetDocumentPortraitOk

`func (o *Attachments) GetDocumentPortraitOk() (*string, bool)`

GetDocumentPortraitOk returns a tuple with the DocumentPortrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPortrait

`func (o *Attachments) SetDocumentPortrait(v string)`

SetDocumentPortrait sets DocumentPortrait field to given value.

### HasDocumentPortrait

`func (o *Attachments) HasDocumentPortrait() bool`

HasDocumentPortrait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



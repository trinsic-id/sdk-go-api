# IdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatingProviderId** | Pointer to **string** | The ID of the integration from which this data originated (eg \&quot;yoti\&quot;, \&quot;clear\&quot;) | [optional] 
**Person** | Pointer to [**PersonData**](PersonData.md) | Identity data of the individual who was verified | [optional] 
**Document** | Pointer to [**DocumentData**](DocumentData.md) | Identity data of the document involved in verification, if relevant | [optional] 
**Attachments** | Pointer to [**Attachments**](Attachments.md) | Access keys for attachments (eg document/selfie images) | [optional] 

## Methods

### NewIdentityData

`func NewIdentityData() *IdentityData`

NewIdentityData instantiates a new IdentityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDataWithDefaults

`func NewIdentityDataWithDefaults() *IdentityData`

NewIdentityDataWithDefaults instantiates a new IdentityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginatingProviderId

`func (o *IdentityData) GetOriginatingProviderId() string`

GetOriginatingProviderId returns the OriginatingProviderId field if non-nil, zero value otherwise.

### GetOriginatingProviderIdOk

`func (o *IdentityData) GetOriginatingProviderIdOk() (*string, bool)`

GetOriginatingProviderIdOk returns a tuple with the OriginatingProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingProviderId

`func (o *IdentityData) SetOriginatingProviderId(v string)`

SetOriginatingProviderId sets OriginatingProviderId field to given value.

### HasOriginatingProviderId

`func (o *IdentityData) HasOriginatingProviderId() bool`

HasOriginatingProviderId returns a boolean if a field has been set.

### GetPerson

`func (o *IdentityData) GetPerson() PersonData`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *IdentityData) GetPersonOk() (*PersonData, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *IdentityData) SetPerson(v PersonData)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *IdentityData) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetDocument

`func (o *IdentityData) GetDocument() DocumentData`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *IdentityData) GetDocumentOk() (*DocumentData, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *IdentityData) SetDocument(v DocumentData)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *IdentityData) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetAttachments

`func (o *IdentityData) GetAttachments() Attachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IdentityData) GetAttachmentsOk() (*Attachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IdentityData) SetAttachments(v Attachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IdentityData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



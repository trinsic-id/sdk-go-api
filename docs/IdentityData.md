# IdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatingProviderId** | Pointer to **NullableString** |  | [optional] 
**OriginatingSubProviderId** | Pointer to **NullableString** |  | [optional] 
**Person** | Pointer to [**NullablePersonData**](PersonData.md) |  | [optional] 
**Document** | Pointer to [**NullableDocumentData**](DocumentData.md) |  | [optional] 
**AttachmentAccessKeys** | Pointer to [**NullableAttachmentAccessKeys**](AttachmentAccessKeys.md) |  | [optional] 

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

### SetOriginatingProviderIdNil

`func (o *IdentityData) SetOriginatingProviderIdNil(b bool)`

 SetOriginatingProviderIdNil sets the value for OriginatingProviderId to be an explicit nil

### UnsetOriginatingProviderId
`func (o *IdentityData) UnsetOriginatingProviderId()`

UnsetOriginatingProviderId ensures that no value is present for OriginatingProviderId, not even an explicit nil
### GetOriginatingSubProviderId

`func (o *IdentityData) GetOriginatingSubProviderId() string`

GetOriginatingSubProviderId returns the OriginatingSubProviderId field if non-nil, zero value otherwise.

### GetOriginatingSubProviderIdOk

`func (o *IdentityData) GetOriginatingSubProviderIdOk() (*string, bool)`

GetOriginatingSubProviderIdOk returns a tuple with the OriginatingSubProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingSubProviderId

`func (o *IdentityData) SetOriginatingSubProviderId(v string)`

SetOriginatingSubProviderId sets OriginatingSubProviderId field to given value.

### HasOriginatingSubProviderId

`func (o *IdentityData) HasOriginatingSubProviderId() bool`

HasOriginatingSubProviderId returns a boolean if a field has been set.

### SetOriginatingSubProviderIdNil

`func (o *IdentityData) SetOriginatingSubProviderIdNil(b bool)`

 SetOriginatingSubProviderIdNil sets the value for OriginatingSubProviderId to be an explicit nil

### UnsetOriginatingSubProviderId
`func (o *IdentityData) UnsetOriginatingSubProviderId()`

UnsetOriginatingSubProviderId ensures that no value is present for OriginatingSubProviderId, not even an explicit nil
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

### SetPersonNil

`func (o *IdentityData) SetPersonNil(b bool)`

 SetPersonNil sets the value for Person to be an explicit nil

### UnsetPerson
`func (o *IdentityData) UnsetPerson()`

UnsetPerson ensures that no value is present for Person, not even an explicit nil
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

### SetDocumentNil

`func (o *IdentityData) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *IdentityData) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetAttachmentAccessKeys

`func (o *IdentityData) GetAttachmentAccessKeys() AttachmentAccessKeys`

GetAttachmentAccessKeys returns the AttachmentAccessKeys field if non-nil, zero value otherwise.

### GetAttachmentAccessKeysOk

`func (o *IdentityData) GetAttachmentAccessKeysOk() (*AttachmentAccessKeys, bool)`

GetAttachmentAccessKeysOk returns a tuple with the AttachmentAccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentAccessKeys

`func (o *IdentityData) SetAttachmentAccessKeys(v AttachmentAccessKeys)`

SetAttachmentAccessKeys sets AttachmentAccessKeys field to given value.

### HasAttachmentAccessKeys

`func (o *IdentityData) HasAttachmentAccessKeys() bool`

HasAttachmentAccessKeys returns a boolean if a field has been set.

### SetAttachmentAccessKeysNil

`func (o *IdentityData) SetAttachmentAccessKeysNil(b bool)`

 SetAttachmentAccessKeysNil sets the value for AttachmentAccessKeys to be an explicit nil

### UnsetAttachmentAccessKeys
`func (o *IdentityData) UnsetAttachmentAccessKeys()`

UnsetAttachmentAccessKeys ensures that no value is present for AttachmentAccessKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



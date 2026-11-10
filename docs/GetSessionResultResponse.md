# GetSessionResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) | The Session in its current state. | 
**IdentityData** | Pointer to [**NullableIdentityData**](IdentityData.md) | The IdentityData which resulted from the Session, if not HPKE-encrypted. | [optional] 
**EncryptedIdentityData** | Pointer to [**NullableTrinsicHPKEBundle**](TrinsicHPKEBundle.md) | The HPKE-encrypted IdentityData which resulted from the Session, if HPKE-encrypted. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`GetSessionResultResponse{FieldName: value}`), or call
`NewGetSessionResultResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewGetSessionResultResponseWithDefaults

`func NewGetSessionResultResponseWithDefaults() *GetSessionResultResponse`

NewGetSessionResultResponseWithDefaults instantiates a new GetSessionResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *GetSessionResultResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetSessionResultResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetSessionResultResponse) SetSession(v Session)`

SetSession sets Session field to given value.


### GetIdentityData

`func (o *GetSessionResultResponse) GetIdentityData() IdentityData`

GetIdentityData returns the IdentityData field if non-nil, zero value otherwise.

### GetIdentityDataOk

`func (o *GetSessionResultResponse) GetIdentityDataOk() (*IdentityData, bool)`

GetIdentityDataOk returns a tuple with the IdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityData

`func (o *GetSessionResultResponse) SetIdentityData(v IdentityData)`

SetIdentityData sets IdentityData field to given value.

### HasIdentityData

`func (o *GetSessionResultResponse) HasIdentityData() bool`

HasIdentityData returns a boolean if a field has been set.

### SetIdentityDataNil

`func (o *GetSessionResultResponse) SetIdentityDataNil(b bool)`

 SetIdentityDataNil sets the value for IdentityData to be an explicit nil

### UnsetIdentityData
`func (o *GetSessionResultResponse) UnsetIdentityData()`

UnsetIdentityData ensures that no value is present for IdentityData, not even an explicit nil
### GetEncryptedIdentityData

`func (o *GetSessionResultResponse) GetEncryptedIdentityData() TrinsicHPKEBundle`

GetEncryptedIdentityData returns the EncryptedIdentityData field if non-nil, zero value otherwise.

### GetEncryptedIdentityDataOk

`func (o *GetSessionResultResponse) GetEncryptedIdentityDataOk() (*TrinsicHPKEBundle, bool)`

GetEncryptedIdentityDataOk returns a tuple with the EncryptedIdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedIdentityData

`func (o *GetSessionResultResponse) SetEncryptedIdentityData(v TrinsicHPKEBundle)`

SetEncryptedIdentityData sets EncryptedIdentityData field to given value.

### HasEncryptedIdentityData

`func (o *GetSessionResultResponse) HasEncryptedIdentityData() bool`

HasEncryptedIdentityData returns a boolean if a field has been set.

### SetEncryptedIdentityDataNil

`func (o *GetSessionResultResponse) SetEncryptedIdentityDataNil(b bool)`

 SetEncryptedIdentityDataNil sets the value for EncryptedIdentityData to be an explicit nil

### UnsetEncryptedIdentityData
`func (o *GetSessionResultResponse) UnsetEncryptedIdentityData()`

UnsetEncryptedIdentityData ensures that no value is present for EncryptedIdentityData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



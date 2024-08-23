# GetSessionResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) |  | 
**IdentityData** | Pointer to [**IdentityData**](IdentityData.md) |  | [optional] 

## Methods

### NewGetSessionResultResponse

`func NewGetSessionResultResponse(session Session, ) *GetSessionResultResponse`

NewGetSessionResultResponse instantiates a new GetSessionResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



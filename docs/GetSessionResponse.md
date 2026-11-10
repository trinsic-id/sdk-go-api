# GetSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) |  | 

## Methods

### Named construction

Use a keyed struct literal (`GetSessionResponse{FieldName: value}`), or call
`NewGetSessionResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewGetSessionResponseWithDefaults

`func NewGetSessionResponseWithDefaults() *GetSessionResponse`

NewGetSessionResponseWithDefaults instantiates a new GetSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *GetSessionResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetSessionResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetSessionResponse) SetSession(v Session)`

SetSession sets Session field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



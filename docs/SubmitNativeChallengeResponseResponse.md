# SubmitNativeChallengeResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) | The Session for which the challenge token was submitted | 

## Methods

### Named construction

Use a keyed struct literal (`SubmitNativeChallengeResponseResponse{FieldName: value}`), or call
`NewSubmitNativeChallengeResponseResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewSubmitNativeChallengeResponseResponseWithDefaults

`func NewSubmitNativeChallengeResponseResponseWithDefaults() *SubmitNativeChallengeResponseResponse`

NewSubmitNativeChallengeResponseResponseWithDefaults instantiates a new SubmitNativeChallengeResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *SubmitNativeChallengeResponseResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SubmitNativeChallengeResponseResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SubmitNativeChallengeResponseResponse) SetSession(v Session)`

SetSession sets Session field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



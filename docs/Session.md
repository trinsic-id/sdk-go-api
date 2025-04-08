# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Done** | **bool** | Whether the Session is in a terminal / final state.              If this is &#x60;true&#x60;, inspect the value of &#x60;Success&#x60; to determine whether the Session was successful. | 
**Success** | **bool** | Whether the Session has completed successfully.              If this is &#x60;false&#x60;, the Session is either not yet done, or has failed. Inspect &#x60;Done&#x60; and &#x60;ErrorCode&#x60; for more information. If this is &#x60;true&#x60;, the Session has completed successfully. | 
**ErrorCode** | Pointer to [**NullableSessionErrorCode**](SessionErrorCode.md) | The reason for the Session&#39;s failure.              Only present if &#x60;Success&#x60; is &#x60;false&#x60;. | [optional] 
**Created** | **int64** | The unix timestamp, in seconds, when this session was created | 
**Updated** | **int64** | The unix timestamp, in seconds, when this session&#39;s state last changed | 

## Methods

### NewSession

`func NewSession(id string, done bool, success bool, created int64, updated int64, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetDone

`func (o *Session) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *Session) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *Session) SetDone(v bool)`

SetDone sets Done field to given value.


### GetSuccess

`func (o *Session) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Session) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Session) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrorCode

`func (o *Session) GetErrorCode() SessionErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Session) GetErrorCodeOk() (*SessionErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Session) SetErrorCode(v SessionErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Session) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *Session) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *Session) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetCreated

`func (o *Session) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Session) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Session) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Session) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Session) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Session) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



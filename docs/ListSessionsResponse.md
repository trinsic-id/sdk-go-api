# ListSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | [**[]Session**](Session.md) |  | 
**Total** | **int32** | The total number of sessions tied to your account | 
**More** | **bool** | Whether there are additional pages of sessions to retrieve | 

## Methods

### NewListSessionsResponse

`func NewListSessionsResponse(sessions []Session, total int32, more bool, ) *ListSessionsResponse`

NewListSessionsResponse instantiates a new ListSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSessionsResponseWithDefaults

`func NewListSessionsResponseWithDefaults() *ListSessionsResponse`

NewListSessionsResponseWithDefaults instantiates a new ListSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ListSessionsResponse) GetSessions() []Session`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListSessionsResponse) GetSessionsOk() (*[]Session, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListSessionsResponse) SetSessions(v []Session)`

SetSessions sets Sessions field to given value.


### GetTotal

`func (o *ListSessionsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListSessionsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListSessionsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMore

`func (o *ListSessionsResponse) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *ListSessionsResponse) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *ListSessionsResponse) SetMore(v bool)`

SetMore sets More field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



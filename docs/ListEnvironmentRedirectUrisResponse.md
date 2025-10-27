# ListEnvironmentRedirectUrisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uris** | [**[]RedirectUriResponse**](RedirectUriResponse.md) | List of redirect uris within the environment. | 
**More** | **bool** | Whether there are additional pages of uris to retrieve | 

## Methods

### NewListEnvironmentRedirectUrisResponse

`func NewListEnvironmentRedirectUrisResponse(uris []RedirectUriResponse, more bool, ) *ListEnvironmentRedirectUrisResponse`

NewListEnvironmentRedirectUrisResponse instantiates a new ListEnvironmentRedirectUrisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentRedirectUrisResponseWithDefaults

`func NewListEnvironmentRedirectUrisResponseWithDefaults() *ListEnvironmentRedirectUrisResponse`

NewListEnvironmentRedirectUrisResponseWithDefaults instantiates a new ListEnvironmentRedirectUrisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUris

`func (o *ListEnvironmentRedirectUrisResponse) GetUris() []RedirectUriResponse`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *ListEnvironmentRedirectUrisResponse) GetUrisOk() (*[]RedirectUriResponse, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *ListEnvironmentRedirectUrisResponse) SetUris(v []RedirectUriResponse)`

SetUris sets Uris field to given value.


### GetMore

`func (o *ListEnvironmentRedirectUrisResponse) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *ListEnvironmentRedirectUrisResponse) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *ListEnvironmentRedirectUrisResponse) SetMore(v bool)`

SetMore sets More field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



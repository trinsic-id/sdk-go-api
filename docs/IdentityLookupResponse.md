# IdentityLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityInNetwork** | **bool** | Whether the given phone number is known to have an identity in the network. | 

## Methods

### NewIdentityLookupResponse

`func NewIdentityLookupResponse(identityInNetwork bool, ) *IdentityLookupResponse`

NewIdentityLookupResponse instantiates a new IdentityLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLookupResponseWithDefaults

`func NewIdentityLookupResponseWithDefaults() *IdentityLookupResponse`

NewIdentityLookupResponseWithDefaults instantiates a new IdentityLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityInNetwork

`func (o *IdentityLookupResponse) GetIdentityInNetwork() bool`

GetIdentityInNetwork returns the IdentityInNetwork field if non-nil, zero value otherwise.

### GetIdentityInNetworkOk

`func (o *IdentityLookupResponse) GetIdentityInNetworkOk() (*bool, bool)`

GetIdentityInNetworkOk returns a tuple with the IdentityInNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityInNetwork

`func (o *IdentityLookupResponse) SetIdentityInNetwork(v bool)`

SetIdentityInNetwork sets IdentityInNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



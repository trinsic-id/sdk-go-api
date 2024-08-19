# ListProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | [**[]ProviderInfo**](ProviderInfo.md) | The list of identity providers available to your account | 

## Methods

### NewListProvidersResponse

`func NewListProvidersResponse(providers []ProviderInfo, ) *ListProvidersResponse`

NewListProvidersResponse instantiates a new ListProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvidersResponseWithDefaults

`func NewListProvidersResponseWithDefaults() *ListProvidersResponse`

NewListProvidersResponseWithDefaults instantiates a new ListProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ListProvidersResponse) GetProviders() []ProviderInfo`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ListProvidersResponse) GetProvidersOk() (*[]ProviderInfo, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ListProvidersResponse) SetProviders(v []ProviderInfo)`

SetProviders sets Providers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



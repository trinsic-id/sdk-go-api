# ListProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | [**[]Provider**](Provider.md) | The list of providers available to this organization | 

## Methods

### Named construction

Use a keyed struct literal (`ListProvidersResponse{FieldName: value}`), or call
`NewListProvidersResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewListProvidersResponseWithDefaults

`func NewListProvidersResponseWithDefaults() *ListProvidersResponse`

NewListProvidersResponseWithDefaults instantiates a new ListProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ListProvidersResponse) GetProviders() []Provider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ListProvidersResponse) GetProvidersOk() (*[]Provider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ListProvidersResponse) SetProviders(v []Provider)`

SetProviders sets Providers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



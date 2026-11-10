# GetProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**Provider**](Provider.md) | The provider available to this organization | 

## Methods

### Named construction

Use a keyed struct literal (`GetProviderResponse{FieldName: value}`), or call
`NewGetProviderResponseWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewGetProviderResponseWithDefaults

`func NewGetProviderResponseWithDefaults() *GetProviderResponse`

NewGetProviderResponseWithDefaults instantiates a new GetProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *GetProviderResponse) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetProviderResponse) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetProviderResponse) SetProvider(v Provider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



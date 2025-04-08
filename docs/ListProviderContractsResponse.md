# ListProviderContractsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderContracts** | [**[]ProviderContract**](ProviderContract.md) | Contracts for all Providers available to your App, filtered by your App&#39;s current test mode setting. | 

## Methods

### NewListProviderContractsResponse

`func NewListProviderContractsResponse(providerContracts []ProviderContract, ) *ListProviderContractsResponse`

NewListProviderContractsResponse instantiates a new ListProviderContractsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProviderContractsResponseWithDefaults

`func NewListProviderContractsResponseWithDefaults() *ListProviderContractsResponse`

NewListProviderContractsResponseWithDefaults instantiates a new ListProviderContractsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderContracts

`func (o *ListProviderContractsResponse) GetProviderContracts() []ProviderContract`

GetProviderContracts returns the ProviderContracts field if non-nil, zero value otherwise.

### GetProviderContractsOk

`func (o *ListProviderContractsResponse) GetProviderContractsOk() (*[]ProviderContract, bool)`

GetProviderContractsOk returns a tuple with the ProviderContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderContracts

`func (o *ListProviderContractsResponse) SetProviderContracts(v []ProviderContract)`

SetProviderContracts sets ProviderContracts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



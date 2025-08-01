# \NetworkAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProviderContracts**](NetworkAPI.md#ListProviderContracts) | **Get** /api/v1/network/providers/contracts | List Provider Contracts
[**ListProviders**](NetworkAPI.md#ListProviders) | **Get** /api/v1/network/providers | List Providers
[**RecommendProviders**](NetworkAPI.md#RecommendProviders) | **Post** /api/v1/network/recommend | Recommend Providers



## ListProviderContracts

> ListProviderContractsResponse ListProviderContracts(ctx).Execute()

List Provider Contracts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.ListProviderContracts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.ListProviderContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviderContracts`: ListProviderContractsResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.ListProviderContracts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProviderContractsRequest struct via the builder pattern


### Return type

[**ListProviderContractsResponse**](ListProviderContractsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> ListProvidersResponse ListProviders(ctx).Health(health).Execute()

List Providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	health := "health_example" // string | Filter providers by health status. Valid values: \"online\", \"offline\", \"all\". Defaults to \"all\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.ListProviders(context.Background()).Health(health).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: ListProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **health** | **string** | Filter providers by health status. Valid values: \&quot;online\&quot;, \&quot;offline\&quot;, \&quot;all\&quot;. Defaults to \&quot;all\&quot;. | 

### Return type

[**ListProvidersResponse**](ListProvidersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendProviders

> RecommendResponse RecommendProviders(ctx).RecommendRequest(recommendRequest).Execute()

Recommend Providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	recommendRequest := *openapiclient.NewRecommendRequest() // RecommendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.RecommendProviders(context.Background()).RecommendRequest(recommendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.RecommendProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendProviders`: RecommendResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.RecommendProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendRequest** | [**RecommendRequest**](RecommendRequest.md) |  | 

### Return type

[**RecommendResponse**](RecommendResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \NetworkAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityLookup**](NetworkAPI.md#IdentityLookup) | **Get** /api/v1/network/identities/{phoneNumber} | Lookup Identity
[**ListProviders**](NetworkAPI.md#ListProviders) | **Get** /api/v1/network/providers | List Identity Providers
[**RecommendProviders**](NetworkAPI.md#RecommendProviders) | **Post** /api/v1/network/recommend | Recommend Providers



## IdentityLookup

> IdentityLookupResponse IdentityLookup(ctx, phoneNumber).Execute()

Lookup Identity



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
	phoneNumber := "phoneNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.IdentityLookup(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.IdentityLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityLookup`: IdentityLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.IdentityLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityLookupResponse**](IdentityLookupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> ListProvidersResponse ListProviders(ctx).Execute()

List Identity Providers



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
	resp, r, err := apiClient.NetworkAPI.ListProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: ListProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


### Return type

[**ListProvidersResponse**](ListProvidersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

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
	recommendRequest := *openapiclient.NewRecommendRequest("PhoneNumber_example") // RecommendRequest |  (optional)

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
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


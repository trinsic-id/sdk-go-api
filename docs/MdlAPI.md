# \MdlAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMdlExchange**](MdlAPI.md#CreateMdlExchange) | **Post** /api/valpha/mdl/exchanges/create | Create mDL Exchange
[**FinalizeMdlExchange**](MdlAPI.md#FinalizeMdlExchange) | **Post** /api/valpha/mdl/exchanges/finalize | Finalize mDL Exchange



## CreateMdlExchange

> CreateMdlExchangeResponse CreateMdlExchange(ctx).CreateMdlExchangeRequest(createMdlExchangeRequest).Execute()

Create mDL Exchange



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
	createMdlExchangeRequest := *openapiclient.NewCreateMdlExchangeRequest("VerificationProfileId_example", "google-wallet", openapiclient.MdlExchangeMechanism("NativeApp"), "org.iso.18013.5.1.mDL", map[string]map[string]bool{"key": map[string]bool{"key": false}}) // CreateMdlExchangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdlAPI.CreateMdlExchange(context.Background()).CreateMdlExchangeRequest(createMdlExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdlAPI.CreateMdlExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdlExchange`: CreateMdlExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `MdlAPI.CreateMdlExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdlExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMdlExchangeRequest** | [**CreateMdlExchangeRequest**](CreateMdlExchangeRequest.md) |  | 

### Return type

[**CreateMdlExchangeResponse**](CreateMdlExchangeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeMdlExchange

> FinalizeMdlExchangeResponse FinalizeMdlExchange(ctx).FinalizeMdlExchangeRequest(finalizeMdlExchangeRequest).Execute()

Finalize mDL Exchange



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
	finalizeMdlExchangeRequest := *openapiclient.NewFinalizeMdlExchangeRequest("VerificationProfileId_example", "ExchangeId_example", "ExchangeContext_example", "ResponseToken_example") // FinalizeMdlExchangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdlAPI.FinalizeMdlExchange(context.Background()).FinalizeMdlExchangeRequest(finalizeMdlExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdlAPI.FinalizeMdlExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinalizeMdlExchange`: FinalizeMdlExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `MdlAPI.FinalizeMdlExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeMdlExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finalizeMdlExchangeRequest** | [**FinalizeMdlExchangeRequest**](FinalizeMdlExchangeRequest.md) |  | 

### Return type

[**FinalizeMdlExchangeResponse**](FinalizeMdlExchangeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


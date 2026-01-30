# \RedirectUrisAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](RedirectUrisAPI.md#Add) | **Post** /api/valpha/redirect-uris | Add Redirect URI
[**Delete**](RedirectUrisAPI.md#Delete) | **Delete** /api/valpha/redirect-uris/{id} | Delete Redirect URI
[**List**](RedirectUrisAPI.md#List) | **Get** /api/valpha/redirect-uris | List Redirect URIs



## Add

> AddRedirectUriResponse Add(ctx).AddRedirectUriRequest(addRedirectUriRequest).Execute()

Add Redirect URI



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
	addRedirectUriRequest := *openapiclient.NewAddRedirectUriRequest("Uri_example") // AddRedirectUriRequest | Request for uri to add to the environment. Must be absolute, not relative. Wildcard \"*\" accepted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectUrisAPI.Add(context.Background()).AddRedirectUriRequest(addRedirectUriRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectUrisAPI.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Add`: AddRedirectUriResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectUrisAPI.Add`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRedirectUriRequest** | [**AddRedirectUriRequest**](AddRedirectUriRequest.md) | Request for uri to add to the environment. Must be absolute, not relative. Wildcard \&quot;*\&quot; accepted. | 

### Return type

[**AddRedirectUriResponse**](AddRedirectUriResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()

Delete Redirect URI



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the redirect uri to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RedirectUrisAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectUrisAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the redirect uri to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListRedirectUrisResponse List(ctx).Page(page).PageSize(pageSize).Execute()

List Redirect URIs



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
	page := int32(56) // int32 | Number of pages of uris to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | Size of the list to be returned. Accepted range from 1 to 100 (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectUrisAPI.List(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectUrisAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListRedirectUrisResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectUrisAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Number of pages of uris to return. | [default to 1]
 **pageSize** | **int32** | Size of the list to be returned. Accepted range from 1 to 100 | [default to 20]

### Return type

[**ListRedirectUrisResponse**](ListRedirectUrisResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


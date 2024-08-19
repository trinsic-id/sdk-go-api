# \SessionsAPI

All URIs are relative to *https://connect.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSession**](SessionsAPI.md#CancelSession) | **Post** /api/v1/sessions/{sessionId}/cancel | Cancel a Session by its ID
[**CreateSession**](SessionsAPI.md#CreateSession) | **Post** /api/v1/sessions | Create a Session to verify a user&#39;s identity
[**ExchangeResultsKey**](SessionsAPI.md#ExchangeResultsKey) | **Post** /api/v1/sessions/{sessionId}/results/exchange | Exchange a Results Access Key for Identity Data
[**GetSession**](SessionsAPI.md#GetSession) | **Get** /api/v1/sessions/{sessionId} | Get a Session by its ID
[**ListSessions**](SessionsAPI.md#ListSessions) | **Get** /api/v1/sessions | List Sessions created by your account
[**RedactSession**](SessionsAPI.md#RedactSession) | **Post** /api/v1/sessions/{sessionId}/redact | Redact a Session, removing all identity data from Trinsic&#39;s servers.                Identity data that a user has chosen to save in their passkey-protected wallet will not be deleted.



## CancelSession

> CancelSessionResponse CancelSession(ctx, sessionId).Execute()

Cancel a Session by its ID

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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CancelSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CancelSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSession`: CancelSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CancelSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelSessionResponse**](CancelSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> CreateSessionResponse CreateSession(ctx).CreateSessionRequest(createSessionRequest).Execute()

Create a Session to verify a user's identity

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
	createSessionRequest := *openapiclient.NewCreateSessionRequest() // CreateSessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CreateSession(context.Background()).CreateSessionRequest(createSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSession`: CreateSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSessionRequest** | [**CreateSessionRequest**](CreateSessionRequest.md) |  | 

### Return type

[**CreateSessionResponse**](CreateSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExchangeResultsKey

> ExchangeResultsKeyResponse ExchangeResultsKey(ctx, sessionId).ExchangeResultsKeyRequest(exchangeResultsKeyRequest).Execute()

Exchange a Results Access Key for Identity Data

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
	sessionId := "sessionId_example" // string | 
	exchangeResultsKeyRequest := *openapiclient.NewExchangeResultsKeyRequest("ResultsAccessKey_example") // ExchangeResultsKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.ExchangeResultsKey(context.Background(), sessionId).ExchangeResultsKeyRequest(exchangeResultsKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.ExchangeResultsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeResultsKey`: ExchangeResultsKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.ExchangeResultsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExchangeResultsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangeResultsKeyRequest** | [**ExchangeResultsKeyRequest**](ExchangeResultsKeyRequest.md) |  | 

### Return type

[**ExchangeResultsKeyResponse**](ExchangeResultsKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> GetSessionResponse GetSession(ctx, sessionId).Execute()

Get a Session by its ID

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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.GetSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.GetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSession`: GetSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSessionResponse**](GetSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> ListSessionsResponse ListSessions(ctx).OrderBy(orderBy).OrderDirection(orderDirection).PageSize(pageSize).Page(page).Execute()

List Sessions created by your account

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
	orderBy := openapiclient.SessionOrdering("Created") // SessionOrdering | The field by which sessions should be ordered (optional)
	orderDirection := openapiclient.OrderDirection("Ascending") // OrderDirection |  (optional)
	pageSize := int32(50) // int32 | The number of items to return per page -- must be between `1` and `50` (optional)
	page := int32(1) // int32 | The page number to return -- starts at `1` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.ListSessions(context.Background()).OrderBy(orderBy).OrderDirection(orderDirection).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.ListSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSessions`: ListSessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.ListSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | [**SessionOrdering**](SessionOrdering.md) | The field by which sessions should be ordered | 
 **orderDirection** | [**OrderDirection**](OrderDirection.md) |  | 
 **pageSize** | **int32** | The number of items to return per page -- must be between &#x60;1&#x60; and &#x60;50&#x60; | 
 **page** | **int32** | The page number to return -- starts at &#x60;1&#x60; | 

### Return type

[**ListSessionsResponse**](ListSessionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedactSession

> RedactSession(ctx, sessionId).Execute()

Redact a Session, removing all identity data from Trinsic's servers.                Identity data that a user has chosen to save in their passkey-protected wallet will not be deleted.

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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsAPI.RedactSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.RedactSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedactSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


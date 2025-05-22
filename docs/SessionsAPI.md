# \SessionsAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSession**](SessionsAPI.md#CancelSession) | **Post** /api/v1/sessions/{sessionId}/cancel | Cancel Session
[**CreateAdvancedProviderSession**](SessionsAPI.md#CreateAdvancedProviderSession) | **Post** /api/v1/sessions/provider/advanced | Create Advanced Provider Session
[**CreateHostedProviderSession**](SessionsAPI.md#CreateHostedProviderSession) | **Post** /api/v1/sessions/provider/hosted | Create Hosted Provider Session
[**CreateWidgetSession**](SessionsAPI.md#CreateWidgetSession) | **Post** /api/v1/sessions/widget | Create Widget Session
[**GetSession**](SessionsAPI.md#GetSession) | **Get** /api/v1/sessions/{sessionId} | Get Session
[**GetSessionResult**](SessionsAPI.md#GetSessionResult) | **Post** /api/v1/sessions/{sessionId}/results | Get Session Results
[**ListSessions**](SessionsAPI.md#ListSessions) | **Get** /api/v1/sessions/list | List Sessions
[**RedactSession**](SessionsAPI.md#RedactSession) | **Post** /api/v1/sessions/{sessionId}/redact | Redact Session
[**RefreshStepContent**](SessionsAPI.md#RefreshStepContent) | **Post** /api/v1/sessions/{acceptanceSessionId}/step/refresh | Refresh Step Content



## CancelSession

> CancelSessionResponse CancelSession(ctx, sessionId).Execute()

Cancel Session



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdvancedProviderSession

> CreateAdvancedProviderSessionResponse CreateAdvancedProviderSession(ctx).CreateAdvancedProviderSessionRequest(createAdvancedProviderSessionRequest).Execute()

Create Advanced Provider Session



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
	createAdvancedProviderSessionRequest := *openapiclient.NewCreateAdvancedProviderSessionRequest("Provider_example", []openapiclient.IntegrationCapability{openapiclient.IntegrationCapability("LaunchBrowser")}) // CreateAdvancedProviderSessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CreateAdvancedProviderSession(context.Background()).CreateAdvancedProviderSessionRequest(createAdvancedProviderSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateAdvancedProviderSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdvancedProviderSession`: CreateAdvancedProviderSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CreateAdvancedProviderSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdvancedProviderSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdvancedProviderSessionRequest** | [**CreateAdvancedProviderSessionRequest**](CreateAdvancedProviderSessionRequest.md) |  | 

### Return type

[**CreateAdvancedProviderSessionResponse**](CreateAdvancedProviderSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostedProviderSession

> CreateHostedProviderSessionResponse CreateHostedProviderSession(ctx).CreateHostedProviderSessionRequest(createHostedProviderSessionRequest).Execute()

Create Hosted Provider Session



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
	createHostedProviderSessionRequest := *openapiclient.NewCreateHostedProviderSessionRequest("Provider_example", "RedirectUrl_example") // CreateHostedProviderSessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CreateHostedProviderSession(context.Background()).CreateHostedProviderSessionRequest(createHostedProviderSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateHostedProviderSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostedProviderSession`: CreateHostedProviderSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CreateHostedProviderSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostedProviderSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHostedProviderSessionRequest** | [**CreateHostedProviderSessionRequest**](CreateHostedProviderSessionRequest.md) |  | 

### Return type

[**CreateHostedProviderSessionResponse**](CreateHostedProviderSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWidgetSession

> CreateWidgetSessionResponse CreateWidgetSession(ctx).CreateWidgetSessionRequest(createWidgetSessionRequest).Execute()

Create Widget Session



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
	createWidgetSessionRequest := *openapiclient.NewCreateWidgetSessionRequest() // CreateWidgetSessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CreateWidgetSession(context.Background()).CreateWidgetSessionRequest(createWidgetSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateWidgetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWidgetSession`: CreateWidgetSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CreateWidgetSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWidgetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWidgetSessionRequest** | [**CreateWidgetSessionRequest**](CreateWidgetSessionRequest.md) |  | 

### Return type

[**CreateWidgetSessionResponse**](CreateWidgetSessionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> GetSessionResponse GetSession(ctx, sessionId).Execute()

Get Session



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
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionResult

> GetSessionResultResponse GetSessionResult(ctx, sessionId).GetSessionResultRequest(getSessionResultRequest).Execute()

Get Session Results

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
	getSessionResultRequest := *openapiclient.NewGetSessionResultRequest("ResultsAccessKey_example") // GetSessionResultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.GetSessionResult(context.Background(), sessionId).GetSessionResultRequest(getSessionResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.GetSessionResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionResult`: GetSessionResultResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.GetSessionResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSessionResultRequest** | [**GetSessionResultRequest**](GetSessionResultRequest.md) |  | 

### Return type

[**GetSessionResultResponse**](GetSessionResultResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> ListSessionsResponse ListSessions(ctx).OrderBy(orderBy).OrderDirection(orderDirection).PageSize(pageSize).Page(page).Execute()

List Sessions



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
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedactSession

> RedactSession(ctx, sessionId).Execute()

Redact Session



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshStepContent

> RefreshStepContentResponse RefreshStepContent(ctx, acceptanceSessionId).RefreshStepContentRequest(refreshStepContentRequest).Execute()

Refresh Step Content



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
	acceptanceSessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	refreshStepContentRequest := *openapiclient.NewRefreshStepContentRequest("ResultsAccessKey_example") // RefreshStepContentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.RefreshStepContent(context.Background(), acceptanceSessionId).RefreshStepContentRequest(refreshStepContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.RefreshStepContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshStepContent`: RefreshStepContentResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.RefreshStepContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acceptanceSessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshStepContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStepContentRequest** | [**RefreshStepContentRequest**](RefreshStepContentRequest.md) |  | 

### Return type

[**RefreshStepContentResponse**](RefreshStepContentResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


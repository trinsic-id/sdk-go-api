# \DeveloperAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestAgainstArbitraryKey**](DeveloperAPI.md#TestAgainstArbitraryKey) | **Post** /api/v1/developer/hpke-test | Test HPKE Against Arbitrary Key
[**TestAgainstVerificationProfile**](DeveloperAPI.md#TestAgainstVerificationProfile) | **Post** /api/v1/developer/hpke-test/{verificationProfileId} | Test HPKE Against Verification Profile Key



## TestAgainstArbitraryKey

> TrinsicHPKEBundle TestAgainstArbitraryKey(ctx).HPKETestRequest(hPKETestRequest).Execute()

Test HPKE Against Arbitrary Key



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
	hPKETestRequest := *openapiclient.NewHPKETestRequest(string(123), "InfoSuffix_example", "PublicKey_example") // HPKETestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeveloperAPI.TestAgainstArbitraryKey(context.Background()).HPKETestRequest(hPKETestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAPI.TestAgainstArbitraryKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestAgainstArbitraryKey`: TrinsicHPKEBundle
	fmt.Fprintf(os.Stdout, "Response from `DeveloperAPI.TestAgainstArbitraryKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestAgainstArbitraryKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hPKETestRequest** | [**HPKETestRequest**](HPKETestRequest.md) |  | 

### Return type

[**TrinsicHPKEBundle**](TrinsicHPKEBundle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestAgainstVerificationProfile

> TrinsicHPKEBundle TestAgainstVerificationProfile(ctx, verificationProfileId).HPKETestVPRequest(hPKETestVPRequest).Execute()

Test HPKE Against Verification Profile Key



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
	verificationProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	hPKETestVPRequest := *openapiclient.NewHPKETestVPRequest(string(123), string(123)) // HPKETestVPRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeveloperAPI.TestAgainstVerificationProfile(context.Background(), verificationProfileId).HPKETestVPRequest(hPKETestVPRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAPI.TestAgainstVerificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestAgainstVerificationProfile`: TrinsicHPKEBundle
	fmt.Fprintf(os.Stdout, "Response from `DeveloperAPI.TestAgainstVerificationProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verificationProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestAgainstVerificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hPKETestVPRequest** | [**HPKETestVPRequest**](HPKETestVPRequest.md) |  | 

### Return type

[**TrinsicHPKEBundle**](TrinsicHPKEBundle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


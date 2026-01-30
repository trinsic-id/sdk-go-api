# \VerificationProfilesAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerificationProfile**](VerificationProfilesAPI.md#CreateVerificationProfile) | **Post** /api/valpha/verification-profiles | Create Verification Profile
[**GetVerificationProfileById**](VerificationProfilesAPI.md#GetVerificationProfileById) | **Get** /api/valpha/verification-profiles/{id} | Get Verification Profile
[**ListVerificationProfiles**](VerificationProfilesAPI.md#ListVerificationProfiles) | **Get** /api/valpha/verification-profiles | List Verification Profiles



## CreateVerificationProfile

> CreateVerificationProfileResponse CreateVerificationProfile(ctx).Alias(alias).BrandName(brandName).PrimaryColor(primaryColor).Providers(providers).Logo(logo).IsProductionUsage(isProductionUsage).Execute()

Create Verification Profile



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
	alias := "alias_example" // string | An alias of the verification profile shown to developers and administrators.
	brandName := "brandName_example" // string | The brand name of the verification profile shown to end-users.
	primaryColor := "primaryColor_example" // string | The primary color of the verification profile. Must be a 6-character hex string prefixed with a '#' character. Example: #000000 (optional)
	providers := []string{"Inner_example"} // []string | The list of providers you'd like to select for this profile. We will not currently enable any providers. (optional)
	logo := os.NewFile(1234, "some_file") // *os.File | The logo of the verification profile. (optional)
	isProductionUsage := true // bool | Whether this profile is for production usage. Only applicable for Live environment profiles. If not specified for Live profiles, defaults to false (Demo). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationProfilesAPI.CreateVerificationProfile(context.Background()).Alias(alias).BrandName(brandName).PrimaryColor(primaryColor).Providers(providers).Logo(logo).IsProductionUsage(isProductionUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationProfilesAPI.CreateVerificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationProfile`: CreateVerificationProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `VerificationProfilesAPI.CreateVerificationProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alias** | **string** | An alias of the verification profile shown to developers and administrators. | 
 **brandName** | **string** | The brand name of the verification profile shown to end-users. | 
 **primaryColor** | **string** | The primary color of the verification profile. Must be a 6-character hex string prefixed with a &#39;#&#39; character. Example: #000000 | 
 **providers** | **[]string** | The list of providers you&#39;d like to select for this profile. We will not currently enable any providers. | 
 **logo** | ***os.File** | The logo of the verification profile. | 
 **isProductionUsage** | **bool** | Whether this profile is for production usage. Only applicable for Live environment profiles. If not specified for Live profiles, defaults to false (Demo). | 

### Return type

[**CreateVerificationProfileResponse**](CreateVerificationProfileResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerificationProfileById

> VerificationProfileResponse GetVerificationProfileById(ctx, id).Execute()

Get Verification Profile



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationProfilesAPI.GetVerificationProfileById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationProfilesAPI.GetVerificationProfileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerificationProfileById`: VerificationProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `VerificationProfilesAPI.GetVerificationProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerificationProfileResponse**](VerificationProfileResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerificationProfiles

> ListVerificationProfilesResponse ListVerificationProfiles(ctx).Page(page).PageSize(pageSize).Execute()

List Verification Profiles



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 | Size of the list to be returned. Accepted range from 1 to 100 (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationProfilesAPI.ListVerificationProfiles(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationProfilesAPI.ListVerificationProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerificationProfiles`: ListVerificationProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `VerificationProfilesAPI.ListVerificationProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** | Size of the list to be returned. Accepted range from 1 to 100 | [default to 20]

### Return type

[**ListVerificationProfilesResponse**](ListVerificationProfilesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


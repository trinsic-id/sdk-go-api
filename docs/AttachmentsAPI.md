# \AttachmentsAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachment**](AttachmentsAPI.md#GetAttachment) | **Post** /api/v1/attachments/attachment | Get Attachment



## GetAttachment

> GetAttachmentResponse GetAttachment(ctx).GetAttachmentRequest(getAttachmentRequest).Execute()

Get Attachment



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
	getAttachmentRequest := *openapiclient.NewGetAttachmentRequest("AttachmentAccessKey_example") // GetAttachmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.GetAttachment(context.Background()).GetAttachmentRequest(getAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: GetAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAttachmentRequest** | [**GetAttachmentRequest**](GetAttachmentRequest.md) |  | 

### Return type

[**GetAttachmentResponse**](GetAttachmentResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


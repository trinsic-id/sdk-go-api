# \AttachmentsAPI

All URIs are relative to *https://api.trinsic.id*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeAttachmentAccessKey**](AttachmentsAPI.md#ExchangeAttachmentAccessKey) | **Post** /api/v1/attachments/exchange | Exchange an Attachment Access Key (from &#x60;IdentityData.Attachments&#x60;) for the raw contents of the attachment.                Use this API to fetch document (front, back, portrait) or other (selfie) images from a verification, if relevant.                In some cases, attachments may not be immediately available after a verification is completed. If so, this endpoint will return an HTTP 202 code, and you should try again later.



## ExchangeAttachmentAccessKey

> ExchangeAttachmentAccessKey(ctx).ExchangeAttachmentAccessKeyRequest(exchangeAttachmentAccessKeyRequest).Execute()

Exchange an Attachment Access Key (from `IdentityData.Attachments`) for the raw contents of the attachment.                Use this API to fetch document (front, back, portrait) or other (selfie) images from a verification, if relevant.                In some cases, attachments may not be immediately available after a verification is completed. If so, this endpoint will return an HTTP 202 code, and you should try again later.

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
	exchangeAttachmentAccessKeyRequest := *openapiclient.NewExchangeAttachmentAccessKeyRequest("AttachmentAccessKey_example") // ExchangeAttachmentAccessKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttachmentsAPI.ExchangeAttachmentAccessKey(context.Background()).ExchangeAttachmentAccessKeyRequest(exchangeAttachmentAccessKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ExchangeAttachmentAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeAttachmentAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeAttachmentAccessKeyRequest** | [**ExchangeAttachmentAccessKeyRequest**](ExchangeAttachmentAccessKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


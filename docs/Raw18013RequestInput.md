# Raw18013RequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentRequests** | [**[]Raw18013DocumentRequest**](Raw18013DocumentRequest.md) | A collection of requests for specific document(s), any of which may be used to satisfy the verification.              At least one request must be provided. No more than 10 requests may be provided. | 

## Methods

### NewRaw18013RequestInput

`func NewRaw18013RequestInput(documentRequests []Raw18013DocumentRequest, ) *Raw18013RequestInput`

NewRaw18013RequestInput instantiates a new Raw18013RequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaw18013RequestInputWithDefaults

`func NewRaw18013RequestInputWithDefaults() *Raw18013RequestInput`

NewRaw18013RequestInputWithDefaults instantiates a new Raw18013RequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentRequests

`func (o *Raw18013RequestInput) GetDocumentRequests() []Raw18013DocumentRequest`

GetDocumentRequests returns the DocumentRequests field if non-nil, zero value otherwise.

### GetDocumentRequestsOk

`func (o *Raw18013RequestInput) GetDocumentRequestsOk() (*[]Raw18013DocumentRequest, bool)`

GetDocumentRequestsOk returns a tuple with the DocumentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRequests

`func (o *Raw18013RequestInput) SetDocumentRequests(v []Raw18013DocumentRequest)`

SetDocumentRequests sets DocumentRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



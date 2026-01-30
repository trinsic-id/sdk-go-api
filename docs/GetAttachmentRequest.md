# GetAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultsAccessKey** | **string** | The Results Access Key for the Session associated with the Attachment being retrieved.              This is returned during Session creation. | 

## Methods

### NewGetAttachmentRequest

`func NewGetAttachmentRequest(resultsAccessKey string, ) *GetAttachmentRequest`

NewGetAttachmentRequest instantiates a new GetAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttachmentRequestWithDefaults

`func NewGetAttachmentRequestWithDefaults() *GetAttachmentRequest`

NewGetAttachmentRequestWithDefaults instantiates a new GetAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultsAccessKey

`func (o *GetAttachmentRequest) GetResultsAccessKey() string`

GetResultsAccessKey returns the ResultsAccessKey field if non-nil, zero value otherwise.

### GetResultsAccessKeyOk

`func (o *GetAttachmentRequest) GetResultsAccessKeyOk() (*string, bool)`

GetResultsAccessKeyOk returns a tuple with the ResultsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAccessKey

`func (o *GetAttachmentRequest) SetResultsAccessKey(v string)`

SetResultsAccessKey sets ResultsAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



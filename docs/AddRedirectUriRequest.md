# AddRedirectUriRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | 

## Methods

### Named construction

Use a keyed struct literal (`AddRedirectUriRequest{FieldName: value}`), or call
`NewAddRedirectUriRequestWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewAddRedirectUriRequestWithDefaults

`func NewAddRedirectUriRequestWithDefaults() *AddRedirectUriRequest`

NewAddRedirectUriRequestWithDefaults instantiates a new AddRedirectUriRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *AddRedirectUriRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AddRedirectUriRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AddRedirectUriRequest) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



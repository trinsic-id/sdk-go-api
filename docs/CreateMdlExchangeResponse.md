# CreateMdlExchangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeId** | **string** | The ID of the mDL exchange which was created. | 
**RequestObjectBase64Url** | **string** | The request object for this mDL exchange.              Pass this into a Trinsic mDL SDK (Web, iOS, Android) exactly as-is to initiate the mDL exchange. | 
**ExchangeContext** | **string** | The encrypted exchange context for this mDL exchange.              This must be passed back to the API during finalization, alongside the response token from the wallet. | 

## Methods

### NewCreateMdlExchangeResponse

`func NewCreateMdlExchangeResponse(exchangeId string, requestObjectBase64Url string, exchangeContext string, ) *CreateMdlExchangeResponse`

NewCreateMdlExchangeResponse instantiates a new CreateMdlExchangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMdlExchangeResponseWithDefaults

`func NewCreateMdlExchangeResponseWithDefaults() *CreateMdlExchangeResponse`

NewCreateMdlExchangeResponseWithDefaults instantiates a new CreateMdlExchangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeId

`func (o *CreateMdlExchangeResponse) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreateMdlExchangeResponse) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreateMdlExchangeResponse) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.


### GetRequestObjectBase64Url

`func (o *CreateMdlExchangeResponse) GetRequestObjectBase64Url() string`

GetRequestObjectBase64Url returns the RequestObjectBase64Url field if non-nil, zero value otherwise.

### GetRequestObjectBase64UrlOk

`func (o *CreateMdlExchangeResponse) GetRequestObjectBase64UrlOk() (*string, bool)`

GetRequestObjectBase64UrlOk returns a tuple with the RequestObjectBase64Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectBase64Url

`func (o *CreateMdlExchangeResponse) SetRequestObjectBase64Url(v string)`

SetRequestObjectBase64Url sets RequestObjectBase64Url field to given value.


### GetExchangeContext

`func (o *CreateMdlExchangeResponse) GetExchangeContext() string`

GetExchangeContext returns the ExchangeContext field if non-nil, zero value otherwise.

### GetExchangeContextOk

`func (o *CreateMdlExchangeResponse) GetExchangeContextOk() (*string, bool)`

GetExchangeContextOk returns a tuple with the ExchangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeContext

`func (o *CreateMdlExchangeResponse) SetExchangeContext(v string)`

SetExchangeContext sets ExchangeContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FinalizeMdlExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationProfileId** | **string** | The ID of the VerificationProfile which was used to create the mDL exchange. | 
**ExchangeId** | **string** | The ID of the mDL exchange to finalize. | 
**ExchangeContext** | **string** | The encrypted exchange context which was returned when the mDL exchange was created. | 
**ResponseToken** | **string** | The response token returned by Trinsic&#39;s SDK after a successful mDL exchange. | 

## Methods

### NewFinalizeMdlExchangeRequest

`func NewFinalizeMdlExchangeRequest(verificationProfileId string, exchangeId string, exchangeContext string, responseToken string, ) *FinalizeMdlExchangeRequest`

NewFinalizeMdlExchangeRequest instantiates a new FinalizeMdlExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinalizeMdlExchangeRequestWithDefaults

`func NewFinalizeMdlExchangeRequestWithDefaults() *FinalizeMdlExchangeRequest`

NewFinalizeMdlExchangeRequestWithDefaults instantiates a new FinalizeMdlExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationProfileId

`func (o *FinalizeMdlExchangeRequest) GetVerificationProfileId() string`

GetVerificationProfileId returns the VerificationProfileId field if non-nil, zero value otherwise.

### GetVerificationProfileIdOk

`func (o *FinalizeMdlExchangeRequest) GetVerificationProfileIdOk() (*string, bool)`

GetVerificationProfileIdOk returns a tuple with the VerificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfileId

`func (o *FinalizeMdlExchangeRequest) SetVerificationProfileId(v string)`

SetVerificationProfileId sets VerificationProfileId field to given value.


### GetExchangeId

`func (o *FinalizeMdlExchangeRequest) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *FinalizeMdlExchangeRequest) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *FinalizeMdlExchangeRequest) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.


### GetExchangeContext

`func (o *FinalizeMdlExchangeRequest) GetExchangeContext() string`

GetExchangeContext returns the ExchangeContext field if non-nil, zero value otherwise.

### GetExchangeContextOk

`func (o *FinalizeMdlExchangeRequest) GetExchangeContextOk() (*string, bool)`

GetExchangeContextOk returns a tuple with the ExchangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeContext

`func (o *FinalizeMdlExchangeRequest) SetExchangeContext(v string)`

SetExchangeContext sets ExchangeContext field to given value.


### GetResponseToken

`func (o *FinalizeMdlExchangeRequest) GetResponseToken() string`

GetResponseToken returns the ResponseToken field if non-nil, zero value otherwise.

### GetResponseTokenOk

`func (o *FinalizeMdlExchangeRequest) GetResponseTokenOk() (*string, bool)`

GetResponseTokenOk returns a tuple with the ResponseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseToken

`func (o *FinalizeMdlExchangeRequest) SetResponseToken(v string)`

SetResponseToken sets ResponseToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



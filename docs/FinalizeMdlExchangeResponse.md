# FinalizeMdlExchangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeId** | **string** | The Exchange ID of the mDL exchange which was just finalized. | 
**CreatedSession** | [**Session**](Session.md) | The AcceptanceSession which was created as a result of finalizing this mDL exchange. | 
**MdlData** | Pointer to [**NullableMdlIdentityData**](MdlIdentityData.md) | The identity data from the mDL exchange, in a semi-normalized format.              Supports all possible fields and namespaces, but does not map them to Trinsic&#39;s common data model. | [optional] 
**NormalizedIdentityData** | Pointer to [**NullableIdentityData**](IdentityData.md) | The identity data from the mDL exchange, normalized into Trinsic&#39;s common data model. | [optional] 

## Methods

### NewFinalizeMdlExchangeResponse

`func NewFinalizeMdlExchangeResponse(exchangeId string, createdSession Session, ) *FinalizeMdlExchangeResponse`

NewFinalizeMdlExchangeResponse instantiates a new FinalizeMdlExchangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinalizeMdlExchangeResponseWithDefaults

`func NewFinalizeMdlExchangeResponseWithDefaults() *FinalizeMdlExchangeResponse`

NewFinalizeMdlExchangeResponseWithDefaults instantiates a new FinalizeMdlExchangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeId

`func (o *FinalizeMdlExchangeResponse) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *FinalizeMdlExchangeResponse) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *FinalizeMdlExchangeResponse) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.


### GetCreatedSession

`func (o *FinalizeMdlExchangeResponse) GetCreatedSession() Session`

GetCreatedSession returns the CreatedSession field if non-nil, zero value otherwise.

### GetCreatedSessionOk

`func (o *FinalizeMdlExchangeResponse) GetCreatedSessionOk() (*Session, bool)`

GetCreatedSessionOk returns a tuple with the CreatedSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedSession

`func (o *FinalizeMdlExchangeResponse) SetCreatedSession(v Session)`

SetCreatedSession sets CreatedSession field to given value.


### GetMdlData

`func (o *FinalizeMdlExchangeResponse) GetMdlData() MdlIdentityData`

GetMdlData returns the MdlData field if non-nil, zero value otherwise.

### GetMdlDataOk

`func (o *FinalizeMdlExchangeResponse) GetMdlDataOk() (*MdlIdentityData, bool)`

GetMdlDataOk returns a tuple with the MdlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdlData

`func (o *FinalizeMdlExchangeResponse) SetMdlData(v MdlIdentityData)`

SetMdlData sets MdlData field to given value.

### HasMdlData

`func (o *FinalizeMdlExchangeResponse) HasMdlData() bool`

HasMdlData returns a boolean if a field has been set.

### SetMdlDataNil

`func (o *FinalizeMdlExchangeResponse) SetMdlDataNil(b bool)`

 SetMdlDataNil sets the value for MdlData to be an explicit nil

### UnsetMdlData
`func (o *FinalizeMdlExchangeResponse) UnsetMdlData()`

UnsetMdlData ensures that no value is present for MdlData, not even an explicit nil
### GetNormalizedIdentityData

`func (o *FinalizeMdlExchangeResponse) GetNormalizedIdentityData() IdentityData`

GetNormalizedIdentityData returns the NormalizedIdentityData field if non-nil, zero value otherwise.

### GetNormalizedIdentityDataOk

`func (o *FinalizeMdlExchangeResponse) GetNormalizedIdentityDataOk() (*IdentityData, bool)`

GetNormalizedIdentityDataOk returns a tuple with the NormalizedIdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedIdentityData

`func (o *FinalizeMdlExchangeResponse) SetNormalizedIdentityData(v IdentityData)`

SetNormalizedIdentityData sets NormalizedIdentityData field to given value.

### HasNormalizedIdentityData

`func (o *FinalizeMdlExchangeResponse) HasNormalizedIdentityData() bool`

HasNormalizedIdentityData returns a boolean if a field has been set.

### SetNormalizedIdentityDataNil

`func (o *FinalizeMdlExchangeResponse) SetNormalizedIdentityDataNil(b bool)`

 SetNormalizedIdentityDataNil sets the value for NormalizedIdentityData to be an explicit nil

### UnsetNormalizedIdentityData
`func (o *FinalizeMdlExchangeResponse) UnsetNormalizedIdentityData()`

UnsetNormalizedIdentityData ensures that no value is present for NormalizedIdentityData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



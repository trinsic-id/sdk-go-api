# HPKETestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The data to encrypt. Must be between 0 and 1024 bytes. | 
**InfoSuffix** | **string** | A 32-character suffix appended to the fixed info prefix to form the HPKE info parameter. | 
**PublicKey** | **string** | The recipient&#39;s public key, as a JWK or PEM string. | 

## Methods

### Named construction

Use a keyed struct literal (`HPKETestRequest{FieldName: value}`), or call
`NewHPKETestRequestWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewHPKETestRequestWithDefaults

`func NewHPKETestRequestWithDefaults() *HPKETestRequest`

NewHPKETestRequestWithDefaults instantiates a new HPKETestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HPKETestRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HPKETestRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HPKETestRequest) SetData(v string)`

SetData sets Data field to given value.


### GetInfoSuffix

`func (o *HPKETestRequest) GetInfoSuffix() string`

GetInfoSuffix returns the InfoSuffix field if non-nil, zero value otherwise.

### GetInfoSuffixOk

`func (o *HPKETestRequest) GetInfoSuffixOk() (*string, bool)`

GetInfoSuffixOk returns a tuple with the InfoSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoSuffix

`func (o *HPKETestRequest) SetInfoSuffix(v string)`

SetInfoSuffix sets InfoSuffix field to given value.


### GetPublicKey

`func (o *HPKETestRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *HPKETestRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *HPKETestRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



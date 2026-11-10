# HPKETestVPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The data to encrypt. Must be between 0 and 1024 bytes. | 
**InfoSuffix** | **string** | A 32-byte suffix appended to the fixed info prefix to form the HPKE info parameter. | 

## Methods

### Named construction

Use a keyed struct literal (`HPKETestVPRequest{FieldName: value}`), or call
`NewHPKETestVPRequestWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewHPKETestVPRequestWithDefaults

`func NewHPKETestVPRequestWithDefaults() *HPKETestVPRequest`

NewHPKETestVPRequestWithDefaults instantiates a new HPKETestVPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HPKETestVPRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HPKETestVPRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HPKETestVPRequest) SetData(v string)`

SetData sets Data field to given value.


### GetInfoSuffix

`func (o *HPKETestVPRequest) GetInfoSuffix() string`

GetInfoSuffix returns the InfoSuffix field if non-nil, zero value otherwise.

### GetInfoSuffixOk

`func (o *HPKETestVPRequest) GetInfoSuffixOk() (*string, bool)`

GetInfoSuffixOk returns a tuple with the InfoSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoSuffix

`func (o *HPKETestVPRequest) SetInfoSuffix(v string)`

SetInfoSuffix sets InfoSuffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



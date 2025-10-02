# ExternalMdlFieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MdlFieldDataType**](MdlFieldDataType.md) | The type of data contained in &#x60;value&#x60;. | 
**Value** | **string** | The string-encoded value of the field. | 

## Methods

### NewExternalMdlFieldData

`func NewExternalMdlFieldData(type_ MdlFieldDataType, value string, ) *ExternalMdlFieldData`

NewExternalMdlFieldData instantiates a new ExternalMdlFieldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMdlFieldDataWithDefaults

`func NewExternalMdlFieldDataWithDefaults() *ExternalMdlFieldData`

NewExternalMdlFieldDataWithDefaults instantiates a new ExternalMdlFieldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalMdlFieldData) GetType() MdlFieldDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalMdlFieldData) GetTypeOk() (*MdlFieldDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalMdlFieldData) SetType(v MdlFieldDataType)`

SetType sets Type field to given value.


### GetValue

`func (o *ExternalMdlFieldData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExternalMdlFieldData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExternalMdlFieldData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



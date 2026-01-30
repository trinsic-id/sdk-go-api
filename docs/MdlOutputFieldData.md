# MdlOutputFieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MdlOutputFieldDataType**](MdlOutputFieldDataType.md) | The type of data contained in &#x60;value&#x60;. | 
**Value** | **string** | The string-encoded value of the field. | 

## Methods

### NewMdlOutputFieldData

`func NewMdlOutputFieldData(type_ MdlOutputFieldDataType, value string, ) *MdlOutputFieldData`

NewMdlOutputFieldData instantiates a new MdlOutputFieldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdlOutputFieldDataWithDefaults

`func NewMdlOutputFieldDataWithDefaults() *MdlOutputFieldData`

NewMdlOutputFieldDataWithDefaults instantiates a new MdlOutputFieldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MdlOutputFieldData) GetType() MdlOutputFieldDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MdlOutputFieldData) GetTypeOk() (*MdlOutputFieldDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MdlOutputFieldData) SetType(v MdlOutputFieldDataType)`

SetType sets Type field to given value.


### GetValue

`func (o *MdlOutputFieldData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MdlOutputFieldData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MdlOutputFieldData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



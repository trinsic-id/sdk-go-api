# MdlOutputFieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableMdlOutputFieldDataType**](MdlOutputFieldDataType.md) | The type of data contained in &#x60;value&#x60;. | [optional] 
**Value** | Pointer to **NullableString** | The string-encoded value of the field. | [optional] 

## Methods

### NewMdlOutputFieldData

`func NewMdlOutputFieldData() *MdlOutputFieldData`

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

### HasType

`func (o *MdlOutputFieldData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MdlOutputFieldData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MdlOutputFieldData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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

### HasValue

`func (o *MdlOutputFieldData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MdlOutputFieldData) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MdlOutputFieldData) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ContractField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field as it appears in verification results. | 
**Outputted** | [**FieldAvailability**](FieldAvailability.md) | Indicates when this field will be present in verification results. | 

## Methods

### NewContractField

`func NewContractField(name string, outputted FieldAvailability, ) *ContractField`

NewContractField instantiates a new ContractField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractFieldWithDefaults

`func NewContractFieldWithDefaults() *ContractField`

NewContractFieldWithDefaults instantiates a new ContractField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractField) SetName(v string)`

SetName sets Name field to given value.


### GetOutputted

`func (o *ContractField) GetOutputted() FieldAvailability`

GetOutputted returns the Outputted field if non-nil, zero value otherwise.

### GetOutputtedOk

`func (o *ContractField) GetOutputtedOk() (*FieldAvailability, bool)`

GetOutputtedOk returns a tuple with the Outputted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputted

`func (o *ContractField) SetOutputted(v FieldAvailability)`

SetOutputted sets Outputted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



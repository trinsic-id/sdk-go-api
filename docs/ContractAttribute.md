# ContractAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The scope of the attribute as it appears in verification results. | 
**Outputted** | [**AttributeAvailability**](AttributeAvailability.md) | Indicates when this attribute will be present in verification results. | 

## Methods

### NewContractAttribute

`func NewContractAttribute(scope string, outputted AttributeAvailability, ) *ContractAttribute`

NewContractAttribute instantiates a new ContractAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractAttributeWithDefaults

`func NewContractAttributeWithDefaults() *ContractAttribute`

NewContractAttributeWithDefaults instantiates a new ContractAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ContractAttribute) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ContractAttribute) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ContractAttribute) SetScope(v string)`

SetScope sets Scope field to given value.


### GetOutputted

`func (o *ContractAttribute) GetOutputted() AttributeAvailability`

GetOutputted returns the Outputted field if non-nil, zero value otherwise.

### GetOutputtedOk

`func (o *ContractAttribute) GetOutputtedOk() (*AttributeAvailability, bool)`

GetOutputtedOk returns a tuple with the Outputted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputted

`func (o *ContractAttribute) SetOutputted(v AttributeAvailability)`

SetOutputted sets Outputted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



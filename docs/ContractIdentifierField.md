# ContractIdentifierField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The scope of the identifier as it appears in verification results. | 
**Outputted** | [**FieldAvailability**](FieldAvailability.md) | Indicates when this field will be present in verification results. | 
**Description** | **string** | A human-readable description of the identifier, written by Trinsic. | 

## Methods

### NewContractIdentifierField

`func NewContractIdentifierField(scope string, outputted FieldAvailability, description string, ) *ContractIdentifierField`

NewContractIdentifierField instantiates a new ContractIdentifierField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractIdentifierFieldWithDefaults

`func NewContractIdentifierFieldWithDefaults() *ContractIdentifierField`

NewContractIdentifierFieldWithDefaults instantiates a new ContractIdentifierField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ContractIdentifierField) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ContractIdentifierField) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ContractIdentifierField) SetScope(v string)`

SetScope sets Scope field to given value.


### GetOutputted

`func (o *ContractIdentifierField) GetOutputted() FieldAvailability`

GetOutputted returns the Outputted field if non-nil, zero value otherwise.

### GetOutputtedOk

`func (o *ContractIdentifierField) GetOutputtedOk() (*FieldAvailability, bool)`

GetOutputtedOk returns a tuple with the Outputted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputted

`func (o *ContractIdentifierField) SetOutputted(v FieldAvailability)`

SetOutputted sets Outputted field to given value.


### GetDescription

`func (o *ContractIdentifierField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractIdentifierField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractIdentifierField) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



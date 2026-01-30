# ContractAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the attachment as it appears in verification results.              Possible values: - \&quot;selfie\&quot; - \&quot;document_front\&quot; - \&quot;document_back\&quot; - \&quot;document_portrait\&quot; - \&quot;document_signature\&quot; - \&quot;document_scan_report\&quot;              Additional attachment types may be defined for specific Providers. | 
**Outputted** | [**AttributeAvailability**](AttributeAvailability.md) | Indicates when this attachment will be present in verification results. | 

## Methods

### NewContractAttachment

`func NewContractAttachment(type_ string, outputted AttributeAvailability, ) *ContractAttachment`

NewContractAttachment instantiates a new ContractAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractAttachmentWithDefaults

`func NewContractAttachmentWithDefaults() *ContractAttachment`

NewContractAttachmentWithDefaults instantiates a new ContractAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContractAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetOutputted

`func (o *ContractAttachment) GetOutputted() AttributeAvailability`

GetOutputted returns the Outputted field if non-nil, zero value otherwise.

### GetOutputtedOk

`func (o *ContractAttachment) GetOutputtedOk() (*AttributeAvailability, bool)`

GetOutputtedOk returns a tuple with the Outputted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputted

`func (o *ContractAttachment) SetOutputted(v AttributeAvailability)`

SetOutputted sets Outputted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



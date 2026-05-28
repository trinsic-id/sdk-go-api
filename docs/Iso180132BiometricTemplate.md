# Iso180132BiometricTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | [**Iso180132BiometricTemplateHeader**](Iso180132BiometricTemplateHeader.md) | Header describing the biometric template. | 
**DataBlock** | **string** | The raw data block, in a biometric type- and format-specific encoding.              For &#x60;face&#x60; and &#x60;signature_usual_mark&#x60;, this is a raw JPEG or JPEG2000 image.              See ISO 18013-2 and 18013-5. | 
**DataIsEncrypted** | **bool** | Whether the data block is encrypted. | 
**BiometricInformationRecordPayload** | Pointer to **NullableString** | Optional Biometric Information Record payload, containing arbitrary domestic data. | [optional] 

## Methods

### NewIso180132BiometricTemplate

`func NewIso180132BiometricTemplate(header Iso180132BiometricTemplateHeader, dataBlock string, dataIsEncrypted bool, ) *Iso180132BiometricTemplate`

NewIso180132BiometricTemplate instantiates a new Iso180132BiometricTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180132BiometricTemplateWithDefaults

`func NewIso180132BiometricTemplateWithDefaults() *Iso180132BiometricTemplate`

NewIso180132BiometricTemplateWithDefaults instantiates a new Iso180132BiometricTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *Iso180132BiometricTemplate) GetHeader() Iso180132BiometricTemplateHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *Iso180132BiometricTemplate) GetHeaderOk() (*Iso180132BiometricTemplateHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *Iso180132BiometricTemplate) SetHeader(v Iso180132BiometricTemplateHeader)`

SetHeader sets Header field to given value.


### GetDataBlock

`func (o *Iso180132BiometricTemplate) GetDataBlock() string`

GetDataBlock returns the DataBlock field if non-nil, zero value otherwise.

### GetDataBlockOk

`func (o *Iso180132BiometricTemplate) GetDataBlockOk() (*string, bool)`

GetDataBlockOk returns a tuple with the DataBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlock

`func (o *Iso180132BiometricTemplate) SetDataBlock(v string)`

SetDataBlock sets DataBlock field to given value.


### GetDataIsEncrypted

`func (o *Iso180132BiometricTemplate) GetDataIsEncrypted() bool`

GetDataIsEncrypted returns the DataIsEncrypted field if non-nil, zero value otherwise.

### GetDataIsEncryptedOk

`func (o *Iso180132BiometricTemplate) GetDataIsEncryptedOk() (*bool, bool)`

GetDataIsEncryptedOk returns a tuple with the DataIsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIsEncrypted

`func (o *Iso180132BiometricTemplate) SetDataIsEncrypted(v bool)`

SetDataIsEncrypted sets DataIsEncrypted field to given value.


### GetBiometricInformationRecordPayload

`func (o *Iso180132BiometricTemplate) GetBiometricInformationRecordPayload() string`

GetBiometricInformationRecordPayload returns the BiometricInformationRecordPayload field if non-nil, zero value otherwise.

### GetBiometricInformationRecordPayloadOk

`func (o *Iso180132BiometricTemplate) GetBiometricInformationRecordPayloadOk() (*string, bool)`

GetBiometricInformationRecordPayloadOk returns a tuple with the BiometricInformationRecordPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricInformationRecordPayload

`func (o *Iso180132BiometricTemplate) SetBiometricInformationRecordPayload(v string)`

SetBiometricInformationRecordPayload sets BiometricInformationRecordPayload field to given value.

### HasBiometricInformationRecordPayload

`func (o *Iso180132BiometricTemplate) HasBiometricInformationRecordPayload() bool`

HasBiometricInformationRecordPayload returns a boolean if a field has been set.

### SetBiometricInformationRecordPayloadNil

`func (o *Iso180132BiometricTemplate) SetBiometricInformationRecordPayloadNil(b bool)`

 SetBiometricInformationRecordPayloadNil sets the value for BiometricInformationRecordPayload to be an explicit nil

### UnsetBiometricInformationRecordPayload
`func (o *Iso180132BiometricTemplate) UnsetBiometricInformationRecordPayload()`

UnsetBiometricInformationRecordPayload ensures that no value is present for BiometricInformationRecordPayload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



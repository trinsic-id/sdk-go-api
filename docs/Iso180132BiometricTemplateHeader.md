# Iso180132BiometricTemplateHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatronHeaderVersion** | Pointer to **NullableInt32** | Patron header version (defaults to 0x0101). | [optional] 
**BiometricType** | Pointer to **NullableInt32** | Biometric type code per ISO 18013-2. | [optional] 
**BiometricSubType** | Pointer to **NullableInt32** | Biometric sub-type code per ISO 18013-2. | [optional] 
**CreationDate** | Pointer to **NullableTime** | Date and time the biometric template was created. | [optional] 
**BiometricInformationRecordCreator** | Pointer to **NullableString** | Name of the Biometric Information Record (BIR) creator. | [optional] 
**ValidityPeriod** | Pointer to [**NullableIso180132BiometricValidityPeriod**](Iso180132BiometricValidityPeriod.md) | Validity period of the biometric data block. | [optional] 
**BiometricDataBlockProduct** | Pointer to [**NullableIso180132BiometricDataBlockProduct**](Iso180132BiometricDataBlockProduct.md) | Owner and type identifying the product that produced the biometric data block. | [optional] 
**BiometricDataBlockFormatOwner** | Pointer to **NullableInt32** | Format owner of the biometric data block. | [optional] 
**BiometricDataBlockFormatType** | Pointer to **NullableInt32** | Format type of the biometric data block. | [optional] 
**BiometricInformationRecordIndex** | Pointer to **NullableString** | Index identifier of the Biometric Information Record, when present. | [optional] 

## Methods

### NewIso180132BiometricTemplateHeader

`func NewIso180132BiometricTemplateHeader() *Iso180132BiometricTemplateHeader`

NewIso180132BiometricTemplateHeader instantiates a new Iso180132BiometricTemplateHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso180132BiometricTemplateHeaderWithDefaults

`func NewIso180132BiometricTemplateHeaderWithDefaults() *Iso180132BiometricTemplateHeader`

NewIso180132BiometricTemplateHeaderWithDefaults instantiates a new Iso180132BiometricTemplateHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatronHeaderVersion

`func (o *Iso180132BiometricTemplateHeader) GetPatronHeaderVersion() int32`

GetPatronHeaderVersion returns the PatronHeaderVersion field if non-nil, zero value otherwise.

### GetPatronHeaderVersionOk

`func (o *Iso180132BiometricTemplateHeader) GetPatronHeaderVersionOk() (*int32, bool)`

GetPatronHeaderVersionOk returns a tuple with the PatronHeaderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronHeaderVersion

`func (o *Iso180132BiometricTemplateHeader) SetPatronHeaderVersion(v int32)`

SetPatronHeaderVersion sets PatronHeaderVersion field to given value.

### HasPatronHeaderVersion

`func (o *Iso180132BiometricTemplateHeader) HasPatronHeaderVersion() bool`

HasPatronHeaderVersion returns a boolean if a field has been set.

### SetPatronHeaderVersionNil

`func (o *Iso180132BiometricTemplateHeader) SetPatronHeaderVersionNil(b bool)`

 SetPatronHeaderVersionNil sets the value for PatronHeaderVersion to be an explicit nil

### UnsetPatronHeaderVersion
`func (o *Iso180132BiometricTemplateHeader) UnsetPatronHeaderVersion()`

UnsetPatronHeaderVersion ensures that no value is present for PatronHeaderVersion, not even an explicit nil
### GetBiometricType

`func (o *Iso180132BiometricTemplateHeader) GetBiometricType() int32`

GetBiometricType returns the BiometricType field if non-nil, zero value otherwise.

### GetBiometricTypeOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricTypeOk() (*int32, bool)`

GetBiometricTypeOk returns a tuple with the BiometricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricType

`func (o *Iso180132BiometricTemplateHeader) SetBiometricType(v int32)`

SetBiometricType sets BiometricType field to given value.

### HasBiometricType

`func (o *Iso180132BiometricTemplateHeader) HasBiometricType() bool`

HasBiometricType returns a boolean if a field has been set.

### SetBiometricTypeNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricTypeNil(b bool)`

 SetBiometricTypeNil sets the value for BiometricType to be an explicit nil

### UnsetBiometricType
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricType()`

UnsetBiometricType ensures that no value is present for BiometricType, not even an explicit nil
### GetBiometricSubType

`func (o *Iso180132BiometricTemplateHeader) GetBiometricSubType() int32`

GetBiometricSubType returns the BiometricSubType field if non-nil, zero value otherwise.

### GetBiometricSubTypeOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricSubTypeOk() (*int32, bool)`

GetBiometricSubTypeOk returns a tuple with the BiometricSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricSubType

`func (o *Iso180132BiometricTemplateHeader) SetBiometricSubType(v int32)`

SetBiometricSubType sets BiometricSubType field to given value.

### HasBiometricSubType

`func (o *Iso180132BiometricTemplateHeader) HasBiometricSubType() bool`

HasBiometricSubType returns a boolean if a field has been set.

### SetBiometricSubTypeNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricSubTypeNil(b bool)`

 SetBiometricSubTypeNil sets the value for BiometricSubType to be an explicit nil

### UnsetBiometricSubType
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricSubType()`

UnsetBiometricSubType ensures that no value is present for BiometricSubType, not even an explicit nil
### GetCreationDate

`func (o *Iso180132BiometricTemplateHeader) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Iso180132BiometricTemplateHeader) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Iso180132BiometricTemplateHeader) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Iso180132BiometricTemplateHeader) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *Iso180132BiometricTemplateHeader) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *Iso180132BiometricTemplateHeader) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetBiometricInformationRecordCreator

`func (o *Iso180132BiometricTemplateHeader) GetBiometricInformationRecordCreator() string`

GetBiometricInformationRecordCreator returns the BiometricInformationRecordCreator field if non-nil, zero value otherwise.

### GetBiometricInformationRecordCreatorOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricInformationRecordCreatorOk() (*string, bool)`

GetBiometricInformationRecordCreatorOk returns a tuple with the BiometricInformationRecordCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricInformationRecordCreator

`func (o *Iso180132BiometricTemplateHeader) SetBiometricInformationRecordCreator(v string)`

SetBiometricInformationRecordCreator sets BiometricInformationRecordCreator field to given value.

### HasBiometricInformationRecordCreator

`func (o *Iso180132BiometricTemplateHeader) HasBiometricInformationRecordCreator() bool`

HasBiometricInformationRecordCreator returns a boolean if a field has been set.

### SetBiometricInformationRecordCreatorNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricInformationRecordCreatorNil(b bool)`

 SetBiometricInformationRecordCreatorNil sets the value for BiometricInformationRecordCreator to be an explicit nil

### UnsetBiometricInformationRecordCreator
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricInformationRecordCreator()`

UnsetBiometricInformationRecordCreator ensures that no value is present for BiometricInformationRecordCreator, not even an explicit nil
### GetValidityPeriod

`func (o *Iso180132BiometricTemplateHeader) GetValidityPeriod() Iso180132BiometricValidityPeriod`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *Iso180132BiometricTemplateHeader) GetValidityPeriodOk() (*Iso180132BiometricValidityPeriod, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *Iso180132BiometricTemplateHeader) SetValidityPeriod(v Iso180132BiometricValidityPeriod)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *Iso180132BiometricTemplateHeader) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### SetValidityPeriodNil

`func (o *Iso180132BiometricTemplateHeader) SetValidityPeriodNil(b bool)`

 SetValidityPeriodNil sets the value for ValidityPeriod to be an explicit nil

### UnsetValidityPeriod
`func (o *Iso180132BiometricTemplateHeader) UnsetValidityPeriod()`

UnsetValidityPeriod ensures that no value is present for ValidityPeriod, not even an explicit nil
### GetBiometricDataBlockProduct

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockProduct() Iso180132BiometricDataBlockProduct`

GetBiometricDataBlockProduct returns the BiometricDataBlockProduct field if non-nil, zero value otherwise.

### GetBiometricDataBlockProductOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockProductOk() (*Iso180132BiometricDataBlockProduct, bool)`

GetBiometricDataBlockProductOk returns a tuple with the BiometricDataBlockProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricDataBlockProduct

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockProduct(v Iso180132BiometricDataBlockProduct)`

SetBiometricDataBlockProduct sets BiometricDataBlockProduct field to given value.

### HasBiometricDataBlockProduct

`func (o *Iso180132BiometricTemplateHeader) HasBiometricDataBlockProduct() bool`

HasBiometricDataBlockProduct returns a boolean if a field has been set.

### SetBiometricDataBlockProductNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockProductNil(b bool)`

 SetBiometricDataBlockProductNil sets the value for BiometricDataBlockProduct to be an explicit nil

### UnsetBiometricDataBlockProduct
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricDataBlockProduct()`

UnsetBiometricDataBlockProduct ensures that no value is present for BiometricDataBlockProduct, not even an explicit nil
### GetBiometricDataBlockFormatOwner

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockFormatOwner() int32`

GetBiometricDataBlockFormatOwner returns the BiometricDataBlockFormatOwner field if non-nil, zero value otherwise.

### GetBiometricDataBlockFormatOwnerOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockFormatOwnerOk() (*int32, bool)`

GetBiometricDataBlockFormatOwnerOk returns a tuple with the BiometricDataBlockFormatOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricDataBlockFormatOwner

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockFormatOwner(v int32)`

SetBiometricDataBlockFormatOwner sets BiometricDataBlockFormatOwner field to given value.

### HasBiometricDataBlockFormatOwner

`func (o *Iso180132BiometricTemplateHeader) HasBiometricDataBlockFormatOwner() bool`

HasBiometricDataBlockFormatOwner returns a boolean if a field has been set.

### SetBiometricDataBlockFormatOwnerNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockFormatOwnerNil(b bool)`

 SetBiometricDataBlockFormatOwnerNil sets the value for BiometricDataBlockFormatOwner to be an explicit nil

### UnsetBiometricDataBlockFormatOwner
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricDataBlockFormatOwner()`

UnsetBiometricDataBlockFormatOwner ensures that no value is present for BiometricDataBlockFormatOwner, not even an explicit nil
### GetBiometricDataBlockFormatType

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockFormatType() int32`

GetBiometricDataBlockFormatType returns the BiometricDataBlockFormatType field if non-nil, zero value otherwise.

### GetBiometricDataBlockFormatTypeOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricDataBlockFormatTypeOk() (*int32, bool)`

GetBiometricDataBlockFormatTypeOk returns a tuple with the BiometricDataBlockFormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricDataBlockFormatType

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockFormatType(v int32)`

SetBiometricDataBlockFormatType sets BiometricDataBlockFormatType field to given value.

### HasBiometricDataBlockFormatType

`func (o *Iso180132BiometricTemplateHeader) HasBiometricDataBlockFormatType() bool`

HasBiometricDataBlockFormatType returns a boolean if a field has been set.

### SetBiometricDataBlockFormatTypeNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricDataBlockFormatTypeNil(b bool)`

 SetBiometricDataBlockFormatTypeNil sets the value for BiometricDataBlockFormatType to be an explicit nil

### UnsetBiometricDataBlockFormatType
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricDataBlockFormatType()`

UnsetBiometricDataBlockFormatType ensures that no value is present for BiometricDataBlockFormatType, not even an explicit nil
### GetBiometricInformationRecordIndex

`func (o *Iso180132BiometricTemplateHeader) GetBiometricInformationRecordIndex() string`

GetBiometricInformationRecordIndex returns the BiometricInformationRecordIndex field if non-nil, zero value otherwise.

### GetBiometricInformationRecordIndexOk

`func (o *Iso180132BiometricTemplateHeader) GetBiometricInformationRecordIndexOk() (*string, bool)`

GetBiometricInformationRecordIndexOk returns a tuple with the BiometricInformationRecordIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricInformationRecordIndex

`func (o *Iso180132BiometricTemplateHeader) SetBiometricInformationRecordIndex(v string)`

SetBiometricInformationRecordIndex sets BiometricInformationRecordIndex field to given value.

### HasBiometricInformationRecordIndex

`func (o *Iso180132BiometricTemplateHeader) HasBiometricInformationRecordIndex() bool`

HasBiometricInformationRecordIndex returns a boolean if a field has been set.

### SetBiometricInformationRecordIndexNil

`func (o *Iso180132BiometricTemplateHeader) SetBiometricInformationRecordIndexNil(b bool)`

 SetBiometricInformationRecordIndexNil sets the value for BiometricInformationRecordIndex to be an explicit nil

### UnsetBiometricInformationRecordIndex
`func (o *Iso180132BiometricTemplateHeader) UnsetBiometricInformationRecordIndex()`

UnsetBiometricInformationRecordIndex ensures that no value is present for BiometricInformationRecordIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



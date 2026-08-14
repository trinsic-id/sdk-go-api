# MdlOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IacaRootCertificate** | Pointer to [**NullableMdlOutputCertificateData**](MdlOutputCertificateData.md) | Information about the IACA Root Certificate which signed (directly or indirectly) the Document Signer Certificate for this mDoc. | [optional] 
**DocumentSignerCertificate** | Pointer to [**NullableMdlOutputCertificateData**](MdlOutputCertificateData.md) | Information about the Document Signer Certificate which signed the mDoc presented by the individual. | [optional] 
**DocumentType** | Pointer to **NullableString** | The document type of the mDoc presented by the individual.              Common values: - \&quot;org.iso.18013.5.1.mDL\&quot; for ISO 18013-5 mDLs - \&quot;com.google.wallet.idcard.1\&quot; for Google Wallet ID Cards | [optional] 
**NameSpaces** | Pointer to [**map[string]map[string]MdlOutputFieldData**](map.md) | The namespaces, and fields within those namespaces, which were present in the processed mDL. | [optional] 

## Methods

### NewMdlOutput

`func NewMdlOutput() *MdlOutput`

NewMdlOutput instantiates a new MdlOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdlOutputWithDefaults

`func NewMdlOutputWithDefaults() *MdlOutput`

NewMdlOutputWithDefaults instantiates a new MdlOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIacaRootCertificate

`func (o *MdlOutput) GetIacaRootCertificate() MdlOutputCertificateData`

GetIacaRootCertificate returns the IacaRootCertificate field if non-nil, zero value otherwise.

### GetIacaRootCertificateOk

`func (o *MdlOutput) GetIacaRootCertificateOk() (*MdlOutputCertificateData, bool)`

GetIacaRootCertificateOk returns a tuple with the IacaRootCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacaRootCertificate

`func (o *MdlOutput) SetIacaRootCertificate(v MdlOutputCertificateData)`

SetIacaRootCertificate sets IacaRootCertificate field to given value.

### HasIacaRootCertificate

`func (o *MdlOutput) HasIacaRootCertificate() bool`

HasIacaRootCertificate returns a boolean if a field has been set.

### SetIacaRootCertificateNil

`func (o *MdlOutput) SetIacaRootCertificateNil(b bool)`

 SetIacaRootCertificateNil sets the value for IacaRootCertificate to be an explicit nil

### UnsetIacaRootCertificate
`func (o *MdlOutput) UnsetIacaRootCertificate()`

UnsetIacaRootCertificate ensures that no value is present for IacaRootCertificate, not even an explicit nil
### GetDocumentSignerCertificate

`func (o *MdlOutput) GetDocumentSignerCertificate() MdlOutputCertificateData`

GetDocumentSignerCertificate returns the DocumentSignerCertificate field if non-nil, zero value otherwise.

### GetDocumentSignerCertificateOk

`func (o *MdlOutput) GetDocumentSignerCertificateOk() (*MdlOutputCertificateData, bool)`

GetDocumentSignerCertificateOk returns a tuple with the DocumentSignerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSignerCertificate

`func (o *MdlOutput) SetDocumentSignerCertificate(v MdlOutputCertificateData)`

SetDocumentSignerCertificate sets DocumentSignerCertificate field to given value.

### HasDocumentSignerCertificate

`func (o *MdlOutput) HasDocumentSignerCertificate() bool`

HasDocumentSignerCertificate returns a boolean if a field has been set.

### SetDocumentSignerCertificateNil

`func (o *MdlOutput) SetDocumentSignerCertificateNil(b bool)`

 SetDocumentSignerCertificateNil sets the value for DocumentSignerCertificate to be an explicit nil

### UnsetDocumentSignerCertificate
`func (o *MdlOutput) UnsetDocumentSignerCertificate()`

UnsetDocumentSignerCertificate ensures that no value is present for DocumentSignerCertificate, not even an explicit nil
### GetDocumentType

`func (o *MdlOutput) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *MdlOutput) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *MdlOutput) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *MdlOutput) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *MdlOutput) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *MdlOutput) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetNameSpaces

`func (o *MdlOutput) GetNameSpaces() map[string]map[string]MdlOutputFieldData`

GetNameSpaces returns the NameSpaces field if non-nil, zero value otherwise.

### GetNameSpacesOk

`func (o *MdlOutput) GetNameSpacesOk() (*map[string]map[string]MdlOutputFieldData, bool)`

GetNameSpacesOk returns a tuple with the NameSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpaces

`func (o *MdlOutput) SetNameSpaces(v map[string]map[string]MdlOutputFieldData)`

SetNameSpaces sets NameSpaces field to given value.

### HasNameSpaces

`func (o *MdlOutput) HasNameSpaces() bool`

HasNameSpaces returns a boolean if a field has been set.

### SetNameSpacesNil

`func (o *MdlOutput) SetNameSpacesNil(b bool)`

 SetNameSpacesNil sets the value for NameSpaces to be an explicit nil

### UnsetNameSpaces
`func (o *MdlOutput) UnsetNameSpaces()`

UnsetNameSpaces ensures that no value is present for NameSpaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



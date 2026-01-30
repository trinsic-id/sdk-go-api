# MdlOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IacaRootCertificate** | [**MdlOutputCertificateData**](MdlOutputCertificateData.md) | Information about the IACA Root Certificate which signed (directly or indirectly) the Document Signer Certificate for this mDL. | 
**DocumentSignerCertificate** | [**MdlOutputCertificateData**](MdlOutputCertificateData.md) | Information about the Document Signer Certificate which signed the mDL presented by the user. | 
**DocumentType** | **string** | The document type of the mDL presented by the user.              Common values: - \&quot;org.iso.18013.5.1.mDL\&quot; for ISO 18013-5 mDLs - \&quot;com.google.wallet.idcard.1\&quot; for Google Wallet ID Cards | 
**NameSpaces** | [**map[string]map[string]MdlOutputFieldData**](map.md) | The namespaces, and fields within those namespaces, which were present in the processed mDL. | 

## Methods

### NewMdlOutput

`func NewMdlOutput(iacaRootCertificate MdlOutputCertificateData, documentSignerCertificate MdlOutputCertificateData, documentType string, nameSpaces map[string]map[string]MdlOutputFieldData, ) *MdlOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



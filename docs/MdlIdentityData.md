# MdlIdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IacaRootCertificate** | [**MdlCertificateData**](MdlCertificateData.md) | Information about the IACA Root Certificate which signed the Issuer Certificate for this mDL. | 
**DocumentSignerCertificate** | [**MdlCertificateData**](MdlCertificateData.md) | Information about the Document Signer Certificate which signed the mDL presented by the user. | 
**NameSpaces** | [**map[string]map[string]ExternalMdlFieldData**](map.md) | The namespaces, and fields within those namespaces, which were present in the processed mDL. | 

## Methods

### NewMdlIdentityData

`func NewMdlIdentityData(iacaRootCertificate MdlCertificateData, documentSignerCertificate MdlCertificateData, nameSpaces map[string]map[string]ExternalMdlFieldData, ) *MdlIdentityData`

NewMdlIdentityData instantiates a new MdlIdentityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdlIdentityDataWithDefaults

`func NewMdlIdentityDataWithDefaults() *MdlIdentityData`

NewMdlIdentityDataWithDefaults instantiates a new MdlIdentityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIacaRootCertificate

`func (o *MdlIdentityData) GetIacaRootCertificate() MdlCertificateData`

GetIacaRootCertificate returns the IacaRootCertificate field if non-nil, zero value otherwise.

### GetIacaRootCertificateOk

`func (o *MdlIdentityData) GetIacaRootCertificateOk() (*MdlCertificateData, bool)`

GetIacaRootCertificateOk returns a tuple with the IacaRootCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacaRootCertificate

`func (o *MdlIdentityData) SetIacaRootCertificate(v MdlCertificateData)`

SetIacaRootCertificate sets IacaRootCertificate field to given value.


### GetDocumentSignerCertificate

`func (o *MdlIdentityData) GetDocumentSignerCertificate() MdlCertificateData`

GetDocumentSignerCertificate returns the DocumentSignerCertificate field if non-nil, zero value otherwise.

### GetDocumentSignerCertificateOk

`func (o *MdlIdentityData) GetDocumentSignerCertificateOk() (*MdlCertificateData, bool)`

GetDocumentSignerCertificateOk returns a tuple with the DocumentSignerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSignerCertificate

`func (o *MdlIdentityData) SetDocumentSignerCertificate(v MdlCertificateData)`

SetDocumentSignerCertificate sets DocumentSignerCertificate field to given value.


### GetNameSpaces

`func (o *MdlIdentityData) GetNameSpaces() map[string]map[string]ExternalMdlFieldData`

GetNameSpaces returns the NameSpaces field if non-nil, zero value otherwise.

### GetNameSpacesOk

`func (o *MdlIdentityData) GetNameSpacesOk() (*map[string]map[string]ExternalMdlFieldData, bool)`

GetNameSpacesOk returns a tuple with the NameSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpaces

`func (o *MdlIdentityData) SetNameSpaces(v map[string]map[string]ExternalMdlFieldData)`

SetNameSpaces sets NameSpaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



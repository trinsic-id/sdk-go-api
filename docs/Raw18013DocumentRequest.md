# Raw18013DocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | **string** | The type of the document being requested, as defined by ISO specifications 18013-5, 18013-7, 23220-2, or a similar or related standard.              Common values: - \&quot;org.iso.18013.5.1.mDL\&quot; -- a Mobile Driver&#39;s License - \&quot;eu.europa.ec.eudi.pid.1\&quot; -- an EUDI Wallet PID - \&quot;com.google.wallet.idcard.1\&quot; -- a Google Wallet ID Pass - \&quot;org.iso.23220.photoid.1\&quot; -- a generic ISO 23220-2 compliant Photo ID (used by Apple Wallet&#39;s ID Pass) | 
**NameSpaces** | **map[string]map[string]bool** | The namespaces and attributes (PII) to request from the document.              This is a map of (nameSpaceName -&gt; (attributeName -&gt; willRetain)), where:              - nameSpaceName is the name of a NameSpace within the document - attributeName is the name of a specific attribute within the NameSpace - willRetain indicates, to the Wallet, whether you as the Relying Party intend to retain the data for longer than the scope of the transaction              Common namespace values: - \&quot;org.iso.18013.5.1\&quot; -- the primary namespace of a Mobile Driver&#39;s License or a Google Wallet ID Pass - \&quot;org.iso.18013.5.1.aamva\&quot; -- the secondary, AAMVA-defined namespace of a Mobile Driver&#39;s License - \&quot;eu.europa.ec.eudi.pid.1\&quot; -- the primary namespace of an EUDI Wallet PID - \&quot;org.iso.23220.1\&quot; -- the primary namespace of an ISO 23220-2 credential or an Apple Wallet ID Pass | 

## Methods

### NewRaw18013DocumentRequest

`func NewRaw18013DocumentRequest(documentType string, nameSpaces map[string]map[string]bool, ) *Raw18013DocumentRequest`

NewRaw18013DocumentRequest instantiates a new Raw18013DocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaw18013DocumentRequestWithDefaults

`func NewRaw18013DocumentRequestWithDefaults() *Raw18013DocumentRequest`

NewRaw18013DocumentRequestWithDefaults instantiates a new Raw18013DocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *Raw18013DocumentRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *Raw18013DocumentRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *Raw18013DocumentRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetNameSpaces

`func (o *Raw18013DocumentRequest) GetNameSpaces() map[string]map[string]bool`

GetNameSpaces returns the NameSpaces field if non-nil, zero value otherwise.

### GetNameSpacesOk

`func (o *Raw18013DocumentRequest) GetNameSpacesOk() (*map[string]map[string]bool, bool)`

GetNameSpacesOk returns a tuple with the NameSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpaces

`func (o *Raw18013DocumentRequest) SetNameSpaces(v map[string]map[string]bool)`

SetNameSpaces sets NameSpaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



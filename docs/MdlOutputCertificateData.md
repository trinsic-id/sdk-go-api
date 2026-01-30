# MdlOutputCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | **string** | The serial number of the certificate | 
**CommonName** | **string** | The common name (CN) of the certificate | 
**StateOrProvinceName** | **string** | The stateOrProvinceName field from the signing certificate.              Per the ISO 18013-5 (mDL) spec, this is an ISO 3166-2:2020 country subdivision code (e.g., \&quot;US-CA\&quot; for California, USA).              May be an empty string for certificates which are not state-specific (e.g., Google Wallet&#39;s ID Pass certificates). | 
**NotBefore** | **time.Time** | The date before which this certificate is not valid. | 
**NotAfter** | **time.Time** | The date after which this certificate is not valid. | 

## Methods

### NewMdlOutputCertificateData

`func NewMdlOutputCertificateData(serialNumber string, commonName string, stateOrProvinceName string, notBefore time.Time, notAfter time.Time, ) *MdlOutputCertificateData`

NewMdlOutputCertificateData instantiates a new MdlOutputCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdlOutputCertificateDataWithDefaults

`func NewMdlOutputCertificateDataWithDefaults() *MdlOutputCertificateData`

NewMdlOutputCertificateDataWithDefaults instantiates a new MdlOutputCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *MdlOutputCertificateData) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MdlOutputCertificateData) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MdlOutputCertificateData) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetCommonName

`func (o *MdlOutputCertificateData) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MdlOutputCertificateData) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MdlOutputCertificateData) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetStateOrProvinceName

`func (o *MdlOutputCertificateData) GetStateOrProvinceName() string`

GetStateOrProvinceName returns the StateOrProvinceName field if non-nil, zero value otherwise.

### GetStateOrProvinceNameOk

`func (o *MdlOutputCertificateData) GetStateOrProvinceNameOk() (*string, bool)`

GetStateOrProvinceNameOk returns a tuple with the StateOrProvinceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvinceName

`func (o *MdlOutputCertificateData) SetStateOrProvinceName(v string)`

SetStateOrProvinceName sets StateOrProvinceName field to given value.


### GetNotBefore

`func (o *MdlOutputCertificateData) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MdlOutputCertificateData) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MdlOutputCertificateData) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *MdlOutputCertificateData) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *MdlOutputCertificateData) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *MdlOutputCertificateData) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



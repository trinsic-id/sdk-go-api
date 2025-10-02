# MdlCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | **string** | The serial number of the certificate | 
**CommonName** | **string** | The common name (CN) of the certificate | 
**StateOrProvinceName** | **string** | The state or province name (ST) of the certificate.              May be an empty string for certificates which are not state-specific (e.g., Google Wallet&#39;s ID Pass certificates). | 
**NotBefore** | **time.Time** | The date before which this certificate is not valid. | 
**NotAfter** | **time.Time** | The date after which this certificate is not valid. | 

## Methods

### NewMdlCertificateData

`func NewMdlCertificateData(serialNumber string, commonName string, stateOrProvinceName string, notBefore time.Time, notAfter time.Time, ) *MdlCertificateData`

NewMdlCertificateData instantiates a new MdlCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdlCertificateDataWithDefaults

`func NewMdlCertificateDataWithDefaults() *MdlCertificateData`

NewMdlCertificateDataWithDefaults instantiates a new MdlCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *MdlCertificateData) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MdlCertificateData) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MdlCertificateData) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetCommonName

`func (o *MdlCertificateData) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MdlCertificateData) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MdlCertificateData) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetStateOrProvinceName

`func (o *MdlCertificateData) GetStateOrProvinceName() string`

GetStateOrProvinceName returns the StateOrProvinceName field if non-nil, zero value otherwise.

### GetStateOrProvinceNameOk

`func (o *MdlCertificateData) GetStateOrProvinceNameOk() (*string, bool)`

GetStateOrProvinceNameOk returns a tuple with the StateOrProvinceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvinceName

`func (o *MdlCertificateData) SetStateOrProvinceName(v string)`

SetStateOrProvinceName sets StateOrProvinceName field to given value.


### GetNotBefore

`func (o *MdlCertificateData) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MdlCertificateData) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MdlCertificateData) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *MdlCertificateData) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *MdlCertificateData) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *MdlCertificateData) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



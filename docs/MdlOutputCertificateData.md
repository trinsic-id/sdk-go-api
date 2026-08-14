# MdlOutputCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **NullableString** | The serial number of the certificate | [optional] 
**CommonName** | Pointer to **NullableString** | The common name (CN) of the certificate | [optional] 
**StateOrProvinceName** | Pointer to **NullableString** | The stateOrProvinceName field from the certificate | [optional] 
**CountryCode** | Pointer to **NullableString** | The countryCode field from the certificate | [optional] 
**NotBefore** | Pointer to **NullableTime** | The date before which this certificate is not valid. | [optional] 
**NotAfter** | Pointer to **NullableTime** | The date after which this certificate is not valid. | [optional] 

## Methods

### NewMdlOutputCertificateData

`func NewMdlOutputCertificateData() *MdlOutputCertificateData`

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

### HasSerialNumber

`func (o *MdlOutputCertificateData) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *MdlOutputCertificateData) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *MdlOutputCertificateData) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
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

### HasCommonName

`func (o *MdlOutputCertificateData) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *MdlOutputCertificateData) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *MdlOutputCertificateData) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
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

### HasStateOrProvinceName

`func (o *MdlOutputCertificateData) HasStateOrProvinceName() bool`

HasStateOrProvinceName returns a boolean if a field has been set.

### SetStateOrProvinceNameNil

`func (o *MdlOutputCertificateData) SetStateOrProvinceNameNil(b bool)`

 SetStateOrProvinceNameNil sets the value for StateOrProvinceName to be an explicit nil

### UnsetStateOrProvinceName
`func (o *MdlOutputCertificateData) UnsetStateOrProvinceName()`

UnsetStateOrProvinceName ensures that no value is present for StateOrProvinceName, not even an explicit nil
### GetCountryCode

`func (o *MdlOutputCertificateData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MdlOutputCertificateData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MdlOutputCertificateData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MdlOutputCertificateData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *MdlOutputCertificateData) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *MdlOutputCertificateData) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
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

### HasNotBefore

`func (o *MdlOutputCertificateData) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *MdlOutputCertificateData) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *MdlOutputCertificateData) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
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

### HasNotAfter

`func (o *MdlOutputCertificateData) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *MdlOutputCertificateData) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *MdlOutputCertificateData) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



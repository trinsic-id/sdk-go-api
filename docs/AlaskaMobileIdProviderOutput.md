# AlaskaMobileIdProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDriversLicense** | Pointer to [**NullableIso180135MobileDriversLicenseCredential**](Iso180135MobileDriversLicenseCredential.md) | A standard 18013-5 Mobile Driver&#39;s License credential, retrieved from the individual&#39;s wallet. | [optional] 
**RawMdlOutput** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the 18013-7 mDL exchange performed through Alaska Mobile ID. | [optional] 

## Methods

### NewAlaskaMobileIdProviderOutput

`func NewAlaskaMobileIdProviderOutput() *AlaskaMobileIdProviderOutput`

NewAlaskaMobileIdProviderOutput instantiates a new AlaskaMobileIdProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlaskaMobileIdProviderOutputWithDefaults

`func NewAlaskaMobileIdProviderOutputWithDefaults() *AlaskaMobileIdProviderOutput`

NewAlaskaMobileIdProviderOutputWithDefaults instantiates a new AlaskaMobileIdProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDriversLicense

`func (o *AlaskaMobileIdProviderOutput) GetMobileDriversLicense() Iso180135MobileDriversLicenseCredential`

GetMobileDriversLicense returns the MobileDriversLicense field if non-nil, zero value otherwise.

### GetMobileDriversLicenseOk

`func (o *AlaskaMobileIdProviderOutput) GetMobileDriversLicenseOk() (*Iso180135MobileDriversLicenseCredential, bool)`

GetMobileDriversLicenseOk returns a tuple with the MobileDriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDriversLicense

`func (o *AlaskaMobileIdProviderOutput) SetMobileDriversLicense(v Iso180135MobileDriversLicenseCredential)`

SetMobileDriversLicense sets MobileDriversLicense field to given value.

### HasMobileDriversLicense

`func (o *AlaskaMobileIdProviderOutput) HasMobileDriversLicense() bool`

HasMobileDriversLicense returns a boolean if a field has been set.

### SetMobileDriversLicenseNil

`func (o *AlaskaMobileIdProviderOutput) SetMobileDriversLicenseNil(b bool)`

 SetMobileDriversLicenseNil sets the value for MobileDriversLicense to be an explicit nil

### UnsetMobileDriversLicense
`func (o *AlaskaMobileIdProviderOutput) UnsetMobileDriversLicense()`

UnsetMobileDriversLicense ensures that no value is present for MobileDriversLicense, not even an explicit nil
### GetRawMdlOutput

`func (o *AlaskaMobileIdProviderOutput) GetRawMdlOutput() MdlOutput`

GetRawMdlOutput returns the RawMdlOutput field if non-nil, zero value otherwise.

### GetRawMdlOutputOk

`func (o *AlaskaMobileIdProviderOutput) GetRawMdlOutputOk() (*MdlOutput, bool)`

GetRawMdlOutputOk returns a tuple with the RawMdlOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMdlOutput

`func (o *AlaskaMobileIdProviderOutput) SetRawMdlOutput(v MdlOutput)`

SetRawMdlOutput sets RawMdlOutput field to given value.

### HasRawMdlOutput

`func (o *AlaskaMobileIdProviderOutput) HasRawMdlOutput() bool`

HasRawMdlOutput returns a boolean if a field has been set.

### SetRawMdlOutputNil

`func (o *AlaskaMobileIdProviderOutput) SetRawMdlOutputNil(b bool)`

 SetRawMdlOutputNil sets the value for RawMdlOutput to be an explicit nil

### UnsetRawMdlOutput
`func (o *AlaskaMobileIdProviderOutput) UnsetRawMdlOutput()`

UnsetRawMdlOutput ensures that no value is present for RawMdlOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SamsungWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDriversLicense** | Pointer to [**NullableIso180135MobileDriversLicenseCredential**](Iso180135MobileDriversLicenseCredential.md) | A standard 18013-5 Mobile Driver&#39;s License credential, retrieved from the individual&#39;s wallet. | [optional] 
**SamsungIdWithClear** | Pointer to [**NullableSamsungIdWithClearCredential**](SamsungIdWithClearCredential.md) | A Samsung ID with CLEAR credential, retrieved from the individual&#39;s Samsung Wallet. | [optional] 
**RawMdlOutput** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the mDL exchange performed through Samsung Wallet. | [optional] 

## Methods

### NewSamsungWalletProviderOutput

`func NewSamsungWalletProviderOutput() *SamsungWalletProviderOutput`

NewSamsungWalletProviderOutput instantiates a new SamsungWalletProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamsungWalletProviderOutputWithDefaults

`func NewSamsungWalletProviderOutputWithDefaults() *SamsungWalletProviderOutput`

NewSamsungWalletProviderOutputWithDefaults instantiates a new SamsungWalletProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDriversLicense

`func (o *SamsungWalletProviderOutput) GetMobileDriversLicense() Iso180135MobileDriversLicenseCredential`

GetMobileDriversLicense returns the MobileDriversLicense field if non-nil, zero value otherwise.

### GetMobileDriversLicenseOk

`func (o *SamsungWalletProviderOutput) GetMobileDriversLicenseOk() (*Iso180135MobileDriversLicenseCredential, bool)`

GetMobileDriversLicenseOk returns a tuple with the MobileDriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDriversLicense

`func (o *SamsungWalletProviderOutput) SetMobileDriversLicense(v Iso180135MobileDriversLicenseCredential)`

SetMobileDriversLicense sets MobileDriversLicense field to given value.

### HasMobileDriversLicense

`func (o *SamsungWalletProviderOutput) HasMobileDriversLicense() bool`

HasMobileDriversLicense returns a boolean if a field has been set.

### SetMobileDriversLicenseNil

`func (o *SamsungWalletProviderOutput) SetMobileDriversLicenseNil(b bool)`

 SetMobileDriversLicenseNil sets the value for MobileDriversLicense to be an explicit nil

### UnsetMobileDriversLicense
`func (o *SamsungWalletProviderOutput) UnsetMobileDriversLicense()`

UnsetMobileDriversLicense ensures that no value is present for MobileDriversLicense, not even an explicit nil
### GetSamsungIdWithClear

`func (o *SamsungWalletProviderOutput) GetSamsungIdWithClear() SamsungIdWithClearCredential`

GetSamsungIdWithClear returns the SamsungIdWithClear field if non-nil, zero value otherwise.

### GetSamsungIdWithClearOk

`func (o *SamsungWalletProviderOutput) GetSamsungIdWithClearOk() (*SamsungIdWithClearCredential, bool)`

GetSamsungIdWithClearOk returns a tuple with the SamsungIdWithClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamsungIdWithClear

`func (o *SamsungWalletProviderOutput) SetSamsungIdWithClear(v SamsungIdWithClearCredential)`

SetSamsungIdWithClear sets SamsungIdWithClear field to given value.

### HasSamsungIdWithClear

`func (o *SamsungWalletProviderOutput) HasSamsungIdWithClear() bool`

HasSamsungIdWithClear returns a boolean if a field has been set.

### SetSamsungIdWithClearNil

`func (o *SamsungWalletProviderOutput) SetSamsungIdWithClearNil(b bool)`

 SetSamsungIdWithClearNil sets the value for SamsungIdWithClear to be an explicit nil

### UnsetSamsungIdWithClear
`func (o *SamsungWalletProviderOutput) UnsetSamsungIdWithClear()`

UnsetSamsungIdWithClear ensures that no value is present for SamsungIdWithClear, not even an explicit nil
### GetRawMdlOutput

`func (o *SamsungWalletProviderOutput) GetRawMdlOutput() MdlOutput`

GetRawMdlOutput returns the RawMdlOutput field if non-nil, zero value otherwise.

### GetRawMdlOutputOk

`func (o *SamsungWalletProviderOutput) GetRawMdlOutputOk() (*MdlOutput, bool)`

GetRawMdlOutputOk returns a tuple with the RawMdlOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMdlOutput

`func (o *SamsungWalletProviderOutput) SetRawMdlOutput(v MdlOutput)`

SetRawMdlOutput sets RawMdlOutput field to given value.

### HasRawMdlOutput

`func (o *SamsungWalletProviderOutput) HasRawMdlOutput() bool`

HasRawMdlOutput returns a boolean if a field has been set.

### SetRawMdlOutputNil

`func (o *SamsungWalletProviderOutput) SetRawMdlOutputNil(b bool)`

 SetRawMdlOutputNil sets the value for RawMdlOutput to be an explicit nil

### UnsetRawMdlOutput
`func (o *SamsungWalletProviderOutput) UnsetRawMdlOutput()`

UnsetRawMdlOutput ensures that no value is present for RawMdlOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AppleWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDriversLicense** | Pointer to [**NullableIso180135MobileDriversLicenseCredential**](Iso180135MobileDriversLicenseCredential.md) | A standard 18013-5 Mobile Driver&#39;s License credential, retrieved from the individual&#39;s wallet. | [optional] 
**DigitalId** | Pointer to [**NullableAppleWalletDigitalIdCredential**](AppleWalletDigitalIdCredential.md) | An Apple Wallet Digital ID credential (&#x60;org.iso.23220.photoid.1&#x60;), retrieved from the individual&#39;s wallet. | [optional] 
**RawMdlOutput** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the mDL exchange performed through Apple Wallet. | [optional] 

## Methods

### NewAppleWalletProviderOutput

`func NewAppleWalletProviderOutput() *AppleWalletProviderOutput`

NewAppleWalletProviderOutput instantiates a new AppleWalletProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleWalletProviderOutputWithDefaults

`func NewAppleWalletProviderOutputWithDefaults() *AppleWalletProviderOutput`

NewAppleWalletProviderOutputWithDefaults instantiates a new AppleWalletProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDriversLicense

`func (o *AppleWalletProviderOutput) GetMobileDriversLicense() Iso180135MobileDriversLicenseCredential`

GetMobileDriversLicense returns the MobileDriversLicense field if non-nil, zero value otherwise.

### GetMobileDriversLicenseOk

`func (o *AppleWalletProviderOutput) GetMobileDriversLicenseOk() (*Iso180135MobileDriversLicenseCredential, bool)`

GetMobileDriversLicenseOk returns a tuple with the MobileDriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDriversLicense

`func (o *AppleWalletProviderOutput) SetMobileDriversLicense(v Iso180135MobileDriversLicenseCredential)`

SetMobileDriversLicense sets MobileDriversLicense field to given value.

### HasMobileDriversLicense

`func (o *AppleWalletProviderOutput) HasMobileDriversLicense() bool`

HasMobileDriversLicense returns a boolean if a field has been set.

### SetMobileDriversLicenseNil

`func (o *AppleWalletProviderOutput) SetMobileDriversLicenseNil(b bool)`

 SetMobileDriversLicenseNil sets the value for MobileDriversLicense to be an explicit nil

### UnsetMobileDriversLicense
`func (o *AppleWalletProviderOutput) UnsetMobileDriversLicense()`

UnsetMobileDriversLicense ensures that no value is present for MobileDriversLicense, not even an explicit nil
### GetDigitalId

`func (o *AppleWalletProviderOutput) GetDigitalId() AppleWalletDigitalIdCredential`

GetDigitalId returns the DigitalId field if non-nil, zero value otherwise.

### GetDigitalIdOk

`func (o *AppleWalletProviderOutput) GetDigitalIdOk() (*AppleWalletDigitalIdCredential, bool)`

GetDigitalIdOk returns a tuple with the DigitalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalId

`func (o *AppleWalletProviderOutput) SetDigitalId(v AppleWalletDigitalIdCredential)`

SetDigitalId sets DigitalId field to given value.

### HasDigitalId

`func (o *AppleWalletProviderOutput) HasDigitalId() bool`

HasDigitalId returns a boolean if a field has been set.

### SetDigitalIdNil

`func (o *AppleWalletProviderOutput) SetDigitalIdNil(b bool)`

 SetDigitalIdNil sets the value for DigitalId to be an explicit nil

### UnsetDigitalId
`func (o *AppleWalletProviderOutput) UnsetDigitalId()`

UnsetDigitalId ensures that no value is present for DigitalId, not even an explicit nil
### GetRawMdlOutput

`func (o *AppleWalletProviderOutput) GetRawMdlOutput() MdlOutput`

GetRawMdlOutput returns the RawMdlOutput field if non-nil, zero value otherwise.

### GetRawMdlOutputOk

`func (o *AppleWalletProviderOutput) GetRawMdlOutputOk() (*MdlOutput, bool)`

GetRawMdlOutputOk returns a tuple with the RawMdlOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMdlOutput

`func (o *AppleWalletProviderOutput) SetRawMdlOutput(v MdlOutput)`

SetRawMdlOutput sets RawMdlOutput field to given value.

### HasRawMdlOutput

`func (o *AppleWalletProviderOutput) HasRawMdlOutput() bool`

HasRawMdlOutput returns a boolean if a field has been set.

### SetRawMdlOutputNil

`func (o *AppleWalletProviderOutput) SetRawMdlOutputNil(b bool)`

 SetRawMdlOutputNil sets the value for RawMdlOutput to be an explicit nil

### UnsetRawMdlOutput
`func (o *AppleWalletProviderOutput) UnsetRawMdlOutput()`

UnsetRawMdlOutput ensures that no value is present for RawMdlOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GoogleWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDriversLicense** | Pointer to [**NullableIso180135MobileDriversLicenseCredential**](Iso180135MobileDriversLicenseCredential.md) | A standard 18013-5 Mobile Driver&#39;s License credential, retrieved from the individual&#39;s wallet. | [optional] 
**IdPass** | Pointer to [**NullableGoogleWalletIdPassCredential**](GoogleWalletIdPassCredential.md) | A Google Wallet ID Pass credential (&#x60;com.google.wallet.idcard.1&#x60;), retrieved from the individual&#39;s wallet. | [optional] 
**RawMdlOutput** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the mDL exchange performed through Google Wallet. | [optional] 

## Methods

### NewGoogleWalletProviderOutput

`func NewGoogleWalletProviderOutput() *GoogleWalletProviderOutput`

NewGoogleWalletProviderOutput instantiates a new GoogleWalletProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWalletProviderOutputWithDefaults

`func NewGoogleWalletProviderOutputWithDefaults() *GoogleWalletProviderOutput`

NewGoogleWalletProviderOutputWithDefaults instantiates a new GoogleWalletProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDriversLicense

`func (o *GoogleWalletProviderOutput) GetMobileDriversLicense() Iso180135MobileDriversLicenseCredential`

GetMobileDriversLicense returns the MobileDriversLicense field if non-nil, zero value otherwise.

### GetMobileDriversLicenseOk

`func (o *GoogleWalletProviderOutput) GetMobileDriversLicenseOk() (*Iso180135MobileDriversLicenseCredential, bool)`

GetMobileDriversLicenseOk returns a tuple with the MobileDriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDriversLicense

`func (o *GoogleWalletProviderOutput) SetMobileDriversLicense(v Iso180135MobileDriversLicenseCredential)`

SetMobileDriversLicense sets MobileDriversLicense field to given value.

### HasMobileDriversLicense

`func (o *GoogleWalletProviderOutput) HasMobileDriversLicense() bool`

HasMobileDriversLicense returns a boolean if a field has been set.

### SetMobileDriversLicenseNil

`func (o *GoogleWalletProviderOutput) SetMobileDriversLicenseNil(b bool)`

 SetMobileDriversLicenseNil sets the value for MobileDriversLicense to be an explicit nil

### UnsetMobileDriversLicense
`func (o *GoogleWalletProviderOutput) UnsetMobileDriversLicense()`

UnsetMobileDriversLicense ensures that no value is present for MobileDriversLicense, not even an explicit nil
### GetIdPass

`func (o *GoogleWalletProviderOutput) GetIdPass() GoogleWalletIdPassCredential`

GetIdPass returns the IdPass field if non-nil, zero value otherwise.

### GetIdPassOk

`func (o *GoogleWalletProviderOutput) GetIdPassOk() (*GoogleWalletIdPassCredential, bool)`

GetIdPassOk returns a tuple with the IdPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPass

`func (o *GoogleWalletProviderOutput) SetIdPass(v GoogleWalletIdPassCredential)`

SetIdPass sets IdPass field to given value.

### HasIdPass

`func (o *GoogleWalletProviderOutput) HasIdPass() bool`

HasIdPass returns a boolean if a field has been set.

### SetIdPassNil

`func (o *GoogleWalletProviderOutput) SetIdPassNil(b bool)`

 SetIdPassNil sets the value for IdPass to be an explicit nil

### UnsetIdPass
`func (o *GoogleWalletProviderOutput) UnsetIdPass()`

UnsetIdPass ensures that no value is present for IdPass, not even an explicit nil
### GetRawMdlOutput

`func (o *GoogleWalletProviderOutput) GetRawMdlOutput() MdlOutput`

GetRawMdlOutput returns the RawMdlOutput field if non-nil, zero value otherwise.

### GetRawMdlOutputOk

`func (o *GoogleWalletProviderOutput) GetRawMdlOutputOk() (*MdlOutput, bool)`

GetRawMdlOutputOk returns a tuple with the RawMdlOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMdlOutput

`func (o *GoogleWalletProviderOutput) SetRawMdlOutput(v MdlOutput)`

SetRawMdlOutput sets RawMdlOutput field to given value.

### HasRawMdlOutput

`func (o *GoogleWalletProviderOutput) HasRawMdlOutput() bool`

HasRawMdlOutput returns a boolean if a field has been set.

### SetRawMdlOutputNil

`func (o *GoogleWalletProviderOutput) SetRawMdlOutputNil(b bool)`

 SetRawMdlOutputNil sets the value for RawMdlOutput to be an explicit nil

### UnsetRawMdlOutput
`func (o *GoogleWalletProviderOutput) UnsetRawMdlOutput()`

UnsetRawMdlOutput ensures that no value is present for RawMdlOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



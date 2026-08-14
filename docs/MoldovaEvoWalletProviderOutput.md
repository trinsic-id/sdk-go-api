# MoldovaEvoWalletProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityCard** | Pointer to [**NullableMoldovaIdentityCardCredential**](MoldovaIdentityCardCredential.md) | A Moldovan Identity Card credential, retrieved from the individual&#39;s wallet. | [optional] 
**DriverLicense** | Pointer to [**NullableMoldovaDriverLicenseCredential**](MoldovaDriverLicenseCredential.md) | A Moldovan Driver License credential, retrieved from the individual&#39;s wallet. | [optional] 
**Pid** | Pointer to [**NullableEudiPidCredential**](EudiPidCredential.md) | An EUDI Person Identification Data (PID) credential, retrieved from the individual&#39;s wallet.              NOTE: Although this credential aligns with the EUDI PID specification, Moldova is not an EU member state. | [optional] 
**VehicleRegistrationCertificate** | Pointer to [**NullableMoldovaVehicleRegistrationCertificateCredential**](MoldovaVehicleRegistrationCertificateCredential.md) | A Moldovan Vehicle Registration Certificate credential, retrieved from the individual&#39;s wallet. | [optional] 
**Raw18013Output** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the 18013-7 exchange performed through EVO Wallet. | [optional] 

## Methods

### NewMoldovaEvoWalletProviderOutput

`func NewMoldovaEvoWalletProviderOutput() *MoldovaEvoWalletProviderOutput`

NewMoldovaEvoWalletProviderOutput instantiates a new MoldovaEvoWalletProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoldovaEvoWalletProviderOutputWithDefaults

`func NewMoldovaEvoWalletProviderOutputWithDefaults() *MoldovaEvoWalletProviderOutput`

NewMoldovaEvoWalletProviderOutputWithDefaults instantiates a new MoldovaEvoWalletProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityCard

`func (o *MoldovaEvoWalletProviderOutput) GetIdentityCard() MoldovaIdentityCardCredential`

GetIdentityCard returns the IdentityCard field if non-nil, zero value otherwise.

### GetIdentityCardOk

`func (o *MoldovaEvoWalletProviderOutput) GetIdentityCardOk() (*MoldovaIdentityCardCredential, bool)`

GetIdentityCardOk returns a tuple with the IdentityCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCard

`func (o *MoldovaEvoWalletProviderOutput) SetIdentityCard(v MoldovaIdentityCardCredential)`

SetIdentityCard sets IdentityCard field to given value.

### HasIdentityCard

`func (o *MoldovaEvoWalletProviderOutput) HasIdentityCard() bool`

HasIdentityCard returns a boolean if a field has been set.

### SetIdentityCardNil

`func (o *MoldovaEvoWalletProviderOutput) SetIdentityCardNil(b bool)`

 SetIdentityCardNil sets the value for IdentityCard to be an explicit nil

### UnsetIdentityCard
`func (o *MoldovaEvoWalletProviderOutput) UnsetIdentityCard()`

UnsetIdentityCard ensures that no value is present for IdentityCard, not even an explicit nil
### GetDriverLicense

`func (o *MoldovaEvoWalletProviderOutput) GetDriverLicense() MoldovaDriverLicenseCredential`

GetDriverLicense returns the DriverLicense field if non-nil, zero value otherwise.

### GetDriverLicenseOk

`func (o *MoldovaEvoWalletProviderOutput) GetDriverLicenseOk() (*MoldovaDriverLicenseCredential, bool)`

GetDriverLicenseOk returns a tuple with the DriverLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicense

`func (o *MoldovaEvoWalletProviderOutput) SetDriverLicense(v MoldovaDriverLicenseCredential)`

SetDriverLicense sets DriverLicense field to given value.

### HasDriverLicense

`func (o *MoldovaEvoWalletProviderOutput) HasDriverLicense() bool`

HasDriverLicense returns a boolean if a field has been set.

### SetDriverLicenseNil

`func (o *MoldovaEvoWalletProviderOutput) SetDriverLicenseNil(b bool)`

 SetDriverLicenseNil sets the value for DriverLicense to be an explicit nil

### UnsetDriverLicense
`func (o *MoldovaEvoWalletProviderOutput) UnsetDriverLicense()`

UnsetDriverLicense ensures that no value is present for DriverLicense, not even an explicit nil
### GetPid

`func (o *MoldovaEvoWalletProviderOutput) GetPid() EudiPidCredential`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *MoldovaEvoWalletProviderOutput) GetPidOk() (*EudiPidCredential, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *MoldovaEvoWalletProviderOutput) SetPid(v EudiPidCredential)`

SetPid sets Pid field to given value.

### HasPid

`func (o *MoldovaEvoWalletProviderOutput) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *MoldovaEvoWalletProviderOutput) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *MoldovaEvoWalletProviderOutput) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetVehicleRegistrationCertificate

`func (o *MoldovaEvoWalletProviderOutput) GetVehicleRegistrationCertificate() MoldovaVehicleRegistrationCertificateCredential`

GetVehicleRegistrationCertificate returns the VehicleRegistrationCertificate field if non-nil, zero value otherwise.

### GetVehicleRegistrationCertificateOk

`func (o *MoldovaEvoWalletProviderOutput) GetVehicleRegistrationCertificateOk() (*MoldovaVehicleRegistrationCertificateCredential, bool)`

GetVehicleRegistrationCertificateOk returns a tuple with the VehicleRegistrationCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleRegistrationCertificate

`func (o *MoldovaEvoWalletProviderOutput) SetVehicleRegistrationCertificate(v MoldovaVehicleRegistrationCertificateCredential)`

SetVehicleRegistrationCertificate sets VehicleRegistrationCertificate field to given value.

### HasVehicleRegistrationCertificate

`func (o *MoldovaEvoWalletProviderOutput) HasVehicleRegistrationCertificate() bool`

HasVehicleRegistrationCertificate returns a boolean if a field has been set.

### SetVehicleRegistrationCertificateNil

`func (o *MoldovaEvoWalletProviderOutput) SetVehicleRegistrationCertificateNil(b bool)`

 SetVehicleRegistrationCertificateNil sets the value for VehicleRegistrationCertificate to be an explicit nil

### UnsetVehicleRegistrationCertificate
`func (o *MoldovaEvoWalletProviderOutput) UnsetVehicleRegistrationCertificate()`

UnsetVehicleRegistrationCertificate ensures that no value is present for VehicleRegistrationCertificate, not even an explicit nil
### GetRaw18013Output

`func (o *MoldovaEvoWalletProviderOutput) GetRaw18013Output() MdlOutput`

GetRaw18013Output returns the Raw18013Output field if non-nil, zero value otherwise.

### GetRaw18013OutputOk

`func (o *MoldovaEvoWalletProviderOutput) GetRaw18013OutputOk() (*MdlOutput, bool)`

GetRaw18013OutputOk returns a tuple with the Raw18013Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw18013Output

`func (o *MoldovaEvoWalletProviderOutput) SetRaw18013Output(v MdlOutput)`

SetRaw18013Output sets Raw18013Output field to given value.

### HasRaw18013Output

`func (o *MoldovaEvoWalletProviderOutput) HasRaw18013Output() bool`

HasRaw18013Output returns a boolean if a field has been set.

### SetRaw18013OutputNil

`func (o *MoldovaEvoWalletProviderOutput) SetRaw18013OutputNil(b bool)`

 SetRaw18013OutputNil sets the value for Raw18013Output to be an explicit nil

### UnsetRaw18013Output
`func (o *MoldovaEvoWalletProviderOutput) UnsetRaw18013Output()`

UnsetRaw18013Output ensures that no value is present for Raw18013Output, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



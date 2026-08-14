# ItsmeDirectDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystem** | Pointer to **NullableString** | The operating system reported for the itsme device.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**OperatingSystemRelease** | Pointer to **NullableString** | The operating system release reported for the itsme device.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**AppName** | Pointer to **NullableString** | The itsme app name reported for the verification.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**AppRelease** | Pointer to **NullableString** | The itsme app release reported for the verification.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**Manufacturer** | Pointer to **NullableString** | The device manufacturer reported for the itsme device.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**DeviceModel** | Pointer to **NullableString** | The device model reported for the itsme device.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**DeviceLabel** | Pointer to **NullableString** | The device label.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**DeviceId** | Pointer to **NullableString** | The device identifier.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**DebugEnabled** | Pointer to **NullableBool** | Whether the itsme device had debugging enabled.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**Rooted** | Pointer to **NullableBool** | Whether the itsme device was reported as rooted.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**DeviceLockLevel** | Pointer to **NullableString** | The device lock level value.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 
**Msisdn** | Pointer to **NullableString** | The Mobile Station International Subscriber Directory Number (MSISDN) value for the device.              Availability by ID document issuing country: best effort for all supported issuing countries. | [optional] 

## Methods

### NewItsmeDirectDevice

`func NewItsmeDirectDevice() *ItsmeDirectDevice`

NewItsmeDirectDevice instantiates a new ItsmeDirectDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItsmeDirectDeviceWithDefaults

`func NewItsmeDirectDeviceWithDefaults() *ItsmeDirectDevice`

NewItsmeDirectDeviceWithDefaults instantiates a new ItsmeDirectDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystem

`func (o *ItsmeDirectDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ItsmeDirectDevice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ItsmeDirectDevice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ItsmeDirectDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *ItsmeDirectDevice) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *ItsmeDirectDevice) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetOperatingSystemRelease

`func (o *ItsmeDirectDevice) GetOperatingSystemRelease() string`

GetOperatingSystemRelease returns the OperatingSystemRelease field if non-nil, zero value otherwise.

### GetOperatingSystemReleaseOk

`func (o *ItsmeDirectDevice) GetOperatingSystemReleaseOk() (*string, bool)`

GetOperatingSystemReleaseOk returns a tuple with the OperatingSystemRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemRelease

`func (o *ItsmeDirectDevice) SetOperatingSystemRelease(v string)`

SetOperatingSystemRelease sets OperatingSystemRelease field to given value.

### HasOperatingSystemRelease

`func (o *ItsmeDirectDevice) HasOperatingSystemRelease() bool`

HasOperatingSystemRelease returns a boolean if a field has been set.

### SetOperatingSystemReleaseNil

`func (o *ItsmeDirectDevice) SetOperatingSystemReleaseNil(b bool)`

 SetOperatingSystemReleaseNil sets the value for OperatingSystemRelease to be an explicit nil

### UnsetOperatingSystemRelease
`func (o *ItsmeDirectDevice) UnsetOperatingSystemRelease()`

UnsetOperatingSystemRelease ensures that no value is present for OperatingSystemRelease, not even an explicit nil
### GetAppName

`func (o *ItsmeDirectDevice) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ItsmeDirectDevice) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ItsmeDirectDevice) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ItsmeDirectDevice) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *ItsmeDirectDevice) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *ItsmeDirectDevice) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppRelease

`func (o *ItsmeDirectDevice) GetAppRelease() string`

GetAppRelease returns the AppRelease field if non-nil, zero value otherwise.

### GetAppReleaseOk

`func (o *ItsmeDirectDevice) GetAppReleaseOk() (*string, bool)`

GetAppReleaseOk returns a tuple with the AppRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRelease

`func (o *ItsmeDirectDevice) SetAppRelease(v string)`

SetAppRelease sets AppRelease field to given value.

### HasAppRelease

`func (o *ItsmeDirectDevice) HasAppRelease() bool`

HasAppRelease returns a boolean if a field has been set.

### SetAppReleaseNil

`func (o *ItsmeDirectDevice) SetAppReleaseNil(b bool)`

 SetAppReleaseNil sets the value for AppRelease to be an explicit nil

### UnsetAppRelease
`func (o *ItsmeDirectDevice) UnsetAppRelease()`

UnsetAppRelease ensures that no value is present for AppRelease, not even an explicit nil
### GetManufacturer

`func (o *ItsmeDirectDevice) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ItsmeDirectDevice) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ItsmeDirectDevice) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ItsmeDirectDevice) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *ItsmeDirectDevice) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *ItsmeDirectDevice) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetDeviceModel

`func (o *ItsmeDirectDevice) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *ItsmeDirectDevice) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *ItsmeDirectDevice) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *ItsmeDirectDevice) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *ItsmeDirectDevice) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *ItsmeDirectDevice) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetDeviceLabel

`func (o *ItsmeDirectDevice) GetDeviceLabel() string`

GetDeviceLabel returns the DeviceLabel field if non-nil, zero value otherwise.

### GetDeviceLabelOk

`func (o *ItsmeDirectDevice) GetDeviceLabelOk() (*string, bool)`

GetDeviceLabelOk returns a tuple with the DeviceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLabel

`func (o *ItsmeDirectDevice) SetDeviceLabel(v string)`

SetDeviceLabel sets DeviceLabel field to given value.

### HasDeviceLabel

`func (o *ItsmeDirectDevice) HasDeviceLabel() bool`

HasDeviceLabel returns a boolean if a field has been set.

### SetDeviceLabelNil

`func (o *ItsmeDirectDevice) SetDeviceLabelNil(b bool)`

 SetDeviceLabelNil sets the value for DeviceLabel to be an explicit nil

### UnsetDeviceLabel
`func (o *ItsmeDirectDevice) UnsetDeviceLabel()`

UnsetDeviceLabel ensures that no value is present for DeviceLabel, not even an explicit nil
### GetDeviceId

`func (o *ItsmeDirectDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ItsmeDirectDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ItsmeDirectDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ItsmeDirectDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *ItsmeDirectDevice) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *ItsmeDirectDevice) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDebugEnabled

`func (o *ItsmeDirectDevice) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *ItsmeDirectDevice) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *ItsmeDirectDevice) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *ItsmeDirectDevice) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### SetDebugEnabledNil

`func (o *ItsmeDirectDevice) SetDebugEnabledNil(b bool)`

 SetDebugEnabledNil sets the value for DebugEnabled to be an explicit nil

### UnsetDebugEnabled
`func (o *ItsmeDirectDevice) UnsetDebugEnabled()`

UnsetDebugEnabled ensures that no value is present for DebugEnabled, not even an explicit nil
### GetRooted

`func (o *ItsmeDirectDevice) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *ItsmeDirectDevice) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *ItsmeDirectDevice) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *ItsmeDirectDevice) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### SetRootedNil

`func (o *ItsmeDirectDevice) SetRootedNil(b bool)`

 SetRootedNil sets the value for Rooted to be an explicit nil

### UnsetRooted
`func (o *ItsmeDirectDevice) UnsetRooted()`

UnsetRooted ensures that no value is present for Rooted, not even an explicit nil
### GetDeviceLockLevel

`func (o *ItsmeDirectDevice) GetDeviceLockLevel() string`

GetDeviceLockLevel returns the DeviceLockLevel field if non-nil, zero value otherwise.

### GetDeviceLockLevelOk

`func (o *ItsmeDirectDevice) GetDeviceLockLevelOk() (*string, bool)`

GetDeviceLockLevelOk returns a tuple with the DeviceLockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLockLevel

`func (o *ItsmeDirectDevice) SetDeviceLockLevel(v string)`

SetDeviceLockLevel sets DeviceLockLevel field to given value.

### HasDeviceLockLevel

`func (o *ItsmeDirectDevice) HasDeviceLockLevel() bool`

HasDeviceLockLevel returns a boolean if a field has been set.

### SetDeviceLockLevelNil

`func (o *ItsmeDirectDevice) SetDeviceLockLevelNil(b bool)`

 SetDeviceLockLevelNil sets the value for DeviceLockLevel to be an explicit nil

### UnsetDeviceLockLevel
`func (o *ItsmeDirectDevice) UnsetDeviceLockLevel()`

UnsetDeviceLockLevel ensures that no value is present for DeviceLockLevel, not even an explicit nil
### GetMsisdn

`func (o *ItsmeDirectDevice) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *ItsmeDirectDevice) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *ItsmeDirectDevice) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *ItsmeDirectDevice) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### SetMsisdnNil

`func (o *ItsmeDirectDevice) SetMsisdnNil(b bool)`

 SetMsisdnNil sets the value for Msisdn to be an explicit nil

### UnsetMsisdn
`func (o *ItsmeDirectDevice) UnsetMsisdn()`

UnsetMsisdn ensures that no value is present for Msisdn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



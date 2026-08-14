# KoreaTelcoMatchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **NullableString** | The phone number submitted for the carrier match. | [optional] 
**TeleType** | Pointer to **NullableString** | The mobile carrier used for the match.              Supported values: Lgu, Skt, Kt. | [optional] 
**ResultCode** | Pointer to **NullableString** | The carrier match result code.              Common result codes: - \&quot;0000\&quot;: Successful match - \&quot;0001\&quot;: Failed - Verification Information Mismatch (General) - \&quot;0002\&quot;: Failed - Unable to Verify Phone Number - \&quot;0004\&quot;: Failed - Date of Birth Verification Error - \&quot;0005\&quot;: Failed - Gender Verification Error - \&quot;0006\&quot;: Failed - Name Verification Error - \&quot;0009\&quot;: Failed - Device OS Mismatch | [optional] 

## Methods

### NewKoreaTelcoMatchProviderOutput

`func NewKoreaTelcoMatchProviderOutput() *KoreaTelcoMatchProviderOutput`

NewKoreaTelcoMatchProviderOutput instantiates a new KoreaTelcoMatchProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKoreaTelcoMatchProviderOutputWithDefaults

`func NewKoreaTelcoMatchProviderOutputWithDefaults() *KoreaTelcoMatchProviderOutput`

NewKoreaTelcoMatchProviderOutputWithDefaults instantiates a new KoreaTelcoMatchProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *KoreaTelcoMatchProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *KoreaTelcoMatchProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *KoreaTelcoMatchProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *KoreaTelcoMatchProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *KoreaTelcoMatchProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *KoreaTelcoMatchProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetTeleType

`func (o *KoreaTelcoMatchProviderOutput) GetTeleType() string`

GetTeleType returns the TeleType field if non-nil, zero value otherwise.

### GetTeleTypeOk

`func (o *KoreaTelcoMatchProviderOutput) GetTeleTypeOk() (*string, bool)`

GetTeleTypeOk returns a tuple with the TeleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeleType

`func (o *KoreaTelcoMatchProviderOutput) SetTeleType(v string)`

SetTeleType sets TeleType field to given value.

### HasTeleType

`func (o *KoreaTelcoMatchProviderOutput) HasTeleType() bool`

HasTeleType returns a boolean if a field has been set.

### SetTeleTypeNil

`func (o *KoreaTelcoMatchProviderOutput) SetTeleTypeNil(b bool)`

 SetTeleTypeNil sets the value for TeleType to be an explicit nil

### UnsetTeleType
`func (o *KoreaTelcoMatchProviderOutput) UnsetTeleType()`

UnsetTeleType ensures that no value is present for TeleType, not even an explicit nil
### GetResultCode

`func (o *KoreaTelcoMatchProviderOutput) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *KoreaTelcoMatchProviderOutput) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *KoreaTelcoMatchProviderOutput) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *KoreaTelcoMatchProviderOutput) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### SetResultCodeNil

`func (o *KoreaTelcoMatchProviderOutput) SetResultCodeNil(b bool)`

 SetResultCodeNil sets the value for ResultCode to be an explicit nil

### UnsetResultCode
`func (o *KoreaTelcoMatchProviderOutput) UnsetResultCode()`

UnsetResultCode ensures that no value is present for ResultCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



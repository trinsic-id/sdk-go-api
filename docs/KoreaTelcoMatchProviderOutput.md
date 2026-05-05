# KoreaTelcoMatchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | The verified phone number as submitted by the relying party. | 
**TeleType** | **string** | The mobile carrier used for verification. Possible values: &lt;list type&#x3D;\&quot;bullet\&quot;&gt;&lt;item&gt;&lt;description&gt;Lgu&lt;/description&gt;&lt;/item&gt;&lt;item&gt;&lt;description&gt;Skt&lt;/description&gt;&lt;/item&gt;&lt;item&gt;&lt;description&gt;Kt&lt;/description&gt;&lt;/item&gt;&lt;/list&gt; | 
**ResultCode** | **string** | The RaonSecure result code returned by the carrier verification system. \&quot;0000\&quot; indicates a successful match.              Common error codes: - \&quot;0000\&quot;: Successful match - \&quot;0001\&quot;: Identity Verification Failed - Verification Information Mismatch (General) - \&quot;0002\&quot;: Identity Verification Failed - Unable to Verify Phone Number - \&quot;0004\&quot;: Identity Verification Failed - Date of Birth Verification Error - \&quot;0005\&quot;: Identity Verification Failed - Gender Verification Error - \&quot;0006\&quot;: Identity Verification Failed - Name Verification Error - \&quot;0009\&quot;: Identity Verification Failed - Device OS Mismatch | 

## Methods

### NewKoreaTelcoMatchProviderOutput

`func NewKoreaTelcoMatchProviderOutput(phoneNumber string, teleType string, resultCode string, ) *KoreaTelcoMatchProviderOutput`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KoreaTelcoMatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Phone number without dashes, e.g. \&quot;010XXXXXXXX\&quot;. | 
**FullName** | **string** | Full legal name (UTF-8), e.g. \&quot;홍길동\&quot; or \&quot;Hong Gildong\&quot;. | 
**DateOfBirth** | **string** | Date of birth. | 
**Sex** | [**KoreaTelcoMatchSex**](KoreaTelcoMatchSex.md) | Sex. Must be Male or Female. | 
**Carrier** | [**MobileCarrier**](MobileCarrier.md) | Mobile carrier. Lgu &#x3D; LG U+, Skt &#x3D; SK Telecom, Kt &#x3D; KT. | 
**OperatingSystem** | [**MobileOperatingSystem**](MobileOperatingSystem.md) | Device operating system. | 

## Methods

### NewKoreaTelcoMatchInput

`func NewKoreaTelcoMatchInput(phoneNumber string, fullName string, dateOfBirth string, sex KoreaTelcoMatchSex, carrier MobileCarrier, operatingSystem MobileOperatingSystem, ) *KoreaTelcoMatchInput`

NewKoreaTelcoMatchInput instantiates a new KoreaTelcoMatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKoreaTelcoMatchInputWithDefaults

`func NewKoreaTelcoMatchInputWithDefaults() *KoreaTelcoMatchInput`

NewKoreaTelcoMatchInputWithDefaults instantiates a new KoreaTelcoMatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *KoreaTelcoMatchInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *KoreaTelcoMatchInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *KoreaTelcoMatchInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetFullName

`func (o *KoreaTelcoMatchInput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *KoreaTelcoMatchInput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *KoreaTelcoMatchInput) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetDateOfBirth

`func (o *KoreaTelcoMatchInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *KoreaTelcoMatchInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *KoreaTelcoMatchInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetSex

`func (o *KoreaTelcoMatchInput) GetSex() KoreaTelcoMatchSex`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *KoreaTelcoMatchInput) GetSexOk() (*KoreaTelcoMatchSex, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *KoreaTelcoMatchInput) SetSex(v KoreaTelcoMatchSex)`

SetSex sets Sex field to given value.


### GetCarrier

`func (o *KoreaTelcoMatchInput) GetCarrier() MobileCarrier`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *KoreaTelcoMatchInput) GetCarrierOk() (*MobileCarrier, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *KoreaTelcoMatchInput) SetCarrier(v MobileCarrier)`

SetCarrier sets Carrier field to given value.


### GetOperatingSystem

`func (o *KoreaTelcoMatchInput) GetOperatingSystem() MobileOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *KoreaTelcoMatchInput) GetOperatingSystemOk() (*MobileOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *KoreaTelcoMatchInput) SetOperatingSystem(v MobileOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



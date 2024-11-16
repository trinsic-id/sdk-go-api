# KnownPersonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **string** | Given (first) name of the individual | [optional] 
**FamilyName** | Pointer to **string** | Family (last) name of the individual | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the individual | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number (with preceding + character and country code) of the individual being verified | [optional] 
**Address** | Pointer to [**KnownAddress**](KnownAddress.md) | The address of the individual being verified | [optional] 
**DateOfBirth** | Pointer to **string** | Date of birth of the individual, in the format \&quot;YYYY-MM-DD\&quot; | [optional] 

## Methods

### NewKnownPersonData

`func NewKnownPersonData() *KnownPersonData`

NewKnownPersonData instantiates a new KnownPersonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownPersonDataWithDefaults

`func NewKnownPersonDataWithDefaults() *KnownPersonData`

NewKnownPersonDataWithDefaults instantiates a new KnownPersonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *KnownPersonData) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *KnownPersonData) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *KnownPersonData) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *KnownPersonData) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *KnownPersonData) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *KnownPersonData) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *KnownPersonData) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *KnownPersonData) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetMiddleName

`func (o *KnownPersonData) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *KnownPersonData) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *KnownPersonData) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *KnownPersonData) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *KnownPersonData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *KnownPersonData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *KnownPersonData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *KnownPersonData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *KnownPersonData) GetAddress() KnownAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KnownPersonData) GetAddressOk() (*KnownAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KnownPersonData) SetAddress(v KnownAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *KnownPersonData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *KnownPersonData) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *KnownPersonData) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *KnownPersonData) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *KnownPersonData) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



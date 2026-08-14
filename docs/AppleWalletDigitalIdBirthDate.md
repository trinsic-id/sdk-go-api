# AppleWalletDigitalIdBirthDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfBirth** | Pointer to **NullableString** | The encoded date of birth. | [optional] 
**ApproximateMask** | Pointer to **NullableString** | An 8-character mask of &#x60;1&#x60;s and &#x60;0&#x60;s, where a &#x60;1&#x60; indicates that the corresponding digit in the birth date (formatted as YYYYMMDD) is uncertain or unknown.              For example, a mask of \&quot;00000011\&quot; indicates that the day-of-month on which the individual was born is not certain. | [optional] 

## Methods

### NewAppleWalletDigitalIdBirthDate

`func NewAppleWalletDigitalIdBirthDate() *AppleWalletDigitalIdBirthDate`

NewAppleWalletDigitalIdBirthDate instantiates a new AppleWalletDigitalIdBirthDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleWalletDigitalIdBirthDateWithDefaults

`func NewAppleWalletDigitalIdBirthDateWithDefaults() *AppleWalletDigitalIdBirthDate`

NewAppleWalletDigitalIdBirthDateWithDefaults instantiates a new AppleWalletDigitalIdBirthDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfBirth

`func (o *AppleWalletDigitalIdBirthDate) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AppleWalletDigitalIdBirthDate) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AppleWalletDigitalIdBirthDate) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AppleWalletDigitalIdBirthDate) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AppleWalletDigitalIdBirthDate) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AppleWalletDigitalIdBirthDate) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetApproximateMask

`func (o *AppleWalletDigitalIdBirthDate) GetApproximateMask() string`

GetApproximateMask returns the ApproximateMask field if non-nil, zero value otherwise.

### GetApproximateMaskOk

`func (o *AppleWalletDigitalIdBirthDate) GetApproximateMaskOk() (*string, bool)`

GetApproximateMaskOk returns a tuple with the ApproximateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMask

`func (o *AppleWalletDigitalIdBirthDate) SetApproximateMask(v string)`

SetApproximateMask sets ApproximateMask field to given value.

### HasApproximateMask

`func (o *AppleWalletDigitalIdBirthDate) HasApproximateMask() bool`

HasApproximateMask returns a boolean if a field has been set.

### SetApproximateMaskNil

`func (o *AppleWalletDigitalIdBirthDate) SetApproximateMaskNil(b bool)`

 SetApproximateMaskNil sets the value for ApproximateMask to be an explicit nil

### UnsetApproximateMask
`func (o *AppleWalletDigitalIdBirthDate) UnsetApproximateMask()`

UnsetApproximateMask ensures that no value is present for ApproximateMask, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



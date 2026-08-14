# CzechBankIdVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustFramework** | Pointer to **NullableString** | The trust framework used for the verification.              This identifies the anti-money-laundering framework used for identity verification. | [optional] 
**Time** | Pointer to **NullableTime** | The UTC date and time when the verification happened.              This value is normalized to UTC. | [optional] 
**VerificationProcess** | Pointer to **NullableString** | The verification process identifier.              This identifies the bank that completed the physical identity verification process. | [optional] 

## Methods

### NewCzechBankIdVerification

`func NewCzechBankIdVerification() *CzechBankIdVerification`

NewCzechBankIdVerification instantiates a new CzechBankIdVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCzechBankIdVerificationWithDefaults

`func NewCzechBankIdVerificationWithDefaults() *CzechBankIdVerification`

NewCzechBankIdVerificationWithDefaults instantiates a new CzechBankIdVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrustFramework

`func (o *CzechBankIdVerification) GetTrustFramework() string`

GetTrustFramework returns the TrustFramework field if non-nil, zero value otherwise.

### GetTrustFrameworkOk

`func (o *CzechBankIdVerification) GetTrustFrameworkOk() (*string, bool)`

GetTrustFrameworkOk returns a tuple with the TrustFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFramework

`func (o *CzechBankIdVerification) SetTrustFramework(v string)`

SetTrustFramework sets TrustFramework field to given value.

### HasTrustFramework

`func (o *CzechBankIdVerification) HasTrustFramework() bool`

HasTrustFramework returns a boolean if a field has been set.

### SetTrustFrameworkNil

`func (o *CzechBankIdVerification) SetTrustFrameworkNil(b bool)`

 SetTrustFrameworkNil sets the value for TrustFramework to be an explicit nil

### UnsetTrustFramework
`func (o *CzechBankIdVerification) UnsetTrustFramework()`

UnsetTrustFramework ensures that no value is present for TrustFramework, not even an explicit nil
### GetTime

`func (o *CzechBankIdVerification) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CzechBankIdVerification) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CzechBankIdVerification) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *CzechBankIdVerification) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *CzechBankIdVerification) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *CzechBankIdVerification) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetVerificationProcess

`func (o *CzechBankIdVerification) GetVerificationProcess() string`

GetVerificationProcess returns the VerificationProcess field if non-nil, zero value otherwise.

### GetVerificationProcessOk

`func (o *CzechBankIdVerification) GetVerificationProcessOk() (*string, bool)`

GetVerificationProcessOk returns a tuple with the VerificationProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProcess

`func (o *CzechBankIdVerification) SetVerificationProcess(v string)`

SetVerificationProcess sets VerificationProcess field to given value.

### HasVerificationProcess

`func (o *CzechBankIdVerification) HasVerificationProcess() bool`

HasVerificationProcess returns a boolean if a field has been set.

### SetVerificationProcessNil

`func (o *CzechBankIdVerification) SetVerificationProcessNil(b bool)`

 SetVerificationProcessNil sets the value for VerificationProcess to be an explicit nil

### UnsetVerificationProcess
`func (o *CzechBankIdVerification) UnsetVerificationProcess()`

UnsetVerificationProcess ensures that no value is present for VerificationProcess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



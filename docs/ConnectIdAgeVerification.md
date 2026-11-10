# ConnectIdAgeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Over16** | Pointer to **NullableBool** | Whether the individual is at least 16 years old. | [optional] 
**Over18** | Pointer to **NullableBool** | Whether the individual is at least 18 years old. | [optional] 
**Over21** | Pointer to **NullableBool** | Whether the individual is at least 21 years old. | [optional] 
**Over25** | Pointer to **NullableBool** | Whether the individual is at least 25 years old. | [optional] 
**Over65** | Pointer to **NullableBool** | Whether the individual is at least 65 years old. | [optional] 
**TrustFramework** | Pointer to **NullableString** | The authority that verified the claim. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`ConnectIdAgeVerification{FieldName: value}`), or call
`NewConnectIdAgeVerificationWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewConnectIdAgeVerificationWithDefaults

`func NewConnectIdAgeVerificationWithDefaults() *ConnectIdAgeVerification`

NewConnectIdAgeVerificationWithDefaults instantiates a new ConnectIdAgeVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOver16

`func (o *ConnectIdAgeVerification) GetOver16() bool`

GetOver16 returns the Over16 field if non-nil, zero value otherwise.

### GetOver16Ok

`func (o *ConnectIdAgeVerification) GetOver16Ok() (*bool, bool)`

GetOver16Ok returns a tuple with the Over16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver16

`func (o *ConnectIdAgeVerification) SetOver16(v bool)`

SetOver16 sets Over16 field to given value.

### HasOver16

`func (o *ConnectIdAgeVerification) HasOver16() bool`

HasOver16 returns a boolean if a field has been set.

### SetOver16Nil

`func (o *ConnectIdAgeVerification) SetOver16Nil(b bool)`

 SetOver16Nil sets the value for Over16 to be an explicit nil

### UnsetOver16
`func (o *ConnectIdAgeVerification) UnsetOver16()`

UnsetOver16 ensures that no value is present for Over16, not even an explicit nil
### GetOver18

`func (o *ConnectIdAgeVerification) GetOver18() bool`

GetOver18 returns the Over18 field if non-nil, zero value otherwise.

### GetOver18Ok

`func (o *ConnectIdAgeVerification) GetOver18Ok() (*bool, bool)`

GetOver18Ok returns a tuple with the Over18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver18

`func (o *ConnectIdAgeVerification) SetOver18(v bool)`

SetOver18 sets Over18 field to given value.

### HasOver18

`func (o *ConnectIdAgeVerification) HasOver18() bool`

HasOver18 returns a boolean if a field has been set.

### SetOver18Nil

`func (o *ConnectIdAgeVerification) SetOver18Nil(b bool)`

 SetOver18Nil sets the value for Over18 to be an explicit nil

### UnsetOver18
`func (o *ConnectIdAgeVerification) UnsetOver18()`

UnsetOver18 ensures that no value is present for Over18, not even an explicit nil
### GetOver21

`func (o *ConnectIdAgeVerification) GetOver21() bool`

GetOver21 returns the Over21 field if non-nil, zero value otherwise.

### GetOver21Ok

`func (o *ConnectIdAgeVerification) GetOver21Ok() (*bool, bool)`

GetOver21Ok returns a tuple with the Over21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver21

`func (o *ConnectIdAgeVerification) SetOver21(v bool)`

SetOver21 sets Over21 field to given value.

### HasOver21

`func (o *ConnectIdAgeVerification) HasOver21() bool`

HasOver21 returns a boolean if a field has been set.

### SetOver21Nil

`func (o *ConnectIdAgeVerification) SetOver21Nil(b bool)`

 SetOver21Nil sets the value for Over21 to be an explicit nil

### UnsetOver21
`func (o *ConnectIdAgeVerification) UnsetOver21()`

UnsetOver21 ensures that no value is present for Over21, not even an explicit nil
### GetOver25

`func (o *ConnectIdAgeVerification) GetOver25() bool`

GetOver25 returns the Over25 field if non-nil, zero value otherwise.

### GetOver25Ok

`func (o *ConnectIdAgeVerification) GetOver25Ok() (*bool, bool)`

GetOver25Ok returns a tuple with the Over25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver25

`func (o *ConnectIdAgeVerification) SetOver25(v bool)`

SetOver25 sets Over25 field to given value.

### HasOver25

`func (o *ConnectIdAgeVerification) HasOver25() bool`

HasOver25 returns a boolean if a field has been set.

### SetOver25Nil

`func (o *ConnectIdAgeVerification) SetOver25Nil(b bool)`

 SetOver25Nil sets the value for Over25 to be an explicit nil

### UnsetOver25
`func (o *ConnectIdAgeVerification) UnsetOver25()`

UnsetOver25 ensures that no value is present for Over25, not even an explicit nil
### GetOver65

`func (o *ConnectIdAgeVerification) GetOver65() bool`

GetOver65 returns the Over65 field if non-nil, zero value otherwise.

### GetOver65Ok

`func (o *ConnectIdAgeVerification) GetOver65Ok() (*bool, bool)`

GetOver65Ok returns a tuple with the Over65 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver65

`func (o *ConnectIdAgeVerification) SetOver65(v bool)`

SetOver65 sets Over65 field to given value.

### HasOver65

`func (o *ConnectIdAgeVerification) HasOver65() bool`

HasOver65 returns a boolean if a field has been set.

### SetOver65Nil

`func (o *ConnectIdAgeVerification) SetOver65Nil(b bool)`

 SetOver65Nil sets the value for Over65 to be an explicit nil

### UnsetOver65
`func (o *ConnectIdAgeVerification) UnsetOver65()`

UnsetOver65 ensures that no value is present for Over65, not even an explicit nil
### GetTrustFramework

`func (o *ConnectIdAgeVerification) GetTrustFramework() string`

GetTrustFramework returns the TrustFramework field if non-nil, zero value otherwise.

### GetTrustFrameworkOk

`func (o *ConnectIdAgeVerification) GetTrustFrameworkOk() (*string, bool)`

GetTrustFrameworkOk returns a tuple with the TrustFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFramework

`func (o *ConnectIdAgeVerification) SetTrustFramework(v string)`

SetTrustFramework sets TrustFramework field to given value.

### HasTrustFramework

`func (o *ConnectIdAgeVerification) HasTrustFramework() bool`

HasTrustFramework returns a boolean if a field has been set.

### SetTrustFrameworkNil

`func (o *ConnectIdAgeVerification) SetTrustFrameworkNil(b bool)`

 SetTrustFrameworkNil sets the value for TrustFramework to be an explicit nil

### UnsetTrustFramework
`func (o *ConnectIdAgeVerification) UnsetTrustFramework()`

UnsetTrustFramework ensures that no value is present for TrustFramework, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



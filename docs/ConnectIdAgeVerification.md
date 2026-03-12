# ConnectIdAgeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Over18** | Pointer to **NullableBool** | If the individual is over 18 years old. | [optional] 
**TrustFramework** | Pointer to **NullableString** | The authority that verified the claim. | [optional] 

## Methods

### NewConnectIdAgeVerification

`func NewConnectIdAgeVerification() *ConnectIdAgeVerification`

NewConnectIdAgeVerification instantiates a new ConnectIdAgeVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectIdAgeVerificationWithDefaults

`func NewConnectIdAgeVerificationWithDefaults() *ConnectIdAgeVerification`

NewConnectIdAgeVerificationWithDefaults instantiates a new ConnectIdAgeVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



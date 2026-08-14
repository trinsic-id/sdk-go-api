# NetherlandsEudiSandboxProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to [**NullableEudiPidCredential**](EudiPidCredential.md) | An EUDI Person Identification Data (PID) credential, retrieved from the individual&#39;s wallet. | [optional] 
**AgeVerification** | Pointer to [**NullableEudiAgeVerificationCredential**](EudiAgeVerificationCredential.md) | An EUDI Age Verification credential, retrieved from the individual&#39;s wallet. | [optional] 

## Methods

### NewNetherlandsEudiSandboxProviderOutput

`func NewNetherlandsEudiSandboxProviderOutput() *NetherlandsEudiSandboxProviderOutput`

NewNetherlandsEudiSandboxProviderOutput instantiates a new NetherlandsEudiSandboxProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetherlandsEudiSandboxProviderOutputWithDefaults

`func NewNetherlandsEudiSandboxProviderOutputWithDefaults() *NetherlandsEudiSandboxProviderOutput`

NewNetherlandsEudiSandboxProviderOutputWithDefaults instantiates a new NetherlandsEudiSandboxProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *NetherlandsEudiSandboxProviderOutput) GetPid() EudiPidCredential`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *NetherlandsEudiSandboxProviderOutput) GetPidOk() (*EudiPidCredential, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *NetherlandsEudiSandboxProviderOutput) SetPid(v EudiPidCredential)`

SetPid sets Pid field to given value.

### HasPid

`func (o *NetherlandsEudiSandboxProviderOutput) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *NetherlandsEudiSandboxProviderOutput) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *NetherlandsEudiSandboxProviderOutput) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetAgeVerification

`func (o *NetherlandsEudiSandboxProviderOutput) GetAgeVerification() EudiAgeVerificationCredential`

GetAgeVerification returns the AgeVerification field if non-nil, zero value otherwise.

### GetAgeVerificationOk

`func (o *NetherlandsEudiSandboxProviderOutput) GetAgeVerificationOk() (*EudiAgeVerificationCredential, bool)`

GetAgeVerificationOk returns a tuple with the AgeVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerification

`func (o *NetherlandsEudiSandboxProviderOutput) SetAgeVerification(v EudiAgeVerificationCredential)`

SetAgeVerification sets AgeVerification field to given value.

### HasAgeVerification

`func (o *NetherlandsEudiSandboxProviderOutput) HasAgeVerification() bool`

HasAgeVerification returns a boolean if a field has been set.

### SetAgeVerificationNil

`func (o *NetherlandsEudiSandboxProviderOutput) SetAgeVerificationNil(b bool)`

 SetAgeVerificationNil sets the value for AgeVerification to be an explicit nil

### UnsetAgeVerification
`func (o *NetherlandsEudiSandboxProviderOutput) UnsetAgeVerification()`

UnsetAgeVerification ensures that no value is present for AgeVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



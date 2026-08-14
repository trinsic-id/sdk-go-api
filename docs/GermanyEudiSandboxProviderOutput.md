# GermanyEudiSandboxProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to [**NullableEudiPidCredential**](EudiPidCredential.md) | An EUDI Person Identification Data (PID) credential, retrieved from the individual&#39;s wallet. | [optional] 
**Raw18013Output** | Pointer to [**NullableMdlOutput**](MdlOutput.md) | The raw output of the 18013-7 exchange performed through the Germany EUDI Sandbox Wallet. | [optional] 

## Methods

### NewGermanyEudiSandboxProviderOutput

`func NewGermanyEudiSandboxProviderOutput() *GermanyEudiSandboxProviderOutput`

NewGermanyEudiSandboxProviderOutput instantiates a new GermanyEudiSandboxProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGermanyEudiSandboxProviderOutputWithDefaults

`func NewGermanyEudiSandboxProviderOutputWithDefaults() *GermanyEudiSandboxProviderOutput`

NewGermanyEudiSandboxProviderOutputWithDefaults instantiates a new GermanyEudiSandboxProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *GermanyEudiSandboxProviderOutput) GetPid() EudiPidCredential`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *GermanyEudiSandboxProviderOutput) GetPidOk() (*EudiPidCredential, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *GermanyEudiSandboxProviderOutput) SetPid(v EudiPidCredential)`

SetPid sets Pid field to given value.

### HasPid

`func (o *GermanyEudiSandboxProviderOutput) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *GermanyEudiSandboxProviderOutput) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *GermanyEudiSandboxProviderOutput) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetRaw18013Output

`func (o *GermanyEudiSandboxProviderOutput) GetRaw18013Output() MdlOutput`

GetRaw18013Output returns the Raw18013Output field if non-nil, zero value otherwise.

### GetRaw18013OutputOk

`func (o *GermanyEudiSandboxProviderOutput) GetRaw18013OutputOk() (*MdlOutput, bool)`

GetRaw18013OutputOk returns a tuple with the Raw18013Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw18013Output

`func (o *GermanyEudiSandboxProviderOutput) SetRaw18013Output(v MdlOutput)`

SetRaw18013Output sets Raw18013Output field to given value.

### HasRaw18013Output

`func (o *GermanyEudiSandboxProviderOutput) HasRaw18013Output() bool`

HasRaw18013Output returns a boolean if a field has been set.

### SetRaw18013OutputNil

`func (o *GermanyEudiSandboxProviderOutput) SetRaw18013OutputNil(b bool)`

 SetRaw18013OutputNil sets the value for Raw18013Output to be an explicit nil

### UnsetRaw18013Output
`func (o *GermanyEudiSandboxProviderOutput) UnsetRaw18013Output()`

UnsetRaw18013Output ensures that no value is present for Raw18013Output, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



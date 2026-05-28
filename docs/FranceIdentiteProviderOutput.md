# FranceIdentiteProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to [**NullableEudiPidCredential**](EudiPidCredential.md) | An EUDI Person Identification Data (PID) credential, retrieved from the individual&#39;s wallet. | [optional] 
**AgeVerification** | Pointer to [**NullableEudiAgeVerificationCredential**](EudiAgeVerificationCredential.md) | An EUDI Age Verification credential, retrieved from the individual&#39;s wallet. | [optional] 
**Raw18013Output** | [**MdlOutput**](MdlOutput.md) | The raw output of the 18013-7 exchange performed through France Identité. | 

## Methods

### NewFranceIdentiteProviderOutput

`func NewFranceIdentiteProviderOutput(raw18013Output MdlOutput, ) *FranceIdentiteProviderOutput`

NewFranceIdentiteProviderOutput instantiates a new FranceIdentiteProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFranceIdentiteProviderOutputWithDefaults

`func NewFranceIdentiteProviderOutputWithDefaults() *FranceIdentiteProviderOutput`

NewFranceIdentiteProviderOutputWithDefaults instantiates a new FranceIdentiteProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *FranceIdentiteProviderOutput) GetPid() EudiPidCredential`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *FranceIdentiteProviderOutput) GetPidOk() (*EudiPidCredential, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *FranceIdentiteProviderOutput) SetPid(v EudiPidCredential)`

SetPid sets Pid field to given value.

### HasPid

`func (o *FranceIdentiteProviderOutput) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *FranceIdentiteProviderOutput) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *FranceIdentiteProviderOutput) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetAgeVerification

`func (o *FranceIdentiteProviderOutput) GetAgeVerification() EudiAgeVerificationCredential`

GetAgeVerification returns the AgeVerification field if non-nil, zero value otherwise.

### GetAgeVerificationOk

`func (o *FranceIdentiteProviderOutput) GetAgeVerificationOk() (*EudiAgeVerificationCredential, bool)`

GetAgeVerificationOk returns a tuple with the AgeVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerification

`func (o *FranceIdentiteProviderOutput) SetAgeVerification(v EudiAgeVerificationCredential)`

SetAgeVerification sets AgeVerification field to given value.

### HasAgeVerification

`func (o *FranceIdentiteProviderOutput) HasAgeVerification() bool`

HasAgeVerification returns a boolean if a field has been set.

### SetAgeVerificationNil

`func (o *FranceIdentiteProviderOutput) SetAgeVerificationNil(b bool)`

 SetAgeVerificationNil sets the value for AgeVerification to be an explicit nil

### UnsetAgeVerification
`func (o *FranceIdentiteProviderOutput) UnsetAgeVerification()`

UnsetAgeVerification ensures that no value is present for AgeVerification, not even an explicit nil
### GetRaw18013Output

`func (o *FranceIdentiteProviderOutput) GetRaw18013Output() MdlOutput`

GetRaw18013Output returns the Raw18013Output field if non-nil, zero value otherwise.

### GetRaw18013OutputOk

`func (o *FranceIdentiteProviderOutput) GetRaw18013OutputOk() (*MdlOutput, bool)`

GetRaw18013OutputOk returns a tuple with the Raw18013Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw18013Output

`func (o *FranceIdentiteProviderOutput) SetRaw18013Output(v MdlOutput)`

SetRaw18013Output sets Raw18013Output field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



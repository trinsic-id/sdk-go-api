# EudiAgeVerificationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeOver** | Pointer to [**[]AgeOverOutput**](AgeOverOutput.md) | Processed age-over claims returned by the credential. | [optional] 

## Methods

### Named construction

Use a keyed struct literal (`EudiAgeVerificationCredential{FieldName: value}`), or call
`NewEudiAgeVerificationCredentialWithDefaults()` and set fields by name. Positional field
factories and unkeyed struct literals are not supported.

### NewEudiAgeVerificationCredentialWithDefaults

`func NewEudiAgeVerificationCredentialWithDefaults() *EudiAgeVerificationCredential`

NewEudiAgeVerificationCredentialWithDefaults instantiates a new EudiAgeVerificationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeOver

`func (o *EudiAgeVerificationCredential) GetAgeOver() []AgeOverOutput`

GetAgeOver returns the AgeOver field if non-nil, zero value otherwise.

### GetAgeOverOk

`func (o *EudiAgeVerificationCredential) GetAgeOverOk() (*[]AgeOverOutput, bool)`

GetAgeOverOk returns a tuple with the AgeOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOver

`func (o *EudiAgeVerificationCredential) SetAgeOver(v []AgeOverOutput)`

SetAgeOver sets AgeOver field to given value.

### HasAgeOver

`func (o *EudiAgeVerificationCredential) HasAgeOver() bool`

HasAgeOver returns a boolean if a field has been set.

### SetAgeOverNil

`func (o *EudiAgeVerificationCredential) SetAgeOverNil(b bool)`

 SetAgeOverNil sets the value for AgeOver to be an explicit nil

### UnsetAgeOver
`func (o *EudiAgeVerificationCredential) UnsetAgeOver()`

UnsetAgeOver ensures that no value is present for AgeOver, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



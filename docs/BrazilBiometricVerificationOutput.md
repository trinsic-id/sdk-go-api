# BrazilBiometricVerificationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelfieVerificationProbability** | Pointer to **NullableString** | The probability of the selfie verification being successful.              Possible values: - VeryLow - Low - High - VeryHigh | [optional] 
**SelfieLivenessOutcome** | Pointer to **NullableString** | The outcome of the selfie verification.              Possible values: - Real - Fake - BadQuality | [optional] 
**DatabaseSelfieAvailable** | Pointer to **NullableBool** | Whether the government database had biometrics available for the individual.              Also false in some cases biometrics are available but not sufficient for selfie verification. | [optional] 
**SelfieSimilarityPercentage** | Pointer to **NullableFloat64** | Similarity match score of selfie verification.              Ranges from 0.0 to 1.0. | [optional] 

## Methods

### NewBrazilBiometricVerificationOutput

`func NewBrazilBiometricVerificationOutput() *BrazilBiometricVerificationOutput`

NewBrazilBiometricVerificationOutput instantiates a new BrazilBiometricVerificationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrazilBiometricVerificationOutputWithDefaults

`func NewBrazilBiometricVerificationOutputWithDefaults() *BrazilBiometricVerificationOutput`

NewBrazilBiometricVerificationOutputWithDefaults instantiates a new BrazilBiometricVerificationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelfieVerificationProbability

`func (o *BrazilBiometricVerificationOutput) GetSelfieVerificationProbability() string`

GetSelfieVerificationProbability returns the SelfieVerificationProbability field if non-nil, zero value otherwise.

### GetSelfieVerificationProbabilityOk

`func (o *BrazilBiometricVerificationOutput) GetSelfieVerificationProbabilityOk() (*string, bool)`

GetSelfieVerificationProbabilityOk returns a tuple with the SelfieVerificationProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieVerificationProbability

`func (o *BrazilBiometricVerificationOutput) SetSelfieVerificationProbability(v string)`

SetSelfieVerificationProbability sets SelfieVerificationProbability field to given value.

### HasSelfieVerificationProbability

`func (o *BrazilBiometricVerificationOutput) HasSelfieVerificationProbability() bool`

HasSelfieVerificationProbability returns a boolean if a field has been set.

### SetSelfieVerificationProbabilityNil

`func (o *BrazilBiometricVerificationOutput) SetSelfieVerificationProbabilityNil(b bool)`

 SetSelfieVerificationProbabilityNil sets the value for SelfieVerificationProbability to be an explicit nil

### UnsetSelfieVerificationProbability
`func (o *BrazilBiometricVerificationOutput) UnsetSelfieVerificationProbability()`

UnsetSelfieVerificationProbability ensures that no value is present for SelfieVerificationProbability, not even an explicit nil
### GetSelfieLivenessOutcome

`func (o *BrazilBiometricVerificationOutput) GetSelfieLivenessOutcome() string`

GetSelfieLivenessOutcome returns the SelfieLivenessOutcome field if non-nil, zero value otherwise.

### GetSelfieLivenessOutcomeOk

`func (o *BrazilBiometricVerificationOutput) GetSelfieLivenessOutcomeOk() (*string, bool)`

GetSelfieLivenessOutcomeOk returns a tuple with the SelfieLivenessOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieLivenessOutcome

`func (o *BrazilBiometricVerificationOutput) SetSelfieLivenessOutcome(v string)`

SetSelfieLivenessOutcome sets SelfieLivenessOutcome field to given value.

### HasSelfieLivenessOutcome

`func (o *BrazilBiometricVerificationOutput) HasSelfieLivenessOutcome() bool`

HasSelfieLivenessOutcome returns a boolean if a field has been set.

### SetSelfieLivenessOutcomeNil

`func (o *BrazilBiometricVerificationOutput) SetSelfieLivenessOutcomeNil(b bool)`

 SetSelfieLivenessOutcomeNil sets the value for SelfieLivenessOutcome to be an explicit nil

### UnsetSelfieLivenessOutcome
`func (o *BrazilBiometricVerificationOutput) UnsetSelfieLivenessOutcome()`

UnsetSelfieLivenessOutcome ensures that no value is present for SelfieLivenessOutcome, not even an explicit nil
### GetDatabaseSelfieAvailable

`func (o *BrazilBiometricVerificationOutput) GetDatabaseSelfieAvailable() bool`

GetDatabaseSelfieAvailable returns the DatabaseSelfieAvailable field if non-nil, zero value otherwise.

### GetDatabaseSelfieAvailableOk

`func (o *BrazilBiometricVerificationOutput) GetDatabaseSelfieAvailableOk() (*bool, bool)`

GetDatabaseSelfieAvailableOk returns a tuple with the DatabaseSelfieAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSelfieAvailable

`func (o *BrazilBiometricVerificationOutput) SetDatabaseSelfieAvailable(v bool)`

SetDatabaseSelfieAvailable sets DatabaseSelfieAvailable field to given value.

### HasDatabaseSelfieAvailable

`func (o *BrazilBiometricVerificationOutput) HasDatabaseSelfieAvailable() bool`

HasDatabaseSelfieAvailable returns a boolean if a field has been set.

### SetDatabaseSelfieAvailableNil

`func (o *BrazilBiometricVerificationOutput) SetDatabaseSelfieAvailableNil(b bool)`

 SetDatabaseSelfieAvailableNil sets the value for DatabaseSelfieAvailable to be an explicit nil

### UnsetDatabaseSelfieAvailable
`func (o *BrazilBiometricVerificationOutput) UnsetDatabaseSelfieAvailable()`

UnsetDatabaseSelfieAvailable ensures that no value is present for DatabaseSelfieAvailable, not even an explicit nil
### GetSelfieSimilarityPercentage

`func (o *BrazilBiometricVerificationOutput) GetSelfieSimilarityPercentage() float64`

GetSelfieSimilarityPercentage returns the SelfieSimilarityPercentage field if non-nil, zero value otherwise.

### GetSelfieSimilarityPercentageOk

`func (o *BrazilBiometricVerificationOutput) GetSelfieSimilarityPercentageOk() (*float64, bool)`

GetSelfieSimilarityPercentageOk returns a tuple with the SelfieSimilarityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieSimilarityPercentage

`func (o *BrazilBiometricVerificationOutput) SetSelfieSimilarityPercentage(v float64)`

SetSelfieSimilarityPercentage sets SelfieSimilarityPercentage field to given value.

### HasSelfieSimilarityPercentage

`func (o *BrazilBiometricVerificationOutput) HasSelfieSimilarityPercentage() bool`

HasSelfieSimilarityPercentage returns a boolean if a field has been set.

### SetSelfieSimilarityPercentageNil

`func (o *BrazilBiometricVerificationOutput) SetSelfieSimilarityPercentageNil(b bool)`

 SetSelfieSimilarityPercentageNil sets the value for SelfieSimilarityPercentage to be an explicit nil

### UnsetSelfieSimilarityPercentage
`func (o *BrazilBiometricVerificationOutput) UnsetSelfieSimilarityPercentage()`

UnsetSelfieSimilarityPercentage ensures that no value is present for SelfieSimilarityPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



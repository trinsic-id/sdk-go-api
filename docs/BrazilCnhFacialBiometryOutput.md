# BrazilCnhFacialBiometryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseFaceAvailable** | **bool** | Whether the government database had facial biometrics available for comparison. | 
**Probability** | Pointer to **NullableString** | Probability bucket returned by Serpro for the facial biometric comparison.              Known values: - VeryLow - Low - High - VeryHigh | [optional] 
**SimilarityScore** | Pointer to **NullableFloat64** | Similarity score returned by Serpro for the facial biometric comparison.              Ranges from 0.0 to 1.0, where 1.0 is a perfect match. | [optional] 
**Liveness** | Pointer to **NullableString** | Liveness result returned by Serpro when liveness is part of the validation response.              Known values: - BAD_QUALITY: The image quality was too low for liveness validation. - FAKE: Liveness validation detected a presentation attack or non-live face. - REAL: Liveness validation detected a live face. | [optional] 

## Methods

### NewBrazilCnhFacialBiometryOutput

`func NewBrazilCnhFacialBiometryOutput(databaseFaceAvailable bool, ) *BrazilCnhFacialBiometryOutput`

NewBrazilCnhFacialBiometryOutput instantiates a new BrazilCnhFacialBiometryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrazilCnhFacialBiometryOutputWithDefaults

`func NewBrazilCnhFacialBiometryOutputWithDefaults() *BrazilCnhFacialBiometryOutput`

NewBrazilCnhFacialBiometryOutputWithDefaults instantiates a new BrazilCnhFacialBiometryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseFaceAvailable

`func (o *BrazilCnhFacialBiometryOutput) GetDatabaseFaceAvailable() bool`

GetDatabaseFaceAvailable returns the DatabaseFaceAvailable field if non-nil, zero value otherwise.

### GetDatabaseFaceAvailableOk

`func (o *BrazilCnhFacialBiometryOutput) GetDatabaseFaceAvailableOk() (*bool, bool)`

GetDatabaseFaceAvailableOk returns a tuple with the DatabaseFaceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseFaceAvailable

`func (o *BrazilCnhFacialBiometryOutput) SetDatabaseFaceAvailable(v bool)`

SetDatabaseFaceAvailable sets DatabaseFaceAvailable field to given value.


### GetProbability

`func (o *BrazilCnhFacialBiometryOutput) GetProbability() string`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *BrazilCnhFacialBiometryOutput) GetProbabilityOk() (*string, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *BrazilCnhFacialBiometryOutput) SetProbability(v string)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *BrazilCnhFacialBiometryOutput) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### SetProbabilityNil

`func (o *BrazilCnhFacialBiometryOutput) SetProbabilityNil(b bool)`

 SetProbabilityNil sets the value for Probability to be an explicit nil

### UnsetProbability
`func (o *BrazilCnhFacialBiometryOutput) UnsetProbability()`

UnsetProbability ensures that no value is present for Probability, not even an explicit nil
### GetSimilarityScore

`func (o *BrazilCnhFacialBiometryOutput) GetSimilarityScore() float64`

GetSimilarityScore returns the SimilarityScore field if non-nil, zero value otherwise.

### GetSimilarityScoreOk

`func (o *BrazilCnhFacialBiometryOutput) GetSimilarityScoreOk() (*float64, bool)`

GetSimilarityScoreOk returns a tuple with the SimilarityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarityScore

`func (o *BrazilCnhFacialBiometryOutput) SetSimilarityScore(v float64)`

SetSimilarityScore sets SimilarityScore field to given value.

### HasSimilarityScore

`func (o *BrazilCnhFacialBiometryOutput) HasSimilarityScore() bool`

HasSimilarityScore returns a boolean if a field has been set.

### SetSimilarityScoreNil

`func (o *BrazilCnhFacialBiometryOutput) SetSimilarityScoreNil(b bool)`

 SetSimilarityScoreNil sets the value for SimilarityScore to be an explicit nil

### UnsetSimilarityScore
`func (o *BrazilCnhFacialBiometryOutput) UnsetSimilarityScore()`

UnsetSimilarityScore ensures that no value is present for SimilarityScore, not even an explicit nil
### GetLiveness

`func (o *BrazilCnhFacialBiometryOutput) GetLiveness() string`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *BrazilCnhFacialBiometryOutput) GetLivenessOk() (*string, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *BrazilCnhFacialBiometryOutput) SetLiveness(v string)`

SetLiveness sets Liveness field to given value.

### HasLiveness

`func (o *BrazilCnhFacialBiometryOutput) HasLiveness() bool`

HasLiveness returns a boolean if a field has been set.

### SetLivenessNil

`func (o *BrazilCnhFacialBiometryOutput) SetLivenessNil(b bool)`

 SetLivenessNil sets the value for Liveness to be an explicit nil

### UnsetLiveness
`func (o *BrazilCnhFacialBiometryOutput) UnsetLiveness()`

UnsetLiveness ensures that no value is present for Liveness, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



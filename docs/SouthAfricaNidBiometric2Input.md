# SouthAfricaNidBiometric2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | Pointer to **NullableString** | The user&#39;s South African National ID number | [optional] 
**LivenessImages** | Pointer to **[]string** | An array of exactly 8 images required for biometric liveness verification. The first 7 images should be liveness frames captured during the liveness detection process, and the last image (8th) should be a selfie of the user. All images must be in JPEG format and each image must be less than 15MB in size. | [optional] 
**TestSelfie** | Pointer to **NullableString** | Test selfie for test environment (optional, maximum 15MB). Must be JPEG format. | [optional] 

## Methods

### NewSouthAfricaNidBiometric2Input

`func NewSouthAfricaNidBiometric2Input() *SouthAfricaNidBiometric2Input`

NewSouthAfricaNidBiometric2Input instantiates a new SouthAfricaNidBiometric2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidBiometric2InputWithDefaults

`func NewSouthAfricaNidBiometric2InputWithDefaults() *SouthAfricaNidBiometric2Input`

NewSouthAfricaNidBiometric2InputWithDefaults instantiates a new SouthAfricaNidBiometric2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *SouthAfricaNidBiometric2Input) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *SouthAfricaNidBiometric2Input) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *SouthAfricaNidBiometric2Input) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *SouthAfricaNidBiometric2Input) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *SouthAfricaNidBiometric2Input) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *SouthAfricaNidBiometric2Input) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetLivenessImages

`func (o *SouthAfricaNidBiometric2Input) GetLivenessImages() []string`

GetLivenessImages returns the LivenessImages field if non-nil, zero value otherwise.

### GetLivenessImagesOk

`func (o *SouthAfricaNidBiometric2Input) GetLivenessImagesOk() (*[]string, bool)`

GetLivenessImagesOk returns a tuple with the LivenessImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessImages

`func (o *SouthAfricaNidBiometric2Input) SetLivenessImages(v []string)`

SetLivenessImages sets LivenessImages field to given value.

### HasLivenessImages

`func (o *SouthAfricaNidBiometric2Input) HasLivenessImages() bool`

HasLivenessImages returns a boolean if a field has been set.

### SetLivenessImagesNil

`func (o *SouthAfricaNidBiometric2Input) SetLivenessImagesNil(b bool)`

 SetLivenessImagesNil sets the value for LivenessImages to be an explicit nil

### UnsetLivenessImages
`func (o *SouthAfricaNidBiometric2Input) UnsetLivenessImages()`

UnsetLivenessImages ensures that no value is present for LivenessImages, not even an explicit nil
### GetTestSelfie

`func (o *SouthAfricaNidBiometric2Input) GetTestSelfie() string`

GetTestSelfie returns the TestSelfie field if non-nil, zero value otherwise.

### GetTestSelfieOk

`func (o *SouthAfricaNidBiometric2Input) GetTestSelfieOk() (*string, bool)`

GetTestSelfieOk returns a tuple with the TestSelfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSelfie

`func (o *SouthAfricaNidBiometric2Input) SetTestSelfie(v string)`

SetTestSelfie sets TestSelfie field to given value.

### HasTestSelfie

`func (o *SouthAfricaNidBiometric2Input) HasTestSelfie() bool`

HasTestSelfie returns a boolean if a field has been set.

### SetTestSelfieNil

`func (o *SouthAfricaNidBiometric2Input) SetTestSelfieNil(b bool)`

 SetTestSelfieNil sets the value for TestSelfie to be an explicit nil

### UnsetTestSelfie
`func (o *SouthAfricaNidBiometric2Input) UnsetTestSelfie()`

UnsetTestSelfie ensures that no value is present for TestSelfie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



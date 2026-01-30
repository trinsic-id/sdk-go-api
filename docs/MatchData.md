# MatchData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | Pointer to [**NullableMatch**](Match.md) | Whether the provided National ID Number matched the information on file for the individual | [optional] 
**FullName** | Pointer to [**NullableMatch**](Match.md) | The match score for the full name of the individual.              Higher values indicate a closer match. | [optional] 
**GivenName** | Pointer to [**NullableMatch**](Match.md) | The match score for the given (first) name of the individual.              Higher values indicate a closer match. | [optional] 
**MiddleName** | Pointer to [**NullableMatch**](Match.md) | The match score for the middle name(s) of the individual.              Higher values indicate a closer match. | [optional] 
**FamilyName** | Pointer to [**NullableMatch**](Match.md) | The match score for the family (last) name of the individual.              Higher values indicate a closer match. | [optional] 
**Sex** | Pointer to [**NullableMatch**](Match.md) | Whether the provided sex of the individual matched the information on file for the individual | [optional] 
**DateOfBirth** | Pointer to [**NullableMatch**](Match.md) | Whether the provided date of birth matched the information on file for the individual | [optional] 
**PhoneNumber** | Pointer to [**NullableMatch**](Match.md) | Whether the provided phone number matched the information on file for the individual | [optional] 
**FaceMatch** | Pointer to [**NullableMatch**](Match.md) | The match score for the face match between the provided selfie image and the biometrics on file for the individual.              Higher values indicate greater match confidence. | [optional] 
**Liveness** | Pointer to [**NullableMatch**](Match.md) | The confidence score for the liveness check performed against the selfie image of the individual.              Higher values indicate lower suspicion. | [optional] 
**ImageAuthenticity** | Pointer to [**NullableMatch**](Match.md) | The confidence score for the image manipulation check performed against the selfie image of the individual.              Higher values indicate lower suspicion of image manipulation. | [optional] 

## Methods

### NewMatchData

`func NewMatchData() *MatchData`

NewMatchData instantiates a new MatchData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchDataWithDefaults

`func NewMatchDataWithDefaults() *MatchData`

NewMatchDataWithDefaults instantiates a new MatchData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *MatchData) GetNationalIdNumber() Match`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *MatchData) GetNationalIdNumberOk() (*Match, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *MatchData) SetNationalIdNumber(v Match)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *MatchData) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *MatchData) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *MatchData) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetFullName

`func (o *MatchData) GetFullName() Match`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MatchData) GetFullNameOk() (*Match, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MatchData) SetFullName(v Match)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MatchData) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *MatchData) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *MatchData) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGivenName

`func (o *MatchData) GetGivenName() Match`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MatchData) GetGivenNameOk() (*Match, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MatchData) SetGivenName(v Match)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MatchData) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MatchData) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MatchData) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *MatchData) GetMiddleName() Match`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MatchData) GetMiddleNameOk() (*Match, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *MatchData) SetMiddleName(v Match)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *MatchData) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *MatchData) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *MatchData) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *MatchData) GetFamilyName() Match`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MatchData) GetFamilyNameOk() (*Match, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MatchData) SetFamilyName(v Match)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MatchData) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *MatchData) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *MatchData) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetSex

`func (o *MatchData) GetSex() Match`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *MatchData) GetSexOk() (*Match, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *MatchData) SetSex(v Match)`

SetSex sets Sex field to given value.

### HasSex

`func (o *MatchData) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *MatchData) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *MatchData) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *MatchData) GetDateOfBirth() Match`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MatchData) GetDateOfBirthOk() (*Match, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MatchData) SetDateOfBirth(v Match)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *MatchData) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *MatchData) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *MatchData) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPhoneNumber

`func (o *MatchData) GetPhoneNumber() Match`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MatchData) GetPhoneNumberOk() (*Match, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MatchData) SetPhoneNumber(v Match)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MatchData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *MatchData) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *MatchData) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetFaceMatch

`func (o *MatchData) GetFaceMatch() Match`

GetFaceMatch returns the FaceMatch field if non-nil, zero value otherwise.

### GetFaceMatchOk

`func (o *MatchData) GetFaceMatchOk() (*Match, bool)`

GetFaceMatchOk returns a tuple with the FaceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceMatch

`func (o *MatchData) SetFaceMatch(v Match)`

SetFaceMatch sets FaceMatch field to given value.

### HasFaceMatch

`func (o *MatchData) HasFaceMatch() bool`

HasFaceMatch returns a boolean if a field has been set.

### SetFaceMatchNil

`func (o *MatchData) SetFaceMatchNil(b bool)`

 SetFaceMatchNil sets the value for FaceMatch to be an explicit nil

### UnsetFaceMatch
`func (o *MatchData) UnsetFaceMatch()`

UnsetFaceMatch ensures that no value is present for FaceMatch, not even an explicit nil
### GetLiveness

`func (o *MatchData) GetLiveness() Match`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *MatchData) GetLivenessOk() (*Match, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *MatchData) SetLiveness(v Match)`

SetLiveness sets Liveness field to given value.

### HasLiveness

`func (o *MatchData) HasLiveness() bool`

HasLiveness returns a boolean if a field has been set.

### SetLivenessNil

`func (o *MatchData) SetLivenessNil(b bool)`

 SetLivenessNil sets the value for Liveness to be an explicit nil

### UnsetLiveness
`func (o *MatchData) UnsetLiveness()`

UnsetLiveness ensures that no value is present for Liveness, not even an explicit nil
### GetImageAuthenticity

`func (o *MatchData) GetImageAuthenticity() Match`

GetImageAuthenticity returns the ImageAuthenticity field if non-nil, zero value otherwise.

### GetImageAuthenticityOk

`func (o *MatchData) GetImageAuthenticityOk() (*Match, bool)`

GetImageAuthenticityOk returns a tuple with the ImageAuthenticity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAuthenticity

`func (o *MatchData) SetImageAuthenticity(v Match)`

SetImageAuthenticity sets ImageAuthenticity field to given value.

### HasImageAuthenticity

`func (o *MatchData) HasImageAuthenticity() bool`

HasImageAuthenticity returns a boolean if a field has been set.

### SetImageAuthenticityNil

`func (o *MatchData) SetImageAuthenticityNil(b bool)`

 SetImageAuthenticityNil sets the value for ImageAuthenticity to be an explicit nil

### UnsetImageAuthenticity
`func (o *MatchData) UnsetImageAuthenticity()`

UnsetImageAuthenticity ensures that no value is present for ImageAuthenticity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



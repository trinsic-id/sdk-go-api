# PolandMobywatelMatchProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | Given name as provided by the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | Current legal family name (nazwisko) as provided by the individual.              In Poland the current legal family name (nazwisko) is a separate concept from your birth family name (nazwisko rodowe). They often match, but they can differ after marriage, adoption, or a court-ordered change. | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality as provided by the individual. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth as provided by the individual. | [optional] 
**PersonalNumber** | Pointer to **NullableString** | Personal number (PESEL) as provided by the individual. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | Expiration date as provided by the individual. | [optional] 
**IsMatchForGivenNames** | Pointer to **NullableBool** | True when the provided given name matches mObywatel wallet data. | [optional] 
**IsMatchForFamilyName** | Pointer to **NullableBool** | True when the provided family name matches mObywatel wallet data. | [optional] 
**IsMatchForExpirationDate** | Pointer to **NullableBool** | True when the provided document expiration date matches mObywatel wallet data. | [optional] 
**IsMatchForPersonalNumber** | Pointer to **NullableBool** | True when the provided personal number (PESEL) matches mObywatel wallet data. | [optional] 
**IsMatchForNationality** | Pointer to **NullableBool** | True when the provided nationality matches mObywatel wallet data. | [optional] 
**IsMatchForBirthDate** | Pointer to **NullableBool** | True when the provided birth date matches mObywatel wallet data. | [optional] 
**IsMatchForSelfie** | Pointer to **NullableBool** | True when the face in the provided selfie matches the face in the document photo in mObywatel wallet. | [optional] 
**IsNotTooSimilarToDocumentPortrait** | Pointer to **NullableBool** | True when the uploaded face and wallet portrait are not suspiciously identical. Often fails when using a document photo as a selfie. | [optional] 
**IsNotDifferentFace** | Pointer to **NullableBool** | True when no other person&#39;s face has been used to verify this document (good). False when another person&#39;s face has been used to verify this document (possible fraud). | [optional] 
**IsSingleFace** | Pointer to **NullableBool** | True when exactly one face was detected on the selfie image (good). False when multiple faces were detected on the selfie image (ambiguous result). | [optional] 
**IsOverAge** | Pointer to **NullableBool** | True when the individual&#39;s age in digital wallet is over 18. | [optional] 
**IsAgeEstimationMatchForSelfie** | Pointer to **NullableBool** | True when the individual&#39;s age estimated from the selfie matches the individual&#39;s age in digital wallet. | [optional] 
**AgeEstimationThreshold** | Pointer to **NullableInt32** | Number of years allowed between the individual&#39;s age in digital wallet and the individual&#39;s age estimated from the selfie. | [optional] 
**EstimatedAgeFromSelfie** | Pointer to **NullableInt32** | Estimated age from the selfie, if provided. | [optional] 

## Methods

### NewPolandMobywatelMatchProviderOutput

`func NewPolandMobywatelMatchProviderOutput() *PolandMobywatelMatchProviderOutput`

NewPolandMobywatelMatchProviderOutput instantiates a new PolandMobywatelMatchProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolandMobywatelMatchProviderOutputWithDefaults

`func NewPolandMobywatelMatchProviderOutputWithDefaults() *PolandMobywatelMatchProviderOutput`

NewPolandMobywatelMatchProviderOutputWithDefaults instantiates a new PolandMobywatelMatchProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *PolandMobywatelMatchProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PolandMobywatelMatchProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PolandMobywatelMatchProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PolandMobywatelMatchProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *PolandMobywatelMatchProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *PolandMobywatelMatchProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *PolandMobywatelMatchProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PolandMobywatelMatchProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PolandMobywatelMatchProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PolandMobywatelMatchProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *PolandMobywatelMatchProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *PolandMobywatelMatchProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetNationality

`func (o *PolandMobywatelMatchProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PolandMobywatelMatchProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PolandMobywatelMatchProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PolandMobywatelMatchProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *PolandMobywatelMatchProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *PolandMobywatelMatchProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDateOfBirth

`func (o *PolandMobywatelMatchProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PolandMobywatelMatchProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PolandMobywatelMatchProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PolandMobywatelMatchProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PolandMobywatelMatchProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PolandMobywatelMatchProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *PolandMobywatelMatchProviderOutput) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.

### SetPersonalNumberNil

`func (o *PolandMobywatelMatchProviderOutput) SetPersonalNumberNil(b bool)`

 SetPersonalNumberNil sets the value for PersonalNumber to be an explicit nil

### UnsetPersonalNumber
`func (o *PolandMobywatelMatchProviderOutput) UnsetPersonalNumber()`

UnsetPersonalNumber ensures that no value is present for PersonalNumber, not even an explicit nil
### GetExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PolandMobywatelMatchProviderOutput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PolandMobywatelMatchProviderOutput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PolandMobywatelMatchProviderOutput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetIsMatchForGivenNames

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForGivenNames() bool`

GetIsMatchForGivenNames returns the IsMatchForGivenNames field if non-nil, zero value otherwise.

### GetIsMatchForGivenNamesOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForGivenNamesOk() (*bool, bool)`

GetIsMatchForGivenNamesOk returns a tuple with the IsMatchForGivenNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForGivenNames

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForGivenNames(v bool)`

SetIsMatchForGivenNames sets IsMatchForGivenNames field to given value.

### HasIsMatchForGivenNames

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForGivenNames() bool`

HasIsMatchForGivenNames returns a boolean if a field has been set.

### SetIsMatchForGivenNamesNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForGivenNamesNil(b bool)`

 SetIsMatchForGivenNamesNil sets the value for IsMatchForGivenNames to be an explicit nil

### UnsetIsMatchForGivenNames
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForGivenNames()`

UnsetIsMatchForGivenNames ensures that no value is present for IsMatchForGivenNames, not even an explicit nil
### GetIsMatchForFamilyName

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForFamilyName() bool`

GetIsMatchForFamilyName returns the IsMatchForFamilyName field if non-nil, zero value otherwise.

### GetIsMatchForFamilyNameOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForFamilyNameOk() (*bool, bool)`

GetIsMatchForFamilyNameOk returns a tuple with the IsMatchForFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForFamilyName

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForFamilyName(v bool)`

SetIsMatchForFamilyName sets IsMatchForFamilyName field to given value.

### HasIsMatchForFamilyName

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForFamilyName() bool`

HasIsMatchForFamilyName returns a boolean if a field has been set.

### SetIsMatchForFamilyNameNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForFamilyNameNil(b bool)`

 SetIsMatchForFamilyNameNil sets the value for IsMatchForFamilyName to be an explicit nil

### UnsetIsMatchForFamilyName
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForFamilyName()`

UnsetIsMatchForFamilyName ensures that no value is present for IsMatchForFamilyName, not even an explicit nil
### GetIsMatchForExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForExpirationDate() bool`

GetIsMatchForExpirationDate returns the IsMatchForExpirationDate field if non-nil, zero value otherwise.

### GetIsMatchForExpirationDateOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForExpirationDateOk() (*bool, bool)`

GetIsMatchForExpirationDateOk returns a tuple with the IsMatchForExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForExpirationDate(v bool)`

SetIsMatchForExpirationDate sets IsMatchForExpirationDate field to given value.

### HasIsMatchForExpirationDate

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForExpirationDate() bool`

HasIsMatchForExpirationDate returns a boolean if a field has been set.

### SetIsMatchForExpirationDateNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForExpirationDateNil(b bool)`

 SetIsMatchForExpirationDateNil sets the value for IsMatchForExpirationDate to be an explicit nil

### UnsetIsMatchForExpirationDate
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForExpirationDate()`

UnsetIsMatchForExpirationDate ensures that no value is present for IsMatchForExpirationDate, not even an explicit nil
### GetIsMatchForPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForPersonalNumber() bool`

GetIsMatchForPersonalNumber returns the IsMatchForPersonalNumber field if non-nil, zero value otherwise.

### GetIsMatchForPersonalNumberOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForPersonalNumberOk() (*bool, bool)`

GetIsMatchForPersonalNumberOk returns a tuple with the IsMatchForPersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForPersonalNumber(v bool)`

SetIsMatchForPersonalNumber sets IsMatchForPersonalNumber field to given value.

### HasIsMatchForPersonalNumber

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForPersonalNumber() bool`

HasIsMatchForPersonalNumber returns a boolean if a field has been set.

### SetIsMatchForPersonalNumberNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForPersonalNumberNil(b bool)`

 SetIsMatchForPersonalNumberNil sets the value for IsMatchForPersonalNumber to be an explicit nil

### UnsetIsMatchForPersonalNumber
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForPersonalNumber()`

UnsetIsMatchForPersonalNumber ensures that no value is present for IsMatchForPersonalNumber, not even an explicit nil
### GetIsMatchForNationality

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForNationality() bool`

GetIsMatchForNationality returns the IsMatchForNationality field if non-nil, zero value otherwise.

### GetIsMatchForNationalityOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForNationalityOk() (*bool, bool)`

GetIsMatchForNationalityOk returns a tuple with the IsMatchForNationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForNationality

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForNationality(v bool)`

SetIsMatchForNationality sets IsMatchForNationality field to given value.

### HasIsMatchForNationality

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForNationality() bool`

HasIsMatchForNationality returns a boolean if a field has been set.

### SetIsMatchForNationalityNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForNationalityNil(b bool)`

 SetIsMatchForNationalityNil sets the value for IsMatchForNationality to be an explicit nil

### UnsetIsMatchForNationality
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForNationality()`

UnsetIsMatchForNationality ensures that no value is present for IsMatchForNationality, not even an explicit nil
### GetIsMatchForBirthDate

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForBirthDate() bool`

GetIsMatchForBirthDate returns the IsMatchForBirthDate field if non-nil, zero value otherwise.

### GetIsMatchForBirthDateOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForBirthDateOk() (*bool, bool)`

GetIsMatchForBirthDateOk returns a tuple with the IsMatchForBirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForBirthDate

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForBirthDate(v bool)`

SetIsMatchForBirthDate sets IsMatchForBirthDate field to given value.

### HasIsMatchForBirthDate

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForBirthDate() bool`

HasIsMatchForBirthDate returns a boolean if a field has been set.

### SetIsMatchForBirthDateNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForBirthDateNil(b bool)`

 SetIsMatchForBirthDateNil sets the value for IsMatchForBirthDate to be an explicit nil

### UnsetIsMatchForBirthDate
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForBirthDate()`

UnsetIsMatchForBirthDate ensures that no value is present for IsMatchForBirthDate, not even an explicit nil
### GetIsMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForSelfie() bool`

GetIsMatchForSelfie returns the IsMatchForSelfie field if non-nil, zero value otherwise.

### GetIsMatchForSelfieOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsMatchForSelfieOk() (*bool, bool)`

GetIsMatchForSelfieOk returns a tuple with the IsMatchForSelfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForSelfie(v bool)`

SetIsMatchForSelfie sets IsMatchForSelfie field to given value.

### HasIsMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) HasIsMatchForSelfie() bool`

HasIsMatchForSelfie returns a boolean if a field has been set.

### SetIsMatchForSelfieNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsMatchForSelfieNil(b bool)`

 SetIsMatchForSelfieNil sets the value for IsMatchForSelfie to be an explicit nil

### UnsetIsMatchForSelfie
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsMatchForSelfie()`

UnsetIsMatchForSelfie ensures that no value is present for IsMatchForSelfie, not even an explicit nil
### GetIsNotTooSimilarToDocumentPortrait

`func (o *PolandMobywatelMatchProviderOutput) GetIsNotTooSimilarToDocumentPortrait() bool`

GetIsNotTooSimilarToDocumentPortrait returns the IsNotTooSimilarToDocumentPortrait field if non-nil, zero value otherwise.

### GetIsNotTooSimilarToDocumentPortraitOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsNotTooSimilarToDocumentPortraitOk() (*bool, bool)`

GetIsNotTooSimilarToDocumentPortraitOk returns a tuple with the IsNotTooSimilarToDocumentPortrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotTooSimilarToDocumentPortrait

`func (o *PolandMobywatelMatchProviderOutput) SetIsNotTooSimilarToDocumentPortrait(v bool)`

SetIsNotTooSimilarToDocumentPortrait sets IsNotTooSimilarToDocumentPortrait field to given value.

### HasIsNotTooSimilarToDocumentPortrait

`func (o *PolandMobywatelMatchProviderOutput) HasIsNotTooSimilarToDocumentPortrait() bool`

HasIsNotTooSimilarToDocumentPortrait returns a boolean if a field has been set.

### SetIsNotTooSimilarToDocumentPortraitNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsNotTooSimilarToDocumentPortraitNil(b bool)`

 SetIsNotTooSimilarToDocumentPortraitNil sets the value for IsNotTooSimilarToDocumentPortrait to be an explicit nil

### UnsetIsNotTooSimilarToDocumentPortrait
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsNotTooSimilarToDocumentPortrait()`

UnsetIsNotTooSimilarToDocumentPortrait ensures that no value is present for IsNotTooSimilarToDocumentPortrait, not even an explicit nil
### GetIsNotDifferentFace

`func (o *PolandMobywatelMatchProviderOutput) GetIsNotDifferentFace() bool`

GetIsNotDifferentFace returns the IsNotDifferentFace field if non-nil, zero value otherwise.

### GetIsNotDifferentFaceOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsNotDifferentFaceOk() (*bool, bool)`

GetIsNotDifferentFaceOk returns a tuple with the IsNotDifferentFace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotDifferentFace

`func (o *PolandMobywatelMatchProviderOutput) SetIsNotDifferentFace(v bool)`

SetIsNotDifferentFace sets IsNotDifferentFace field to given value.

### HasIsNotDifferentFace

`func (o *PolandMobywatelMatchProviderOutput) HasIsNotDifferentFace() bool`

HasIsNotDifferentFace returns a boolean if a field has been set.

### SetIsNotDifferentFaceNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsNotDifferentFaceNil(b bool)`

 SetIsNotDifferentFaceNil sets the value for IsNotDifferentFace to be an explicit nil

### UnsetIsNotDifferentFace
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsNotDifferentFace()`

UnsetIsNotDifferentFace ensures that no value is present for IsNotDifferentFace, not even an explicit nil
### GetIsSingleFace

`func (o *PolandMobywatelMatchProviderOutput) GetIsSingleFace() bool`

GetIsSingleFace returns the IsSingleFace field if non-nil, zero value otherwise.

### GetIsSingleFaceOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsSingleFaceOk() (*bool, bool)`

GetIsSingleFaceOk returns a tuple with the IsSingleFace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleFace

`func (o *PolandMobywatelMatchProviderOutput) SetIsSingleFace(v bool)`

SetIsSingleFace sets IsSingleFace field to given value.

### HasIsSingleFace

`func (o *PolandMobywatelMatchProviderOutput) HasIsSingleFace() bool`

HasIsSingleFace returns a boolean if a field has been set.

### SetIsSingleFaceNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsSingleFaceNil(b bool)`

 SetIsSingleFaceNil sets the value for IsSingleFace to be an explicit nil

### UnsetIsSingleFace
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsSingleFace()`

UnsetIsSingleFace ensures that no value is present for IsSingleFace, not even an explicit nil
### GetIsOverAge

`func (o *PolandMobywatelMatchProviderOutput) GetIsOverAge() bool`

GetIsOverAge returns the IsOverAge field if non-nil, zero value otherwise.

### GetIsOverAgeOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsOverAgeOk() (*bool, bool)`

GetIsOverAgeOk returns a tuple with the IsOverAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverAge

`func (o *PolandMobywatelMatchProviderOutput) SetIsOverAge(v bool)`

SetIsOverAge sets IsOverAge field to given value.

### HasIsOverAge

`func (o *PolandMobywatelMatchProviderOutput) HasIsOverAge() bool`

HasIsOverAge returns a boolean if a field has been set.

### SetIsOverAgeNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsOverAgeNil(b bool)`

 SetIsOverAgeNil sets the value for IsOverAge to be an explicit nil

### UnsetIsOverAge
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsOverAge()`

UnsetIsOverAge ensures that no value is present for IsOverAge, not even an explicit nil
### GetIsAgeEstimationMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) GetIsAgeEstimationMatchForSelfie() bool`

GetIsAgeEstimationMatchForSelfie returns the IsAgeEstimationMatchForSelfie field if non-nil, zero value otherwise.

### GetIsAgeEstimationMatchForSelfieOk

`func (o *PolandMobywatelMatchProviderOutput) GetIsAgeEstimationMatchForSelfieOk() (*bool, bool)`

GetIsAgeEstimationMatchForSelfieOk returns a tuple with the IsAgeEstimationMatchForSelfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAgeEstimationMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) SetIsAgeEstimationMatchForSelfie(v bool)`

SetIsAgeEstimationMatchForSelfie sets IsAgeEstimationMatchForSelfie field to given value.

### HasIsAgeEstimationMatchForSelfie

`func (o *PolandMobywatelMatchProviderOutput) HasIsAgeEstimationMatchForSelfie() bool`

HasIsAgeEstimationMatchForSelfie returns a boolean if a field has been set.

### SetIsAgeEstimationMatchForSelfieNil

`func (o *PolandMobywatelMatchProviderOutput) SetIsAgeEstimationMatchForSelfieNil(b bool)`

 SetIsAgeEstimationMatchForSelfieNil sets the value for IsAgeEstimationMatchForSelfie to be an explicit nil

### UnsetIsAgeEstimationMatchForSelfie
`func (o *PolandMobywatelMatchProviderOutput) UnsetIsAgeEstimationMatchForSelfie()`

UnsetIsAgeEstimationMatchForSelfie ensures that no value is present for IsAgeEstimationMatchForSelfie, not even an explicit nil
### GetAgeEstimationThreshold

`func (o *PolandMobywatelMatchProviderOutput) GetAgeEstimationThreshold() int32`

GetAgeEstimationThreshold returns the AgeEstimationThreshold field if non-nil, zero value otherwise.

### GetAgeEstimationThresholdOk

`func (o *PolandMobywatelMatchProviderOutput) GetAgeEstimationThresholdOk() (*int32, bool)`

GetAgeEstimationThresholdOk returns a tuple with the AgeEstimationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeEstimationThreshold

`func (o *PolandMobywatelMatchProviderOutput) SetAgeEstimationThreshold(v int32)`

SetAgeEstimationThreshold sets AgeEstimationThreshold field to given value.

### HasAgeEstimationThreshold

`func (o *PolandMobywatelMatchProviderOutput) HasAgeEstimationThreshold() bool`

HasAgeEstimationThreshold returns a boolean if a field has been set.

### SetAgeEstimationThresholdNil

`func (o *PolandMobywatelMatchProviderOutput) SetAgeEstimationThresholdNil(b bool)`

 SetAgeEstimationThresholdNil sets the value for AgeEstimationThreshold to be an explicit nil

### UnsetAgeEstimationThreshold
`func (o *PolandMobywatelMatchProviderOutput) UnsetAgeEstimationThreshold()`

UnsetAgeEstimationThreshold ensures that no value is present for AgeEstimationThreshold, not even an explicit nil
### GetEstimatedAgeFromSelfie

`func (o *PolandMobywatelMatchProviderOutput) GetEstimatedAgeFromSelfie() int32`

GetEstimatedAgeFromSelfie returns the EstimatedAgeFromSelfie field if non-nil, zero value otherwise.

### GetEstimatedAgeFromSelfieOk

`func (o *PolandMobywatelMatchProviderOutput) GetEstimatedAgeFromSelfieOk() (*int32, bool)`

GetEstimatedAgeFromSelfieOk returns a tuple with the EstimatedAgeFromSelfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAgeFromSelfie

`func (o *PolandMobywatelMatchProviderOutput) SetEstimatedAgeFromSelfie(v int32)`

SetEstimatedAgeFromSelfie sets EstimatedAgeFromSelfie field to given value.

### HasEstimatedAgeFromSelfie

`func (o *PolandMobywatelMatchProviderOutput) HasEstimatedAgeFromSelfie() bool`

HasEstimatedAgeFromSelfie returns a boolean if a field has been set.

### SetEstimatedAgeFromSelfieNil

`func (o *PolandMobywatelMatchProviderOutput) SetEstimatedAgeFromSelfieNil(b bool)`

 SetEstimatedAgeFromSelfieNil sets the value for EstimatedAgeFromSelfie to be an explicit nil

### UnsetEstimatedAgeFromSelfie
`func (o *PolandMobywatelMatchProviderOutput) UnsetEstimatedAgeFromSelfie()`

UnsetEstimatedAgeFromSelfie ensures that no value is present for EstimatedAgeFromSelfie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



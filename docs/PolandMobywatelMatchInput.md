# PolandMobywatelMatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | Pointer to **NullableString** | Given name as it appears on mobile ID (mDowód). | [optional] 
**FamilyName** | Pointer to **NullableString** | Current legal family name (nazwisko) as it appears on mobile ID (mDowód).              In Poland the current legal family name (nazwisko) is a separate idea from your birth family name (nazwisko rodowe). They often match, but they can differ after marriage, adoption, or a court-ordered change. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth as it appears on mobile ID (mDowód). | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality as it appears on mobile ID (mDowód). | [optional] 
**PersonalNumber** | Pointer to **NullableString** | Polish national identification number (PESEL) as it appears on mobile ID (mDowód).              NOTE: The provided example value is a randomly generated, but valid PESEL number that does not correspond to a real person. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | Document expiration date as it appears on mobile ID (mDowód). | [optional] 
**SelfieBytes** | Pointer to **NullableString** | The raw bytes of the selfie image collected from the user. | [optional] 
**SelfieImageMimeType** | Pointer to **NullableString** | The MIME Type of the file contained in SelfieBytes.              Only JPEG or PNG formats are supported. | [optional] 

## Methods

### NewPolandMobywatelMatchInput

`func NewPolandMobywatelMatchInput() *PolandMobywatelMatchInput`

NewPolandMobywatelMatchInput instantiates a new PolandMobywatelMatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolandMobywatelMatchInputWithDefaults

`func NewPolandMobywatelMatchInputWithDefaults() *PolandMobywatelMatchInput`

NewPolandMobywatelMatchInputWithDefaults instantiates a new PolandMobywatelMatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *PolandMobywatelMatchInput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PolandMobywatelMatchInput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PolandMobywatelMatchInput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PolandMobywatelMatchInput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *PolandMobywatelMatchInput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *PolandMobywatelMatchInput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *PolandMobywatelMatchInput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PolandMobywatelMatchInput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PolandMobywatelMatchInput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PolandMobywatelMatchInput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *PolandMobywatelMatchInput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *PolandMobywatelMatchInput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetDateOfBirth

`func (o *PolandMobywatelMatchInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PolandMobywatelMatchInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PolandMobywatelMatchInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PolandMobywatelMatchInput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *PolandMobywatelMatchInput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *PolandMobywatelMatchInput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationality

`func (o *PolandMobywatelMatchInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PolandMobywatelMatchInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PolandMobywatelMatchInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PolandMobywatelMatchInput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *PolandMobywatelMatchInput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *PolandMobywatelMatchInput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetPersonalNumber

`func (o *PolandMobywatelMatchInput) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *PolandMobywatelMatchInput) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *PolandMobywatelMatchInput) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *PolandMobywatelMatchInput) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.

### SetPersonalNumberNil

`func (o *PolandMobywatelMatchInput) SetPersonalNumberNil(b bool)`

 SetPersonalNumberNil sets the value for PersonalNumber to be an explicit nil

### UnsetPersonalNumber
`func (o *PolandMobywatelMatchInput) UnsetPersonalNumber()`

UnsetPersonalNumber ensures that no value is present for PersonalNumber, not even an explicit nil
### GetExpirationDate

`func (o *PolandMobywatelMatchInput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PolandMobywatelMatchInput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PolandMobywatelMatchInput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PolandMobywatelMatchInput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PolandMobywatelMatchInput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PolandMobywatelMatchInput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetSelfieBytes

`func (o *PolandMobywatelMatchInput) GetSelfieBytes() string`

GetSelfieBytes returns the SelfieBytes field if non-nil, zero value otherwise.

### GetSelfieBytesOk

`func (o *PolandMobywatelMatchInput) GetSelfieBytesOk() (*string, bool)`

GetSelfieBytesOk returns a tuple with the SelfieBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieBytes

`func (o *PolandMobywatelMatchInput) SetSelfieBytes(v string)`

SetSelfieBytes sets SelfieBytes field to given value.

### HasSelfieBytes

`func (o *PolandMobywatelMatchInput) HasSelfieBytes() bool`

HasSelfieBytes returns a boolean if a field has been set.

### SetSelfieBytesNil

`func (o *PolandMobywatelMatchInput) SetSelfieBytesNil(b bool)`

 SetSelfieBytesNil sets the value for SelfieBytes to be an explicit nil

### UnsetSelfieBytes
`func (o *PolandMobywatelMatchInput) UnsetSelfieBytes()`

UnsetSelfieBytes ensures that no value is present for SelfieBytes, not even an explicit nil
### GetSelfieImageMimeType

`func (o *PolandMobywatelMatchInput) GetSelfieImageMimeType() string`

GetSelfieImageMimeType returns the SelfieImageMimeType field if non-nil, zero value otherwise.

### GetSelfieImageMimeTypeOk

`func (o *PolandMobywatelMatchInput) GetSelfieImageMimeTypeOk() (*string, bool)`

GetSelfieImageMimeTypeOk returns a tuple with the SelfieImageMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfieImageMimeType

`func (o *PolandMobywatelMatchInput) SetSelfieImageMimeType(v string)`

SetSelfieImageMimeType sets SelfieImageMimeType field to given value.

### HasSelfieImageMimeType

`func (o *PolandMobywatelMatchInput) HasSelfieImageMimeType() bool`

HasSelfieImageMimeType returns a boolean if a field has been set.

### SetSelfieImageMimeTypeNil

`func (o *PolandMobywatelMatchInput) SetSelfieImageMimeTypeNil(b bool)`

 SetSelfieImageMimeTypeNil sets the value for SelfieImageMimeType to be an explicit nil

### UnsetSelfieImageMimeType
`func (o *PolandMobywatelMatchInput) UnsetSelfieImageMimeType()`

UnsetSelfieImageMimeType ensures that no value is present for SelfieImageMimeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



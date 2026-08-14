# ClearProviderOutputWatchlistHitDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** | Names associated with the watchlist hit. | [optional] 
**Alias** | Pointer to **[]string** | Aliases associated with the watchlist hit. | [optional] 
**Country** | Pointer to **[]string** | ISO 3166-1 alpha-2 country codes associated with the watchlist hit. | [optional] 
**Address** | Pointer to [**[]ClearProviderOutputAddress**](ClearProviderOutputAddress.md) | Addresses associated with the watchlist hit. | [optional] 
**DateOfBirth** | Pointer to [**[]ClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | Dates of birth associated with the watchlist hit. | [optional] 
**DateOfDeath** | Pointer to [**[]ClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | Dates of death associated with the watchlist hit. | [optional] 
**PlaceOfBirth** | Pointer to **[]string** | Places of birth associated with the watchlist hit. | [optional] 
**Gender** | Pointer to **[]string** | Genders associated with the watchlist hit, as represented in the source lists. | [optional] 
**Nationality** | Pointer to **[]string** | ISO 3166-1 alpha-2 nationalities associated with the watchlist hit. | [optional] 
**Position** | Pointer to **[]string** | Positions associated with the watchlist hit. | [optional] 
**PassportNumber** | Pointer to **[]string** | Passport numbers associated with the watchlist hit. | [optional] 
**IdNumber** | Pointer to **[]string** | Identification numbers associated with the watchlist hit. | [optional] 
**Notes** | Pointer to **[]string** | Notes from the watchlist source data. | [optional] 
**CreatedAt** | Pointer to [**[]ClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | Dates when the source entries were created. | [optional] 
**ModifiedAt** | Pointer to [**[]ClearProviderOutputPartialDate**](ClearProviderOutputPartialDate.md) | Dates when the source entries were last modified. | [optional] 
**SourceUrls** | Pointer to **[]string** | Source URLs for the watchlist hit. | [optional] 

## Methods

### NewClearProviderOutputWatchlistHitDetails

`func NewClearProviderOutputWatchlistHitDetails() *ClearProviderOutputWatchlistHitDetails`

NewClearProviderOutputWatchlistHitDetails instantiates a new ClearProviderOutputWatchlistHitDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearProviderOutputWatchlistHitDetailsWithDefaults

`func NewClearProviderOutputWatchlistHitDetailsWithDefaults() *ClearProviderOutputWatchlistHitDetails`

NewClearProviderOutputWatchlistHitDetailsWithDefaults instantiates a new ClearProviderOutputWatchlistHitDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClearProviderOutputWatchlistHitDetails) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClearProviderOutputWatchlistHitDetails) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ClearProviderOutputWatchlistHitDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAlias

`func (o *ClearProviderOutputWatchlistHitDetails) GetAlias() []string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetAliasOk() (*[]string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ClearProviderOutputWatchlistHitDetails) SetAlias(v []string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ClearProviderOutputWatchlistHitDetails) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetCountry

`func (o *ClearProviderOutputWatchlistHitDetails) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClearProviderOutputWatchlistHitDetails) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ClearProviderOutputWatchlistHitDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetAddress

`func (o *ClearProviderOutputWatchlistHitDetails) GetAddress() []ClearProviderOutputAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetAddressOk() (*[]ClearProviderOutputAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClearProviderOutputWatchlistHitDetails) SetAddress(v []ClearProviderOutputAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClearProviderOutputWatchlistHitDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDateOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) GetDateOfBirth() []ClearProviderOutputPartialDate`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetDateOfBirthOk() (*[]ClearProviderOutputPartialDate, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) SetDateOfBirth(v []ClearProviderOutputPartialDate)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetDateOfDeath

`func (o *ClearProviderOutputWatchlistHitDetails) GetDateOfDeath() []ClearProviderOutputPartialDate`

GetDateOfDeath returns the DateOfDeath field if non-nil, zero value otherwise.

### GetDateOfDeathOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetDateOfDeathOk() (*[]ClearProviderOutputPartialDate, bool)`

GetDateOfDeathOk returns a tuple with the DateOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDeath

`func (o *ClearProviderOutputWatchlistHitDetails) SetDateOfDeath(v []ClearProviderOutputPartialDate)`

SetDateOfDeath sets DateOfDeath field to given value.

### HasDateOfDeath

`func (o *ClearProviderOutputWatchlistHitDetails) HasDateOfDeath() bool`

HasDateOfDeath returns a boolean if a field has been set.

### SetDateOfDeathNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetDateOfDeathNil(b bool)`

 SetDateOfDeathNil sets the value for DateOfDeath to be an explicit nil

### UnsetDateOfDeath
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetDateOfDeath()`

UnsetDateOfDeath ensures that no value is present for DateOfDeath, not even an explicit nil
### GetPlaceOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) GetPlaceOfBirth() []string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetPlaceOfBirthOk() (*[]string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) SetPlaceOfBirth(v []string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ClearProviderOutputWatchlistHitDetails) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetGender

`func (o *ClearProviderOutputWatchlistHitDetails) GetGender() []string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetGenderOk() (*[]string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ClearProviderOutputWatchlistHitDetails) SetGender(v []string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ClearProviderOutputWatchlistHitDetails) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetNationality

`func (o *ClearProviderOutputWatchlistHitDetails) GetNationality() []string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetNationalityOk() (*[]string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ClearProviderOutputWatchlistHitDetails) SetNationality(v []string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ClearProviderOutputWatchlistHitDetails) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetPosition

`func (o *ClearProviderOutputWatchlistHitDetails) GetPosition() []string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetPositionOk() (*[]string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ClearProviderOutputWatchlistHitDetails) SetPosition(v []string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ClearProviderOutputWatchlistHitDetails) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetPassportNumber

`func (o *ClearProviderOutputWatchlistHitDetails) GetPassportNumber() []string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetPassportNumberOk() (*[]string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *ClearProviderOutputWatchlistHitDetails) SetPassportNumber(v []string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *ClearProviderOutputWatchlistHitDetails) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### SetPassportNumberNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetPassportNumberNil(b bool)`

 SetPassportNumberNil sets the value for PassportNumber to be an explicit nil

### UnsetPassportNumber
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetPassportNumber()`

UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
### GetIdNumber

`func (o *ClearProviderOutputWatchlistHitDetails) GetIdNumber() []string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetIdNumberOk() (*[]string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *ClearProviderOutputWatchlistHitDetails) SetIdNumber(v []string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *ClearProviderOutputWatchlistHitDetails) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### SetIdNumberNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetIdNumberNil(b bool)`

 SetIdNumberNil sets the value for IdNumber to be an explicit nil

### UnsetIdNumber
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetIdNumber()`

UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
### GetNotes

`func (o *ClearProviderOutputWatchlistHitDetails) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ClearProviderOutputWatchlistHitDetails) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ClearProviderOutputWatchlistHitDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCreatedAt

`func (o *ClearProviderOutputWatchlistHitDetails) GetCreatedAt() []ClearProviderOutputPartialDate`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetCreatedAtOk() (*[]ClearProviderOutputPartialDate, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClearProviderOutputWatchlistHitDetails) SetCreatedAt(v []ClearProviderOutputPartialDate)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClearProviderOutputWatchlistHitDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedAt

`func (o *ClearProviderOutputWatchlistHitDetails) GetModifiedAt() []ClearProviderOutputPartialDate`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetModifiedAtOk() (*[]ClearProviderOutputPartialDate, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClearProviderOutputWatchlistHitDetails) SetModifiedAt(v []ClearProviderOutputPartialDate)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ClearProviderOutputWatchlistHitDetails) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetSourceUrls

`func (o *ClearProviderOutputWatchlistHitDetails) GetSourceUrls() []string`

GetSourceUrls returns the SourceUrls field if non-nil, zero value otherwise.

### GetSourceUrlsOk

`func (o *ClearProviderOutputWatchlistHitDetails) GetSourceUrlsOk() (*[]string, bool)`

GetSourceUrlsOk returns a tuple with the SourceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrls

`func (o *ClearProviderOutputWatchlistHitDetails) SetSourceUrls(v []string)`

SetSourceUrls sets SourceUrls field to given value.

### HasSourceUrls

`func (o *ClearProviderOutputWatchlistHitDetails) HasSourceUrls() bool`

HasSourceUrls returns a boolean if a field has been set.

### SetSourceUrlsNil

`func (o *ClearProviderOutputWatchlistHitDetails) SetSourceUrlsNil(b bool)`

 SetSourceUrlsNil sets the value for SourceUrls to be an explicit nil

### UnsetSourceUrls
`func (o *ClearProviderOutputWatchlistHitDetails) UnsetSourceUrls()`

UnsetSourceUrls ensures that no value is present for SourceUrls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



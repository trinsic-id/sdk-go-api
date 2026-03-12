# UnitedKingdomEvisaProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name of the individual as recorded on the eVisa. | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the individual as recorded on the eVisa. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**ShareCode** | Pointer to **NullableString** | The 9-character share code used for the verification. | [optional] 
**ReferenceNumber** | Pointer to **NullableString** | The eVisa reference number issued by UK Visas and Immigration (UKVI). | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date of the eVisa. Only present for immigration_status verifications. | [optional] 
**ValidFrom** | Pointer to **NullableString** | The start date from which the eVisa is valid. Only returned for immigration_status and right_to_rent verifications — not returned for right_to_work. | [optional] 
**Nationality** | Pointer to **NullableString** | The nationality of the individual as recorded on the eVisa. Only present for immigration_status verifications. | [optional] 
**ImmigrationStatus** | Pointer to **NullableString** | The immigration status of the individual. Only present for immigration_status verifications. This is a personalized record for each individual, and there is no single definitive list. Some common values: - Settled - Pre-settled - Student - Graduate - Youth Mobility - Global Talent - Skilled Worker - Limited leave | [optional] 
**Outcome** | Pointer to **NullableString** | Overall verification outcome: \&quot;pass\&quot; or \&quot;fail\&quot;. | [optional] 
**EvidenceType** | Pointer to **NullableString** | Type of verification. List of possible values: - immigration_status - right_to_work - right_to_rent | [optional] 

## Methods

### NewUnitedKingdomEvisaProviderOutput

`func NewUnitedKingdomEvisaProviderOutput() *UnitedKingdomEvisaProviderOutput`

NewUnitedKingdomEvisaProviderOutput instantiates a new UnitedKingdomEvisaProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitedKingdomEvisaProviderOutputWithDefaults

`func NewUnitedKingdomEvisaProviderOutputWithDefaults() *UnitedKingdomEvisaProviderOutput`

NewUnitedKingdomEvisaProviderOutputWithDefaults instantiates a new UnitedKingdomEvisaProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UnitedKingdomEvisaProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnitedKingdomEvisaProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnitedKingdomEvisaProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnitedKingdomEvisaProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UnitedKingdomEvisaProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UnitedKingdomEvisaProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UnitedKingdomEvisaProviderOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnitedKingdomEvisaProviderOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnitedKingdomEvisaProviderOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnitedKingdomEvisaProviderOutput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UnitedKingdomEvisaProviderOutput) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UnitedKingdomEvisaProviderOutput) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDateOfBirth

`func (o *UnitedKingdomEvisaProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UnitedKingdomEvisaProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UnitedKingdomEvisaProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UnitedKingdomEvisaProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *UnitedKingdomEvisaProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *UnitedKingdomEvisaProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetShareCode

`func (o *UnitedKingdomEvisaProviderOutput) GetShareCode() string`

GetShareCode returns the ShareCode field if non-nil, zero value otherwise.

### GetShareCodeOk

`func (o *UnitedKingdomEvisaProviderOutput) GetShareCodeOk() (*string, bool)`

GetShareCodeOk returns a tuple with the ShareCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareCode

`func (o *UnitedKingdomEvisaProviderOutput) SetShareCode(v string)`

SetShareCode sets ShareCode field to given value.

### HasShareCode

`func (o *UnitedKingdomEvisaProviderOutput) HasShareCode() bool`

HasShareCode returns a boolean if a field has been set.

### SetShareCodeNil

`func (o *UnitedKingdomEvisaProviderOutput) SetShareCodeNil(b bool)`

 SetShareCodeNil sets the value for ShareCode to be an explicit nil

### UnsetShareCode
`func (o *UnitedKingdomEvisaProviderOutput) UnsetShareCode()`

UnsetShareCode ensures that no value is present for ShareCode, not even an explicit nil
### GetReferenceNumber

`func (o *UnitedKingdomEvisaProviderOutput) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *UnitedKingdomEvisaProviderOutput) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *UnitedKingdomEvisaProviderOutput) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *UnitedKingdomEvisaProviderOutput) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *UnitedKingdomEvisaProviderOutput) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *UnitedKingdomEvisaProviderOutput) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetExpirationDate

`func (o *UnitedKingdomEvisaProviderOutput) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UnitedKingdomEvisaProviderOutput) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UnitedKingdomEvisaProviderOutput) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UnitedKingdomEvisaProviderOutput) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *UnitedKingdomEvisaProviderOutput) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *UnitedKingdomEvisaProviderOutput) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetValidFrom

`func (o *UnitedKingdomEvisaProviderOutput) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *UnitedKingdomEvisaProviderOutput) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *UnitedKingdomEvisaProviderOutput) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *UnitedKingdomEvisaProviderOutput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFromNil

`func (o *UnitedKingdomEvisaProviderOutput) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *UnitedKingdomEvisaProviderOutput) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetNationality

`func (o *UnitedKingdomEvisaProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UnitedKingdomEvisaProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UnitedKingdomEvisaProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UnitedKingdomEvisaProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *UnitedKingdomEvisaProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *UnitedKingdomEvisaProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetImmigrationStatus

`func (o *UnitedKingdomEvisaProviderOutput) GetImmigrationStatus() string`

GetImmigrationStatus returns the ImmigrationStatus field if non-nil, zero value otherwise.

### GetImmigrationStatusOk

`func (o *UnitedKingdomEvisaProviderOutput) GetImmigrationStatusOk() (*string, bool)`

GetImmigrationStatusOk returns a tuple with the ImmigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmigrationStatus

`func (o *UnitedKingdomEvisaProviderOutput) SetImmigrationStatus(v string)`

SetImmigrationStatus sets ImmigrationStatus field to given value.

### HasImmigrationStatus

`func (o *UnitedKingdomEvisaProviderOutput) HasImmigrationStatus() bool`

HasImmigrationStatus returns a boolean if a field has been set.

### SetImmigrationStatusNil

`func (o *UnitedKingdomEvisaProviderOutput) SetImmigrationStatusNil(b bool)`

 SetImmigrationStatusNil sets the value for ImmigrationStatus to be an explicit nil

### UnsetImmigrationStatus
`func (o *UnitedKingdomEvisaProviderOutput) UnsetImmigrationStatus()`

UnsetImmigrationStatus ensures that no value is present for ImmigrationStatus, not even an explicit nil
### GetOutcome

`func (o *UnitedKingdomEvisaProviderOutput) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *UnitedKingdomEvisaProviderOutput) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *UnitedKingdomEvisaProviderOutput) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *UnitedKingdomEvisaProviderOutput) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *UnitedKingdomEvisaProviderOutput) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *UnitedKingdomEvisaProviderOutput) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetEvidenceType

`func (o *UnitedKingdomEvisaProviderOutput) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *UnitedKingdomEvisaProviderOutput) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *UnitedKingdomEvisaProviderOutput) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *UnitedKingdomEvisaProviderOutput) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *UnitedKingdomEvisaProviderOutput) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *UnitedKingdomEvisaProviderOutput) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



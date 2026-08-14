# SouthAfricaNidLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | Pointer to **NullableString** | The South African National Identity Number (13 digits).              Issued for life by the Department of Home Affairs (DHA) and stored in the HANIS (Home Affairs National Identification System) database. The same number is mandatory for banking, employment, taxation, and voting, and is printed on both the legacy green ID book and the Smart ID Card (rolled out from 2013 onward).              Format: - YYMMDD G(4) C A Z - YYMMDD is the date of birth - G(4) is the gender code (below 5000 female, 5000 or above male) - C is the citizenship indicator (0 citizen, 1 permanent resident) - A is reserved (it had a politically sensitive meaning in the past, but is currently   semantically meaningless) - Z is a Luhn check digit | [optional] 
**CitizenshipStatus** | Pointer to **NullableString** | Citizenship status of the individual.              Possible values: - Citizen - PermanentResident - Refugee | [optional] 
**FirstName** | Pointer to **NullableString** | First name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | Family name of the individual. | [optional] 
**FullName** | Pointer to **NullableString** | Full name of the individual. | [optional] 
**Sex** | Pointer to **NullableString** | Sex of the individual.              Possible values: - Male - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | Date of birth of the individual. | [optional] 
**Nationality** | Pointer to **NullableString** | Nationality of the individual as an ISO 3166-1 alpha-2 country code.              Set to ZA when CitizenshipStatus is Citizen. Null when not a South African citizen (Permanent Resident or Refugee). | [optional] 
**DateOfDeath** | Pointer to **NullableString** | Date of death of the individual, if recorded by DHA (Department of Home Affairs).              Null if the individual is alive, or the data is missing from the registry. | [optional] 
**BirthCountry** | Pointer to **NullableString** | Country of birth of the individual as an ISO 3166-1 alpha-2 country code. | [optional] 

## Methods

### NewSouthAfricaNidLookup2ProviderOutput

`func NewSouthAfricaNidLookup2ProviderOutput() *SouthAfricaNidLookup2ProviderOutput`

NewSouthAfricaNidLookup2ProviderOutput instantiates a new SouthAfricaNidLookup2ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSouthAfricaNidLookup2ProviderOutputWithDefaults

`func NewSouthAfricaNidLookup2ProviderOutputWithDefaults() *SouthAfricaNidLookup2ProviderOutput`

NewSouthAfricaNidLookup2ProviderOutputWithDefaults instantiates a new SouthAfricaNidLookup2ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdNumber

`func (o *SouthAfricaNidLookup2ProviderOutput) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *SouthAfricaNidLookup2ProviderOutput) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *SouthAfricaNidLookup2ProviderOutput) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetCitizenshipStatus

`func (o *SouthAfricaNidLookup2ProviderOutput) GetCitizenshipStatus() string`

GetCitizenshipStatus returns the CitizenshipStatus field if non-nil, zero value otherwise.

### GetCitizenshipStatusOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetCitizenshipStatusOk() (*string, bool)`

GetCitizenshipStatusOk returns a tuple with the CitizenshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipStatus

`func (o *SouthAfricaNidLookup2ProviderOutput) SetCitizenshipStatus(v string)`

SetCitizenshipStatus sets CitizenshipStatus field to given value.

### HasCitizenshipStatus

`func (o *SouthAfricaNidLookup2ProviderOutput) HasCitizenshipStatus() bool`

HasCitizenshipStatus returns a boolean if a field has been set.

### SetCitizenshipStatusNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetCitizenshipStatusNil(b bool)`

 SetCitizenshipStatusNil sets the value for CitizenshipStatus to be an explicit nil

### UnsetCitizenshipStatus
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetCitizenshipStatus()`

UnsetCitizenshipStatus ensures that no value is present for CitizenshipStatus, not even an explicit nil
### GetFirstName

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SouthAfricaNidLookup2ProviderOutput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetFamilyName

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *SouthAfricaNidLookup2ProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFullName

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SouthAfricaNidLookup2ProviderOutput) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetSex

`func (o *SouthAfricaNidLookup2ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SouthAfricaNidLookup2ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SouthAfricaNidLookup2ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *SouthAfricaNidLookup2ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SouthAfricaNidLookup2ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SouthAfricaNidLookup2ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationality

`func (o *SouthAfricaNidLookup2ProviderOutput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *SouthAfricaNidLookup2ProviderOutput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *SouthAfricaNidLookup2ProviderOutput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetDateOfDeath

`func (o *SouthAfricaNidLookup2ProviderOutput) GetDateOfDeath() string`

GetDateOfDeath returns the DateOfDeath field if non-nil, zero value otherwise.

### GetDateOfDeathOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetDateOfDeathOk() (*string, bool)`

GetDateOfDeathOk returns a tuple with the DateOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDeath

`func (o *SouthAfricaNidLookup2ProviderOutput) SetDateOfDeath(v string)`

SetDateOfDeath sets DateOfDeath field to given value.

### HasDateOfDeath

`func (o *SouthAfricaNidLookup2ProviderOutput) HasDateOfDeath() bool`

HasDateOfDeath returns a boolean if a field has been set.

### SetDateOfDeathNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetDateOfDeathNil(b bool)`

 SetDateOfDeathNil sets the value for DateOfDeath to be an explicit nil

### UnsetDateOfDeath
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetDateOfDeath()`

UnsetDateOfDeath ensures that no value is present for DateOfDeath, not even an explicit nil
### GetBirthCountry

`func (o *SouthAfricaNidLookup2ProviderOutput) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *SouthAfricaNidLookup2ProviderOutput) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *SouthAfricaNidLookup2ProviderOutput) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *SouthAfricaNidLookup2ProviderOutput) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *SouthAfricaNidLookup2ProviderOutput) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *SouthAfricaNidLookup2ProviderOutput) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



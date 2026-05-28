# SouthAfricaNidLookup2ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdNumber** | **string** | The South African National Identity Number (13 digits).              Issued for life by the Department of Home Affairs (DHA) and stored in the HANIS (Home Affairs National Identification System) database. The same number is mandatory for banking, employment, taxation, and voting, and is printed on both the legacy green ID book and the Smart ID Card (rolled out from 2013 onward).              Format: - YYMMDD G(4) C A Z - YYMMDD is the date of birth - G(4) is the gender code (below 5000 female, 5000 or above male) - C is the citizenship indicator (0 citizen, 1 permanent resident) - A is reserved (it had a politically sensitive meaning in the past, but is currently   semantically meaningless) - Z is a Luhn check digit | 
**CitizenshipStatus** | **string** | Citizenship status of the individual.              Possible values: - Citizen - PermanentResident | 
**FirstName** | **string** | First name of the individual. | 
**FamilyName** | **string** | Family name of the individual. | 
**FullName** | **string** | Full name of the individual. | 
**Sex** | **string** | Sex of the individual.              Possible values: - Male - Female | 
**DateOfBirth** | **string** | Date of birth of the individual. | 
**Nationality** | **string** | Nationality of the individual as an ISO 3166-1 alpha-2 country code. | 
**DateOfDeath** | Pointer to **NullableString** | Date of death of the individual, if recorded by DHA (Department of Home Affairs).              Null if the individual is alive, or the data is missing from the registry. | [optional] 
**BirthCountry** | Pointer to **NullableString** | Country of birth of the individual as an ISO 3166-1 alpha-2 country code. | [optional] 

## Methods

### NewSouthAfricaNidLookup2ProviderOutput

`func NewSouthAfricaNidLookup2ProviderOutput(nationalIdNumber string, citizenshipStatus string, firstName string, familyName string, fullName string, sex string, dateOfBirth string, nationality string, ) *SouthAfricaNidLookup2ProviderOutput`

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



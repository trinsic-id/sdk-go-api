# NigeriaNinLookup3ProviderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NationalIdentityNumber** | Pointer to **NullableString** | National Identification Number (NIN).              This is a unique, permanent identifier assigned by the National Identity Management Commission (NIMC) upon enrollment.              Format: - 11 numeric digits - No publicly known encoding scheme is used to encode personal information in the NIN - Last digit is a checksum using the Verhoeff algorithm | [optional] 
**GivenName** | Pointer to **NullableString** | The given name of the individual. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the individual. | [optional] 
**FamilyName** | Pointer to **NullableString** | The family name of the individual. | [optional] 
**Sex** | Pointer to **NullableString** | The sex of the individual.              Possible values: - Male - Female | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The date of birth of the individual. | [optional] 
**BirthCountry** | Pointer to **NullableString** | Country of birth as an ISO 3166-1 alpha-2 code. | [optional] 
**BirthState** | Pointer to **NullableString** | State of birth as recorded by National Identity Management Commission (NIMC). | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Phone number registered with National Identity Management Commission (NIMC).              Format: - International E.164 | [optional] 
**StreetAddress** | Pointer to **NullableString** | Street or residence line registered with National Identity Management Commission (NIMC).              This is the street-level address line only. Town, Local Government Area, and state are returned separately. | [optional] 
**ResidenceTown** | Pointer to **NullableString** | Town of residence registered with National Identity Management Commission (NIMC). | [optional] 
**LocalGovernmentArea** | Pointer to **NullableString** | Local Government Area of residence.              Nigeria is divided into 774 Local Government Areas (LGAs), which are the third-tier administrative divisions below states and the Federal Capital Territory. LGAs are roughly equivalent to counties or municipalities in other countries. | [optional] 
**State** | Pointer to **NullableString** | State of residence registered with National Identity Management Commission (NIMC). | [optional] 
**NextOfKinFirstName** | Pointer to **NullableString** | First name of the individual&#39;s next of kin. | [optional] 
**NextOfKinStreetAddress** | Pointer to **NullableString** | Street or residence line of the individual&#39;s next of kin. | [optional] 
**NextOfKinLocalGovernmentArea** | Pointer to **NullableString** | Local Government Area of the individual&#39;s next of kin. | [optional] 
**NextOfKinTown** | Pointer to **NullableString** | Town of residence of the individual&#39;s next of kin. | [optional] 

## Methods

### NewNigeriaNinLookup3ProviderOutput

`func NewNigeriaNinLookup3ProviderOutput() *NigeriaNinLookup3ProviderOutput`

NewNigeriaNinLookup3ProviderOutput instantiates a new NigeriaNinLookup3ProviderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNigeriaNinLookup3ProviderOutputWithDefaults

`func NewNigeriaNinLookup3ProviderOutputWithDefaults() *NigeriaNinLookup3ProviderOutput`

NewNigeriaNinLookup3ProviderOutputWithDefaults instantiates a new NigeriaNinLookup3ProviderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationalIdentityNumber

`func (o *NigeriaNinLookup3ProviderOutput) GetNationalIdentityNumber() string`

GetNationalIdentityNumber returns the NationalIdentityNumber field if non-nil, zero value otherwise.

### GetNationalIdentityNumberOk

`func (o *NigeriaNinLookup3ProviderOutput) GetNationalIdentityNumberOk() (*string, bool)`

GetNationalIdentityNumberOk returns a tuple with the NationalIdentityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentityNumber

`func (o *NigeriaNinLookup3ProviderOutput) SetNationalIdentityNumber(v string)`

SetNationalIdentityNumber sets NationalIdentityNumber field to given value.

### HasNationalIdentityNumber

`func (o *NigeriaNinLookup3ProviderOutput) HasNationalIdentityNumber() bool`

HasNationalIdentityNumber returns a boolean if a field has been set.

### SetNationalIdentityNumberNil

`func (o *NigeriaNinLookup3ProviderOutput) SetNationalIdentityNumberNil(b bool)`

 SetNationalIdentityNumberNil sets the value for NationalIdentityNumber to be an explicit nil

### UnsetNationalIdentityNumber
`func (o *NigeriaNinLookup3ProviderOutput) UnsetNationalIdentityNumber()`

UnsetNationalIdentityNumber ensures that no value is present for NationalIdentityNumber, not even an explicit nil
### GetGivenName

`func (o *NigeriaNinLookup3ProviderOutput) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *NigeriaNinLookup3ProviderOutput) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *NigeriaNinLookup3ProviderOutput) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *NigeriaNinLookup3ProviderOutput) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *NigeriaNinLookup3ProviderOutput) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *NigeriaNinLookup3ProviderOutput) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMiddleName

`func (o *NigeriaNinLookup3ProviderOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *NigeriaNinLookup3ProviderOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *NigeriaNinLookup3ProviderOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *NigeriaNinLookup3ProviderOutput) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *NigeriaNinLookup3ProviderOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *NigeriaNinLookup3ProviderOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetFamilyName

`func (o *NigeriaNinLookup3ProviderOutput) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *NigeriaNinLookup3ProviderOutput) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *NigeriaNinLookup3ProviderOutput) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *NigeriaNinLookup3ProviderOutput) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### SetFamilyNameNil

`func (o *NigeriaNinLookup3ProviderOutput) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *NigeriaNinLookup3ProviderOutput) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetSex

`func (o *NigeriaNinLookup3ProviderOutput) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *NigeriaNinLookup3ProviderOutput) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *NigeriaNinLookup3ProviderOutput) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *NigeriaNinLookup3ProviderOutput) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *NigeriaNinLookup3ProviderOutput) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *NigeriaNinLookup3ProviderOutput) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetDateOfBirth

`func (o *NigeriaNinLookup3ProviderOutput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *NigeriaNinLookup3ProviderOutput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *NigeriaNinLookup3ProviderOutput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *NigeriaNinLookup3ProviderOutput) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *NigeriaNinLookup3ProviderOutput) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *NigeriaNinLookup3ProviderOutput) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetBirthCountry

`func (o *NigeriaNinLookup3ProviderOutput) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *NigeriaNinLookup3ProviderOutput) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *NigeriaNinLookup3ProviderOutput) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *NigeriaNinLookup3ProviderOutput) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### SetBirthCountryNil

`func (o *NigeriaNinLookup3ProviderOutput) SetBirthCountryNil(b bool)`

 SetBirthCountryNil sets the value for BirthCountry to be an explicit nil

### UnsetBirthCountry
`func (o *NigeriaNinLookup3ProviderOutput) UnsetBirthCountry()`

UnsetBirthCountry ensures that no value is present for BirthCountry, not even an explicit nil
### GetBirthState

`func (o *NigeriaNinLookup3ProviderOutput) GetBirthState() string`

GetBirthState returns the BirthState field if non-nil, zero value otherwise.

### GetBirthStateOk

`func (o *NigeriaNinLookup3ProviderOutput) GetBirthStateOk() (*string, bool)`

GetBirthStateOk returns a tuple with the BirthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthState

`func (o *NigeriaNinLookup3ProviderOutput) SetBirthState(v string)`

SetBirthState sets BirthState field to given value.

### HasBirthState

`func (o *NigeriaNinLookup3ProviderOutput) HasBirthState() bool`

HasBirthState returns a boolean if a field has been set.

### SetBirthStateNil

`func (o *NigeriaNinLookup3ProviderOutput) SetBirthStateNil(b bool)`

 SetBirthStateNil sets the value for BirthState to be an explicit nil

### UnsetBirthState
`func (o *NigeriaNinLookup3ProviderOutput) UnsetBirthState()`

UnsetBirthState ensures that no value is present for BirthState, not even an explicit nil
### GetPhoneNumber

`func (o *NigeriaNinLookup3ProviderOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NigeriaNinLookup3ProviderOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NigeriaNinLookup3ProviderOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NigeriaNinLookup3ProviderOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *NigeriaNinLookup3ProviderOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *NigeriaNinLookup3ProviderOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *NigeriaNinLookup3ProviderOutput) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *NigeriaNinLookup3ProviderOutput) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *NigeriaNinLookup3ProviderOutput) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetResidenceTown

`func (o *NigeriaNinLookup3ProviderOutput) GetResidenceTown() string`

GetResidenceTown returns the ResidenceTown field if non-nil, zero value otherwise.

### GetResidenceTownOk

`func (o *NigeriaNinLookup3ProviderOutput) GetResidenceTownOk() (*string, bool)`

GetResidenceTownOk returns a tuple with the ResidenceTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceTown

`func (o *NigeriaNinLookup3ProviderOutput) SetResidenceTown(v string)`

SetResidenceTown sets ResidenceTown field to given value.

### HasResidenceTown

`func (o *NigeriaNinLookup3ProviderOutput) HasResidenceTown() bool`

HasResidenceTown returns a boolean if a field has been set.

### SetResidenceTownNil

`func (o *NigeriaNinLookup3ProviderOutput) SetResidenceTownNil(b bool)`

 SetResidenceTownNil sets the value for ResidenceTown to be an explicit nil

### UnsetResidenceTown
`func (o *NigeriaNinLookup3ProviderOutput) UnsetResidenceTown()`

UnsetResidenceTown ensures that no value is present for ResidenceTown, not even an explicit nil
### GetLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) GetLocalGovernmentArea() string`

GetLocalGovernmentArea returns the LocalGovernmentArea field if non-nil, zero value otherwise.

### GetLocalGovernmentAreaOk

`func (o *NigeriaNinLookup3ProviderOutput) GetLocalGovernmentAreaOk() (*string, bool)`

GetLocalGovernmentAreaOk returns a tuple with the LocalGovernmentArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) SetLocalGovernmentArea(v string)`

SetLocalGovernmentArea sets LocalGovernmentArea field to given value.

### HasLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) HasLocalGovernmentArea() bool`

HasLocalGovernmentArea returns a boolean if a field has been set.

### SetLocalGovernmentAreaNil

`func (o *NigeriaNinLookup3ProviderOutput) SetLocalGovernmentAreaNil(b bool)`

 SetLocalGovernmentAreaNil sets the value for LocalGovernmentArea to be an explicit nil

### UnsetLocalGovernmentArea
`func (o *NigeriaNinLookup3ProviderOutput) UnsetLocalGovernmentArea()`

UnsetLocalGovernmentArea ensures that no value is present for LocalGovernmentArea, not even an explicit nil
### GetState

`func (o *NigeriaNinLookup3ProviderOutput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NigeriaNinLookup3ProviderOutput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NigeriaNinLookup3ProviderOutput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NigeriaNinLookup3ProviderOutput) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *NigeriaNinLookup3ProviderOutput) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *NigeriaNinLookup3ProviderOutput) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetNextOfKinFirstName

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinFirstName() string`

GetNextOfKinFirstName returns the NextOfKinFirstName field if non-nil, zero value otherwise.

### GetNextOfKinFirstNameOk

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinFirstNameOk() (*string, bool)`

GetNextOfKinFirstNameOk returns a tuple with the NextOfKinFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOfKinFirstName

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinFirstName(v string)`

SetNextOfKinFirstName sets NextOfKinFirstName field to given value.

### HasNextOfKinFirstName

`func (o *NigeriaNinLookup3ProviderOutput) HasNextOfKinFirstName() bool`

HasNextOfKinFirstName returns a boolean if a field has been set.

### SetNextOfKinFirstNameNil

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinFirstNameNil(b bool)`

 SetNextOfKinFirstNameNil sets the value for NextOfKinFirstName to be an explicit nil

### UnsetNextOfKinFirstName
`func (o *NigeriaNinLookup3ProviderOutput) UnsetNextOfKinFirstName()`

UnsetNextOfKinFirstName ensures that no value is present for NextOfKinFirstName, not even an explicit nil
### GetNextOfKinStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinStreetAddress() string`

GetNextOfKinStreetAddress returns the NextOfKinStreetAddress field if non-nil, zero value otherwise.

### GetNextOfKinStreetAddressOk

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinStreetAddressOk() (*string, bool)`

GetNextOfKinStreetAddressOk returns a tuple with the NextOfKinStreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOfKinStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinStreetAddress(v string)`

SetNextOfKinStreetAddress sets NextOfKinStreetAddress field to given value.

### HasNextOfKinStreetAddress

`func (o *NigeriaNinLookup3ProviderOutput) HasNextOfKinStreetAddress() bool`

HasNextOfKinStreetAddress returns a boolean if a field has been set.

### SetNextOfKinStreetAddressNil

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinStreetAddressNil(b bool)`

 SetNextOfKinStreetAddressNil sets the value for NextOfKinStreetAddress to be an explicit nil

### UnsetNextOfKinStreetAddress
`func (o *NigeriaNinLookup3ProviderOutput) UnsetNextOfKinStreetAddress()`

UnsetNextOfKinStreetAddress ensures that no value is present for NextOfKinStreetAddress, not even an explicit nil
### GetNextOfKinLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinLocalGovernmentArea() string`

GetNextOfKinLocalGovernmentArea returns the NextOfKinLocalGovernmentArea field if non-nil, zero value otherwise.

### GetNextOfKinLocalGovernmentAreaOk

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinLocalGovernmentAreaOk() (*string, bool)`

GetNextOfKinLocalGovernmentAreaOk returns a tuple with the NextOfKinLocalGovernmentArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOfKinLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinLocalGovernmentArea(v string)`

SetNextOfKinLocalGovernmentArea sets NextOfKinLocalGovernmentArea field to given value.

### HasNextOfKinLocalGovernmentArea

`func (o *NigeriaNinLookup3ProviderOutput) HasNextOfKinLocalGovernmentArea() bool`

HasNextOfKinLocalGovernmentArea returns a boolean if a field has been set.

### SetNextOfKinLocalGovernmentAreaNil

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinLocalGovernmentAreaNil(b bool)`

 SetNextOfKinLocalGovernmentAreaNil sets the value for NextOfKinLocalGovernmentArea to be an explicit nil

### UnsetNextOfKinLocalGovernmentArea
`func (o *NigeriaNinLookup3ProviderOutput) UnsetNextOfKinLocalGovernmentArea()`

UnsetNextOfKinLocalGovernmentArea ensures that no value is present for NextOfKinLocalGovernmentArea, not even an explicit nil
### GetNextOfKinTown

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinTown() string`

GetNextOfKinTown returns the NextOfKinTown field if non-nil, zero value otherwise.

### GetNextOfKinTownOk

`func (o *NigeriaNinLookup3ProviderOutput) GetNextOfKinTownOk() (*string, bool)`

GetNextOfKinTownOk returns a tuple with the NextOfKinTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOfKinTown

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinTown(v string)`

SetNextOfKinTown sets NextOfKinTown field to given value.

### HasNextOfKinTown

`func (o *NigeriaNinLookup3ProviderOutput) HasNextOfKinTown() bool`

HasNextOfKinTown returns a boolean if a field has been set.

### SetNextOfKinTownNil

`func (o *NigeriaNinLookup3ProviderOutput) SetNextOfKinTownNil(b bool)`

 SetNextOfKinTownNil sets the value for NextOfKinTown to be an explicit nil

### UnsetNextOfKinTown
`func (o *NigeriaNinLookup3ProviderOutput) UnsetNextOfKinTown()`

UnsetNextOfKinTown ensures that no value is present for NextOfKinTown, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


